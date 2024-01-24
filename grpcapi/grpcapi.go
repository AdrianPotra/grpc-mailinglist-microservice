/* Author: Adrian Potra
   Version: 1.0
*/

package grpcapi

import (
	"context"
	"database/sql"
	"log"
	"net"
	mdb "server/maildb"
	pb "server/proto"
	"time"

	"google.golang.org/grpc"
)

// we will use type embedding on the mail server and when we embed the type, we can run it with the grpc code that was generated
type MailServer struct {
	pb.UnimplementedMailingListServiceServer // code that was generated as part of the protoc command
	db                                       *sql.DB
}

// func that converts a protoc buffer message into a data struct
func pbEntryToMdbEntry(pbEntry *pb.EmailEntry) mdb.EmailEntry {
	// we need to convert time - in protoc structure we have the time as type int, so we need to convert it to an actual time, in our case unix time
	t := time.Unix(pbEntry.ConfirmedAt, 0)
	// we create the data struct
	return mdb.EmailEntry{
		Id:          pbEntry.Id,
		Email:       pbEntry.Email,
		ConfirmedAt: &t,
		OptOut:      pbEntry.OptOut,
	}

}

// inverse func - converting an mdb entry into a protoc buffer message
func mdbEntryToPbEntry(mdbEntry *mdb.EmailEntry) pb.EmailEntry {
	return pb.EmailEntry{
		Id:          mdbEntry.Id,
		Email:       mdbEntry.Email,
		ConfirmedAt: mdbEntry.ConfirmedAt.Unix(), // it's of type int so we need to use the unix method to convert
		OptOut:      mdbEntry.OptOut,
	}
}

// creating endpoints for grpc api

// util func
func emailResponse(db *sql.DB, email string) (*pb.EmailResponse, error) {
	entry, err := mdb.GetEmail(db, email)
	if err != nil { // if error, return empty response along with the error
		return &pb.EmailResponse{}, err
	}

	if entry == nil { // if no email, we return an epty response
		return &pb.EmailResponse{}, nil
	}
	// convert our mdb entry email into a protoc buffer entry
	res := mdbEntryToPbEntry(entry)
	// we wrap our res into our email response type and we set the email entry to the response and we have nil for error
	return &pb.EmailResponse{EmailEntry: &res}, nil
}

// now we start to implement our interface - we will have to include a context as part of the specifications, but we don't actually use it here - it should contain metadata info about the request and its current status; usage could be for example for a long running response from the server, the context can be checked for status
func (s *MailServer) GetEmail(ctx context.Context, req *pb.GetEmailRequest) (*pb.EmailResponse, error) {
	log.Printf("gRPC GetEmail: %v\n", req)
	return emailResponse(s.db, req.EmailAddr)
}

func (s *MailServer) GetEmailBatch(ctx context.Context, req *pb.GetEmailBatchRequest) (*pb.GetEmailBatchResponse, error) {
	log.Printf("gRPC GetEmaiBatch: %v\n", req)
	// create query parameters
	params := mdb.GetEmailBatchQueryParams{
		Page:  int(req.Page),
		Count: int(req.Count),
	}
	// create the db
	mdbEntries, err := mdb.GetEmailBatch(s.db, params)
	if err != nil {
		return &pb.GetEmailBatchResponse{}, err
	}
	// we need to convert the mdb email entry struct into a protoc buffer entry - since we get back a slice of it, we need to loop through it to convert each slice element
	pbEntries := make([]*pb.EmailEntry, 0, len(mdbEntries))
	for i := 0; i < len(mdbEntries); i++ {
		entry := mdbEntryToPbEntry(&mdbEntries[i])
		pbEntries = append(pbEntries, &entry)
	}
	// now we return the entries
	return &pb.GetEmailBatchResponse{EmailEntries: pbEntries}, nil
}

func (s *MailServer) CreateEmail(ctx context.Context, req *pb.CreateEmailRequest) (*pb.EmailResponse, error) {
	log.Printf("gRPC CreateEmail: %v\n", req)
	// create email
	err := mdb.CreateEmail(s.db, req.EmailAddr)
	if err != nil {
		return &pb.EmailResponse{}, err
	}
	return emailResponse(s.db, req.EmailAddr)
}

func (s *MailServer) UpdateEmail(ctx context.Context, req *pb.UpdateEmailRequest) (*pb.EmailResponse, error) {
	log.Printf("gRPC UpdateEmail: %v\n", req)
	// when we do this request, we will have a protoc buffer type, so we need to convert to a mdb type to send to db
	entry := pbEntryToMdbEntry(req.EmailEntry)

	// update email
	err := mdb.UpdateEmail(s.db, entry)
	if err != nil {
		return &pb.EmailResponse{}, err
	}
	return emailResponse(s.db, entry.Email)
}

func (s *MailServer) DeleteEmail(ctx context.Context, req *pb.DeleteEmailRequest) (*pb.EmailResponse, error) {
	log.Printf("gRPC DeleteEmail: %v\n", req)

	// update email
	err := mdb.DeleteEmail(s.db, req.EmailAddr)
	if err != nil {
		return &pb.EmailResponse{}, err
	}
	return emailResponse(s.db, req.EmailAddr)
}

// func to start the server - we will run the net.listen method and that will listen to the bind address and it's goinf to bind an address using the TCP protoc

func Serve(db *sql.DB, bind string) {

	listener, err := net.Listen("tcp", bind)
	if err != nil { // if it fails, we're going to terminate the program because we won't be able to continue at all with the grpc server
		log.Fatalf("gRPC server error: failure to bind %v\n", bind)
	}
	// create the server
	grpcServer := grpc.NewServer()
	// create email server - or creaye our mail server struct which defined above
	mailServer := MailServer{db: db}

	pb.RegisterMailingListServiceServer(grpcServer, &mailServer) // we are using this generated function and what it does is it's taking the grpc server and tells it that we're using the email server as our grpc handler

	log.Printf("gRPC API server listening on %v\n", bind)
	// we start the server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("gRPC server error: %v\n", err)
	}
}
