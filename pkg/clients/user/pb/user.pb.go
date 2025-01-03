// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: user.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Response_Status int32

const (
	Response_OK    Response_Status = 0
	Response_ERROR Response_Status = 1
)

// Enum value maps for Response_Status.
var (
	Response_Status_name = map[int32]string{
		0: "OK",
		1: "ERROR",
	}
	Response_Status_value = map[string]int32{
		"OK":    0,
		"ERROR": 1,
	}
)

func (x Response_Status) Enum() *Response_Status {
	p := new(Response_Status)
	*p = x
	return p
}

func (x Response_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Response_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[0].Descriptor()
}

func (Response_Status) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[0]
}

func (x Response_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Response_Status.Descriptor instead.
func (Response_Status) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7, 0}
}

type SubscriptionStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscriptionStatsRequest) Reset() {
	*x = SubscriptionStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionStatsRequest) ProtoMessage() {}

func (x *SubscriptionStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionStatsRequest.ProtoReflect.Descriptor instead.
func (*SubscriptionStatsRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

type SubscriptionStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BasicPlanSubscribers         int32  `protobuf:"varint,1,opt,name=basic_plan_subscribers,json=basicPlanSubscribers,proto3" json:"basic_plan_subscribers,omitempty"`
	StandardPlanSubscribers      int32  `protobuf:"varint,2,opt,name=standard_plan_subscribers,json=standardPlanSubscribers,proto3" json:"standard_plan_subscribers,omitempty"`
	PremiumPlanSubscribers       int32  `protobuf:"varint,3,opt,name=premium_plan_subscribers,json=premiumPlanSubscribers,proto3" json:"premium_plan_subscribers,omitempty"`
	TotalAmountCollectedLifetime uint32 `protobuf:"varint,4,opt,name=total_amount_collected_lifetime,json=totalAmountCollectedLifetime,proto3" json:"total_amount_collected_lifetime,omitempty"`
	TotalAmountCollectedWeekly   uint32 `protobuf:"varint,5,opt,name=total_amount_collected_weekly,json=totalAmountCollectedWeekly,proto3" json:"total_amount_collected_weekly,omitempty"`
	TotalAmountCollectedMonthly  uint32 `protobuf:"varint,6,opt,name=total_amount_collected_monthly,json=totalAmountCollectedMonthly,proto3" json:"total_amount_collected_monthly,omitempty"`
	TotalAmountCollectedYearly   uint32 `protobuf:"varint,7,opt,name=total_amount_collected_yearly,json=totalAmountCollectedYearly,proto3" json:"total_amount_collected_yearly,omitempty"`
}

func (x *SubscriptionStatsResponse) Reset() {
	*x = SubscriptionStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionStatsResponse) ProtoMessage() {}

func (x *SubscriptionStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionStatsResponse.ProtoReflect.Descriptor instead.
func (*SubscriptionStatsResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *SubscriptionStatsResponse) GetBasicPlanSubscribers() int32 {
	if x != nil {
		return x.BasicPlanSubscribers
	}
	return 0
}

func (x *SubscriptionStatsResponse) GetStandardPlanSubscribers() int32 {
	if x != nil {
		return x.StandardPlanSubscribers
	}
	return 0
}

func (x *SubscriptionStatsResponse) GetPremiumPlanSubscribers() int32 {
	if x != nil {
		return x.PremiumPlanSubscribers
	}
	return 0
}

func (x *SubscriptionStatsResponse) GetTotalAmountCollectedLifetime() uint32 {
	if x != nil {
		return x.TotalAmountCollectedLifetime
	}
	return 0
}

func (x *SubscriptionStatsResponse) GetTotalAmountCollectedWeekly() uint32 {
	if x != nil {
		return x.TotalAmountCollectedWeekly
	}
	return 0
}

func (x *SubscriptionStatsResponse) GetTotalAmountCollectedMonthly() uint32 {
	if x != nil {
		return x.TotalAmountCollectedMonthly
	}
	return 0
}

func (x *SubscriptionStatsResponse) GetTotalAmountCollectedYearly() uint32 {
	if x != nil {
		return x.TotalAmountCollectedYearly
	}
	return 0
}

type UserStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserStatsRequest) Reset() {
	*x = UserStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStatsRequest) ProtoMessage() {}

func (x *UserStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStatsRequest.ProtoReflect.Descriptor instead.
func (*UserStatsRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

type UserStatsProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalUsers    int32 `protobuf:"varint,1,opt,name=total_users,json=totalUsers,proto3" json:"total_users,omitempty"`
	ActiveUsers   int32 `protobuf:"varint,2,opt,name=active_users,json=activeUsers,proto3" json:"active_users,omitempty"`
	BlockedUsers  int32 `protobuf:"varint,3,opt,name=blocked_users,json=blockedUsers,proto3" json:"blocked_users,omitempty"`
	PrimeUsers    int32 `protobuf:"varint,4,opt,name=prime_users,json=primeUsers,proto3" json:"prime_users,omitempty"`
	NonPrimeUsers int32 `protobuf:"varint,5,opt,name=non_prime_users,json=nonPrimeUsers,proto3" json:"non_prime_users,omitempty"`
}

func (x *UserStatsProfileResponse) Reset() {
	*x = UserStatsProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStatsProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStatsProfileResponse) ProtoMessage() {}

func (x *UserStatsProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStatsProfileResponse.ProtoReflect.Descriptor instead.
func (*UserStatsProfileResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserStatsProfileResponse) GetTotalUsers() int32 {
	if x != nil {
		return x.TotalUsers
	}
	return 0
}

func (x *UserStatsProfileResponse) GetActiveUsers() int32 {
	if x != nil {
		return x.ActiveUsers
	}
	return 0
}

func (x *UserStatsProfileResponse) GetBlockedUsers() int32 {
	if x != nil {
		return x.BlockedUsers
	}
	return 0
}

func (x *UserStatsProfileResponse) GetPrimeUsers() int32 {
	if x != nil {
		return x.PrimeUsers
	}
	return 0
}

func (x *UserStatsProfileResponse) GetNonPrimeUsers() int32 {
	if x != nil {
		return x.NonPrimeUsers
	}
	return 0
}

type UPlanList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plans []*USubscription `protobuf:"bytes,1,rep,name=plans,proto3" json:"plans,omitempty"`
}

func (x *UPlanList) Reset() {
	*x = UPlanList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UPlanList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UPlanList) ProtoMessage() {}

func (x *UPlanList) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UPlanList.ProtoReflect.Descriptor instead.
func (*UPlanList) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *UPlanList) GetPlans() []*USubscription {
	if x != nil {
		return x.Plans
	}
	return nil
}

type USubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         uint32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Plan       string  `protobuf:"bytes,2,opt,name=plan,proto3" json:"plan,omitempty"`
	Duration   string  `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Price      float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Gst        float64 `protobuf:"fixed64,5,opt,name=gst,proto3" json:"gst,omitempty"`
	TotalPrice uint32  `protobuf:"varint,6,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *USubscription) Reset() {
	*x = USubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *USubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*USubscription) ProtoMessage() {}

