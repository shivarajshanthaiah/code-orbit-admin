package service

import (
	userpb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/clients/user/pb"
	interRepo "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/repo/interfaces"
	"github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/service/interfaces"
)

type AdminService struct {
	Repo       interRepo.AdminRepoInter
	UserClient userpb.UserServiceClient
}

func NewAdminService(repo interRepo.AdminRepoInter, userClient userpb.UserServiceClient) interfaces.AdminServiceInter {
	return &AdminService{
		Repo:       repo,
		UserClient: userClient,
	}
}
