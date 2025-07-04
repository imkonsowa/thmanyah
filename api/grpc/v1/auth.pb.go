// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: v1/auth.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_v1_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,proto3" json:"access_token,omitempty"`
	User          *User                  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_v1_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type RegisterRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Email           string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password        string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	ConfirmPassword string                 `protobuf:"bytes,3,opt,name=confirm_password,proto3" json:"confirm_password,omitempty"`
	Name            string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_v1_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetConfirmPassword() string {
	if x != nil {
		return x.ConfirmPassword
	}
	return ""
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,proto3" json:"access_token,omitempty"`
	User          *User                  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	mi := &file_v1_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RegisterResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type RefreshTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RefreshToken  string                 `protobuf:"bytes,1,opt,name=refresh_token,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshTokenRequest) Reset() {
	*x = RefreshTokenRequest{}
	mi := &file_v1_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRequest) ProtoMessage() {}

func (x *RefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{4}
}

func (x *RefreshTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RefreshTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,proto3" json:"access_token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refresh_token,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshTokenResponse) Reset() {
	*x = RefreshTokenResponse{}
	mi := &file_v1_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenResponse) ProtoMessage() {}

func (x *RefreshTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenResponse.ProtoReflect.Descriptor instead.
func (*RefreshTokenResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{5}
}

func (x *RefreshTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RefreshTokenResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type Socials struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Twitter       string                 `protobuf:"bytes,1,opt,name=twitter,proto3" json:"twitter,omitempty"`
	Github        string                 `protobuf:"bytes,2,opt,name=github,proto3" json:"github,omitempty"`
	Linkedin      string                 `protobuf:"bytes,3,opt,name=linkedin,proto3" json:"linkedin,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Socials) Reset() {
	*x = Socials{}
	mi := &file_v1_auth_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Socials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Socials) ProtoMessage() {}

func (x *Socials) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Socials.ProtoReflect.Descriptor instead.
func (*Socials) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{6}
}

func (x *Socials) GetTwitter() string {
	if x != nil {
		return x.Twitter
	}
	return ""
}

func (x *Socials) GetGithub() string {
	if x != nil {
		return x.Github
	}
	return ""
}

func (x *Socials) GetLinkedin() string {
	if x != nil {
		return x.Linkedin
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,4,opt,name=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,5,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_v1_auth_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{7}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *User) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UserProfileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Positions     string                 `protobuf:"bytes,2,opt,name=positions,proto3" json:"positions,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Socials       *Socials               `protobuf:"bytes,4,opt,name=socials,proto3" json:"socials,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserProfileRequest) Reset() {
	*x = UserProfileRequest{}
	mi := &file_v1_auth_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfileRequest) ProtoMessage() {}

func (x *UserProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfileRequest.ProtoReflect.Descriptor instead.
func (*UserProfileRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{8}
}

func (x *UserProfileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserProfileRequest) GetPositions() string {
	if x != nil {
		return x.Positions
	}
	return ""
}

func (x *UserProfileRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserProfileRequest) GetSocials() *Socials {
	if x != nil {
		return x.Socials
	}
	return nil
}

type UserProfileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserProfileResponse) Reset() {
	*x = UserProfileResponse{}
	mi := &file_v1_auth_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfileResponse) ProtoMessage() {}

func (x *UserProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfileResponse.ProtoReflect.Descriptor instead.
func (*UserProfileResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{9}
}

func (x *UserProfileResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Position      string                 `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Socials       *Socials               `protobuf:"bytes,3,opt,name=socials,proto3" json:"socials,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_v1_auth_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserRequest) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *UpdateUserRequest) GetSocials() *Socials {
	if x != nil {
		return x.Socials
	}
	return nil
}

type UpdateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserResponse) Reset() {
	*x = UpdateUserResponse{}
	mi := &file_v1_auth_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResponse) ProtoMessage() {}

