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

func (a *AdminHandler) AdminGetAllUsers(ctx context.Context, p *pb.AdNoParam) (*pb.AdUserList, error) {
	response, err := a.SVC.GetAllUsersService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (a *AdminHandler) AdminFindUserByID(ctx context.Context, p *pb.AdID) (*pb.AdUserProfile, error) {
	response, err := a.SVC.GetUserByIDService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