func (x *USubscription) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use USubscription.ProtoReflect.Descriptor instead.
func (*USubscription) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *USubscription) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *USubscription) GetPlan() string {
	if x != nil {
		return x.Plan
	}
	return ""
}

func (x *USubscription) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *USubscription) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *USubscription) GetGst() float64 {
	if x != nil {
		return x.Gst
	}
	return 0
}

func (x *USubscription) GetTotalPrice() uint32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *ID) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  Response_Status `protobuf:"varint,1,opt,name=status,proto3,enum=pb.Response_Status" json:"status,omitempty"`
	Message string          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Types that are assignable to Payload:
	//
	//	*Response_Error
	//	*Response_Data
	Payload isResponse_Payload `protobuf_oneof:"payload"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *Response) GetStatus() Response_Status {
	if x != nil {
		return x.Status
	}
	return Response_OK
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (m *Response) GetPayload() isResponse_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Response) GetError() string {
	if x, ok := x.GetPayload().(*Response_Error); ok {
		return x.Error
	}
	return ""
}

func (x *Response) GetData() string {
	if x, ok := x.GetPayload().(*Response_Data); ok {
		return x.Data
	}
	return ""
}

type isResponse_Payload interface {
	isResponse_Payload()
}

type Response_Error struct {
	Error string `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