func (x *UpdateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_v1_auth_proto protoreflect.FileDescriptor

const file_v1_auth_proto_rawDesc = "" +
	"\n" +
	"\rv1/auth.proto\x12\vthmanyah.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x17validate/validate.proto\x1a\x1bgoogle/protobuf/empty.proto\"\x89\x01\n" +
	"\fLoginRequest\x12R\n" +
	"\x05email\x18\x01 \x01(\tB<\xfaB9r7\x10\x01\x18\x80\x0120^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$R\x05email\x12%\n" +
	"\bpassword\x18\x02 \x01(\tB\t\xfaB\x06r\x04\x10\x06\x18 R\bpassword\"Z\n" +
	"\rLoginResponse\x12\"\n" +
	"\faccess_token\x18\x01 \x01(\tR\faccess_token\x12%\n" +
	"\x04user\x18\x02 \x01(\v2\x11.thmanyah.v1.UserR\x04user\"\xe3\x01\n" +
	"\x0fRegisterRequest\x12R\n" +
	"\x05email\x18\x01 \x01(\tB<\xfaB9r7\x10\x01\x18\x80\x0120^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$R\x05email\x12%\n" +
	"\bpassword\x18\x02 \x01(\tB\t\xfaB\x06r\x04\x10\x06\x18 R\bpassword\x125\n" +
	"\x10confirm_password\x18\x03 \x01(\tB\t\xfaB\x06r\x04\x10\x06\x18 R\x10confirm_password\x12\x1e\n" +
	"\x04name\x18\x04 \x01(\tB\n" +
	"\xfaB\ar\x05\x10\x01\x18\x80\x01R\x04name\"]\n" +
	"\x10RegisterResponse\x12\"\n" +
	"\faccess_token\x18\x01 \x01(\tR\faccess_token\x12%\n" +
	"\x04user\x18\x02 \x01(\v2\x11.thmanyah.v1.UserR\x04user\";\n" +
	"\x13RefreshTokenRequest\x12$\n" +
	"\rrefresh_token\x18\x01 \x01(\tR\rrefresh_token\"`\n" +
	"\x14RefreshTokenResponse\x12\"\n" +
	"\faccess_token\x18\x01 \x01(\tR\faccess_token\x12$\n" +
	"\rrefresh_token\x18\x02 \x01(\tR\rrefresh_token\"W\n" +
	"\aSocials\x12\x18\n" +
	"\atwitter\x18\x01 \x01(\tR\atwitter\x12\x16\n" +
	"\x06github\x18\x02 \x01(\tR\x06github\x12\x1a\n" +
	"\blinkedin\x18\x03 \x01(\tR\blinkedin\"\x80\x01\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x1e\n" +
	"\n" +
	"created_at\x18\x04 \x01(\tR\n" +
	"created_at\x12\x1e\n" +
	"\n" +
	"updated_at\x18\x05 \x01(\tR\n" +
	"updated_at\"\x8c\x01\n" +
	"\x12UserProfileRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1c\n" +
	"\tpositions\x18\x02 \x01(\tR\tpositions\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12.\n" +
	"\asocials\x18\x04 \x01(\v2\x14.thmanyah.v1.SocialsR\asocials\"<\n" +
	"\x13UserProfileResponse\x12%\n" +
	"\x04user\x18\x01 \x01(\v2\x11.thmanyah.v1.UserR\x04user\"s\n" +
	"\x11UpdateUserRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1a\n" +
	"\bposition\x18\x02 \x01(\tR\bposition\x12.\n" +
	"\asocials\x18\x03 \x01(\v2\x14.thmanyah.v1.SocialsR\asocials\";\n" +
	"\x12UpdateUserResponse\x12%\n" +
	"\x04user\x18\x01 \x01(\v2\x11.thmanyah.v1.UserR\x04user2\xb4\x04\n" +
	"\vAuthService\x12]\n" +
	"\x05Login\x12\x19.thmanyah.v1.LoginRequest\x1a\x1a.thmanyah.v1.LoginResponse\"\x1d\x82\xd3\xe4\x93\x02\x17:\x01*\"\x12/api/v1/auth/login\x12i\n" +
	"\bRegister\x12\x1c.thmanyah.v1.RegisterRequest\x1a\x1d.thmanyah.v1.RegisterResponse\" \x82\xd3\xe4\x93\x02\x1a:\x01*\"\x15/api/v1/auth/register\x12z\n" +
	"\fRefreshToken\x12 .thmanyah.v1.RefreshTokenRequest\x1a!.thmanyah.v1.RefreshTokenResponse\"%\x82\xd3\xe4\x93\x02\x1f:\x01*\"\x1a/api/v1/auth/refresh-token\x12h\n" +
	"\x0eGetUserProfile\x12\x16.google.protobuf.Empty\x1a .thmanyah.v1.UserProfileResponse\"\x1c\x82\xd3\xe4\x93\x02\x16\x12\x14/api/v1/auth/profile\x12u\n" +
	"\x11UpdateUserProfile\x12\x1e.thmanyah.v1.UpdateUserRequest\x1a\x1f.thmanyah.v1.UpdateUserResponse\"\x1f\x82\xd3\xe4\x93\x02\x19:\x01*\x1a\x14/api/v1/auth/profileB\x14Z\x12thmanyah/api/v1;v1b\x06proto3"

var (
	file_v1_auth_proto_rawDescOnce sync.Once
	file_v1_auth_proto_rawDescData []byte
)

func file_v1_auth_proto_rawDescGZIP() []byte {
	file_v1_auth_proto_rawDescOnce.Do(func() {
		file_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_auth_proto_rawDesc), len(file_v1_auth_proto_rawDesc)))
	})
	return file_v1_auth_proto_rawDescData
}

