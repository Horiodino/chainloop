//
// Copyright 2023 The Chainloop Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: attestation/v1/crafting_state.proto

package v1

import (
	v1 "github.com/chainloop-dev/chainloop/app/controlplane/api/workflowcontract/v1"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Attestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InitializedAt *timestamppb.Timestamp           `protobuf:"bytes,1,opt,name=initialized_at,json=initializedAt,proto3" json:"initialized_at,omitempty"`
	FinishedAt    *timestamppb.Timestamp           `protobuf:"bytes,2,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	Workflow      *WorkflowMetadata                `protobuf:"bytes,3,opt,name=workflow,proto3" json:"workflow,omitempty"`
	Materials     map[string]*Attestation_Material `protobuf:"bytes,4,rep,name=materials,proto3" json:"materials,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// List of env variables
	EnvVars    map[string]string                   `protobuf:"bytes,6,rep,name=env_vars,json=envVars,proto3" json:"env_vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RunnerUrl  string                              `protobuf:"bytes,7,opt,name=runner_url,json=runnerUrl,proto3" json:"runner_url,omitempty"`
	RunnerType v1.CraftingSchema_Runner_RunnerType `protobuf:"varint,8,opt,name=runner_type,json=runnerType,proto3,enum=workflowcontract.v1.CraftingSchema_Runner_RunnerType" json:"runner_type,omitempty"`
}

func (x *Attestation) Reset() {
	*x = Attestation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attestation_v1_crafting_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attestation) ProtoMessage() {}

func (x *Attestation) ProtoReflect() protoreflect.Message {
	mi := &file_attestation_v1_crafting_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attestation.ProtoReflect.Descriptor instead.
func (*Attestation) Descriptor() ([]byte, []int) {
	return file_attestation_v1_crafting_state_proto_rawDescGZIP(), []int{0}
}

func (x *Attestation) GetInitializedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.InitializedAt
	}
	return nil
}

func (x *Attestation) GetFinishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

func (x *Attestation) GetWorkflow() *WorkflowMetadata {
	if x != nil {
		return x.Workflow
	}
	return nil
}

func (x *Attestation) GetMaterials() map[string]*Attestation_Material {
	if x != nil {
		return x.Materials
	}
	return nil
}

func (x *Attestation) GetEnvVars() map[string]string {
	if x != nil {
		return x.EnvVars
	}
	return nil
}

func (x *Attestation) GetRunnerUrl() string {
	if x != nil {
		return x.RunnerUrl
	}
	return ""
}

func (x *Attestation) GetRunnerType() v1.CraftingSchema_Runner_RunnerType {
	if x != nil {
		return x.RunnerType
	}
	return v1.CraftingSchema_Runner_RunnerType(0)
}

// Intermediate information that will get stored in the system while the run is being executed
type CraftingState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputSchema *v1.CraftingSchema `protobuf:"bytes,1,opt,name=input_schema,json=inputSchema,proto3" json:"input_schema,omitempty"`
	Attestation *Attestation       `protobuf:"bytes,2,opt,name=attestation,proto3" json:"attestation,omitempty"`
	DryRun      bool               `protobuf:"varint,3,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
}

func (x *CraftingState) Reset() {
	*x = CraftingState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attestation_v1_crafting_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CraftingState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CraftingState) ProtoMessage() {}

