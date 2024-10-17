package service

import (
	"context"

	userpb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/clients/user/pb"
	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/proto"
)

func (a *AdminService) AdminGetAllUserStatsService(p *pb.AdUserStatsRequest) (*pb.AdUserStatsResponse, error) {
	ctx := context.Background()

	result, err := a.UserClient.GetUserProfileStats(ctx, &userpb.UserStatsRequest{})
	if err != nil {
		return nil, err
	}

	return &pb.AdUserStatsResponse{
		TotalUsers:    result.TotalUsers,
		ActiveUsers:   result.ActiveUsers,
		BlockedUsers:  result.BlockedUsers,
		PrimeUsers:    result.PrimeUsers,
		NonPrimeUsers: result.NonPrimeUsers,
	}, nil
}

func (a *AdminService) AdminGetAllSubscriptionStatsService(p *pb.AdSubscriptionStatsRequest) (*pb.AdSubscriptionStatsResponse, error) {
	ctx := context.Background()

	result, err := a.UserClient.GetSubscriptionStats(ctx, &userpb.SubscriptionStatsRequest{})
	if err != nil {
		return nil, err
	}

	return &pb.AdSubscriptionStatsResponse{
		BasicPlanSubscribers:         result.BasicPlanSubscribers,
		StandardPlanSubscribers:      result.StandardPlanSubscribers,
		PremiumPlanSubscribers:       result.PremiumPlanSubscribers,
		TotalAmountCollectedLifetime: result.TotalAmountCollectedLifetime,
		TotalAmountCollectedWeekly:   result.TotalAmountCollectedWeekly,
		TotalAmountCollectedMonthly:  result.TotalAmountCollectedMonthly,
		TotalAmountCollectedYearly:   result.TotalAmountCollectedYearly,
	}, nil
}