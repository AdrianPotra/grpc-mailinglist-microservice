/*
Author: Adrian Potra
Version: 1.0.

import the "github.com/alexflint/go-arg" package
*/

package main

import (
	"database/sql"
	"log"
	"server/grpcapi"
	mdb "server/maildb"
	"sync"

	"github.com/alexflint/go-arg"
)

var args struct {
	DBPath   string `arg:"env:MAILINGLIST_DB"`
	BindGrpc string `arg:"env:MAILINGLIST_BIND_GRPC"`
}

func main() {
	// parsing arguments
	arg.MustParse(&args)
	//setting defaults
	if args.DBPath == "" {
		args.DBPath = "list.db"
	}

	if args.BindGrpc == "" {
		args.BindGrpc = ":8081" // if we don't send ip address, it will default to localhost
	}
	log.Printf("using database '%v'\n", args.DBPath)
	db, err := sql.Open("sqlite3", args.DBPath)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	mdb.TryCreate(db)

	// creating go routine - starting our gRPC API server
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Printf("starting gRPC API server...\n")
		grpcapi.Serve(db, args.BindGrpc)
		wg.Done()
	}()
	wg.Wait()
}
