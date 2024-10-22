package interfaces

import "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/model"

type AdminRepoInter interface {
	FindAdminByEmail(email string) (*model.Admin, error)
}
