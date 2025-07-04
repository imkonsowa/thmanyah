// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.2
// source: conf/conf.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type Bootstrap struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Server        *Server                `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data          *Data                  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	mi := &file_conf_conf_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type Server struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Http          *Server_HTTP           `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc          *Server_GRPC           `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server) Reset() {
	*x = Server{}
	mi := &file_conf_conf_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Server) GetHttp() *Server_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type Database struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Host          string                 `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	User          string                 `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Dbname        string                 `protobuf:"bytes,4,opt,name=dbname,proto3" json:"dbname,omitempty"`
	Port          int64                  `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Database) Reset() {
	*x = Database{}
	mi := &file_conf_conf_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Database) ProtoMessage() {}

func (x *Database) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Database.ProtoReflect.Descriptor instead.
func (*Database) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Database) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Database) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Database) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Database) GetDbname() string {
	if x != nil {
		return x.Dbname
	}
	return ""
}

func (x *Database) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

type S3 struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Host           string                 `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	AccessKey      string                 `protobuf:"bytes,2,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey      string                 `protobuf:"bytes,3,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	Region         string                 `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	InitialBuckets []string               `protobuf:"bytes,5,rep,name=initial_buckets,json=initialBuckets,proto3" json:"initial_buckets,omitempty"`
	FilesHost      string                 `protobuf:"bytes,6,opt,name=files_host,json=filesHost,proto3" json:"files_host,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *S3) Reset() {
	*x = S3{}
	mi := &file_conf_conf_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *S3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3) ProtoMessage() {}

func (x *S3) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3.ProtoReflect.Descriptor instead.
func (*S3) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{3}
}

func (x *S3) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *S3) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *S3) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *S3) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *S3) GetInitialBuckets() []string {
	if x != nil {
		return x.InitialBuckets
	}
	return nil
}

func (x *S3) GetFilesHost() string {
	if x != nil {
		return x.FilesHost
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Postgres      *Database              `protobuf:"bytes,1,opt,name=postgres,proto3" json:"postgres,omitempty"`
	S3            *S3                    `protobuf:"bytes,3,opt,name=s3,proto3" json:"s3,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data) Reset() {
	*x = Data{}
	mi := &file_conf_conf_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{4}
}

func (x *Data) GetPostgres() *Database {
	if x != nil {
		return x.Postgres
	}
	return nil
}

func (x *Data) GetS3() *S3 {
	if x != nil {
		return x.S3
	}
	return nil
}

type Server_HTTP struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout       *durationpb.Duration   `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	mi := &file_conf_conf_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Server_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Server_GRPC struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout       *durationpb.Duration   `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server_GRPC) Reset() {
	*x = Server_GRPC{}
	mi := &file_conf_conf_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_GRPC) ProtoMessage() {}

