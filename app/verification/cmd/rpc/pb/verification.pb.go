// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: verification.proto

package pb

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

//req resp
type VerifyEmailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *VerifyEmailReq) Reset() {
	*x = VerifyEmailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyEmailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyEmailReq) ProtoMessage() {}

func (x *VerifyEmailReq) ProtoReflect() protoreflect.Message {
	mi := &file_verification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyEmailReq.ProtoReflect.Descriptor instead.
func (*VerifyEmailReq) Descriptor() ([]byte, []int) {
	return file_verification_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyEmailReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type VerifyEmailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VerifyEmailResp) Reset() {
	*x = VerifyEmailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyEmailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyEmailResp) ProtoMessage() {}

func (x *VerifyEmailResp) ProtoReflect() protoreflect.Message {
	mi := &file_verification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyEmailResp.ProtoReflect.Descriptor instead.
func (*VerifyEmailResp) Descriptor() ([]byte, []int) {
	return file_verification_proto_rawDescGZIP(), []int{1}
}

type VerifyImageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VerifyImageReq) Reset() {
	*x = VerifyImageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyImageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyImageReq) ProtoMessage() {}

func (x *VerifyImageReq) ProtoReflect() protoreflect.Message {
	mi := &file_verification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyImageReq.ProtoReflect.Descriptor instead.
func (*VerifyImageReq) Descriptor() ([]byte, []int) {
	return file_verification_proto_rawDescGZIP(), []int{2}
}

type VerifyImageResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageUrl string `protobuf:"bytes,1,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
}

func (x *VerifyImageResp) Reset() {
	*x = VerifyImageResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyImageResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyImageResp) ProtoMessage() {}

func (x *VerifyImageResp) ProtoReflect() protoreflect.Message {
	mi := &file_verification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyImageResp.ProtoReflect.Descriptor instead.
func (*VerifyImageResp) Descriptor() ([]byte, []int) {
	return file_verification_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyImageResp) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

var File_verification_proto protoreflect.FileDescriptor

var file_verification_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x26, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x22, 0x11, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x10, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x22, 0x2d, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x32, 0x7e, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x0b,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x62,
	0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_verification_proto_rawDescOnce sync.Once
	file_verification_proto_rawDescData = file_verification_proto_rawDesc
)

func file_verification_proto_rawDescGZIP() []byte {
	file_verification_proto_rawDescOnce.Do(func() {
		file_verification_proto_rawDescData = protoimpl.X.CompressGZIP(file_verification_proto_rawDescData)
	})
	return file_verification_proto_rawDescData
}

var file_verification_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_verification_proto_goTypes = []interface{}{
	(*VerifyEmailReq)(nil),  // 0: pb.VerifyEmailReq
	(*VerifyEmailResp)(nil), // 1: pb.VerifyEmailResp
	(*VerifyImageReq)(nil),  // 2: pb.VerifyImageReq
	(*VerifyImageResp)(nil), // 3: pb.VerifyImageResp
}
var file_verification_proto_depIdxs = []int32{
	0, // 0: pb.verification.verifyEmail:input_type -> pb.VerifyEmailReq
	2, // 1: pb.verification.verifyImage:input_type -> pb.VerifyImageReq
	1, // 2: pb.verification.verifyEmail:output_type -> pb.VerifyEmailResp
	3, // 3: pb.verification.verifyImage:output_type -> pb.VerifyImageResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_verification_proto_init() }
func file_verification_proto_init() {
	if File_verification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_verification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyEmailReq); i {
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
		file_verification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyEmailResp); i {
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
		file_verification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyImageReq); i {
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
		file_verification_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyImageResp); i {
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
			RawDescriptor: file_verification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_verification_proto_goTypes,
		DependencyIndexes: file_verification_proto_depIdxs,
		MessageInfos:      file_verification_proto_msgTypes,
	}.Build()
	File_verification_proto = out.File
	file_verification_proto_rawDesc = nil
	file_verification_proto_goTypes = nil
	file_verification_proto_depIdxs = nil
}
