// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: oauthcredentials/v1/oauth.proto

package oauthcredentials

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListOAuthsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *ListOAuthsRequest) Reset() {
	*x = ListOAuthsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOAuthsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOAuthsRequest) ProtoMessage() {}

func (x *ListOAuthsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOAuthsRequest.ProtoReflect.Descriptor instead.
func (*ListOAuthsRequest) Descriptor() ([]byte, []int) {
	return file_oauthcredentials_v1_oauth_proto_rawDescGZIP(), []int{0}
}

func (x *ListOAuthsRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type ListOAuthsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oauths []*OAuth `protobuf:"bytes,1,rep,name=oauths,proto3" json:"oauths,omitempty"`
}

func (x *ListOAuthsResponse) Reset() {
	*x = ListOAuthsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOAuthsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOAuthsResponse) ProtoMessage() {}

func (x *ListOAuthsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOAuthsResponse.ProtoReflect.Descriptor instead.
func (*ListOAuthsResponse) Descriptor() ([]byte, []int) {
	return file_oauthcredentials_v1_oauth_proto_rawDescGZIP(), []int{1}
}

func (x *ListOAuthsResponse) GetOauths() []*OAuth {
	if x != nil {
		return x.Oauths
	}
	return nil
}

type GetOAuthByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOAuthByIDRequest) Reset() {
	*x = GetOAuthByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOAuthByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOAuthByIDRequest) ProtoMessage() {}

func (x *GetOAuthByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOAuthByIDRequest.ProtoReflect.Descriptor instead.
func (*GetOAuthByIDRequest) Descriptor() ([]byte, []int) {
	return file_oauthcredentials_v1_oauth_proto_rawDescGZIP(), []int{2}
}

func (x *GetOAuthByIDRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *GetOAuthByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetOAuthByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oauth *OAuth `protobuf:"bytes,1,opt,name=oauth,proto3" json:"oauth,omitempty"`
}

func (x *GetOAuthByIDResponse) Reset() {
	*x = GetOAuthByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOAuthByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOAuthByIDResponse) ProtoMessage() {}

func (x *GetOAuthByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOAuthByIDResponse.ProtoReflect.Descriptor instead.
func (*GetOAuthByIDResponse) Descriptor() ([]byte, []int) {
	return file_oauthcredentials_v1_oauth_proto_rawDescGZIP(), []int{3}
}

func (x *GetOAuthByIDResponse) GetOauth() *OAuth {
	if x != nil {
		return x.Oauth
	}
	return nil
}

type GetOAuthByProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner    string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *GetOAuthByProviderRequest) Reset() {
	*x = GetOAuthByProviderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOAuthByProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOAuthByProviderRequest) ProtoMessage() {}

func (x *GetOAuthByProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOAuthByProviderRequest.ProtoReflect.Descriptor instead.
func (*GetOAuthByProviderRequest) Descriptor() ([]byte, []int) {
	return file_oauthcredentials_v1_oauth_proto_rawDescGZIP(), []int{4}
}

func (x *GetOAuthByProviderRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *GetOAuthByProviderRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

type GetOAuthByProviderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oauth *OAuth `protobuf:"bytes,1,opt,name=oauth,proto3" json:"oauth,omitempty"`
}

func (x *GetOAuthByProviderResponse) Reset() {
	*x = GetOAuthByProviderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOAuthByProviderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOAuthByProviderResponse) ProtoMessage() {}

func (x *GetOAuthByProviderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOAuthByProviderResponse.ProtoReflect.Descriptor instead.
func (*GetOAuthByProviderResponse) Descriptor() ([]byte, []int) {
	return file_oauthcredentials_v1_oauth_proto_rawDescGZIP(), []int{5}
}

func (x *GetOAuthByProviderResponse) GetOauth() *OAuth {
	if x != nil {
		return x.Oauth
	}
	return nil
}

type GetOAuthCredentialByProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner    string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *GetOAuthCredentialByProviderRequest) Reset() {
	*x = GetOAuthCredentialByProviderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOAuthCredentialByProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOAuthCredentialByProviderRequest) ProtoMessage() {}

func (x *GetOAuthCredentialByProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOAuthCredentialByProviderRequest.ProtoReflect.Descriptor instead.
func (*GetOAuthCredentialByProviderRequest) Descriptor() ([]byte, []int) {
	return file_oauthcredentials_v1_oauth_proto_rawDescGZIP(), []int{6}
}

func (x *GetOAuthCredentialByProviderRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *GetOAuthCredentialByProviderRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

type GetOAuthCredentialByProviderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *GetOAuthCredentialByProviderResponse) Reset() {
	*x = GetOAuthCredentialByProviderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOAuthCredentialByProviderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOAuthCredentialByProviderResponse) ProtoMessage() {}

func (x *GetOAuthCredentialByProviderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOAuthCredentialByProviderResponse.ProtoReflect.Descriptor instead.
func (*GetOAuthCredentialByProviderResponse) Descriptor() ([]byte, []int) {
	return file_oauthcredentials_v1_oauth_proto_rawDescGZIP(), []int{7}
}

func (x *GetOAuthCredentialByProviderResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type OAuthTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner       string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Provider    string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Code        string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	RedirectUri string `protobuf:"bytes,4,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
}

func (x *OAuthTokenRequest) Reset() {
	*x = OAuthTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthTokenRequest) ProtoMessage() {}

func (x *OAuthTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthTokenRequest.ProtoReflect.Descriptor instead.
func (*OAuthTokenRequest) Descriptor() ([]byte, []int) {
	return file_oauthcredentials_v1_oauth_proto_rawDescGZIP(), []int{8}
}

func (x *OAuthTokenRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *OAuthTokenRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *OAuthTokenRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *OAuthTokenRequest) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

type ExchangeCodeForTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner    string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Code     string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	State    string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *ExchangeCodeForTokenRequest) Reset() {
	*x = ExchangeCodeForTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeCodeForTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeCodeForTokenRequest) ProtoMessage() {}

func (x *ExchangeCodeForTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeCodeForTokenRequest.ProtoReflect.Descriptor instead.
func (*ExchangeCodeForTokenRequest) Descriptor() ([]byte, []int) {
	return file_oauthcredentials_v1_oauth_proto_rawDescGZIP(), []int{9}
}

func (x *ExchangeCodeForTokenRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ExchangeCodeForTokenRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *ExchangeCodeForTokenRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ExchangeCodeForTokenRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type ExchangeCodeForTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExchangeCodeForTokenResponse) Reset() {
	*x = ExchangeCodeForTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeCodeForTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeCodeForTokenResponse) ProtoMessage() {}

func (x *ExchangeCodeForTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeCodeForTokenResponse.ProtoReflect.Descriptor instead.
func (*ExchangeCodeForTokenResponse) Descriptor() ([]byte, []int) {
	return file_oauthcredentials_v1_oauth_proto_rawDescGZIP(), []int{10}
}

type OAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner    string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Provider string   `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	Scopes   []string `protobuf:"bytes,6,rep,name=scopes,proto3" json:"scopes,omitempty"`
}

func (x *OAuth) Reset() {
	*x = OAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth) ProtoMessage() {}

func (x *OAuth) ProtoReflect() protoreflect.Message {
	mi := &file_oauthcredentials_v1_oauth_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuth.ProtoReflect.Descriptor instead.
func (*OAuth) Descriptor() ([]byte, []int) {
	return file_oauthcredentials_v1_oauth_proto_rawDescGZIP(), []int{11}
}

func (x *OAuth) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OAuth) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *OAuth) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *OAuth) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

var File_oauthcredentials_v1_oauth_proto protoreflect.FileDescriptor

var file_oauthcredentials_v1_oauth_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x29, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x22, 0x48, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x06, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x73, 0x22, 0x3b, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x48, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f,
	0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x05, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x52, 0x05, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x22, 0x4d, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x22, 0x4e, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x05, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x52, 0x05, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x22, 0x57, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x49, 0x0a, 0x24, 0x47, 0x65,
	0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x42, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7c, 0x0a, 0x11, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x55, 0x72, 0x69, 0x22, 0x79, 0x0a, 0x1b, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x1e,
	0x0a, 0x1c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x6f,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x61,
	0x0a, 0x05, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x73, 0x32, 0xe8, 0x04, 0x0a, 0x0c, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x61, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x73,
	0x12, 0x26, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x65, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74,
	0x68, 0x42, 0x79, 0x49, 0x44, 0x12, 0x28, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f,
	0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x12, 0x2e, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74,
	0x68, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74,
	0x68, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x95, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75,
	0x74, 0x68, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x79, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x38, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x42,
	0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x39, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7d, 0x0a,
	0x14, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x30, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x61, 0x5a, 0x5f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x2d, 0x69, 0x6f, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oauthcredentials_v1_oauth_proto_rawDescOnce sync.Once
	file_oauthcredentials_v1_oauth_proto_rawDescData = file_oauthcredentials_v1_oauth_proto_rawDesc
)

func file_oauthcredentials_v1_oauth_proto_rawDescGZIP() []byte {
	file_oauthcredentials_v1_oauth_proto_rawDescOnce.Do(func() {
		file_oauthcredentials_v1_oauth_proto_rawDescData = protoimpl.X.CompressGZIP(file_oauthcredentials_v1_oauth_proto_rawDescData)
	})
	return file_oauthcredentials_v1_oauth_proto_rawDescData
}

var file_oauthcredentials_v1_oauth_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_oauthcredentials_v1_oauth_proto_goTypes = []interface{}{
	(*ListOAuthsRequest)(nil),                    // 0: oauthcredentials.v1.ListOAuthsRequest
	(*ListOAuthsResponse)(nil),                   // 1: oauthcredentials.v1.ListOAuthsResponse
	(*GetOAuthByIDRequest)(nil),                  // 2: oauthcredentials.v1.GetOAuthByIDRequest
	(*GetOAuthByIDResponse)(nil),                 // 3: oauthcredentials.v1.GetOAuthByIDResponse
	(*GetOAuthByProviderRequest)(nil),            // 4: oauthcredentials.v1.GetOAuthByProviderRequest
	(*GetOAuthByProviderResponse)(nil),           // 5: oauthcredentials.v1.GetOAuthByProviderResponse
	(*GetOAuthCredentialByProviderRequest)(nil),  // 6: oauthcredentials.v1.GetOAuthCredentialByProviderRequest
	(*GetOAuthCredentialByProviderResponse)(nil), // 7: oauthcredentials.v1.GetOAuthCredentialByProviderResponse
	(*OAuthTokenRequest)(nil),                    // 8: oauthcredentials.v1.OAuthTokenRequest
	(*ExchangeCodeForTokenRequest)(nil),          // 9: oauthcredentials.v1.ExchangeCodeForTokenRequest
	(*ExchangeCodeForTokenResponse)(nil),         // 10: oauthcredentials.v1.ExchangeCodeForTokenResponse
	(*OAuth)(nil),                                // 11: oauthcredentials.v1.OAuth
}
var file_oauthcredentials_v1_oauth_proto_depIdxs = []int32{
	11, // 0: oauthcredentials.v1.ListOAuthsResponse.oauths:type_name -> oauthcredentials.v1.OAuth
	11, // 1: oauthcredentials.v1.GetOAuthByIDResponse.oauth:type_name -> oauthcredentials.v1.OAuth
	11, // 2: oauthcredentials.v1.GetOAuthByProviderResponse.oauth:type_name -> oauthcredentials.v1.OAuth
	0,  // 3: oauthcredentials.v1.OAuthService.ListOAuths:input_type -> oauthcredentials.v1.ListOAuthsRequest
	2,  // 4: oauthcredentials.v1.OAuthService.GetOAuthByID:input_type -> oauthcredentials.v1.GetOAuthByIDRequest
	4,  // 5: oauthcredentials.v1.OAuthService.GetOAuthByProvider:input_type -> oauthcredentials.v1.GetOAuthByProviderRequest
	6,  // 6: oauthcredentials.v1.OAuthService.GetOAuthCredentialByProvider:input_type -> oauthcredentials.v1.GetOAuthCredentialByProviderRequest
	9,  // 7: oauthcredentials.v1.OAuthService.ExchangeCodeForToken:input_type -> oauthcredentials.v1.ExchangeCodeForTokenRequest
	1,  // 8: oauthcredentials.v1.OAuthService.ListOAuths:output_type -> oauthcredentials.v1.ListOAuthsResponse
	3,  // 9: oauthcredentials.v1.OAuthService.GetOAuthByID:output_type -> oauthcredentials.v1.GetOAuthByIDResponse
	5,  // 10: oauthcredentials.v1.OAuthService.GetOAuthByProvider:output_type -> oauthcredentials.v1.GetOAuthByProviderResponse
	7,  // 11: oauthcredentials.v1.OAuthService.GetOAuthCredentialByProvider:output_type -> oauthcredentials.v1.GetOAuthCredentialByProviderResponse
	10, // 12: oauthcredentials.v1.OAuthService.ExchangeCodeForToken:output_type -> oauthcredentials.v1.ExchangeCodeForTokenResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_oauthcredentials_v1_oauth_proto_init() }
func file_oauthcredentials_v1_oauth_proto_init() {
	if File_oauthcredentials_v1_oauth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oauthcredentials_v1_oauth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOAuthsRequest); i {
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
		file_oauthcredentials_v1_oauth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOAuthsResponse); i {
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
		file_oauthcredentials_v1_oauth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOAuthByIDRequest); i {
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
		file_oauthcredentials_v1_oauth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOAuthByIDResponse); i {
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
		file_oauthcredentials_v1_oauth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOAuthByProviderRequest); i {
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
		file_oauthcredentials_v1_oauth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOAuthByProviderResponse); i {
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
		file_oauthcredentials_v1_oauth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOAuthCredentialByProviderRequest); i {
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
		file_oauthcredentials_v1_oauth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOAuthCredentialByProviderResponse); i {
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
		file_oauthcredentials_v1_oauth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthTokenRequest); i {
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
		file_oauthcredentials_v1_oauth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeCodeForTokenRequest); i {
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
		file_oauthcredentials_v1_oauth_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeCodeForTokenResponse); i {
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
		file_oauthcredentials_v1_oauth_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuth); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_oauthcredentials_v1_oauth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oauthcredentials_v1_oauth_proto_goTypes,
		DependencyIndexes: file_oauthcredentials_v1_oauth_proto_depIdxs,
		MessageInfos:      file_oauthcredentials_v1_oauth_proto_msgTypes,
	}.Build()
	File_oauthcredentials_v1_oauth_proto = out.File
	file_oauthcredentials_v1_oauth_proto_rawDesc = nil
	file_oauthcredentials_v1_oauth_proto_goTypes = nil
	file_oauthcredentials_v1_oauth_proto_depIdxs = nil
}
