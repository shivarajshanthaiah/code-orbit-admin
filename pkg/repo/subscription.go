package repo

import "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/model"

func (a *AdminRepository) CreateSubscription(plan *model.Subscription) error {
	if err := a.DB.Create(&plan).Error; err != nil {
		return err
	}
	return nil
}

func (a *AdminRepository) GetAllPlans() (*[]model.Subscription, error) {
	var plans []model.Subscription
	if err := a.DB.Find(&plans).Error; err != nil {
		return nil, err
	}
	return &plans, nil
}

func (a *AdminRepository) GetPlanByID(planID uint) (*model.Subscription, error) {
	var plan model.Subscription
	err := a.DB.Where("id = ?", planID).First(&plan).Error
	if err != nil {
		return nil, err
	}
	return &plan, nil
}

func (a *AdminRepository) UpdatePlan(plan *model.Subscription) error {
	if err := a.DB.Save(&plan).Error; err != nil {
		return err
	}
	return nil
}
