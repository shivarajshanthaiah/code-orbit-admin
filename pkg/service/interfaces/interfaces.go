package interfaces

import (
	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/proto"
)

type AdminServiceInter interface {
	LoginService(p *pb.AdminLogin) (*pb.AdminResponse, error)

	BlockUserService(p *pb.UserId) (*pb.AdminResponse, error)
	UnBlockUserService(p *pb.UserId) (*pb.AdminResponse, error)
	GetAllUsersService(p *pb.AdNoParam) (*pb.AdUserList, error)
	GetUserByIDService(p *pb.AdID) (*pb.AdUserProfile, error)

	InsertProblemService(p *pb.AdProblem) (*pb.AdminResponse, error)
}
