package problem

import (
	"log"

	"github.com/shivaraj-shanthaiah/code_orbit_admin/config"
	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/clients/problem/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.ProblemServiceClient, error) {
	grpc, err := grpc.Dial(":"+cfg.ProblemPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc problem client: %s, ", cfg.ProblemPort)
		return nil, err
	}
	log.Printf("Succesfully connected to problem client at port: %v", cfg.ProblemPort)
	return pb.NewProblemServiceClient(grpc), nil
}
