package main

import (
	"context"
	client_grpc "example"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
)


func runRPCs(id int){
	Host := "localhost:8080"
	switch(id) {
	case 1:
		// unary Chat Service
		conn, err := grpc.Dial(Host, grpc.WithInsecure())
		if err != nil {
			log.Println(err.Error())
		}
		defer conn.Close()
		// get Response
		c := client_grpc.NewChatServiceClient(conn)
		msg := client_grpc.UnaryNormalMessage{}
		msg.Body = "unary Chat Data from Client"
		msg.Language = "EN-US"
		resp, err := c.UnaryComms(context.Background(), &msg)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(resp)
		return;
	case 2:
		// server Streamed Chat Service
		conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		c := client_grpc.NewChatServiceClient(conn)

		// generate random data
		data := make([]string, 10)
		for i := range data {
		data[i] = strconv.Itoa(0 + i)
		}
		// send message
		msg := client_grpc.StreamNormalMessage{Body: data}
		stream, err := c.ServerStreamComms(context.Background(), &msg)
		if err != nil {
			panic(err)
		}
		for {
			data, err := stream.Recv()
			if err == io.EOF {
			break
			}
			if err != nil {
			log.Fatalf("%v.SayHello(_) = _, %v", c, err)
			}
			fmt.Printf("Number : %s, Language : %s\n",data.GetBody(), data.GetLanguage())
		}
		
		defer conn.Close()
	case 3:
		// server Streamed Chat Service
		conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		c := client_grpc.NewChatServiceClient(conn)
		stream, err := c.ClientStreamComms(context.Background())
		if err != nil {
			panic(err)
		}
		// generate random data
		data := make([]string, 10)
		for i := range data {
			numStr := strconv.Itoa(0 + i)
			time.Sleep(time.Second/4)
			stream.Send(&client_grpc.StreamNormalMessage{Body: []string{numStr}, Language: "EN-US"})
		}

		reply, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
		}
		log.Println(reply.Body, reply.Language)
	case 4:
		// server Streamed Chat Service
		conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		c := client_grpc.NewChatServiceClient(conn)
		stream, err := c.BiDirectionStreamComms(context.Background())
		if err != nil {
			panic(err)
		}
		waitc := make(chan struct{})
		go func() {
			for {
				in, err := stream.Recv()
				if err == io.EOF {
					// read done.
					close(waitc)
					return
				}
				if err != nil {
					log.Fatalf("Failed to receive a note : %v", err)
				}
				log.Println(in)
			}
		}()
		data := make([]string, 10)
		for i := range data {
			numStr := strconv.Itoa(0 + i)
			time.Sleep(time.Second/4)
			stream.Send(&client_grpc.StreamNormalMessage{Body: []string{numStr}, Language: "EN-US"})
			log.Println("REQUEST SENT")
		}
		err = stream.CloseSend()
		if err != nil {
			log.Panicln(err)
		}
		<-waitc
		break;
		
	default:
		fmt.Println("No RPC defined")
	}
}

func main() {

	// if len(os.Args) < 2 {
	// 	log.Fatal("Please provide an ID")
	// }

	id := "1"
	
	digit, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Cannot convert Argument to int")
	}
	runRPCs(digit)	
}