func (x *Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_GRPC.ProtoReflect.Descriptor instead.
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Server_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

var File_conf_conf_proto protoreflect.FileDescriptor

const file_conf_conf_proto_rawDesc = "" +
	"\n" +
	"\x0fconf/conf.proto\x12\n" +
	"kratos.api\x1a\x1egoogle/protobuf/duration.proto\"]\n" +
	"\tBootstrap\x12*\n" +
	"\x06server\x18\x01 \x01(\v2\x12.kratos.api.ServerR\x06server\x12$\n" +
	"\x04data\x18\x02 \x01(\v2\x10.kratos.api.DataR\x04data\"\xb8\x02\n" +
	"\x06Server\x12+\n" +
	"\x04http\x18\x01 \x01(\v2\x17.kratos.api.Server.HTTPR\x04http\x12+\n" +
	"\x04grpc\x18\x02 \x01(\v2\x17.kratos.api.Server.GRPCR\x04grpc\x1ai\n" +
	"\x04HTTP\x12\x18\n" +
	"\anetwork\x18\x01 \x01(\tR\anetwork\x12\x12\n" +
	"\x04addr\x18\x02 \x01(\tR\x04addr\x123\n" +
	"\atimeout\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\atimeout\x1ai\n" +
	"\x04GRPC\x12\x18\n" +
	"\anetwork\x18\x01 \x01(\tR\anetwork\x12\x12\n" +
	"\x04addr\x18\x02 \x01(\tR\x04addr\x123\n" +
	"\atimeout\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\atimeout\"z\n" +
	"\bDatabase\x12\x12\n" +
	"\x04host\x18\x01 \x01(\tR\x04host\x12\x12\n" +
	"\x04user\x18\x02 \x01(\tR\x04user\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\x12\x16\n" +
	"\x06dbname\x18\x04 \x01(\tR\x06dbname\x12\x12\n" +
	"\x04port\x18\x05 \x01(\x03R\x04port\"\xb6\x01\n" +
	"\x02S3\x12\x12\n" +
	"\x04host\x18\x01 \x01(\tR\x04host\x12\x1d\n" +
	"\n" +
	"access_key\x18\x02 \x01(\tR\taccessKey\x12\x1d\n" +
	"\n" +
	"secret_key\x18\x03 \x01(\tR\tsecretKey\x12\x16\n" +
	"\x06region\x18\x04 \x01(\tR\x06region\x12'\n" +
	"\x0finitial_buckets\x18\x05 \x03(\tR\x0einitialBuckets\x12\x1d\n" +
	"\n" +
	"files_host\x18\x06 \x01(\tR\tfilesHost\"X\n" +
	"\x04Data\x120\n" +
	"\bpostgres\x18\x01 \x01(\v2\x14.kratos.api.DatabaseR\bpostgres\x12\x1e\n" +
	"\x02s3\x18\x03 \x01(\v2\x0e.kratos.api.S3R\x02s3B\x1fZ\x1dgeeksquest/internal/conf;confb\x06proto3"

var (
	file_conf_conf_proto_rawDescOnce sync.Once
	file_conf_conf_proto_rawDescData []byte
)

func file_conf_conf_proto_rawDescGZIP() []byte {
	file_conf_conf_proto_rawDescOnce.Do(func() {
		file_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_conf_conf_proto_rawDesc), len(file_conf_conf_proto_rawDesc)))
	})
	return file_conf_conf_proto_rawDescData
}

var file_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_conf_conf_proto_goTypes = []any{
	(*Bootstrap)(nil),           // 0: kratos.api.Bootstrap
	(*Server)(nil),              // 1: kratos.api.Server
	(*Database)(nil),            // 2: kratos.api.Database
	(*S3)(nil),                  // 3: kratos.api.S3
	(*Data)(nil),                // 4: kratos.api.Data
	(*Server_HTTP)(nil),         // 5: kratos.api.Server.HTTP
	(*Server_GRPC)(nil),         // 6: kratos.api.Server.GRPC
	(*durationpb.Duration)(nil), // 7: google.protobuf.Duration
}
var file_conf_conf_proto_depIdxs = []int32{
	1, // 0: kratos.api.Bootstrap.server:type_name -> kratos.api.Server
	4, // 1: kratos.api.Bootstrap.data:type_name -> kratos.api.Data
	5, // 2: kratos.api.Server.http:type_name -> kratos.api.Server.HTTP
	6, // 3: kratos.api.Server.grpc:type_name -> kratos.api.Server.GRPC
	2, // 4: kratos.api.Data.postgres:type_name -> kratos.api.Database
	3, // 5: kratos.api.Data.s3:type_name -> kratos.api.S3
	7, // 6: kratos.api.Server.HTTP.timeout:type_name -> google.protobuf.Duration
	7, // 7: kratos.api.Server.GRPC.timeout:type_name -> google.protobuf.Duration
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_conf_conf_proto_init() }
func file_conf_conf_proto_init() {
	if File_conf_conf_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_conf_conf_proto_rawDesc), len(file_conf_conf_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_conf_proto_goTypes,
		DependencyIndexes: file_conf_conf_proto_depIdxs,
		MessageInfos:      file_conf_conf_proto_msgTypes,
	}.Build()
	File_conf_conf_proto = out.File
	file_conf_conf_proto_goTypes = nil
	file_conf_conf_proto_depIdxs = nil
}