func (x *CraftingState) ProtoReflect() protoreflect.Message {
	mi := &file_attestation_v1_crafting_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CraftingState.ProtoReflect.Descriptor instead.
func (*CraftingState) Descriptor() ([]byte, []int) {
	return file_attestation_v1_crafting_state_proto_rawDescGZIP(), []int{1}
}

func (x *CraftingState) GetInputSchema() *v1.CraftingSchema {
	if x != nil {
		return x.InputSchema
	}
	return nil
}

func (x *CraftingState) GetAttestation() *Attestation {
	if x != nil {
		return x.Attestation
	}
	return nil
}

func (x *CraftingState) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

type WorkflowMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Project        string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	Team           string `protobuf:"bytes,3,opt,name=team,proto3" json:"team,omitempty"`
	WorkflowId     string `protobuf:"bytes,5,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	WorkflowRunId  string `protobuf:"bytes,6,opt,name=workflow_run_id,json=workflowRunId,proto3" json:"workflow_run_id,omitempty"` // Not required since we might be doing a dry-run
	SchemaRevision string `protobuf:"bytes,7,opt,name=schema_revision,json=schemaRevision,proto3" json:"schema_revision,omitempty"`
}

func (x *WorkflowMetadata) Reset() {
	*x = WorkflowMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attestation_v1_crafting_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowMetadata) ProtoMessage() {}

func (x *WorkflowMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_attestation_v1_crafting_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowMetadata.ProtoReflect.Descriptor instead.
func (*WorkflowMetadata) Descriptor() ([]byte, []int) {
	return file_attestation_v1_crafting_state_proto_rawDescGZIP(), []int{2}
}

func (x *WorkflowMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WorkflowMetadata) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *WorkflowMetadata) GetTeam() string {
	if x != nil {
		return x.Team
	}
	return ""
}

func (x *WorkflowMetadata) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *WorkflowMetadata) GetWorkflowRunId() string {
	if x != nil {
		return x.WorkflowRunId
	}
	return ""
}

func (x *WorkflowMetadata) GetSchemaRevision() string {
	if x != nil {
		return x.SchemaRevision
	}
	return ""
}

type Attestation_Material struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to M:
	//
	//	*Attestation_Material_String_
	//	*Attestation_Material_ContainerImage_
	//	*Attestation_Material_Artifact_
	M            isAttestation_Material_M                `protobuf_oneof:"m"`
	AddedAt      *timestamppb.Timestamp                  `protobuf:"bytes,5,opt,name=added_at,json=addedAt,proto3" json:"added_at,omitempty"`
	MaterialType v1.CraftingSchema_Material_MaterialType `protobuf:"varint,6,opt,name=material_type,json=materialType,proto3,enum=workflowcontract.v1.CraftingSchema_Material_MaterialType" json:"material_type,omitempty"`
}

func (x *Attestation_Material) Reset() {
	*x = Attestation_Material{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attestation_v1_crafting_state_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attestation_Material) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attestation_Material) ProtoMessage() {}

func (x *Attestation_Material) ProtoReflect() protoreflect.Message {
	mi := &file_attestation_v1_crafting_state_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attestation_Material.ProtoReflect.Descriptor instead.
func (*Attestation_Material) Descriptor() ([]byte, []int) {
	return file_attestation_v1_crafting_state_proto_rawDescGZIP(), []int{0, 1}
}

func (m *Attestation_Material) GetM() isAttestation_Material_M {
	if m != nil {
		return m.M
	}
	return nil
}

func (x *Attestation_Material) GetString_() *Attestation_Material_KeyVal {
	if x, ok := x.GetM().(*Attestation_Material_String_); ok {
		return x.String_
	}
	return nil
}

func (x *Attestation_Material) GetContainerImage() *Attestation_Material_ContainerImage {
	if x, ok := x.GetM().(*Attestation_Material_ContainerImage_); ok {
		return x.ContainerImage
	}
	return nil
}

func (x *Attestation_Material) GetArtifact() *Attestation_Material_Artifact {
	if x, ok := x.GetM().(*Attestation_Material_Artifact_); ok {
		return x.Artifact
	}
	return nil
}

func (x *Attestation_Material) GetAddedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.AddedAt
	}
	return nil
}

func (x *Attestation_Material) GetMaterialType() v1.CraftingSchema_Material_MaterialType {
	if x != nil {
		return x.MaterialType
	}
	return v1.CraftingSchema_Material_MaterialType(0)
}

type isAttestation_Material_M interface {
	isAttestation_Material_M()
}

type Attestation_Material_String_ struct {
	String_ *Attestation_Material_KeyVal `protobuf:"bytes,1,opt,name=string,proto3,oneof"`
}

type Attestation_Material_ContainerImage_ struct {
	ContainerImage *Attestation_Material_ContainerImage `protobuf:"bytes,2,opt,name=container_image,json=containerImage,proto3,oneof"`
}

type Attestation_Material_Artifact_ struct {
	Artifact *Attestation_Material_Artifact `protobuf:"bytes,3,opt,name=artifact,proto3,oneof"`
}

func (*Attestation_Material_String_) isAttestation_Material_M() {}

func (*Attestation_Material_ContainerImage_) isAttestation_Material_M() {}

func (*Attestation_Material_Artifact_) isAttestation_Material_M() {}

type Attestation_Material_KeyVal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Attestation_Material_KeyVal) Reset() {
	*x = Attestation_Material_KeyVal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attestation_v1_crafting_state_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attestation_Material_KeyVal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attestation_Material_KeyVal) ProtoMessage() {}

func (x *Attestation_Material_KeyVal) ProtoReflect() protoreflect.Message {
	mi := &file_attestation_v1_crafting_state_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attestation_Material_KeyVal.ProtoReflect.Descriptor instead.
func (*Attestation_Material_KeyVal) Descriptor() ([]byte, []int) {
	return file_attestation_v1_crafting_state_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *Attestation_Material_KeyVal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Attestation_Material_KeyVal) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Attestation_Material_ContainerImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Digest    string `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	IsSubject bool   `protobuf:"varint,4,opt,name=is_subject,json=isSubject,proto3" json:"is_subject,omitempty"`
}

