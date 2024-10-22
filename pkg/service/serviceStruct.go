package service

import (
	problempb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/clients/problem/pb"
	userpb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/clients/user/pb"
	interRepo "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/repo/interfaces"
	"github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/service/interfaces"
)

type AdminService struct {
	Repo          interRepo.AdminRepoInter
	UserClient    userpb.UserServiceClient
	ProblemClient problempb.ProblemServiceClient
}

func NewAdminService(repo interRepo.AdminRepoInter, userClient userpb.UserServiceClient, problemClient problempb.ProblemServiceClient) interfaces.AdminServiceInter {
	return &AdminService{
		Repo:          repo,
		UserClient:    userClient,
		ProblemClient: problemClient,
	}
}
