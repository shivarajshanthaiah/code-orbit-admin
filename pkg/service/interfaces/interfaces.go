package interfaces

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/proto"
)

type AdminServiceInter interface {
	LoginService(p *pb.AdminLogin) (*pb.AdminResponse, error)

	BlockUserService(p *pb.UserId) (*pb.AdminResponse, error)
	UnBlockUserService(p *pb.UserId) (*pb.AdminResponse, error)
	GetAllUsersService(p *pb.AdNoParam) (*pb.AdUserList, error)
	GetUserByIDService(p *pb.AdID) (*pb.AdUserProfile, error)

	InsertProblemService(p *pb.AdProblem) (*pb.AdminResponse, error)
	GetAllProblemsService(p *pb.AdNoParam) (*pb.AdProblemList, error)
	EditProblemService(p *pb.AdProblem) (*pb.AdProblem, error)
	UpgradeProblemService(p *pb.AdProblemId) (*pb.AdminResponse, error)

	InsertTestCaseService(ctx context.Context, req *pb.AdTestCaseRequest) (*pb.AdminResponse, error)
	UpdateTestCaseService(ctx context.Context, req *pb.AdUpdateTestCaseRequest) (*pb.AdminResponse, error)
	GetProblemWithTestCasesService(ctx context.Context, req *pb.AdProblemId) (*pb.AdminTestcaseResponse, error)
}
