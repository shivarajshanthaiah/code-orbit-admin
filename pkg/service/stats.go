package service

import (
	"context"

	problempb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/clients/problem/pb"
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

func (a *AdminService) AdminGetProblemStatsService(req *pb.AdProblemStatsRequest) (*pb.AdProblemStatsResponse, error) {
	ctx := context.Background()

	result, err := a.ProblemClient.GetProblemStats(ctx, &problempb.ProblemStatsRequest{})
	if err != nil {
		return nil, err
	}

	return &pb.AdProblemStatsResponse{
		TotalProblems:      result.TotalProblems,
		EasyProblems:       result.EasyProblems,
		MediumProblems:     result.MediumProblems,
		HardProblems:       result.HardProblems,
		TypeProblemCount:   result.TypeProblemCount,
		PremiumProblems:    result.PremiumProblems,
		NonPremiumProblems: result.NonPremiumProblems,
	}, nil
}

func (a *AdminService) AdminGetLeaderboardStatsService(req *pb.AdLeaderboardRequest) (*pb.AdLeaderboardResponse, error) {
	ctx := context.Background()
    response, err := a.ProblemClient.GetLeaderboard(ctx, &problempb.LeaderboardRequest{})
    if err != nil {
        return nil, err
    }

    var leaderboardEntries []*pb.AdLeaderboardEntry
    for _, entry := range response.Leaderboard {
        leaderboardEntries = append(leaderboardEntries, &pb.AdLeaderboardEntry{
            UserId:      entry.UserId,
            SolvedCount: entry.SolvedCount,
			Rank: entry.Rank,
        })
    }

    return &pb.AdLeaderboardResponse{
        Leaderboard: leaderboardEntries,
    }, nil
}
