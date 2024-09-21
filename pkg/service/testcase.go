package service

import (
	"context"

	problempb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/clients/problem/pb"
	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/proto"
)

// InsertTestCaseService implements interfaces.AdminServiceInter.
func (a *AdminService) InsertTestCaseService(ctx context.Context, req *pb.AdTestCaseRequest) (*pb.AdminResponse, error) {
	// Convert AdTestCase to TestCase expected by the Problem Service
	var problemTestCases []*problempb.TestCase
	for _, adTestCase := range req.TestCases {
		problemTestCases = append(problemTestCases, &problempb.TestCase{
			Input:          adTestCase.Input,
			ExpectedOutput: adTestCase.ExpectedOutput,
		})
	}

	// Prepare the request for the Problem Service
	testcaseRequest := &problempb.TestCaseRequest{
		ProblemId: req.ProblemId,
		TestCases: problemTestCases, // Pass the converted test cases
	}

	// Call the Problem Service
	result, err := a.ProblemClient.InsertTestCases(ctx, testcaseRequest)
	if err != nil {
		return &pb.AdminResponse{
			Status:  pb.AdminResponse_ERROR,
			Message: "Failed to create test cases",
			Payload: &pb.AdminResponse_Error{
				Error: err.Error(),
			},
		}, err
	}

	// Return success response
	return &pb.AdminResponse{
		Status:  pb.AdminResponse_Status(result.Status),
		Message: result.Message,
	}, nil
}

func (a *AdminService) UpdateTestCaseService(ctx context.Context, req *pb.AdUpdateTestCaseRequest) (*pb.AdminResponse, error) {
	// Convert AdTestCase to TestCase expected by the Problem Service
	var problemTestCases []*problempb.TestCase
	for _, adTestCase := range req.TestCases {
		problemTestCases = append(problemTestCases, &problempb.TestCase{
			Input:          adTestCase.Input,
			ExpectedOutput: adTestCase.ExpectedOutput,
		})
	}

	// Prepare the request for the Problem Service
	updateRequest := &problempb.UpdateTestCaseRequest{
		TestCaseId: req.TestCaseId,
		ProblemId:  req.ProblemId,
		TestCases:  problemTestCases,
	}

	// Call the Problem Service to update the test cases
	result, err := a.ProblemClient.UpdateTestCases(ctx, updateRequest)
	if err != nil {
		return &pb.AdminResponse{
			Status:  pb.AdminResponse_ERROR,
			Message: "Failed to update test cases",
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
