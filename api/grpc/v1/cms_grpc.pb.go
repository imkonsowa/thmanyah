// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: v1/cms.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CmsService_CreateProgram_FullMethodName      = "/thmanyah.v1.CmsService/CreateProgram"
	CmsService_UpdateProgram_FullMethodName      = "/thmanyah.v1.CmsService/UpdateProgram"
	CmsService_DeleteProgram_FullMethodName      = "/thmanyah.v1.CmsService/DeleteProgram"
	CmsService_GetProgram_FullMethodName         = "/thmanyah.v1.CmsService/GetProgram"
	CmsService_ListPrograms_FullMethodName       = "/thmanyah.v1.CmsService/ListPrograms"
	CmsService_CreateCategory_FullMethodName     = "/thmanyah.v1.CmsService/CreateCategory"
	CmsService_UpdateCategory_FullMethodName     = "/thmanyah.v1.CmsService/UpdateCategory"
	CmsService_DeleteCategory_FullMethodName     = "/thmanyah.v1.CmsService/DeleteCategory"
	CmsService_GetCategory_FullMethodName        = "/thmanyah.v1.CmsService/GetCategory"
	CmsService_ListCategories_FullMethodName     = "/thmanyah.v1.CmsService/ListCategories"
	CmsService_CreateEpisode_FullMethodName      = "/thmanyah.v1.CmsService/CreateEpisode"
	CmsService_UpdateEpisode_FullMethodName      = "/thmanyah.v1.CmsService/UpdateEpisode"
	CmsService_DeleteEpisode_FullMethodName      = "/thmanyah.v1.CmsService/DeleteEpisode"
	CmsService_GetEpisode_FullMethodName         = "/thmanyah.v1.CmsService/GetEpisode"
	CmsService_ListEpisodes_FullMethodName       = "/thmanyah.v1.CmsService/ListEpisodes"
	CmsService_ImportData_FullMethodName         = "/thmanyah.v1.CmsService/ImportData"
	CmsService_BulkUpdatePrograms_FullMethodName = "/thmanyah.v1.CmsService/BulkUpdatePrograms"
	CmsService_BulkDeletePrograms_FullMethodName = "/thmanyah.v1.CmsService/BulkDeletePrograms"
)

