package service

import (
	"errors"

	"github.com/shivaraj-shanthaiah/code_orbit_admin/config"
	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/proto"
	"github.com/shivaraj-shanthaiah/code_orbit_admin/utils"
)

// LoginService implements interfaces.AdminServiceInter.
func (a *AdminService) LoginService(p *pb.AdminLogin) (*pb.AdminResponse, error) {
	admin, err := a.Repo.FindAdminByEmail(p.Email)
	if err != nil {
		return nil, err
	}

	if admin.Password != p.Password {
		return &pb.AdminResponse{
			Status:  pb.AdminResponse_ERROR,
			Message: "Password is incorrect",
			Payload: &pb.AdminResponse_Error{
				Error: "incorect Password",
			},
		}, errors.New("incorrect Password")
	}

	jwtToken, err := utils.GenerateToken(config.LoadConfig().SECRETKEY, admin.Email)
	if err != nil {
		return &pb.AdminResponse{
			Status:  pb.AdminResponse_ERROR,
			Message: "error in generating token",
			Payload: &pb.AdminResponse_Error{
				Error: err.Error()},
		}, errors.New("error generating token")
	}

	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "Login successful",
		Payload: &pb.AdminResponse_Data{
			Data: jwtToken,
		},
	}, nil
}
