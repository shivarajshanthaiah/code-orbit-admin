package service

import (
	"context"

	userpb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/clients/user/pb"
	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/proto"
)

func (a *AdminService) BlockUserService(p *pb.UserId) (*pb.AdminResponse, error) {
	ctx := context.Background()
	user := &userpb.ID{
		ID: p.ID,
	}

	_, err := a.UserClient.BlockUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "User blocked Successfully",
	}, nil
}

func (a *AdminService) UnBlockUserService(p *pb.UserId) (*pb.AdminResponse, error){
	ctx := context.Background()
	user := &userpb.ID{
		ID: p.ID,
	}

	_, err := a.UserClient.UnBlockUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "User Unblocked Successfully",
	}, nil
}
