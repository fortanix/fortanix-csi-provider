package v1alpha1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Version of the Secrets Store CSI Driver
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *VersionRequest) Reset() {
	*x = VersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_v1alpha1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionRequest) ProtoMessage() {}

func (x *VersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provider_v1alpha1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionRequest.ProtoReflect.Descriptor instead.
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return file_provider_v1alpha1_service_proto_rawDescGZIP(), []int{0}
}

func (x *VersionRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Version of the Secrets Store CSI Driver
	Version string `protobuf:"bytes,1,opt,name=version,proto3"                             json:"version,omitempty"`
	// Name of the Secrets Store CSI Driver Provider
	RuntimeName string `protobuf:"bytes,2,opt,name=runtime_name,json=runtimeName,proto3"       json:"runtime_name,omitempty"`
	// Version of the Secrets Store CSI Driver Provider. The string must be semver-compatible.
	RuntimeVersion string `protobuf:"bytes,3,opt,name=runtime_version,json=runtimeVersion,proto3" json:"runtime_version,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_v1alpha1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_provider_v1alpha1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_provider_v1alpha1_service_proto_rawDescGZIP(), []int{1}
}

func (x *VersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *VersionResponse) GetRuntimeName() string {
	if x != nil {
		return x.RuntimeName
	}
	return ""
}

func (x *VersionResponse) GetRuntimeVersion() string {
	if x != nil {
		return x.RuntimeVersion
	}
	return ""
}

type MountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Attributes is the parameters field defined in the SecretProviderClass
	Attributes string `protobuf:"bytes,1,opt,name=attributes,proto3"                                       json:"attributes,omitempty"`
	// Secrets is the secret content referenced in nodePublishSecretRef secret data
	Secrets string `protobuf:"bytes,2,opt,name=secrets,proto3"                                          json:"secrets,omitempty"`
	// TargetPath is the path to which the volume will be published
	// TODO(tam7t): deprecate
	TargetPath string `protobuf:"bytes,3,opt,name=target_path,json=targetPath,proto3"                      json:"target_path,omitempty"`
	// Permission is the file permissions
	// TODO(tam7t): deprecate
	Permission string `protobuf:"bytes,4,opt,name=permission,proto3"                                       json:"permission,omitempty"`
	// CurrentObjectVersion is the list of objects and their versions that's
	// currently mounted in the pod
	CurrentObjectVersion []*ObjectVersion `protobuf:"bytes,5,rep,name=current_object_version,json=currentObjectVersion,proto3" json:"current_object_version,omitempty"`
}

func (x *MountRequest) Reset() {
	*x = MountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_v1alpha1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountRequest) ProtoMessage() {}

func (x *MountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provider_v1alpha1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountRequest.ProtoReflect.Descriptor instead.
func (*MountRequest) Descriptor() ([]byte, []int) {
	return file_provider_v1alpha1_service_proto_rawDescGZIP(), []int{2}
}

func (x *MountRequest) GetAttributes() string {
	if x != nil {
		return x.Attributes
	}
	return ""
}

func (x *MountRequest) GetSecrets() string {
	if x != nil {
		return x.Secrets
	}
	return ""
}

func (x *MountRequest) GetTargetPath() string {
	if x != nil {
		return x.TargetPath
	}
	return ""
}

func (x *MountRequest) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *MountRequest) GetCurrentObjectVersion() []*ObjectVersion {
	if x != nil {
		return x.CurrentObjectVersion
	}
	return nil
}

type MountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectVersion []*ObjectVersion `protobuf:"bytes,1,rep,name=object_version,json=objectVersion,proto3" json:"object_version,omitempty"`
	Error         *Error           `protobuf:"bytes,2,opt,name=error,proto3"                             json:"error,omitempty"`
	// files contains the entire mount volume filesystem.
	//
	// The total size of all files should not exceed 1MiB or syncing to
	// Kubernetes Secrets will fail. If the contents of all files exceeds
	// 4MiB then requests could fail unless MaxCallRecvMsgSize is increased.
	Files []*File `protobuf:"bytes,3,rep,name=files,proto3"                             json:"files,omitempty"`
}

func (x *MountResponse) Reset() {
	*x = MountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_v1alpha1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountResponse) ProtoMessage() {}

func (x *MountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_provider_v1alpha1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountResponse.ProtoReflect.Descriptor instead.
func (*MountResponse) Descriptor() ([]byte, []int) {
	return file_provider_v1alpha1_service_proto_rawDescGZIP(), []int{3}
}

func (x *MountResponse) GetObjectVersion() []*ObjectVersion {
	if x != nil {
		return x.ObjectVersion
	}
	return nil
}

func (x *MountResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *MountResponse) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

// File holds secret file contents and location in the mount path to write the
// file.
type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The relative path of the file within the mount.
	// May not be an absolute path.
	// May not contain the path element '..'.
	// May not start with the string '..'.
	Path string `protobuf:"bytes,1,opt,name=path,proto3"     json:"path,omitempty"`
	// The mode bits used to set permissions on this file.
	// Must be a decimal value between 0 and 511.
	Mode int32 `protobuf:"varint,2,opt,name=mode,proto3"    json:"mode,omitempty"`
	// The file contents.
	Contents []byte `protobuf:"bytes,3,opt,name=contents,proto3" json:"contents,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_v1alpha1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_provider_v1alpha1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_provider_v1alpha1_service_proto_rawDescGZIP(), []int{4}
}

