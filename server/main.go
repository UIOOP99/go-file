package main

import (
	"context"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/mjafari98/go-file/models"
	"github.com/mjafari98/go-file/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"strings"
)

const (
	GRPCPort = ":50052"
	RESTPort = ":9091"
)

var DB = models.ConnectAndMigrate()

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
	return
}

// allowCORS allows Cross Origin Resource Sharing from any origin.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
			preflightHandler(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	fileServer := FilesServer{}

	// start REST server
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	dialOptions := []grpc.DialOption{grpc.WithInsecure()}
	_ = pb.RegisterFilesHandlerFromEndpoint(ctx, mux, "0.0.0.0:9091", dialOptions)

	log.Printf(
		"server REST started in localhost%s (Wait 60 second before making http requests) ...\n",
		RESTPort,
	)
	go func() {
		err := http.ListenAndServe(RESTPort, allowCORS(mux))
		if err != nil {
			log.Fatal("cannot start REST server: ", err)
		}
	}()
	// end of REST server

	// start gRPC server
	listener, err := net.Listen("tcp", GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterFilesServer(grpcServer, &fileServer)
	reflection.Register(grpcServer)

	log.Printf("server gRPC is starting in localhost%s ...\n", GRPCPort)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start GRPC server: ", err)
	}
	// end of gRPC server
}
