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

func (a *AdminService) UnBlockUserService(p *pb.UserId) (*pb.AdminResponse, error) {
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

func (a *AdminService) GetAllUsersService(p *pb.AdNoParam) (*pb.AdUserList, error) {
	ctx := context.Background()

	result, err := a.UserClient.GetAllUsers(ctx, &userpb.NoParam{})
	if err != nil {
		return nil, err
	}

	var userList pb.AdUserList
	for _, user := range result.Users {
		userList.Users = append(userList.Users, &pb.AdUserProfile{
			User_ID:                user.User_ID,
			User_Name:              user.User_Name,
			Email:                  user.Email,
			Phone:                  user.Phone,
			Is_Prime_Member:        user.Is_Prime_Member,
			Is_Blocked:             user.Is_Blocked,
			Membership_Expiry_Date: user.Membership_Expiry_Date,
		})
	}

	return &userList, nil
}

func (a *AdminService) GetUserByIDService(p *pb.AdID) (*pb.AdUserProfile, error) {
	ctx := context.Background()

	user := &userpb.ID{
		ID: p.ID,
	}

	result, err := a.UserClient.ViewProfile(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.AdUserProfile{
		User_ID:                result.User_ID,
		User_Name:              result.User_Name,
		Email:                  result.Email,
		Phone:                  result.Phone,
		Is_Prime_Member:        result.Is_Prime_Member,
		Is_Blocked:             result.Is_Blocked,
		Membership_Expiry_Date: result.Membership_Expiry_Date,
	}, nil
}