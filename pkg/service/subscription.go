package service

import (
	"context"

	userpb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/clients/user/pb"

	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/proto"
)

func (a *AdminService) InsertPlanService(p *pb.AdSubscription) (*pb.AdminResponse, error) {
	ctx := context.Background()
	plan := &userpb.USubscription{
		Plan:       p.Plan,
		Duration:   p.Duration,
		Price:      p.Price,
		Gst:        p.Gst,
		TotalPrice: p.TotalPrice,
	}

	result, err := a.UserClient.AddSubscriptionPlan(ctx, plan)
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
		Message: result.Message,
	}, nil
}

func (a *AdminService) UpdatePlanService(p *pb.AdSubscription) (*pb.AdSubscription, error) {
	ctx := context.Background()

	plan, err := a.UserClient.UpdatePlan(ctx, &userpb.USubscription{
		ID: p.ID,
	})

	if err != nil {
		return nil, err
	}

	plan.Plan = p.Plan
	plan.Duration = p.Duration
	plan.Price = p.Price
	plan.Gst = p.Gst
	plan.TotalPrice = p.TotalPrice

	_, err = a.UserClient.UpdatePlan(ctx, plan)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (a *AdminService) FindAllPlansService(p *pb.AdNoParam) (*pb.AdPlanList, error) {
	ctx := context.Background()

	plan, err := a.UserClient.GetAllPlans(ctx, &userpb.NoParam{})
	if err != nil {
		return nil, err
	}

	var planList pb.AdPlanList
	for _, plan := range plan.Plans {
		pbPlan := &pb.AdSubscription{
			ID:         uint32(plan.ID),
			Plan:       plan.Plan,
			Duration:   plan.Duration,
			Price:      plan.Price,
			Gst:        plan.Gst,
			TotalPrice: plan.TotalPrice,
		}
		planList.Plans = append(planList.Plans, pbPlan)
	}
	return &planList, nil
}

