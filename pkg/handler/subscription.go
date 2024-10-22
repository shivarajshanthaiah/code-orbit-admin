package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/proto"
)

func (a *AdminHandler) AdminAddSubscriptionPlan(ctx context.Context, p *pb.AdSubscription) (*pb.AdminResponse, error) {
	response, err := a.SVC.InsertPlanService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (a *AdminHandler) AdminUpdatePlan(ctx context.Context, p *pb.AdSubscription) (*pb.AdSubscription, error) {
	response, err := a.SVC.UpdatePlanService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (a *AdminHandler) AdminGetAllPlans(ctx context.Context, p *pb.AdNoParam) (*pb.AdPlanList, error) {
	response, err := a.SVC.FindAllPlansService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
