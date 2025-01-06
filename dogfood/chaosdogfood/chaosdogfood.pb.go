// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2025 Datadog, Inc.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: chaosdogfood.proto

package chaosdogfood

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Animal string `protobuf:"bytes,1,opt,name=animal,proto3" json:"animal,omitempty"`
}

func (x *FoodRequest) Reset() {
	*x = FoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaosdogfood_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodRequest) ProtoMessage() {}

func (x *FoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chaosdogfood_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodRequest.ProtoReflect.Descriptor instead.
func (*FoodRequest) Descriptor() ([]byte, []int) {
	return file_chaosdogfood_proto_rawDescGZIP(), []int{0}
}

func (x *FoodRequest) GetAnimal() string {
	if x != nil {
		return x.Animal
	}
	return ""
}

type FoodReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message        string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ConfirmationId int32  `protobuf:"varint,2,opt,name=confirmation_id,json=confirmationId,proto3" json:"confirmation_id,omitempty"`
}

func (x *FoodReply) Reset() {
	*x = FoodReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaosdogfood_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodReply) ProtoMessage() {}

func (x *FoodReply) ProtoReflect() protoreflect.Message {
	mi := &file_chaosdogfood_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodReply.ProtoReflect.Descriptor instead.
func (*FoodReply) Descriptor() ([]byte, []int) {
	return file_chaosdogfood_proto_rawDescGZIP(), []int{1}
}

func (x *FoodReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *FoodReply) GetConfirmationId() int32 {
	if x != nil {
		return x.ConfirmationId
	}
	return 0
}

type CatalogReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*CatalogItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CatalogReply) Reset() {
	*x = CatalogReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaosdogfood_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CatalogReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatalogReply) ProtoMessage() {}

func (x *CatalogReply) ProtoReflect() protoreflect.Message {
	mi := &file_chaosdogfood_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatalogReply.ProtoReflect.Descriptor instead.
func (*CatalogReply) Descriptor() ([]byte, []int) {
	return file_chaosdogfood_proto_rawDescGZIP(), []int{2}
}

func (x *CatalogReply) GetItems() []*CatalogItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type CatalogItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Animal string `protobuf:"bytes,1,opt,name=animal,proto3" json:"animal,omitempty"`
	Food   string `protobuf:"bytes,2,opt,name=food,proto3" json:"food,omitempty"`
}

func (x *CatalogItem) Reset() {
	*x = CatalogItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaosdogfood_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CatalogItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatalogItem) ProtoMessage() {}

func (x *CatalogItem) ProtoReflect() protoreflect.Message {
	mi := &file_chaosdogfood_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatalogItem.ProtoReflect.Descriptor instead.
func (*CatalogItem) Descriptor() ([]byte, []int) {
	return file_chaosdogfood_proto_rawDescGZIP(), []int{3}
}

func (x *CatalogItem) GetAnimal() string {
	if x != nil {
		return x.Animal
	}
	return ""
}

func (x *CatalogItem) GetFood() string {
	if x != nil {
		return x.Food
	}
	return ""
}

var File_chaosdogfood_proto protoreflect.FileDescriptor

var file_chaosdogfood_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x64, 0x6f, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x64, 0x6f, 0x67, 0x66, 0x6f,
	0x6f, 0x64, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x25, 0x0a, 0x0b, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x22, 0x4e, 0x0a, 0x09, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x0c, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x64, 0x6f, 0x67,
	0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x39, 0x0a, 0x0b, 0x43, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x6f, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x6f,
	0x6f, 0x64, 0x32, 0x91, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6f, 0x73, 0x44, 0x6f, 0x67, 0x66,
	0x6f, 0x6f, 0x64, 0x12, 0x3d, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x63,
	0x68, 0x61, 0x6f, 0x73, 0x64, 0x6f, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x64,
	0x6f, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x42, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73,
	0x64, 0x6f, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x63, 0x68, 0x61, 0x6f,
	0x73, 0x64, 0x6f, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chaosdogfood_proto_rawDescOnce sync.Once
	file_chaosdogfood_proto_rawDescData = file_chaosdogfood_proto_rawDesc
)

func file_chaosdogfood_proto_rawDescGZIP() []byte {
	file_chaosdogfood_proto_rawDescOnce.Do(func() {
		file_chaosdogfood_proto_rawDescData = protoimpl.X.CompressGZIP(file_chaosdogfood_proto_rawDescData)
	})
	return file_chaosdogfood_proto_rawDescData
}

var file_chaosdogfood_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chaosdogfood_proto_goTypes = []interface{}{
	(*FoodRequest)(nil),   // 0: chaosdogfood.FoodRequest
	(*FoodReply)(nil),     // 1: chaosdogfood.FoodReply
	(*CatalogReply)(nil),  // 2: chaosdogfood.CatalogReply
	(*CatalogItem)(nil),   // 3: chaosdogfood.CatalogItem
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_chaosdogfood_proto_depIdxs = []int32{
	3, // 0: chaosdogfood.CatalogReply.items:type_name -> chaosdogfood.CatalogItem
	0, // 1: chaosdogfood.ChaosDogfood.order:input_type -> chaosdogfood.FoodRequest
	4, // 2: chaosdogfood.ChaosDogfood.getCatalog:input_type -> google.protobuf.Empty
	1, // 3: chaosdogfood.ChaosDogfood.order:output_type -> chaosdogfood.FoodReply
	2, // 4: chaosdogfood.ChaosDogfood.getCatalog:output_type -> chaosdogfood.CatalogReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chaosdogfood_proto_init() }
func file_chaosdogfood_proto_init() {
	if File_chaosdogfood_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chaosdogfood_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodRequest); i {
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
		file_chaosdogfood_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodReply); i {
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
		file_chaosdogfood_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CatalogReply); i {
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
		file_chaosdogfood_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CatalogItem); i {
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
			RawDescriptor: file_chaosdogfood_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chaosdogfood_proto_goTypes,
		DependencyIndexes: file_chaosdogfood_proto_depIdxs,
		MessageInfos:      file_chaosdogfood_proto_msgTypes,
	}.Build()
	File_chaosdogfood_proto = out.File
	file_chaosdogfood_proto_rawDesc = nil
	file_chaosdogfood_proto_goTypes = nil
	file_chaosdogfood_proto_depIdxs = nil
}