func (x *File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *File) GetMode() int32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *File) GetContents() []byte {
	if x != nil {
		return x.Contents
	}
	return nil
}

type ObjectVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id is the object UID that is fetched from external secrets store
	// The Id should be unique. If multiple objects fetched from the secrets
	// store contain the same name, the provider should return a uid. This will
	// be populated in the SecretProviderClassPodStatus and sent back to the
	// provider as part of rotation reconcile
	// Example: secret/secret1, key/secret1, projects/$PROJECT_ID/secrets/secret1
	Id string `protobuf:"bytes,1,opt,name=id,proto3"      json:"id,omitempty"`
	// Version is the object version that is fetched from external secrets store
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ObjectVersion) Reset() {
	*x = ObjectVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_v1alpha1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectVersion) ProtoMessage() {}

func (x *ObjectVersion) ProtoReflect() protoreflect.Message {
	mi := &file_provider_v1alpha1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectVersion.ProtoReflect.Descriptor instead.
func (*ObjectVersion) Descriptor() ([]byte, []int) {
	return file_provider_v1alpha1_service_proto_rawDescGZIP(), []int{5}
}

func (x *ObjectVersion) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ObjectVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Code is the error code that the provider can return which will be used for publishing metrics
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_v1alpha1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_provider_v1alpha1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_provider_v1alpha1_service_proto_rawDescGZIP(), []int{6}
}

func (x *Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_provider_v1alpha1_service_proto protoreflect.FileDescriptor

var file_provider_v1alpha1_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x2a, 0x0a, 0x0e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x77, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0xd8, 0x01, 0x0a, 0x0c, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x16,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x9c, 0x01, 0x0a, 0x0d,
	0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x0e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0d,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x4a, 0x0a, 0x04, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x39, 0x0a, 0x0d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x1b, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x91,
	0x01, 0x0a, 0x11, 0x43, 0x53, 0x49, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x05, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x16, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_provider_v1alpha1_service_proto_rawDescOnce sync.Once
	file_provider_v1alpha1_service_proto_rawDescData = file_provider_v1alpha1_service_proto_rawDesc
)

func file_provider_v1alpha1_service_proto_rawDescGZIP() []byte {
	file_provider_v1alpha1_service_proto_rawDescOnce.Do(func() {
		file_provider_v1alpha1_service_proto_rawDescData = protoimpl.X.CompressGZIP(
			file_provider_v1alpha1_service_proto_rawDescData,
		)
	})
	return file_provider_v1alpha1_service_proto_rawDescData
}

var (
	file_provider_v1alpha1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
	file_provider_v1alpha1_service_proto_goTypes  = []interface{}{
		(*VersionRequest)(nil),  // 0: v1alpha1.VersionRequest
		(*VersionResponse)(nil), // 1: v1alpha1.VersionResponse
		(*MountRequest)(nil),    // 2: v1alpha1.MountRequest
		(*MountResponse)(nil),   // 3: v1alpha1.MountResponse
		(*File)(nil),            // 4: v1alpha1.File
		(*ObjectVersion)(nil),   // 5: v1alpha1.ObjectVersion
		(*Error)(nil),           // 6: v1alpha1.Error
	}
)
var file_provider_v1alpha1_service_proto_depIdxs = []int32{
	5, // 0: v1alpha1.MountRequest.current_object_version:type_name -> v1alpha1.ObjectVersion
	5, // 1: v1alpha1.MountResponse.object_version:type_name -> v1alpha1.ObjectVersion
	6, // 2: v1alpha1.MountResponse.error:type_name -> v1alpha1.Error
	4, // 3: v1alpha1.MountResponse.files:type_name -> v1alpha1.File
	0, // 4: v1alpha1.CSIDriverProvider.Version:input_type -> v1alpha1.VersionRequest
	2, // 5: v1alpha1.CSIDriverProvider.Mount:input_type -> v1alpha1.MountRequest
	1, // 6: v1alpha1.CSIDriverProvider.Version:output_type -> v1alpha1.VersionResponse
	3, // 7: v1alpha1.CSIDriverProvider.Mount:output_type -> v1alpha1.MountResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_provider_v1alpha1_service_proto_init() }
func file_provider_v1alpha1_service_proto_init() {
	if File_provider_v1alpha1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_provider_v1alpha1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionRequest); i {
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
		file_provider_v1alpha1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
		file_provider_v1alpha1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountRequest); i {
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
		file_provider_v1alpha1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountResponse); i {
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
		file_provider_v1alpha1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_provider_v1alpha1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectVersion); i {
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
		file_provider_v1alpha1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_provider_v1alpha1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_provider_v1alpha1_service_proto_goTypes,
		DependencyIndexes: file_provider_v1alpha1_service_proto_depIdxs,
		MessageInfos:      file_provider_v1alpha1_service_proto_msgTypes,
	}.Build()
	File_provider_v1alpha1_service_proto = out.File
	file_provider_v1alpha1_service_proto_rawDesc = nil
	file_provider_v1alpha1_service_proto_goTypes = nil
	file_provider_v1alpha1_service_proto_depIdxs = nil
}
