// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: api/common.proto

package scql

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

type RequestHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomHeaders map[string]string `protobuf:"bytes,1,rep,name=custom_headers,json=customHeaders,proto3" json:"custom_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RequestHeader) Reset() {
	*x = RequestHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHeader) ProtoMessage() {}

func (x *RequestHeader) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHeader.ProtoReflect.Descriptor instead.
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return file_api_common_proto_rawDescGZIP(), []int{0}
}

func (x *RequestHeader) GetCustomHeaders() map[string]string {
	if x != nil {
		return x.CustomHeaders
	}
	return nil
}

type ReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status          *Status   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	OutColumns      []*Tensor `protobuf:"bytes,2,rep,name=out_columns,json=outColumns,proto3" json:"out_columns,omitempty"`
	DagId           int32     `protobuf:"varint,3,opt,name=dag_id,json=dagId,proto3" json:"dag_id,omitempty"`
	SessionId       string    `protobuf:"bytes,4,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	PartyCode       string    `protobuf:"bytes,5,opt,name=party_code,json=partyCode,proto3" json:"party_code,omitempty"`
	NumRowsAffected int64     `protobuf:"varint,6,opt,name=num_rows_affected,json=numRowsAffected,proto3" json:"num_rows_affected,omitempty"`
}

func (x *ReportRequest) Reset() {
	*x = ReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRequest) ProtoMessage() {}

func (x *ReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRequest.ProtoReflect.Descriptor instead.
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return file_api_common_proto_rawDescGZIP(), []int{1}
}

func (x *ReportRequest) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ReportRequest) GetOutColumns() []*Tensor {
	if x != nil {
		return x.OutColumns
	}
	return nil
}

func (x *ReportRequest) GetDagId() int32 {
	if x != nil {
		return x.DagId
	}
	return 0
}

func (x *ReportRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ReportRequest) GetPartyCode() string {
	if x != nil {
		return x.PartyCode
	}
	return ""
}

func (x *ReportRequest) GetNumRowsAffected() int64 {
	if x != nil {
		return x.NumRowsAffected
	}
	return 0
}

var File_api_common_proto protoreflect.FileDescriptor

var file_api_common_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x1a, 0x0e, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01,
	0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x50, 0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x1a, 0x40, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xeb, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30,
	0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73,
	0x12, 0x15, 0x0a, 0x06, 0x64, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x64, 0x61, 0x67, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x6f, 0x77,
	0x73, 0x5f, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x52, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x73,
	0x63, 0x71, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_common_proto_rawDescOnce sync.Once
	file_api_common_proto_rawDescData = file_api_common_proto_rawDesc
)

func file_api_common_proto_rawDescGZIP() []byte {
	file_api_common_proto_rawDescOnce.Do(func() {
		file_api_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_common_proto_rawDescData)
	})
	return file_api_common_proto_rawDescData
}

var file_api_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_common_proto_goTypes = []interface{}{
	(*RequestHeader)(nil), // 0: scql.pb.RequestHeader
	(*ReportRequest)(nil), // 1: scql.pb.ReportRequest
	nil,                   // 2: scql.pb.RequestHeader.CustomHeadersEntry
	(*Status)(nil),        // 3: scql.pb.Status
	(*Tensor)(nil),        // 4: scql.pb.Tensor
}
var file_api_common_proto_depIdxs = []int32{
	2, // 0: scql.pb.RequestHeader.custom_headers:type_name -> scql.pb.RequestHeader.CustomHeadersEntry
	3, // 1: scql.pb.ReportRequest.status:type_name -> scql.pb.Status
	4, // 2: scql.pb.ReportRequest.out_columns:type_name -> scql.pb.Tensor
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_common_proto_init() }
func file_api_common_proto_init() {
	if File_api_common_proto != nil {
		return
	}
	file_api_core_proto_init()
	file_api_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHeader); i {
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
		file_api_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportRequest); i {
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
			RawDescriptor: file_api_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_common_proto_goTypes,
		DependencyIndexes: file_api_common_proto_depIdxs,
		MessageInfos:      file_api_common_proto_msgTypes,
	}.Build()
	File_api_common_proto = out.File
	file_api_common_proto_rawDesc = nil
	file_api_common_proto_goTypes = nil
	file_api_common_proto_depIdxs = nil
}