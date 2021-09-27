package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	chat "example"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)
type server struct{
	chat.UnimplementedChatServiceServer
}
type server2 struct{
	chat.UnimplementedChatServiceTwoServer
}




// Unary Server code
func (*server)UnaryComms(ctx context.Context, message *chat.UnaryNormalMessage) (*chat.UnaryServerMessage, error){
	fmt.Printf("Recieved Message body from client: %s\n", message.Body)
	return &chat.UnaryServerMessage{Body: "unary Chat Data from Server", Language: "ESPN"}, nil
}

// Server Stream code
func (*server)ServerStreamComms(message *chat.StreamNormalMessage, stream chat.ChatService_ServerStreamCommsServer) (error){
	fmt.Printf("Recieved Message body from client: %s\n", message.Body)
	for _, number := range message.Body {
		data := []string{number}
		if err := stream.Send(&chat.StreamServerMessage{Body: data, Language: "EN-US"}); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}

// Client Stream Code
func (*server)ClientStreamComms(stream chat.ChatService_ClientStreamCommsServer) (error){
	// fmt.Printf("Recieved Message body from client: %s\n", message.Body)
	arr:= make([]string, 0)
	for {
		data, err := stream.Recv()
		
		if err == io.EOF {
			return stream.SendAndClose(&chat.StreamServerMessage{
				Body: arr,
				Language: "EU",
			})
		}
		if err != nil {
		  return err
		}
		arr = append(arr, data.Body[0])
		fmt.Printf("Number : %s, Language : %s\n",data.GetBody(), data.GetLanguage())

  }
}
// BiDirectional Stream Code
func (*server)BiDirectionStreamComms(stream chat.ChatService_BiDirectionStreamCommsServer) (error){
	go func() {
		

		for {
			in, err := stream.Recv()
			if err == io.EOF {
				log.Println("NO MORE DATA FROM CLIENT")
				break;
			}
			if err != nil {
				log.Panicln(err)
				return;
			}
			log.Println(in)

			
		}
	}()
	data := make([]string, 10)
	for i := range data {
		numStr := "SERVER : "
		numStr += strconv.Itoa(0 + i)
		time.Sleep(time.Second/4)
		err := stream.Send(&chat.StreamServerMessage{Body: []string{"numStr"}, Language: "EU"})
		if err != nil {
			log.Panicln(err)
			return err;
		}	
		log.Println("REQUEST SENT")
		time.Sleep(time.Second/4)
	}
	
	return nil
}


func main() {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Failed to listen to port 8080: %v", err)
	}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &server{})
	chat.RegisterChatServiceTwoServer(grpcServer, &server2{})
	// to serialize an deserialise data
	reflection.Register(grpcServer)
	fmt.Println("Server working on 8080")
	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("Failed to serve gRPC server over port 8080: %v", err)
		return
	}else{
		fmt.Println("Started gRPC server over port 8080")
	}

	
}