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