var file_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_v1_auth_proto_goTypes = []any{
	(*LoginRequest)(nil),         // 0: thmanyah.v1.LoginRequest
	(*LoginResponse)(nil),        // 1: thmanyah.v1.LoginResponse
	(*RegisterRequest)(nil),      // 2: thmanyah.v1.RegisterRequest
	(*RegisterResponse)(nil),     // 3: thmanyah.v1.RegisterResponse
	(*RefreshTokenRequest)(nil),  // 4: thmanyah.v1.RefreshTokenRequest
	(*RefreshTokenResponse)(nil), // 5: thmanyah.v1.RefreshTokenResponse
	(*Socials)(nil),              // 6: thmanyah.v1.Socials
	(*User)(nil),                 // 7: thmanyah.v1.User
	(*UserProfileRequest)(nil),   // 8: thmanyah.v1.UserProfileRequest
	(*UserProfileResponse)(nil),  // 9: thmanyah.v1.UserProfileResponse
	(*UpdateUserRequest)(nil),    // 10: thmanyah.v1.UpdateUserRequest
	(*UpdateUserResponse)(nil),   // 11: thmanyah.v1.UpdateUserResponse
	(*emptypb.Empty)(nil),        // 12: google.protobuf.Empty
}
var file_v1_auth_proto_depIdxs = []int32{
	7,  // 0: thmanyah.v1.LoginResponse.user:type_name -> thmanyah.v1.User
	7,  // 1: thmanyah.v1.RegisterResponse.user:type_name -> thmanyah.v1.User
	6,  // 2: thmanyah.v1.UserProfileRequest.socials:type_name -> thmanyah.v1.Socials
	7,  // 3: thmanyah.v1.UserProfileResponse.user:type_name -> thmanyah.v1.User
	6,  // 4: thmanyah.v1.UpdateUserRequest.socials:type_name -> thmanyah.v1.Socials
	7,  // 5: thmanyah.v1.UpdateUserResponse.user:type_name -> thmanyah.v1.User
	0,  // 6: thmanyah.v1.AuthService.Login:input_type -> thmanyah.v1.LoginRequest
	2,  // 7: thmanyah.v1.AuthService.Register:input_type -> thmanyah.v1.RegisterRequest
	4,  // 8: thmanyah.v1.AuthService.RefreshToken:input_type -> thmanyah.v1.RefreshTokenRequest
	12, // 9: thmanyah.v1.AuthService.GetUserProfile:input_type -> google.protobuf.Empty
	10, // 10: thmanyah.v1.AuthService.UpdateUserProfile:input_type -> thmanyah.v1.UpdateUserRequest
	1,  // 11: thmanyah.v1.AuthService.Login:output_type -> thmanyah.v1.LoginResponse
	3,  // 12: thmanyah.v1.AuthService.Register:output_type -> thmanyah.v1.RegisterResponse
	5,  // 13: thmanyah.v1.AuthService.RefreshToken:output_type -> thmanyah.v1.RefreshTokenResponse
	9,  // 14: thmanyah.v1.AuthService.GetUserProfile:output_type -> thmanyah.v1.UserProfileResponse
	11, // 15: thmanyah.v1.AuthService.UpdateUserProfile:output_type -> thmanyah.v1.UpdateUserResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_v1_auth_proto_init() }
func file_v1_auth_proto_init() {
	if File_v1_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_auth_proto_rawDesc), len(file_v1_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_auth_proto_goTypes,
		DependencyIndexes: file_v1_auth_proto_depIdxs,
		MessageInfos:      file_v1_auth_proto_msgTypes,
	}.Build()
	File_v1_auth_proto = out.File
	file_v1_auth_proto_goTypes = nil
	file_v1_auth_proto_depIdxs = nil
}
