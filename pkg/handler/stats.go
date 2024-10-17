package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/proto"
)

func (a *AdminHandler) AdminGetUserStats(ctx context.Context, p *pb.AdUserStatsRequest) (*pb.AdUserStatsResponse, error) {
	response, err := a.SVC.AdminGetAllUserStatsService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (a *AdminHandler) AdminGetSubscriptionStats(ctx context.Context, p *pb.AdSubscriptionStatsRequest) (*pb.AdSubscriptionStatsResponse, error) {
	response, err := a.SVC.AdminGetAllSubscriptionStatsService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
