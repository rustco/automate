// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.19.0
// source: interservice/cfgmgmt/request/actions.proto

package request

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

type PolicyRevision struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RevisionId    string                 `protobuf:"bytes,1,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty" toml:"revision_id,omitempty" mapstructure:"revision_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PolicyRevision) Reset() {
	*x = PolicyRevision{}
	mi := &file_interservice_cfgmgmt_request_actions_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PolicyRevision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyRevision) ProtoMessage() {}

func (x *PolicyRevision) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_cfgmgmt_request_actions_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyRevision.ProtoReflect.Descriptor instead.
func (*PolicyRevision) Descriptor() ([]byte, []int) {
	return file_interservice_cfgmgmt_request_actions_proto_rawDescGZIP(), []int{0}
}

func (x *PolicyRevision) GetRevisionId() string {
	if x != nil {
		return x.RevisionId
	}
	return ""
}

type PolicyUpdateAction struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	PolicyName         string                 `protobuf:"bytes,1,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty" toml:"policy_name,omitempty" mapstructure:"policy_name,omitempty"`
	PolicyGroup        string                 `protobuf:"bytes,2,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty" toml:"policy_group,omitempty" mapstructure:"policy_group,omitempty"`
	PolicyRevisionId   string                 `protobuf:"bytes,3,opt,name=policy_revision_id,json=policyRevisionId,proto3" json:"policy_revision_id,omitempty" toml:"policy_revision_id,omitempty" mapstructure:"policy_revision_id,omitempty"`
	ChefServerFqdn     string                 `protobuf:"bytes,4,opt,name=chef_server_fqdn,json=chefServerFqdn,proto3" json:"chef_server_fqdn,omitempty" toml:"chef_server_fqdn,omitempty" mapstructure:"chef_server_fqdn,omitempty"`
	ChefServerOrgname  string                 `protobuf:"bytes,5,opt,name=chef_server_orgname,json=chefServerOrgname,proto3" json:"chef_server_orgname,omitempty" toml:"chef_server_orgname,omitempty" mapstructure:"chef_server_orgname,omitempty"`
	ChefServerUsername string                 `protobuf:"bytes,6,opt,name=chef_server_username,json=chefServerUsername,proto3" json:"chef_server_username,omitempty" toml:"chef_server_username,omitempty" mapstructure:"chef_server_username,omitempty"`
	PolicyfileContent  string                 `protobuf:"bytes,7,opt,name=policyfile_content,json=policyfileContent,proto3" json:"policyfile_content,omitempty" toml:"policyfile_content,omitempty" mapstructure:"policyfile_content,omitempty"`
	Cookbooks          []*PolicyCookbookLock  `protobuf:"bytes,8,rep,name=cookbooks,proto3" json:"cookbooks,omitempty" toml:"cookbooks,omitempty" mapstructure:"cookbooks,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *PolicyUpdateAction) Reset() {
	*x = PolicyUpdateAction{}
	mi := &file_interservice_cfgmgmt_request_actions_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PolicyUpdateAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyUpdateAction) ProtoMessage() {}

func (x *PolicyUpdateAction) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_cfgmgmt_request_actions_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyUpdateAction.ProtoReflect.Descriptor instead.
func (*PolicyUpdateAction) Descriptor() ([]byte, []int) {
	return file_interservice_cfgmgmt_request_actions_proto_rawDescGZIP(), []int{1}
}

func (x *PolicyUpdateAction) GetPolicyName() string {
	if x != nil {
		return x.PolicyName
	}
	return ""
}

func (x *PolicyUpdateAction) GetPolicyGroup() string {
	if x != nil {
		return x.PolicyGroup
	}
	return ""
}

func (x *PolicyUpdateAction) GetPolicyRevisionId() string {
	if x != nil {
		return x.PolicyRevisionId
	}
	return ""
}

func (x *PolicyUpdateAction) GetChefServerFqdn() string {
	if x != nil {
		return x.ChefServerFqdn
	}
	return ""
}

func (x *PolicyUpdateAction) GetChefServerOrgname() string {
	if x != nil {
		return x.ChefServerOrgname
	}
	return ""
}

func (x *PolicyUpdateAction) GetChefServerUsername() string {
	if x != nil {
		return x.ChefServerUsername
	}
	return ""
}

func (x *PolicyUpdateAction) GetPolicyfileContent() string {
	if x != nil {
		return x.PolicyfileContent
	}
	return ""
}

func (x *PolicyUpdateAction) GetCookbooks() []*PolicyCookbookLock {
	if x != nil {
		return x.Cookbooks
	}
	return nil
}

