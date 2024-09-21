package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/proto"
)

func (a *AdminHandler) InsertTestCases(ctx context.Context, p *pb.AdTestCaseRequest) (*pb.AdminResponse, error) {
	response, err := a.SVC.InsertTestCaseService(ctx, p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (a *AdminHandler) UpdateTestCases(ctx context.Context, p *pb.AdUpdateTestCaseRequest) (*pb.AdminResponse, error) {
	response, err := a.SVC.UpdateTestCaseService(ctx, p)
	if err != nil {
		return response, err
	}
	return response, nil
}
