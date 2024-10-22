package repo

import (
	"github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/repo/interfaces"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepoInter {
	return &AdminRepository{
		DB: db,
	}
}
