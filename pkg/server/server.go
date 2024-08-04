package server

import (
	"fmt"
	"log"
	"net"

	"github.com/shivaraj-shanthaiah/godoc/docservice/pkg/handler"
	pb "github.com/shivaraj-shanthaiah/godoc/docservice/pkg/pb"
	"google.golang.org/grpc"
)

func NewGrpcServer(handlr *handler.DoctorsHandler) {
	log.Println("connecting to gRPC server")
	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		fmt.Println(err)
		log.Fatal("error creating listener on port 8084")
	}

	grp := grpc.NewServer()
	pb.RegisterDoctorServicesServer(grp, handlr)

	log.Printf("listening on gRPC server 8084")
	if err := grp.Serve(lis); err != nil {
		log.Fatal("error connecting to gRPC server")

	}
}
