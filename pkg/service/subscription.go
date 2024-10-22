package service

import (
	"github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/model"
	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/proto"
)

func (a *AdminService) InsertPlanService(p *pb.AdSubscription) (*pb.AdminResponse, error) {

	plan := model.Subscription{
		Plan:       p.Plan,
		Duration:   p.Duration,
		Price:      p.Price,
		GST:        p.Gst,
		TotalPrice: p.TotalPrice,
	}

	err := a.Repo.CreateSubscription(&plan)
	if err != nil {
		return &pb.AdminResponse{
			Status:  pb.AdminResponse_ERROR,
			Message: "Error creating subscription plan",
			Payload: &pb.AdminResponse_Error{
				Error: err.Error(),
			},
		}, err
	}

	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "Subscription plan created sucessfully",
	}, nil
}

func (a *AdminService) UpdatePlanService(p *pb.AdSubscription) (*pb.AdSubscription, error) {
	plan, err := a.Repo.GetPlanByID(uint(p.ID))
	if err != nil {
		return nil, err
	}

	plan.Plan = p.Plan
	plan.Duration = p.Duration
	plan.Price = p.Price
	plan.GST = p.Gst
	plan.TotalPrice = p.TotalPrice

	err = a.Repo.UpdatePlan(plan)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (a *AdminService) FindAllPlansService(p *pb.AdNoParam) (*pb.AdPlanList, error) {
	result, err := a.Repo.GetAllPlans()
	if err != nil {
		return nil, err
	}

	if result == nil {
		return &pb.AdPlanList{}, nil
	}

	var planList pb.AdPlanList
	for _, plan := range *result {
		pbPlan := &pb.AdSubscription{
			ID:         uint32(plan.ID),
			Plan:       plan.Plan,
			Duration:   plan.Duration,
			Price:      plan.Price,
			Gst:        plan.GST,
			TotalPrice: plan.TotalPrice,
		}
		planList.Plans = append(planList.Plans, pbPlan)
	}
	return &planList, nil
}

func (a *AdminService) GetSubscriptionByIDService(p *pb.SubscriptionID) (*pb.AdSubscription, error) {
	result, err := a.Repo.GetPlanByID(uint(p.ID))
	if err != nil {
		return nil, err
	}

	subscription := &pb.AdSubscription{
		ID:         uint32(result.ID),
		Plan:       result.Plan,
		Duration:   result.Duration,
		Price:      result.Price,
		Gst:        result.GST,
		TotalPrice: result.TotalPrice,
	}

	return subscription, nil
}
