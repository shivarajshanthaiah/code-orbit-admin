package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/proto"
)

func (a *AdminHandler) AdminBlockUser(ctx context.Context, p *pb.UserId) (*pb.AdminResponse, error) {
	response, err := a.SVC.BlockUserService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (a *AdminHandler) AdminUnBlockUser(ctx context.Context, p *pb.UserId) (*pb.AdminResponse, error) {
	response, err := a.SVC.UnBlockUserService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