func (x *Attestation_Material_ContainerImage) Reset() {
	*x = Attestation_Material_ContainerImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attestation_v1_crafting_state_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attestation_Material_ContainerImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attestation_Material_ContainerImage) ProtoMessage() {}

func (x *Attestation_Material_ContainerImage) ProtoReflect() protoreflect.Message {
	mi := &file_attestation_v1_crafting_state_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attestation_Material_ContainerImage.ProtoReflect.Descriptor instead.
func (*Attestation_Material_ContainerImage) Descriptor() ([]byte, []int) {
	return file_attestation_v1_crafting_state_proto_rawDescGZIP(), []int{0, 1, 1}
}

func (x *Attestation_Material_ContainerImage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Attestation_Material_ContainerImage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Attestation_Material_ContainerImage) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *Attestation_Material_ContainerImage) GetIsSubject() bool {
	if x != nil {
		return x.IsSubject
	}
	return false
}

type Attestation_Material_Artifact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the artifact
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// filename, use for record purposes
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// the digest is enough to retrieve the artifact since it's stored in a CAS
	// which also has annotated the fileName
	Digest    string `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	IsSubject bool   `protobuf:"varint,4,opt,name=is_subject,json=isSubject,proto3" json:"is_subject,omitempty"`
}

func (x *Attestation_Material_Artifact) Reset() {
	*x = Attestation_Material_Artifact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attestation_v1_crafting_state_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attestation_Material_Artifact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attestation_Material_Artifact) ProtoMessage() {}

func (x *Attestation_Material_Artifact) ProtoReflect() protoreflect.Message {
	mi := &file_attestation_v1_crafting_state_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attestation_Material_Artifact.ProtoReflect.Descriptor instead.
func (*Attestation_Material_Artifact) Descriptor() ([]byte, []int) {
	return file_attestation_v1_crafting_state_proto_rawDescGZIP(), []int{0, 1, 2}
}

func (x *Attestation_Material_Artifact) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Attestation_Material_Artifact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Attestation_Material_Artifact) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *Attestation_Material_Artifact) GetIsSubject() bool {
	if x != nil {
		return x.IsSubject
	}
	return false
}

var File_attestation_v1_crafting_state_proto protoreflect.FileDescriptor

var file_attestation_v1_crafting_state_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x72, 0x61, 0x66, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x29, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x61, 0x66, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x0a, 0x0a, 0x0b, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x0e, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0xb2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x46, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x48, 0x0a, 0x09,
	0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x43, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61,
	0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x56, 0x0a, 0x0b, 0x72, 0x75,
	0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x35, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x61, 0x66, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x1a, 0x62, 0x0a, 0x0e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xe8, 0x05, 0x0a, 0x08, 0x4d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c,
	0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x5e, 0x0a, 0x0f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x48, 0x00, 0x52, 0x08, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x61, 0x64, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12, 0x5e,
	0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x61, 0x66,
	0x74, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0c, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x40,
	0x0a, 0x06, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x86, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73,
	0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x80, 0x01, 0x0a, 0x08, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x06,
	0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x03, 0x0a, 0x01,
	0x6d, 0x1a, 0x3a, 0x0a, 0x0c, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xaf, 0x01,
	0x0a, 0x0d, 0x43, 0x72, 0x61, 0x66, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x46, 0x0a, 0x0c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x61, 0x66,
	0x74, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x0b, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x3d, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x72, 0x79, 0x5f, 0x72, 0x75,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x22,
	0xe1, 0x01, 0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x12,
	0x28, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x75, 0x6e, 0x49,
	0x64, 0x12, 0x30, 0x0a, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x72, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x6f, 0x6f, 0x70, 0x2d, 0x64, 0x65, 0x76, 0x2f,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x6f, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6c,
	0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_attestation_v1_crafting_state_proto_rawDescOnce sync.Once
	file_attestation_v1_crafting_state_proto_rawDescData = file_attestation_v1_crafting_state_proto_rawDesc
)

func file_attestation_v1_crafting_state_proto_rawDescGZIP() []byte {
	file_attestation_v1_crafting_state_proto_rawDescOnce.Do(func() {
		file_attestation_v1_crafting_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_attestation_v1_crafting_state_proto_rawDescData)
	})
	return file_attestation_v1_crafting_state_proto_rawDescData
}

var file_attestation_v1_crafting_state_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_attestation_v1_crafting_state_proto_goTypes = []interface{}{
	(*Attestation)(nil),                 // 0: attestation.v1.Attestation
	(*CraftingState)(nil),               // 1: attestation.v1.CraftingState
	(*WorkflowMetadata)(nil),            // 2: attestation.v1.WorkflowMetadata
	nil,                                 // 3: attestation.v1.Attestation.MaterialsEntry
	(*Attestation_Material)(nil),        // 4: attestation.v1.Attestation.Material
	nil,                                 // 5: attestation.v1.Attestation.EnvVarsEntry
	(*Attestation_Material_KeyVal)(nil), // 6: attestation.v1.Attestation.Material.KeyVal
	(*Attestation_Material_ContainerImage)(nil),  // 7: attestation.v1.Attestation.Material.ContainerImage
	(*Attestation_Material_Artifact)(nil),        // 8: attestation.v1.Attestation.Material.Artifact
	(*timestamppb.Timestamp)(nil),                // 9: google.protobuf.Timestamp
	(v1.CraftingSchema_Runner_RunnerType)(0),     // 10: workflowcontract.v1.CraftingSchema.Runner.RunnerType
	(*v1.CraftingSchema)(nil),                    // 11: workflowcontract.v1.CraftingSchema
	(v1.CraftingSchema_Material_MaterialType)(0), // 12: workflowcontract.v1.CraftingSchema.Material.MaterialType
}
var file_attestation_v1_crafting_state_proto_depIdxs = []int32{
	9,  // 0: attestation.v1.Attestation.initialized_at:type_name -> google.protobuf.Timestamp
	9,  // 1: attestation.v1.Attestation.finished_at:type_name -> google.protobuf.Timestamp
	2,  // 2: attestation.v1.Attestation.workflow:type_name -> attestation.v1.WorkflowMetadata
	3,  // 3: attestation.v1.Attestation.materials:type_name -> attestation.v1.Attestation.MaterialsEntry
	5,  // 4: attestation.v1.Attestation.env_vars:type_name -> attestation.v1.Attestation.EnvVarsEntry
	10, // 5: attestation.v1.Attestation.runner_type:type_name -> workflowcontract.v1.CraftingSchema.Runner.RunnerType
	11, // 6: attestation.v1.CraftingState.input_schema:type_name -> workflowcontract.v1.CraftingSchema
	0,  // 7: attestation.v1.CraftingState.attestation:type_name -> attestation.v1.Attestation
	4,  // 8: attestation.v1.Attestation.MaterialsEntry.value:type_name -> attestation.v1.Attestation.Material
	6,  // 9: attestation.v1.Attestation.Material.string:type_name -> attestation.v1.Attestation.Material.KeyVal
	7,  // 10: attestation.v1.Attestation.Material.container_image:type_name -> attestation.v1.Attestation.Material.ContainerImage
	8,  // 11: attestation.v1.Attestation.Material.artifact:type_name -> attestation.v1.Attestation.Material.Artifact
	9,  // 12: attestation.v1.Attestation.Material.added_at:type_name -> google.protobuf.Timestamp
	12, // 13: attestation.v1.Attestation.Material.material_type:type_name -> workflowcontract.v1.CraftingSchema.Material.MaterialType
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_attestation_v1_crafting_state_proto_init() }
func file_attestation_v1_crafting_state_proto_init() {
	if File_attestation_v1_crafting_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_attestation_v1_crafting_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attestation); i {
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
		file_attestation_v1_crafting_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CraftingState); i {
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
		file_attestation_v1_crafting_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowMetadata); i {
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
		file_attestation_v1_crafting_state_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attestation_Material); i {
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
		file_attestation_v1_crafting_state_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attestation_Material_KeyVal); i {
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
		file_attestation_v1_crafting_state_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attestation_Material_ContainerImage); i {
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
		file_attestation_v1_crafting_state_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attestation_Material_Artifact); i {
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
	file_attestation_v1_crafting_state_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Attestation_Material_String_)(nil),
		(*Attestation_Material_ContainerImage_)(nil),
		(*Attestation_Material_Artifact_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_attestation_v1_crafting_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_attestation_v1_crafting_state_proto_goTypes,
		DependencyIndexes: file_attestation_v1_crafting_state_proto_depIdxs,
		MessageInfos:      file_attestation_v1_crafting_state_proto_msgTypes,
	}.Build()
	File_attestation_v1_crafting_state_proto = out.File
	file_attestation_v1_crafting_state_proto_rawDesc = nil
	file_attestation_v1_crafting_state_proto_goTypes = nil
	file_attestation_v1_crafting_state_proto_depIdxs = nil
}
