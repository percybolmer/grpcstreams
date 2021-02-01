package main

import (
	"context"
	"fmt"
	"log"
	"time"

	hardwaremonitoring "github.com/percybolmer/grpcstreams/proto"
	"google.golang.org/grpc"
)

func main() {

	// Create our context
	ctx := context.Background()
	// Setup connection
	conn, err := grpc.DialContext(ctx, "localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	// Close connection when we are done
	defer conn.Close()
	// Use the generated NewHardwareMonitorClient method and pass our Connection
	client := hardwaremonitoring.NewHardwareMonitorClient(conn)

	// Call Monitor to receive the Stream of data
	// With an empty request
	emptyreq := &hardwaremonitoring.EmptyRequest{}
	// call Monitor function, this will return a stream of data
	stream, err := client.Monitor(ctx, emptyreq)
	if err != nil {
		panic(err)
	}
	// Create a timer to cancel
	stop := time.NewTicker(7 * time.Second)
	// Itterate stream
	for {
		select {
		case <-stop.C:
			// Tell the Server to close this Stream, used to clean up running on the server
			err := stream.CloseSend()
			if err != nil {
				log.Fatal("Failed to close stream: ", err.Error())
			}
			return
		default:
			// Recieve on the stream
			res, err := stream.Recv()
			if err != nil {
				panic(err)
			}
			fmt.Println("New Hardware state receieved")
			fmt.Println("CPU Usage: ", res.Cpu)
			fmt.Println("Memory Used: ", res.MemoryUsed)
			fmt.Println("Memory Free: ", res.MemoryFree)
		}
	}
}
