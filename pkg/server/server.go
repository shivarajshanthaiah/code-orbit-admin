package server

import (
	"fmt"
	"log"
	"net"

	"github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/handler"
	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/proto"
	"google.golang.org/grpc"
)

func NewGrpcAdminServer(port string, hdlr *handler.AdminHandler) error {
	log.Println("connecting to gRPC server")
	addr := fmt.Sprintf(":%s", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("error creating listener on %v", addr)
		return err
	}

	grpc := grpc.NewServer()
	pb.RegisterAdminServiceServer(grpc, hdlr)

	log.Printf("listening on gRPC server %v", port)
	err = grpc.Serve(lis)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	return nil
}