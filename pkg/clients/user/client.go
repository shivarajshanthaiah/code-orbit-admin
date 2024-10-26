package user

import (
	"log"

	"github.com/shivaraj-shanthaiah/code_orbit_admin/config"
	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/clients/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.UserServiceClient, error) {
	address := "user-service:" + cfg.UserPort
	grpc, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc user client: %s, ", cfg.UserPort)
		return nil, err
	}
	log.Printf("Succesfully connected to user client at port: %v", cfg.UserPort)
	return pb.NewUserServiceClient(grpc), nil
}
