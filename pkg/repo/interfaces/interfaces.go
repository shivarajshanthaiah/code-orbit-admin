package interfaces

import "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/model"

type AdminRepoInter interface {
	FindAdminByEmail(email string) (*model.Admin, error)

	CreateSubscription(plan *model.Subscription) error
	GetAllPlans() (*[]model.Subscription, error)
	GetPlanByID(planID uint) (*model.Subscription, error)
	UpdatePlan(plan *model.Subscription) error
}