type PolicyCookbookLock struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CookbookName  string                 `protobuf:"bytes,1,opt,name=cookbook_name,json=cookbookName,proto3" json:"cookbook_name,omitempty" toml:"cookbook_name,omitempty" mapstructure:"cookbook_name,omitempty"`
	PolicyId      string                 `protobuf:"bytes,2,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty" toml:"policy_id,omitempty" mapstructure:"policy_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PolicyCookbookLock) Reset() {
	*x = PolicyCookbookLock{}
	mi := &file_interservice_cfgmgmt_request_actions_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PolicyCookbookLock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyCookbookLock) ProtoMessage() {}

func (x *PolicyCookbookLock) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_cfgmgmt_request_actions_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyCookbookLock.ProtoReflect.Descriptor instead.
func (*PolicyCookbookLock) Descriptor() ([]byte, []int) {
	return file_interservice_cfgmgmt_request_actions_proto_rawDescGZIP(), []int{2}
}

func (x *PolicyCookbookLock) GetCookbookName() string {
	if x != nil {
		return x.CookbookName
	}
	return ""
}

func (x *PolicyCookbookLock) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

var File_interservice_cfgmgmt_request_actions_proto protoreflect.FileDescriptor

var file_interservice_cfgmgmt_request_actions_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63,
	0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x31, 0x0a, 0x0e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x99, 0x03, 0x0a, 0x12, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x2c, 0x0a, 0x12, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x28,
	0x0a, 0x10, 0x63, 0x68, 0x65, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x66, 0x71,
	0x64, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x68, 0x65, 0x66, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x46, 0x71, 0x64, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x68, 0x65, 0x66,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6f, 0x72, 0x67, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x68, 0x65, 0x66, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4f, 0x72, 0x67, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x68, 0x65, 0x66,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x68, 0x65, 0x66, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x66, 0x69,
	0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x56, 0x0a, 0x09, 0x63, 0x6f, 0x6f,
	0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6f, 0x6b, 0x62, 0x6f,
	0x6f, 0x6b, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x09, 0x63, 0x6f, 0x6f, 0x6b, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x22, 0x56, 0x0a, 0x12, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6f, 0x6b, 0x62,
	0x6f, 0x6f, 0x6b, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6f, 0x6b, 0x62,
	0x6f, 0x6f, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x6f, 0x6f, 0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interservice_cfgmgmt_request_actions_proto_rawDescOnce sync.Once
	file_interservice_cfgmgmt_request_actions_proto_rawDescData = file_interservice_cfgmgmt_request_actions_proto_rawDesc
)

func file_interservice_cfgmgmt_request_actions_proto_rawDescGZIP() []byte {
	file_interservice_cfgmgmt_request_actions_proto_rawDescOnce.Do(func() {
		file_interservice_cfgmgmt_request_actions_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_cfgmgmt_request_actions_proto_rawDescData)
	})
	return file_interservice_cfgmgmt_request_actions_proto_rawDescData
}

var file_interservice_cfgmgmt_request_actions_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_interservice_cfgmgmt_request_actions_proto_goTypes = []any{
	(*PolicyRevision)(nil),     // 0: chef.automate.domain.cfgmgmt.request.PolicyRevision
	(*PolicyUpdateAction)(nil), // 1: chef.automate.domain.cfgmgmt.request.PolicyUpdateAction
	(*PolicyCookbookLock)(nil), // 2: chef.automate.domain.cfgmgmt.request.PolicyCookbookLock
}
var file_interservice_cfgmgmt_request_actions_proto_depIdxs = []int32{
	2, // 0: chef.automate.domain.cfgmgmt.request.PolicyUpdateAction.cookbooks:type_name -> chef.automate.domain.cfgmgmt.request.PolicyCookbookLock
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_interservice_cfgmgmt_request_actions_proto_init() }
func file_interservice_cfgmgmt_request_actions_proto_init() {
	if File_interservice_cfgmgmt_request_actions_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_interservice_cfgmgmt_request_actions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_interservice_cfgmgmt_request_actions_proto_goTypes,
		DependencyIndexes: file_interservice_cfgmgmt_request_actions_proto_depIdxs,
		MessageInfos:      file_interservice_cfgmgmt_request_actions_proto_msgTypes,
	}.Build()
	File_interservice_cfgmgmt_request_actions_proto = out.File
	file_interservice_cfgmgmt_request_actions_proto_rawDesc = nil
	file_interservice_cfgmgmt_request_actions_proto_goTypes = nil
	file_interservice_cfgmgmt_request_actions_proto_depIdxs = nil
}
