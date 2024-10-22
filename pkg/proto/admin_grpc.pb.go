// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: admin.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	AdminLoginRequest(ctx context.Context, in *AdminLogin, opts ...grpc.CallOption) (*AdminResponse, error)
	AdminBlockUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*AdminResponse, error)
	AdminUnBlockUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*AdminResponse, error)
	AdminGetAllUsers(ctx context.Context, in *AdNoParam, opts ...grpc.CallOption) (*AdUserList, error)
	AdminFindUserByID(ctx context.Context, in *AdID, opts ...grpc.CallOption) (*AdUserProfile, error)
	InsertProblem(ctx context.Context, in *AdProblem, opts ...grpc.CallOption) (*AdminResponse, error)
	AdminGetAllProblems(ctx context.Context, in *AdNoParam, opts ...grpc.CallOption) (*AdProblemList, error)
	AdminEditProblem(ctx context.Context, in *AdProblem, opts ...grpc.CallOption) (*AdProblem, error)
	AdminUpgradeProbem(ctx context.Context, in *AdProblemId, opts ...grpc.CallOption) (*AdminResponse, error)
	InsertTestCases(ctx context.Context, in *AdTestCaseRequest, opts ...grpc.CallOption) (*AdminResponse, error)
	UpdateTestCases(ctx context.Context, in *AdUpdateTestCaseRequest, opts ...grpc.CallOption) (*AdminResponse, error)
	GetProblemWithTestCases(ctx context.Context, in *AdProblemId, opts ...grpc.CallOption) (*AdminTestcaseResponse, error)
	AddSubscriptionPlan(ctx context.Context, in *AdSubscription, opts ...grpc.CallOption) (*AdminResponse, error)
	GetAllPlans(ctx context.Context, in *AdNoParam, opts ...grpc.CallOption) (*AdPlanList, error)
	AdminUpdatePlan(ctx context.Context, in *AdSubscription, opts ...grpc.CallOption) (*AdSubscription, error)
	GetSubscriptionByID(ctx context.Context, in *SubscriptionID, opts ...grpc.CallOption) (*AdSubscription, error)
	AdminGetUserStats(ctx context.Context, in *AdUserStatsRequest, opts ...grpc.CallOption) (*AdUserStatsResponse, error)
	AdminGetSubscriptionStats(ctx context.Context, in *AdSubscriptionStatsRequest, opts ...grpc.CallOption) (*AdSubscriptionStatsResponse, error)
	AdminGetProblemStats(ctx context.Context, in *AdProblemStatsRequest, opts ...grpc.CallOption) (*AdProblemStatsResponse, error)
	AdminGetLeaderboardStats(ctx context.Context, in *AdLeaderboardRequest, opts ...grpc.CallOption) (*AdLeaderboardResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) AdminLoginRequest(ctx context.Context, in *AdminLogin, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, "/pb.AdminService/AdminLoginRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminBlockUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, "/pb.AdminService/AdminBlockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminUnBlockUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, "/pb.AdminService/AdminUnBlockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminGetAllUsers(ctx context.Context, in *AdNoParam, opts ...grpc.CallOption) (*AdUserList, error) {
	out := new(AdUserList)
	err := c.cc.Invoke(ctx, "/pb.AdminService/AdminGetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminFindUserByID(ctx context.Context, in *AdID, opts ...grpc.CallOption) (*AdUserProfile, error) {
	out := new(AdUserProfile)
	err := c.cc.Invoke(ctx, "/pb.AdminService/AdminFindUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) InsertProblem(ctx context.Context, in *AdProblem, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, "/pb.AdminService/InsertProblem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminGetAllProblems(ctx context.Context, in *AdNoParam, opts ...grpc.CallOption) (*AdProblemList, error) {
	out := new(AdProblemList)
	err := c.cc.Invoke(ctx, "/pb.AdminService/AdminGetAllProblems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminEditProblem(ctx context.Context, in *AdProblem, opts ...grpc.CallOption) (*AdProblem, error) {
	out := new(AdProblem)
	err := c.cc.Invoke(ctx, "/pb.AdminService/AdminEditProblem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminUpgradeProbem(ctx context.Context, in *AdProblemId, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, "/pb.AdminService/AdminUpgradeProbem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) InsertTestCases(ctx context.Context, in *AdTestCaseRequest, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, "/pb.AdminService/InsertTestCases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UpdateTestCases(ctx context.Context, in *AdUpdateTestCaseRequest, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, "/pb.AdminService/UpdateTestCases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetProblemWithTestCases(ctx context.Context, in *AdProblemId, opts ...grpc.CallOption) (*AdminTestcaseResponse, error) {
	out := new(AdminTestcaseResponse)
	err := c.cc.Invoke(ctx, "/pb.AdminService/GetProblemWithTestCases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AddSubscriptionPlan(ctx context.Context, in *AdSubscription, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, "/pb.AdminService/AddSubscriptionPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetAllPlans(ctx context.Context, in *AdNoParam, opts ...grpc.CallOption) (*AdPlanList, error) {
	out := new(AdPlanList)
	err := c.cc.Invoke(ctx, "/pb.AdminService/GetAllPlans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminUpdatePlan(ctx context.Context, in *AdSubscription, opts ...grpc.CallOption) (*AdSubscription, error) {
	out := new(AdSubscription)
	err := c.cc.Invoke(ctx, "/pb.AdminService/AdminUpdatePlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetSubscriptionByID(ctx context.Context, in *SubscriptionID, opts ...grpc.CallOption) (*AdSubscription, error) {
	out := new(AdSubscription)
	err := c.cc.Invoke(ctx, "/pb.AdminService/GetSubscriptionByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminGetUserStats(ctx context.Context, in *AdUserStatsRequest, opts ...grpc.CallOption) (*AdUserStatsResponse, error) {
	out := new(AdUserStatsResponse)
	err := c.cc.Invoke(ctx, "/pb.AdminService/AdminGetUserStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminGetSubscriptionStats(ctx context.Context, in *AdSubscriptionStatsRequest, opts ...grpc.CallOption) (*AdSubscriptionStatsResponse, error) {
	out := new(AdSubscriptionStatsResponse)
	err := c.cc.Invoke(ctx, "/pb.AdminService/AdminGetSubscriptionStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminGetProblemStats(ctx context.Context, in *AdProblemStatsRequest, opts ...grpc.CallOption) (*AdProblemStatsResponse, error) {
	out := new(AdProblemStatsResponse)
	err := c.cc.Invoke(ctx, "/pb.AdminService/AdminGetProblemStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminGetLeaderboardStats(ctx context.Context, in *AdLeaderboardRequest, opts ...grpc.CallOption) (*AdLeaderboardResponse, error) {
	out := new(AdLeaderboardResponse)
	err := c.cc.Invoke(ctx, "/pb.AdminService/AdminGetLeaderboardStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	AdminLoginRequest(context.Context, *AdminLogin) (*AdminResponse, error)
	AdminBlockUser(context.Context, *UserId) (*AdminResponse, error)
	AdminUnBlockUser(context.Context, *UserId) (*AdminResponse, error)
	AdminGetAllUsers(context.Context, *AdNoParam) (*AdUserList, error)
	AdminFindUserByID(context.Context, *AdID) (*AdUserProfile, error)
	InsertProblem(context.Context, *AdProblem) (*AdminResponse, error)
	AdminGetAllProblems(context.Context, *AdNoParam) (*AdProblemList, error)
	AdminEditProblem(context.Context, *AdProblem) (*AdProblem, error)
	AdminUpgradeProbem(context.Context, *AdProblemId) (*AdminResponse, error)
	InsertTestCases(context.Context, *AdTestCaseRequest) (*AdminResponse, error)
	UpdateTestCases(context.Context, *AdUpdateTestCaseRequest) (*AdminResponse, error)
	GetProblemWithTestCases(context.Context, *AdProblemId) (*AdminTestcaseResponse, error)
	AddSubscriptionPlan(context.Context, *AdSubscription) (*AdminResponse, error)
	GetAllPlans(context.Context, *AdNoParam) (*AdPlanList, error)
	AdminUpdatePlan(context.Context, *AdSubscription) (*AdSubscription, error)
	GetSubscriptionByID(context.Context, *SubscriptionID) (*AdSubscription, error)
	AdminGetUserStats(context.Context, *AdUserStatsRequest) (*AdUserStatsResponse, error)
	AdminGetSubscriptionStats(context.Context, *AdSubscriptionStatsRequest) (*AdSubscriptionStatsResponse, error)
	AdminGetProblemStats(context.Context, *AdProblemStatsRequest) (*AdProblemStatsResponse, error)
	AdminGetLeaderboardStats(context.Context, *AdLeaderboardRequest) (*AdLeaderboardResponse, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) AdminLoginRequest(context.Context, *AdminLogin) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLoginRequest not implemented")
}
func (UnimplementedAdminServiceServer) AdminBlockUser(context.Context, *UserId) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminBlockUser not implemented")
}
func (UnimplementedAdminServiceServer) AdminUnBlockUser(context.Context, *UserId) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUnBlockUser not implemented")
}
func (UnimplementedAdminServiceServer) AdminGetAllUsers(context.Context, *AdNoParam) (*AdUserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetAllUsers not implemented")
}
func (UnimplementedAdminServiceServer) AdminFindUserByID(context.Context, *AdID) (*AdUserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminFindUserByID not implemented")
}
func (UnimplementedAdminServiceServer) InsertProblem(context.Context, *AdProblem) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertProblem not implemented")
}
func (UnimplementedAdminServiceServer) AdminGetAllProblems(context.Context, *AdNoParam) (*AdProblemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetAllProblems not implemented")
}
func (UnimplementedAdminServiceServer) AdminEditProblem(context.Context, *AdProblem) (*AdProblem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminEditProblem not implemented")
}
func (UnimplementedAdminServiceServer) AdminUpgradeProbem(context.Context, *AdProblemId) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpgradeProbem not implemented")
}
func (UnimplementedAdminServiceServer) InsertTestCases(context.Context, *AdTestCaseRequest) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertTestCases not implemented")
}
func (UnimplementedAdminServiceServer) UpdateTestCases(context.Context, *AdUpdateTestCaseRequest) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTestCases not implemented")
}
func (UnimplementedAdminServiceServer) GetProblemWithTestCases(context.Context, *AdProblemId) (*AdminTestcaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblemWithTestCases not implemented")
}
func (UnimplementedAdminServiceServer) AddSubscriptionPlan(context.Context, *AdSubscription) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSubscriptionPlan not implemented")
}
func (UnimplementedAdminServiceServer) GetAllPlans(context.Context, *AdNoParam) (*AdPlanList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPlans not implemented")
}
func (UnimplementedAdminServiceServer) AdminUpdatePlan(context.Context, *AdSubscription) (*AdSubscription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdatePlan not implemented")
}
func (UnimplementedAdminServiceServer) GetSubscriptionByID(context.Context, *SubscriptionID) (*AdSubscription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptionByID not implemented")
}
func (UnimplementedAdminServiceServer) AdminGetUserStats(context.Context, *AdUserStatsRequest) (*AdUserStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetUserStats not implemented")
}
func (UnimplementedAdminServiceServer) AdminGetSubscriptionStats(context.Context, *AdSubscriptionStatsRequest) (*AdSubscriptionStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetSubscriptionStats not implemented")
}
func (UnimplementedAdminServiceServer) AdminGetProblemStats(context.Context, *AdProblemStatsRequest) (*AdProblemStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetProblemStats not implemented")
}
func (UnimplementedAdminServiceServer) AdminGetLeaderboardStats(context.Context, *AdLeaderboardRequest) (*AdLeaderboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetLeaderboardStats not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_AdminLoginRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminLoginRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/AdminLoginRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminLoginRequest(ctx, req.(*AdminLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminBlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminBlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/AdminBlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminBlockUser(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminUnBlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminUnBlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/AdminUnBlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminUnBlockUser(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminGetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdNoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminGetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/AdminGetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminGetAllUsers(ctx, req.(*AdNoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminFindUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminFindUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/AdminFindUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminFindUserByID(ctx, req.(*AdID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_InsertProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdProblem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).InsertProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/InsertProblem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).InsertProblem(ctx, req.(*AdProblem))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminGetAllProblems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdNoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminGetAllProblems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/AdminGetAllProblems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminGetAllProblems(ctx, req.(*AdNoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminEditProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdProblem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminEditProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/AdminEditProblem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminEditProblem(ctx, req.(*AdProblem))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminUpgradeProbem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdProblemId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminUpgradeProbem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/AdminUpgradeProbem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminUpgradeProbem(ctx, req.(*AdProblemId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_InsertTestCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdTestCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).InsertTestCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/InsertTestCases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).InsertTestCases(ctx, req.(*AdTestCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UpdateTestCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdUpdateTestCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UpdateTestCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/UpdateTestCases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UpdateTestCases(ctx, req.(*AdUpdateTestCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetProblemWithTestCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdProblemId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetProblemWithTestCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/GetProblemWithTestCases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetProblemWithTestCases(ctx, req.(*AdProblemId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AddSubscriptionPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdSubscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AddSubscriptionPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/AddSubscriptionPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AddSubscriptionPlan(ctx, req.(*AdSubscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetAllPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdNoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetAllPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/GetAllPlans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetAllPlans(ctx, req.(*AdNoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminUpdatePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdSubscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminUpdatePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/AdminUpdatePlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminUpdatePlan(ctx, req.(*AdSubscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetSubscriptionByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetSubscriptionByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/GetSubscriptionByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetSubscriptionByID(ctx, req.(*SubscriptionID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminGetUserStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdUserStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminGetUserStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/AdminGetUserStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminGetUserStats(ctx, req.(*AdUserStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminGetSubscriptionStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdSubscriptionStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminGetSubscriptionStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/AdminGetSubscriptionStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminGetSubscriptionStats(ctx, req.(*AdSubscriptionStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminGetProblemStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdProblemStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminGetProblemStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/AdminGetProblemStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminGetProblemStats(ctx, req.(*AdProblemStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminGetLeaderboardStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdLeaderboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminGetLeaderboardStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AdminService/AdminGetLeaderboardStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminGetLeaderboardStats(ctx, req.(*AdLeaderboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminLoginRequest",
			Handler:    _AdminService_AdminLoginRequest_Handler,
		},
		{
			MethodName: "AdminBlockUser",
			Handler:    _AdminService_AdminBlockUser_Handler,
		},
		{
			MethodName: "AdminUnBlockUser",
			Handler:    _AdminService_AdminUnBlockUser_Handler,
		},
		{
			MethodName: "AdminGetAllUsers",
			Handler:    _AdminService_AdminGetAllUsers_Handler,
		},
		{
			MethodName: "AdminFindUserByID",
			Handler:    _AdminService_AdminFindUserByID_Handler,
		},
		{
			MethodName: "InsertProblem",
			Handler:    _AdminService_InsertProblem_Handler,
		},
		{
			MethodName: "AdminGetAllProblems",
			Handler:    _AdminService_AdminGetAllProblems_Handler,
		},
		{
			MethodName: "AdminEditProblem",
			Handler:    _AdminService_AdminEditProblem_Handler,
		},
		{
			MethodName: "AdminUpgradeProbem",
			Handler:    _AdminService_AdminUpgradeProbem_Handler,
		},
		{
			MethodName: "InsertTestCases",
			Handler:    _AdminService_InsertTestCases_Handler,
		},
		{
			MethodName: "UpdateTestCases",
			Handler:    _AdminService_UpdateTestCases_Handler,
		},
		{
			MethodName: "GetProblemWithTestCases",
			Handler:    _AdminService_GetProblemWithTestCases_Handler,
		},
		{
			MethodName: "AddSubscriptionPlan",
			Handler:    _AdminService_AddSubscriptionPlan_Handler,
		},
		{
			MethodName: "GetAllPlans",
			Handler:    _AdminService_GetAllPlans_Handler,
		},
		{
			MethodName: "AdminUpdatePlan",
			Handler:    _AdminService_AdminUpdatePlan_Handler,
		},
		{
			MethodName: "GetSubscriptionByID",
			Handler:    _AdminService_GetSubscriptionByID_Handler,
		},
		{
			MethodName: "AdminGetUserStats",
			Handler:    _AdminService_AdminGetUserStats_Handler,
		},
		{
			MethodName: "AdminGetSubscriptionStats",
			Handler:    _AdminService_AdminGetSubscriptionStats_Handler,
		},
		{
			MethodName: "AdminGetProblemStats",
			Handler:    _AdminService_AdminGetProblemStats_Handler,
		},
		{
			MethodName: "AdminGetLeaderboardStats",
			Handler:    _AdminService_AdminGetLeaderboardStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}
