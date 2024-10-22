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

// func (a *AdminService) GetProblemWithTestCasesService(ctx context.Context, req *pb.AdProblemId) (*pb.AdminTestcaseResponse, error) {
//     // Call the Problem Service
//     result, err := a.ProblemClient.GetProblemWithTestCases(ctx, &problempb.ProblemId{
//         ID: req.ID,
//     })
//     if err != nil {
//         return &pb.AdminTestcaseResponse{
//             Status:  pb.AdminTestcaseResponse_ERROR,
//             Message: "Failed to fetch problem and test cases",
//             Payload: &pb.AdminTestcaseResponse_Error{
//                 Error: err.Error(),
//             },
//         }, err
//     }

//     // Map the response from ProblemService to AdminService response
//     return &pb.AdminTestcaseResponse{
//         Status:  pb.AdminTestcaseResponse_Status(result.Status),
//         Message: result.Message,
//         Payload: result.Payload, // Directly forward the payload
//     }, nil
// }

func (a *AdminService) GetProblemWithTestCasesService(ctx context.Context, req *pb.AdProblemId) (*pb.AdminTestcaseResponse, error) {
	// Call the Problem Service
	result, err := a.ProblemClient.GetProblemWithTestCases(ctx, &problempb.ProblemId{
		ID: req.ID,
	})
	if err != nil {
		return &pb.AdminTestcaseResponse{
			Status:  pb.AdminTestcaseResponse_ERROR,
			Message: "Failed to fetch problem and test cases",
			Payload: &pb.AdminTestcaseResponse_Error{
				Error: err.Error(),
			},
		}, err
	}

	// Check if there was an error in the ProblemService response
	if result.Status == problempb.GetProblemResponse_ERROR {
		return &pb.AdminTestcaseResponse{
			Status:  pb.AdminTestcaseResponse_ERROR,
			Message: result.Message,
			Payload: &pb.AdminTestcaseResponse_Error{
				Error: result.GetError(),
			},
		}, nil
	}

	// Map the ProblemService response to the AdminService response format
	problem := result.GetData().Problem
	testCases := result.GetData().TestCases

	var adminTestCases []*pb.AdTestCase
	for _, tc := range testCases {
		adminTestCases = append(adminTestCases, &pb.AdTestCase{
			TestCaseId:     tc.TestCaseId,
			Input:          tc.Input,
			ExpectedOutput: tc.ExpectedOutput,
		})
	}

	// Prepare the response for the AdminService
	return &pb.AdminTestcaseResponse{
		Status:  pb.AdminTestcaseResponse_OK,
		Message: "Problem and test cases fetched successfully",
		Payload: &pb.AdminTestcaseResponse_Data{
			Data: &pb.AdProblemWithTestCases{
				Problem: &pb.AdProblem{
					ID:          problem.ID,
					Title:       problem.Title,
					Discription: problem.Discription,
					Difficulty:  problem.Difficulty,
					Type:        problem.Type,
					IsPremium:   problem.IsPremium,
				},
				TestCases: adminTestCases,
			},
		},
	}, nil
}
