# grpc-mailinglist-microservice

## Description
Microservice that manages an email list. 
It's using a gRPC server and a gRPC client to communicate with the server and it utilizes protocol buffers as communication format and utilizes SQLite as data storage. 
There is a small list in the project containing various email address samples to work with, especially when using GET requests. 

## WHY 
I'm familiar with REST API's using JSON as data format and recently I've heard about gRPC service and protocol buffers, which is supposed to be even faster over the wire than JSON, so it was a motivation for me to learn and to try to implement these technologies into a project. Last year at the workplace there were a few funny scenarios around the mass email communication topics, so that stuck with me somehow in my head and gave me the idea to actually implement a mailing list service in my own project. Funny enough, there was someone else in the coding community having kind of the same idea and he kind of outlined a small blueprint for the project, which was my inspiration on the steps I need to take for my own project. 

## Quick Start
There are a few requirements and/pr prerequisites for this project, outlined below: 
### Setup
This project requires a `gcc` compiler installed and the `protobuf` code generation tools.
### Install protobuf compiler
Install the `protoc` tool using the instructions available at [https://grpc.io/docs/protoc-installation/](https://grpc.io/docs/protoc-installation/).
Alternatively you can download a pre-built binary from [https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases) and placing the extracted binary somewhere in your `$PATH`.
### Install Go protobuf codegen tools
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Generate Go code from .proto files

```
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  Proto/mail.proto
```
On Windows 10 and 11, you will probably need to remove the backslashes

And as usual, run a 
```
go mod tidy
```
to be sure that all package dependencies are downloaded

## Usage

We perform CRUD operations on the database, like adding email, getting the email, updating email and delete, although in this scenario delete is not exactly a delete per se, since the action is assigned to set the opt-out flag to true (default is false), so that the users in the mailing list won't receive further communication when sending a mass email to the mailing list. At least this was the initial idea behind this project. 

Ideally you should split your terminal into 2, one for the server and one for the client. 
For the server, you can do a **go run ./server**  and you should notice the messages: 
  using database 'list.db'
  starting gRPC API server...
  gRPC API server listening on :8081
  
By default it runs on localhost on port 8081. 

Now you can run the client using  **go run ./client** in the other terminal, but before you do that, make sure that in the **client.go** file, on line 121 - newEmail := CreateEmail(client, "1329999@9999,99")  you put some other email address, as that one might be included in the database and the email address should be unique, so it won't let you add the same email twice and at the moment this is hardcoded on the client, along with some other parameters that I will mention below. 
If you do have a unique email, you should see in the server response how the email is created, updated, deleted (which is an opt-out flag in this case) and also to get multiple emails in a batch.
You can have pagination here in the batch, for example on lines 128 till 130 in the same **client.go** file, you will notice these lines of code: 
 <br>
getEmailBatch(client, 5, 1) // 5 emails on page 1
 <br>
getEmailBatch(client, 5, 2) // 5 emails on page 2
 <br>
getEmailBatch(client, 5, 3) // 5 emails on page 3

You can play around with the pagination and result rows per page. 


