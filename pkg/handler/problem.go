package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/proto"
)

func (a *AdminHandler) InsertProblem(ctx context.Context, p *pb.AdProblem) (*pb.AdminResponse, error) {
	response, err := a.SVC.InsertProblemService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (a *AdminHandler) AdminGetAllProblems(ctx context.Context, p *pb.AdNoParam) (*pb.AdProblemList, error) {
	response, err := a.SVC.GetAllProblemsService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}


// func (a *AdminHandler) AdminEditProblem(ctx context.Context, p *pb.AdProblem) (*pb.AdProblem, error){

// }

func (a *AdminHandler) GetProblemWithTestCases(ctx context.Context, p *pb.AdProblemId) (*pb.AdminTestcaseResponse, error) {
    response, err := a.SVC.GetProblemWithTestCasesService(ctx, p)
    if err != nil {
        return response, err
    }
    return response, nil
}
