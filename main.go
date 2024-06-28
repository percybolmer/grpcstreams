package main

import (
	"fmt"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	hardwaremonitoring "github.com/percybolmer/grpcstreams/proto"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Welcome to streaming HW monitoring")
	// Setup a tcp connection to port 7777
	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		panic(err)
	}
	// Create a gRPC server
	gRPCserver := grpc.NewServer()

	// Create a server object of the type we created in server.go
	s := &Server{}

	// Regiser our server as a gRPC server
	hardwaremonitoring.RegisterHardwareMonitorServer(gRPCserver, s)

	// Host the regular gRPC api on a goroutine
	go func() {
		log.Fatal(gRPCserver.Serve(lis))
	}()

	// We need to wrap the gRPC server with a multiplexer to enable
	// the usage of http2 over http1
	grpcWebServer := grpcweb.WrapServer(gRPCserver)

	multiplex := grpcMultiplexer{
		grpcWebServer,
	}

	// a regular http router
	r := http.NewServeMux()
	// Load our React application
	webapp := http.FileServer(http.Dir("webapp/hwmonitor/build"))
	// Host the web app at / and wrap it in a multiplexer
	r.Handle("/", multiplex.Handler(webapp))

	// create a http server with some defaults
	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8383",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// host it
	log.Fatal(srv.ListenAndServe())
}

// grpcMultiplexer enables HTTP requests and gRPC requests to multiple on the same channel
// this is needed since browsers dont fully support http2 yet
type grpcMultiplexer struct {
	*grpcweb.WrappedGrpcServer
}

// Handler is used to route requests to either grpc or to regular http
func (m *grpcMultiplexer) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.IsGrpcWebRequest(r) {
			m.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Server is our struct that will handle the Hardware monitoring Logic
// It will fulfill the gRPC interface generated
type Server struct {
	hardwaremonitoring.UnimplementedHardwareMonitorServer
}

// Monitor is used to start a stream of HardwareStats
func (s *Server) Monitor(req *hardwaremonitoring.EmptyRequest, stream hardwaremonitoring.HardwareMonitor_MonitorServer) error {
	// Start a ticker that executes each 2 seconds
	timer := time.NewTicker(2 * time.Second)

	for {
		select {
		// Exit on stream context done
		case <-stream.Context().Done():
			return nil
		case <-timer.C:
			// Grab stats and output
			hwStats, err := s.GetStats()
			if err != nil {
				log.Println(err.Error())
			} else {

			}
			// Send the Hardware stats on the stream
			err = stream.Send(hwStats)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}

// GetStats will extract system stats and output a Hardware Object, or an error
// if extraction fails
func (s *Server) GetStats() (*hardwaremonitoring.HardwareStats, error) {
	// Extarcyt Memory statas
	mem, err := memory.Get()
	if err != nil {
		return nil, err
	}
	// Extract CPU stats
	cpu, err := cpu.Get()
	if err != nil {
		return nil, err
	}
	// Create our response object
	hwStats := &hardwaremonitoring.HardwareStats{
		Cpu:        int32(cpu.Total),
		MemoryFree: int32(mem.Free),
		MemoryUsed: int32(mem.Used),
	}

	return hwStats, nil
}