// CmsServiceClient is the client API for CmsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CmsServiceClient interface {
	CreateProgram(ctx context.Context, in *CreateProgramRequest, opts ...grpc.CallOption) (*CreateProgramResponse, error)
	UpdateProgram(ctx context.Context, in *UpdateProgramRequest, opts ...grpc.CallOption) (*UpdateProgramResponse, error)
	DeleteProgram(ctx context.Context, in *DeleteProgramRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetProgram(ctx context.Context, in *GetProgramRequest, opts ...grpc.CallOption) (*GetProgramResponse, error)
	ListPrograms(ctx context.Context, in *ListProgramsRequest, opts ...grpc.CallOption) (*ListProgramsResponse, error)
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error)
	UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*UpdateCategoryResponse, error)
	DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*GetCategoryResponse, error)
	ListCategories(ctx context.Context, in *ListCategoriesRequest, opts ...grpc.CallOption) (*ListCategoriesResponse, error)
	CreateEpisode(ctx context.Context, in *CreateEpisodeRequest, opts ...grpc.CallOption) (*CreateEpisodeResponse, error)
	UpdateEpisode(ctx context.Context, in *UpdateEpisodeRequest, opts ...grpc.CallOption) (*UpdateEpisodeResponse, error)
	DeleteEpisode(ctx context.Context, in *DeleteEpisodeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetEpisode(ctx context.Context, in *GetEpisodeRequest, opts ...grpc.CallOption) (*GetEpisodeResponse, error)
	ListEpisodes(ctx context.Context, in *ListEpisodesRequest, opts ...grpc.CallOption) (*ListEpisodesResponse, error)
	ImportData(ctx context.Context, in *ImportDataRequest, opts ...grpc.CallOption) (*ImportDataResponse, error)
	BulkUpdatePrograms(ctx context.Context, in *BulkUpdateProgramsRequest, opts ...grpc.CallOption) (*BulkUpdateProgramsResponse, error)
	BulkDeletePrograms(ctx context.Context, in *BulkDeleteProgramsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type cmsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCmsServiceClient(cc grpc.ClientConnInterface) CmsServiceClient {
	return &cmsServiceClient{cc}
}

func (c *cmsServiceClient) CreateProgram(ctx context.Context, in *CreateProgramRequest, opts ...grpc.CallOption) (*CreateProgramResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProgramResponse)
	err := c.cc.Invoke(ctx, CmsService_CreateProgram_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) UpdateProgram(ctx context.Context, in *UpdateProgramRequest, opts ...grpc.CallOption) (*UpdateProgramResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProgramResponse)
	err := c.cc.Invoke(ctx, CmsService_UpdateProgram_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) DeleteProgram(ctx context.Context, in *DeleteProgramRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CmsService_DeleteProgram_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) GetProgram(ctx context.Context, in *GetProgramRequest, opts ...grpc.CallOption) (*GetProgramResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProgramResponse)
	err := c.cc.Invoke(ctx, CmsService_GetProgram_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) ListPrograms(ctx context.Context, in *ListProgramsRequest, opts ...grpc.CallOption) (*ListProgramsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProgramsResponse)
	err := c.cc.Invoke(ctx, CmsService_ListPrograms_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCategoryResponse)
	err := c.cc.Invoke(ctx, CmsService_CreateCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*UpdateCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCategoryResponse)
	err := c.cc.Invoke(ctx, CmsService_UpdateCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CmsService_DeleteCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*GetCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCategoryResponse)
	err := c.cc.Invoke(ctx, CmsService_GetCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) ListCategories(ctx context.Context, in *ListCategoriesRequest, opts ...grpc.CallOption) (*ListCategoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCategoriesResponse)
	err := c.cc.Invoke(ctx, CmsService_ListCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) CreateEpisode(ctx context.Context, in *CreateEpisodeRequest, opts ...grpc.CallOption) (*CreateEpisodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateEpisodeResponse)
	err := c.cc.Invoke(ctx, CmsService_CreateEpisode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) UpdateEpisode(ctx context.Context, in *UpdateEpisodeRequest, opts ...grpc.CallOption) (*UpdateEpisodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateEpisodeResponse)
	err := c.cc.Invoke(ctx, CmsService_UpdateEpisode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) DeleteEpisode(ctx context.Context, in *DeleteEpisodeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CmsService_DeleteEpisode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) GetEpisode(ctx context.Context, in *GetEpisodeRequest, opts ...grpc.CallOption) (*GetEpisodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEpisodeResponse)
	err := c.cc.Invoke(ctx, CmsService_GetEpisode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) ListEpisodes(ctx context.Context, in *ListEpisodesRequest, opts ...grpc.CallOption) (*ListEpisodesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListEpisodesResponse)
	err := c.cc.Invoke(ctx, CmsService_ListEpisodes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) ImportData(ctx context.Context, in *ImportDataRequest, opts ...grpc.CallOption) (*ImportDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ImportDataResponse)
	err := c.cc.Invoke(ctx, CmsService_ImportData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) BulkUpdatePrograms(ctx context.Context, in *BulkUpdateProgramsRequest, opts ...grpc.CallOption) (*BulkUpdateProgramsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BulkUpdateProgramsResponse)
	err := c.cc.Invoke(ctx, CmsService_BulkUpdatePrograms_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsServiceClient) BulkDeletePrograms(ctx context.Context, in *BulkDeleteProgramsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CmsService_BulkDeletePrograms_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CmsServiceServer is the server API for CmsService service.
// All implementations must embed UnimplementedCmsServiceServer
// for forward compatibility.
type CmsServiceServer interface {
	CreateProgram(context.Context, *CreateProgramRequest) (*CreateProgramResponse, error)
	UpdateProgram(context.Context, *UpdateProgramRequest) (*UpdateProgramResponse, error)
	DeleteProgram(context.Context, *DeleteProgramRequest) (*emptypb.Empty, error)
	GetProgram(context.Context, *GetProgramRequest) (*GetProgramResponse, error)
	ListPrograms(context.Context, *ListProgramsRequest) (*ListProgramsResponse, error)
	CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error)
	UpdateCategory(context.Context, *UpdateCategoryRequest) (*UpdateCategoryResponse, error)
	DeleteCategory(context.Context, *DeleteCategoryRequest) (*emptypb.Empty, error)
	GetCategory(context.Context, *GetCategoryRequest) (*GetCategoryResponse, error)
	ListCategories(context.Context, *ListCategoriesRequest) (*ListCategoriesResponse, error)
	CreateEpisode(context.Context, *CreateEpisodeRequest) (*CreateEpisodeResponse, error)
	UpdateEpisode(context.Context, *UpdateEpisodeRequest) (*UpdateEpisodeResponse, error)
	DeleteEpisode(context.Context, *DeleteEpisodeRequest) (*emptypb.Empty, error)
	GetEpisode(context.Context, *GetEpisodeRequest) (*GetEpisodeResponse, error)
	ListEpisodes(context.Context, *ListEpisodesRequest) (*ListEpisodesResponse, error)
	ImportData(context.Context, *ImportDataRequest) (*ImportDataResponse, error)
	BulkUpdatePrograms(context.Context, *BulkUpdateProgramsRequest) (*BulkUpdateProgramsResponse, error)
	BulkDeletePrograms(context.Context, *BulkDeleteProgramsRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCmsServiceServer()
}

// UnimplementedCmsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCmsServiceServer struct{}

func (UnimplementedCmsServiceServer) CreateProgram(context.Context, *CreateProgramRequest) (*CreateProgramResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProgram not implemented")
}
func (UnimplementedCmsServiceServer) UpdateProgram(context.Context, *UpdateProgramRequest) (*UpdateProgramResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProgram not implemented")
}
func (UnimplementedCmsServiceServer) DeleteProgram(context.Context, *DeleteProgramRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProgram not implemented")
}
func (UnimplementedCmsServiceServer) GetProgram(context.Context, *GetProgramRequest) (*GetProgramResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProgram not implemented")
}
func (UnimplementedCmsServiceServer) ListPrograms(context.Context, *ListProgramsRequest) (*ListProgramsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPrograms not implemented")
}
func (UnimplementedCmsServiceServer) CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedCmsServiceServer) UpdateCategory(context.Context, *UpdateCategoryRequest) (*UpdateCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedCmsServiceServer) DeleteCategory(context.Context, *DeleteCategoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedCmsServiceServer) GetCategory(context.Context, *GetCategoryRequest) (*GetCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedCmsServiceServer) ListCategories(context.Context, *ListCategoriesRequest) (*ListCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCategories not implemented")
}
func (UnimplementedCmsServiceServer) CreateEpisode(context.Context, *CreateEpisodeRequest) (*CreateEpisodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEpisode not implemented")
}
func (UnimplementedCmsServiceServer) UpdateEpisode(context.Context, *UpdateEpisodeRequest) (*UpdateEpisodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEpisode not implemented")
}
func (UnimplementedCmsServiceServer) DeleteEpisode(context.Context, *DeleteEpisodeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEpisode not implemented")
}
func (UnimplementedCmsServiceServer) GetEpisode(context.Context, *GetEpisodeRequest) (*GetEpisodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEpisode not implemented")
}
func (UnimplementedCmsServiceServer) ListEpisodes(context.Context, *ListEpisodesRequest) (*ListEpisodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEpisodes not implemented")
}
func (UnimplementedCmsServiceServer) ImportData(context.Context, *ImportDataRequest) (*ImportDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportData not implemented")
}
func (UnimplementedCmsServiceServer) BulkUpdatePrograms(context.Context, *BulkUpdateProgramsRequest) (*BulkUpdateProgramsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkUpdatePrograms not implemented")
}
func (UnimplementedCmsServiceServer) BulkDeletePrograms(context.Context, *BulkDeleteProgramsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkDeletePrograms not implemented")
}
func (UnimplementedCmsServiceServer) mustEmbedUnimplementedCmsServiceServer() {}
func (UnimplementedCmsServiceServer) testEmbeddedByValue()                    {}

// UnsafeCmsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CmsServiceServer will
// result in compilation errors.
type UnsafeCmsServiceServer interface {
	mustEmbedUnimplementedCmsServiceServer()
}

func RegisterCmsServiceServer(s grpc.ServiceRegistrar, srv CmsServiceServer) {
	// If the following call pancis, it indicates UnimplementedCmsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CmsService_ServiceDesc, srv)
}

func _CmsService_CreateProgram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProgramRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).CreateProgram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_CreateProgram_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).CreateProgram(ctx, req.(*CreateProgramRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_UpdateProgram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProgramRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).UpdateProgram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_UpdateProgram_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).UpdateProgram(ctx, req.(*UpdateProgramRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_DeleteProgram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProgramRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).DeleteProgram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_DeleteProgram_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).DeleteProgram(ctx, req.(*DeleteProgramRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_GetProgram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProgramRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).GetProgram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_GetProgram_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).GetProgram(ctx, req.(*GetProgramRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_ListPrograms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProgramsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).ListPrograms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_ListPrograms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).ListPrograms(ctx, req.(*ListProgramsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_CreateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_UpdateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).UpdateCategory(ctx, req.(*UpdateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_DeleteCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).DeleteCategory(ctx, req.(*DeleteCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_GetCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).GetCategory(ctx, req.(*GetCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_ListCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).ListCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_ListCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).ListCategories(ctx, req.(*ListCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_CreateEpisode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEpisodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).CreateEpisode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_CreateEpisode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).CreateEpisode(ctx, req.(*CreateEpisodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_UpdateEpisode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEpisodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).UpdateEpisode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_UpdateEpisode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).UpdateEpisode(ctx, req.(*UpdateEpisodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_DeleteEpisode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEpisodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).DeleteEpisode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_DeleteEpisode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).DeleteEpisode(ctx, req.(*DeleteEpisodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_GetEpisode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEpisodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).GetEpisode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_GetEpisode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).GetEpisode(ctx, req.(*GetEpisodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_ListEpisodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEpisodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).ListEpisodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_ListEpisodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).ListEpisodes(ctx, req.(*ListEpisodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_ImportData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).ImportData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_ImportData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).ImportData(ctx, req.(*ImportDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_BulkUpdatePrograms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkUpdateProgramsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).BulkUpdatePrograms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_BulkUpdatePrograms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).BulkUpdatePrograms(ctx, req.(*BulkUpdateProgramsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsService_BulkDeletePrograms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkDeleteProgramsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServiceServer).BulkDeletePrograms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CmsService_BulkDeletePrograms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServiceServer).BulkDeletePrograms(ctx, req.(*BulkDeleteProgramsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CmsService_ServiceDesc is the grpc.ServiceDesc for CmsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CmsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thmanyah.v1.CmsService",
	HandlerType: (*CmsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProgram",
			Handler:    _CmsService_CreateProgram_Handler,
		},
		{
			MethodName: "UpdateProgram",
			Handler:    _CmsService_UpdateProgram_Handler,
		},
		{
			MethodName: "DeleteProgram",
			Handler:    _CmsService_DeleteProgram_Handler,
		},
		{
			MethodName: "GetProgram",
			Handler:    _CmsService_GetProgram_Handler,
		},
		{
			MethodName: "ListPrograms",
			Handler:    _CmsService_ListPrograms_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _CmsService_CreateCategory_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _CmsService_UpdateCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _CmsService_DeleteCategory_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _CmsService_GetCategory_Handler,
		},
		{
			MethodName: "ListCategories",
			Handler:    _CmsService_ListCategories_Handler,
		},
		{
			MethodName: "CreateEpisode",
			Handler:    _CmsService_CreateEpisode_Handler,
		},
		{
			MethodName: "UpdateEpisode",
			Handler:    _CmsService_UpdateEpisode_Handler,
		},
		{
			MethodName: "DeleteEpisode",
			Handler:    _CmsService_DeleteEpisode_Handler,
		},
		{
			MethodName: "GetEpisode",
			Handler:    _CmsService_GetEpisode_Handler,
		},
		{
			MethodName: "ListEpisodes",
			Handler:    _CmsService_ListEpisodes_Handler,
		},
		{
			MethodName: "ImportData",
			Handler:    _CmsService_ImportData_Handler,
		},
		{
			MethodName: "BulkUpdatePrograms",
			Handler:    _CmsService_BulkUpdatePrograms_Handler,
		},
		{
			MethodName: "BulkDeletePrograms",
			Handler:    _CmsService_BulkDeletePrograms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/cms.proto",
}
