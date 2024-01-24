package main

import (
	"context"
	"log"
	pb "server/proto"
	"time"

	"github.com/alexflint/go-arg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// utility func to log the responses from server - server sends mostly email responses back
func logResponse(res *pb.EmailResponse, err error) {
	if err != nil {
		log.Fatalf("   error: %v", err)
	}
	if res.EmailEntry == nil {
		log.Printf("   email not found")
	} else {

		log.Printf("   response: %v", res.EmailEntry)
	}

}

func CreateEmail(client pb.MailingListServiceClient, addr string) *pb.EmailEntry { // we supply the addr, which is email to the generated pb mailing list service client grpc service

	log.Println("create email")
	// grpc server requires that we have a context with the request that we create - with timeout will cancel requests after a certain amount of time, and we're using time.second, so the request has 1 second to complete
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() // if it completes before 1 second, we will free up the resources
	// when we run protocl buffers file, it created both server and client functions to utilize - here we're using the cloent func to make a request of the server
	res, err := client.CreateEmail(ctx, &pb.CreateEmailRequest{EmailAddr: addr})
	logResponse(res, err)
	return res.EmailEntry
}

func getEmail(client pb.MailingListServiceClient, addr string) *pb.EmailEntry { // we supply the addr, which is email to the generated pb mailing list service client grpc service

	log.Println("get email")
	// grpc server requires that we have a context with the request that we create - with timeout will cancel requests after a certain amount of time, and we're using time.second, so the request has 1 second to complete
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() // if it completes before 1 second, we will free up the resources
	// when we run protocl buffers file, it created both server and client functions to utilize - here we're using the cloent func to make a request of the server
	res, err := client.GetEmail(ctx, &pb.GetEmailRequest{EmailAddr: addr})
	logResponse(res, err)
	return res.EmailEntry
}

func getEmailBatch(client pb.MailingListServiceClient, count int, page int) {

	log.Println("get email batch")
	// grpc server requires that we have a context with the request that we create - with timeout will cancel requests after a certain amount of time, and we're using time.second, so the request has 1 second to complete
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() // if it completes before 1 second, we will free up the resources
	// when we run protocl buffers file, it created both server and client functions to utilize - here we're using the cloent func to make a request of the server
	res, err := client.GetEmailBatch(ctx, &pb.GetEmailBatchRequest{Count: int32(count), Page: int32(page)})
	if err != nil {
		log.Fatalf("   error: %v", err)
	}
	log.Println("response:")
	for i := 0; i < len(res.EmailEntries); i++ {
		log.Printf("   item [%v of %v]: %s", i+1, len(res.EmailEntries), res.EmailEntries[i]) // we're printing which number it is, the total that we're getting back and the actual response itself
	}

}

func UpdateEmail(client pb.MailingListServiceClient, entry pb.EmailEntry) *pb.EmailEntry {

	log.Println("update email")
	// grpc server requires that we have a context with the request that we create - with timeout will cancel requests after a certain amount of time, and we're using time.second, so the request has 1 second to complete
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() // if it completes before 1 second, we will free up the resources
	// when we run protocl buffers file, it created both server and client functions to utilize - here we're using the cloent func to make a request of the server
	res, err := client.UpdateEmail(ctx, &pb.UpdateEmailRequest{EmailEntry: &entry})
	logResponse(res, err)
	return res.EmailEntry
}

func DeleteEmail(client pb.MailingListServiceClient, addr string) *pb.EmailEntry { // we supply the addr, which is email to the generated pb mailing list service client grpc service

	log.Println("delete email")
	// grpc server requires that we have a context with the request that we create - with timeout will cancel requests after a certain amount of time, and we're using time.second, so the request has 1 second to complete
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() // if it completes before 1 second, we will free up the resources
	// when we run protocl buffers file, it created both server and client functions to utilize - here we're using the cloent func to make a request of the server
	res, err := client.DeleteEmail(ctx, &pb.DeleteEmailRequest{EmailAddr: addr})
	logResponse(res, err)
	return res.EmailEntry
}

// now we create the cli args
var args struct {
	GrpcAddr string `arg:"env:MAILINGLIST_GRPC_ADDR"`
}

func main() {
	// parse the args
	arg.MustParse(&args)
	// set default
	if args.GrpcAddr == "" {
		args.GrpcAddr = ":8081"
	}
	//connect to the server
	conn, err := grpc.Dial(args.GrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials())) // we are connecting with an insecure connection here = no auth, no credentials etc
	if err != nil {
		log.Fatalf("did not connect: %v", err)

	}
	defer conn.Close()
	// create a new client - we use the below func to tell the existing grpc connection that the client is associated with the rpc messages that we defined within our mailing list service
	client := pb.NewMailingListServiceClient(conn)

	// now we create requests

	// TO MODIFY BELOW TO ACCEPT VALUES FROM USER INPUT AND NOT HARDCODED

	//create email
	newEmail := CreateEmail(client, "1329999@9999,99")
	// update it
	newEmail.ConfirmedAt = 10000
	UpdateEmail(client, *newEmail)
	//delete - or in our case modify opt out param
	DeleteEmail(client, newEmail.Email)
	// batch
	getEmailBatch(client, 5, 1) // 5 emails on page 1
	getEmailBatch(client, 5, 2) // 5 emails on page 2
	getEmailBatch(client, 5, 3) // 5 emails on page 3

}
