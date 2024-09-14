package interfaces

import (
	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/proto"
)

type AdminServiceInter interface {
	LoginService(p *pb.AdminLogin) (*pb.AdminResponse, error)
	BlockUserService(p *pb.UserId) (*pb.AdminResponse, error)
	UnBlockUserService(p *pb.UserId) (*pb.AdminResponse, error)
}
