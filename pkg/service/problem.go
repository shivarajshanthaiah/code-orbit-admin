package service

import (
	"context"

	problempb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/clients/problem/pb"
	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/proto"
)

// InsertProblemService implements interfaces.AdminServiceInter.
func (a *AdminService) InsertProblemService(p *pb.AdProblem) (*pb.AdminResponse, error) {
	ctx := context.Background()

	problem := &problempb.Problem{
		Title:       p.Title,
		Discription: p.Discription,
		Difficulty:  p.Difficulty,
		Tags:        p.Tags,
	}

	result, err := a.ProblemClient.InsertProblem(ctx, problem)
	if err != nil {
		return &pb.AdminResponse{
			Status:  pb.AdminResponse_ERROR,
			Message: "Failed to create problem",
			Payload: &pb.AdminResponse_Error{
				Error: err.Error(),
			},
		}, err
	}

	return &pb.AdminResponse{
		Status:  pb.AdminResponse_Status(result.Status),
		Message: result.Message,
	}, nil
}