type Response_Data struct {
	Data string `protobuf:"bytes,4,opt,name=data,proto3,oneof"`
}

func (*Response_Error) isResponse_Payload() {}

func (*Response_Data) isResponse_Payload() {}

type NoParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoParam) Reset() {
	*x = NoParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoParam) ProtoMessage() {}

func (x *NoParam) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoParam.ProtoReflect.Descriptor instead.
func (*NoParam) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

type UserList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*Profile `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UserList) Reset() {
	*x = UserList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserList) ProtoMessage() {}

func (x *UserList) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserList.ProtoReflect.Descriptor instead.
func (*UserList) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *UserList) GetUsers() []*Profile {
	if x != nil {
		return x.Users
	}
	return nil
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User_ID                string                 `protobuf:"bytes,1,opt,name=User_ID,json=UserID,proto3" json:"User_ID,omitempty"`
	User_Name              string                 `protobuf:"bytes,2,opt,name=User_Name,json=UserName,proto3" json:"User_Name,omitempty"`
	Email                  string                 `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone                  string                 `protobuf:"bytes,4,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Is_Prime_Member        bool                   `protobuf:"varint,5,opt,name=is_Prime_Member,json=isPrimeMember,proto3" json:"is_Prime_Member,omitempty"`
	Is_Blocked             bool                   `protobuf:"varint,6,opt,name=Is_Blocked,json=IsBlocked,proto3" json:"Is_Blocked,omitempty"`
	Membership_Expiry_Date *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=membership_Expiry_Date,json=membershipExpiryDate,proto3" json:"membership_Expiry_Date,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{10}
}

func (x *Profile) GetUser_ID() string {
	if x != nil {
		return x.User_ID
	}
	return ""
}

func (x *Profile) GetUser_Name() string {
	if x != nil {
		return x.User_Name
	}
	return ""
}

func (x *Profile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Profile) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Profile) GetIs_Prime_Member() bool {
	if x != nil {
		return x.Is_Prime_Member
	}
	return false
}

func (x *Profile) GetIs_Blocked() bool {
	if x != nil {
		return x.Is_Blocked
	}
	return false
}

func (x *Profile) GetMembership_Expiry_Date() *timestamppb.Timestamp {
	if x != nil {
		return x.Membership_Expiry_Date
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x1a, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xd9, 0x03,
	0x0a, 0x19, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x50, 0x6c, 0x61, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x73, 0x12, 0x3a, 0x0a, 0x19, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x6c,
	0x61, 0x6e, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x17, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x50, 0x6c,
	0x61, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x12, 0x38, 0x0a,
	0x18, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x16, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x50, 0x6c, 0x61, 0x6e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x12, 0x45, 0x0a, 0x1f, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x1c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x41,
	0x0a, 0x1d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x57, 0x65, 0x65, 0x6b, 0x6c,
	0x79, 0x12, 0x43, 0x0a, 0x1e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x6e, 0x74,
	0x68, 0x6c, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x12, 0x41, 0x0a, 0x1d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x5f, 0x79, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x59, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xcc, 0x01,
	0x0a, 0x18, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x6d,
	0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e,
	0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x34, 0x0a, 0x09,
	0x55, 0x50, 0x6c, 0x61, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x70, 0x6c, 0x61,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x70, 0x6c, 0x61,
	0x6e, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x0d, 0x55, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x73,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x67, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x14, 0x0a,
	0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x22, 0xa7, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x09, 0x0a,
	0x07, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x2d, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x84, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x09,
	0x55, 0x73, 0x65, 0x72, 0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x50, 0x72, 0x69, 0x6d,
	0x65, 0x5f, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x69, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x49, 0x73, 0x5f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x49, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x50, 0x0a, 0x16,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x79, 0x5f, 0x44, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x14, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x32, 0xda,
	0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x21,
	0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x12, 0x06, 0x2e, 0x70, 0x62,
	0x2e, 0x49, 0x44, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x23, 0x0a, 0x0b, 0x55, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x44, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x0b, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x36, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x11, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x50,
	0x6c, 0x61, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1c,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_user_proto_goTypes = []interface{}{
	(Response_Status)(0),              // 0: pb.Response.Status
	(*SubscriptionStatsRequest)(nil),  // 1: pb.SubscriptionStatsRequest
	(*SubscriptionStatsResponse)(nil), // 2: pb.SubscriptionStatsResponse
	(*UserStatsRequest)(nil),          // 3: pb.UserStatsRequest
	(*UserStatsProfileResponse)(nil),  // 4: pb.UserStatsProfileResponse
	(*UPlanList)(nil),                 // 5: pb.UPlanList
	(*USubscription)(nil),             // 6: pb.USubscription
	(*ID)(nil),                        // 7: pb.ID
	(*Response)(nil),                  // 8: pb.Response
	(*NoParam)(nil),                   // 9: pb.NoParam
	(*UserList)(nil),                  // 10: pb.UserList
	(*Profile)(nil),                   // 11: pb.Profile
	(*timestamppb.Timestamp)(nil),     // 12: google.protobuf.Timestamp
}
var file_user_proto_depIdxs = []int32{
	6,  // 0: pb.UPlanList.plans:type_name -> pb.USubscription
	0,  // 1: pb.Response.status:type_name -> pb.Response.Status
	11, // 2: pb.UserList.users:type_name -> pb.Profile
	12, // 3: pb.Profile.membership_Expiry_Date:type_name -> google.protobuf.Timestamp
	7,  // 4: pb.UserService.BlockUser:input_type -> pb.ID
	7,  // 5: pb.UserService.UnBlockUser:input_type -> pb.ID
	9,  // 6: pb.UserService.GetAllUsers:input_type -> pb.NoParam
	7,  // 7: pb.UserService.ViewProfile:input_type -> pb.ID
	6,  // 8: pb.UserService.AddSubscriptionPlan:input_type -> pb.USubscription
	9,  // 9: pb.UserService.GetAllPlans:input_type -> pb.NoParam
	6,  // 10: pb.UserService.UpdatePlan:input_type -> pb.USubscription
	3,  // 11: pb.UserService.GetUserProfileStats:input_type -> pb.UserStatsRequest
	1,  // 12: pb.UserService.GetSubscriptionStats:input_type -> pb.SubscriptionStatsRequest
	8,  // 13: pb.UserService.BlockUser:output_type -> pb.Response
	8,  // 14: pb.UserService.UnBlockUser:output_type -> pb.Response
	10, // 15: pb.UserService.GetAllUsers:output_type -> pb.UserList
	11, // 16: pb.UserService.ViewProfile:output_type -> pb.Profile
	8,  // 17: pb.UserService.AddSubscriptionPlan:output_type -> pb.Response
	5,  // 18: pb.UserService.GetAllPlans:output_type -> pb.UPlanList
	6,  // 19: pb.UserService.UpdatePlan:output_type -> pb.USubscription
	4,  // 20: pb.UserService.GetUserProfileStats:output_type -> pb.UserStatsProfileResponse
	2,  // 21: pb.UserService.GetSubscriptionStats:output_type -> pb.SubscriptionStatsResponse
	13, // [13:22] is the sub-list for method output_type
	4,  // [4:13] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionStatsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionStatsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserStatsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserStatsProfileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UPlanList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*USubscription); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_user_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*Response_Error)(nil),
		(*Response_Data)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		EnumInfos:         file_user_proto_enumTypes,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
