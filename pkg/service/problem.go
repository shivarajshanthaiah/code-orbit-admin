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
		Type:        p.Type,
		IsPremium:   p.IsPremium,
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

func (a *AdminService) GetAllProblemsService(p *pb.AdNoParam) (*pb.AdProblemList, error) {
	ctx := context.Background()

	result, err := a.ProblemClient.GetAllProblems(ctx, &problempb.ProbNoParam{})
	if err != nil {
		return nil, err
	}

	var problemList pb.AdProblemList
	for _, problem := range result.Problems {
		adProblem := &pb.AdProblem{
			ID:          uint32(problem.ID),
			Title:       problem.Title,
			Discription: problem.Discription,
			Difficulty:  problem.Difficulty,
			Type:        problem.Type,
			IsPremium:   problem.IsPremium,
		}
		problemList.Problems = append(problemList.Problems, adProblem)
	}

	return &problemList, nil
}

func (a *AdminService) EditProblemService(p *pb.AdProblem) (*pb.AdProblem, error) {
	ctx := context.Background()

	problem, err := a.ProblemClient.FindProblemByID(ctx, &problempb.ProblemId{
		ID: int32(p.ID),
	})
	if err != nil {
		return nil, err
	}

	problem.Title = p.Title
	problem.Discription = p.Discription
	problem.Difficulty = p.Difficulty
	problem.Type = p.Type
	problem.IsPremium = p.IsPremium

	_, err = a.ProblemClient.EditProblem(ctx, problem)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (a *AdminService) UpgradeProblemService(p *pb.AdProblemId) (*pb.AdminResponse, error){
	ctx := context.Background()

	problem := &problempb.ProblemId{
		ID: p.ID,
	}
	 _, err := a.ProblemClient.AdminUpgradeProbem(ctx, problem)
	 if err != nil {
		return nil, err
	}

	return &pb.AdminResponse{
		Status: pb.AdminResponse_OK,
		Message: "Problem upgraded successfully",
	}, nil
}
