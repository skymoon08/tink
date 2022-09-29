// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: third_party/tink/proto/hpke.proto

package hpke_proto

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

type HpkeKem int32

const (
	HpkeKem_KEM_UNKNOWN              HpkeKem = 0
	HpkeKem_DHKEM_X25519_HKDF_SHA256 HpkeKem = 1
	HpkeKem_DHKEM_P256_HKDF_SHA256   HpkeKem = 2
	HpkeKem_DHKEM_P384_HKDF_SHA384   HpkeKem = 3
	HpkeKem_DHKEM_P521_HKDF_SHA512   HpkeKem = 4
)

// Enum value maps for HpkeKem.
var (
	HpkeKem_name = map[int32]string{
		0: "KEM_UNKNOWN",
		1: "DHKEM_X25519_HKDF_SHA256",
		2: "DHKEM_P256_HKDF_SHA256",
		3: "DHKEM_P384_HKDF_SHA384",
		4: "DHKEM_P521_HKDF_SHA512",
	}
	HpkeKem_value = map[string]int32{
		"KEM_UNKNOWN":              0,
		"DHKEM_X25519_HKDF_SHA256": 1,
		"DHKEM_P256_HKDF_SHA256":   2,
		"DHKEM_P384_HKDF_SHA384":   3,
		"DHKEM_P521_HKDF_SHA512":   4,
	}
)

func (x HpkeKem) Enum() *HpkeKem {
	p := new(HpkeKem)
	*p = x
	return p
}

func (x HpkeKem) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HpkeKem) Descriptor() protoreflect.EnumDescriptor {
	return file_third_party_tink_proto_hpke_proto_enumTypes[0].Descriptor()
}

func (HpkeKem) Type() protoreflect.EnumType {
	return &file_third_party_tink_proto_hpke_proto_enumTypes[0]
}

func (x HpkeKem) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HpkeKem.Descriptor instead.
func (HpkeKem) EnumDescriptor() ([]byte, []int) {
	return file_third_party_tink_proto_hpke_proto_rawDescGZIP(), []int{0}
}

type HpkeKdf int32

const (
	HpkeKdf_KDF_UNKNOWN HpkeKdf = 0
	HpkeKdf_HKDF_SHA256 HpkeKdf = 1
	HpkeKdf_HKDF_SHA384 HpkeKdf = 2
	HpkeKdf_HKDF_SHA512 HpkeKdf = 3
)

// Enum value maps for HpkeKdf.
var (
	HpkeKdf_name = map[int32]string{
		0: "KDF_UNKNOWN",
		1: "HKDF_SHA256",
		2: "HKDF_SHA384",
		3: "HKDF_SHA512",
	}
	HpkeKdf_value = map[string]int32{
		"KDF_UNKNOWN": 0,
		"HKDF_SHA256": 1,
		"HKDF_SHA384": 2,
		"HKDF_SHA512": 3,
	}
)

func (x HpkeKdf) Enum() *HpkeKdf {
	p := new(HpkeKdf)
	*p = x
	return p
}

func (x HpkeKdf) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HpkeKdf) Descriptor() protoreflect.EnumDescriptor {
	return file_third_party_tink_proto_hpke_proto_enumTypes[1].Descriptor()
}

func (HpkeKdf) Type() protoreflect.EnumType {
	return &file_third_party_tink_proto_hpke_proto_enumTypes[1]
}

func (x HpkeKdf) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HpkeKdf.Descriptor instead.
func (HpkeKdf) EnumDescriptor() ([]byte, []int) {
	return file_third_party_tink_proto_hpke_proto_rawDescGZIP(), []int{1}
}

type HpkeAead int32

const (
	HpkeAead_AEAD_UNKNOWN      HpkeAead = 0
	HpkeAead_AES_128_GCM       HpkeAead = 1
	HpkeAead_AES_256_GCM       HpkeAead = 2
	HpkeAead_CHACHA20_POLY1305 HpkeAead = 3
)

// Enum value maps for HpkeAead.
var (
	HpkeAead_name = map[int32]string{
		0: "AEAD_UNKNOWN",
		1: "AES_128_GCM",
		2: "AES_256_GCM",
		3: "CHACHA20_POLY1305",
	}
	HpkeAead_value = map[string]int32{
		"AEAD_UNKNOWN":      0,
		"AES_128_GCM":       1,
		"AES_256_GCM":       2,
		"CHACHA20_POLY1305": 3,
	}
)

func (x HpkeAead) Enum() *HpkeAead {
	p := new(HpkeAead)
	*p = x
	return p
}

func (x HpkeAead) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HpkeAead) Descriptor() protoreflect.EnumDescriptor {
	return file_third_party_tink_proto_hpke_proto_enumTypes[2].Descriptor()
}

func (HpkeAead) Type() protoreflect.EnumType {
	return &file_third_party_tink_proto_hpke_proto_enumTypes[2]
}

func (x HpkeAead) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HpkeAead.Descriptor instead.
func (HpkeAead) EnumDescriptor() ([]byte, []int) {
	return file_third_party_tink_proto_hpke_proto_rawDescGZIP(), []int{2}
}

type HpkeParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kem  HpkeKem  `protobuf:"varint,1,opt,name=kem,proto3,enum=google.crypto.tink.HpkeKem" json:"kem,omitempty"`
	Kdf  HpkeKdf  `protobuf:"varint,2,opt,name=kdf,proto3,enum=google.crypto.tink.HpkeKdf" json:"kdf,omitempty"`
	Aead HpkeAead `protobuf:"varint,3,opt,name=aead,proto3,enum=google.crypto.tink.HpkeAead" json:"aead,omitempty"`
}

func (x *HpkeParams) Reset() {
	*x = HpkeParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_hpke_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HpkeParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HpkeParams) ProtoMessage() {}

func (x *HpkeParams) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_hpke_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HpkeParams.ProtoReflect.Descriptor instead.
func (*HpkeParams) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_hpke_proto_rawDescGZIP(), []int{0}
}

func (x *HpkeParams) GetKem() HpkeKem {
	if x != nil {
		return x.Kem
	}
	return HpkeKem_KEM_UNKNOWN
}

func (x *HpkeParams) GetKdf() HpkeKdf {
	if x != nil {
		return x.Kdf
	}
	return HpkeKdf_KDF_UNKNOWN
}

func (x *HpkeParams) GetAead() HpkeAead {
	if x != nil {
		return x.Aead
	}
	return HpkeAead_AEAD_UNKNOWN
}

type HpkePublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint32      `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Params  *HpkeParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	// KEM-encoding of public key (i.e., SerializePublicKey() ) as described in
	// https://www.rfc-editor.org/rfc/rfc9180.html#name-cryptographic-dependencies.
	PublicKey []byte `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *HpkePublicKey) Reset() {
	*x = HpkePublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_hpke_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HpkePublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HpkePublicKey) ProtoMessage() {}

func (x *HpkePublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_hpke_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HpkePublicKey.ProtoReflect.Descriptor instead.
func (*HpkePublicKey) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_hpke_proto_rawDescGZIP(), []int{1}
}

func (x *HpkePublicKey) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *HpkePublicKey) GetParams() *HpkeParams {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *HpkePublicKey) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

type HpkePrivateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   uint32         `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	PublicKey *HpkePublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// KEM-encoding of private key (i.e., SerializePrivateKey() ) as described in
	// https://www.rfc-editor.org/rfc/rfc9180.html#name-cryptographic-dependencies.
	PrivateKey []byte `protobuf:"bytes,3,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
}

func (x *HpkePrivateKey) Reset() {
	*x = HpkePrivateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_hpke_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HpkePrivateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HpkePrivateKey) ProtoMessage() {}

func (x *HpkePrivateKey) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_hpke_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HpkePrivateKey.ProtoReflect.Descriptor instead.
func (*HpkePrivateKey) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_hpke_proto_rawDescGZIP(), []int{2}
}

func (x *HpkePrivateKey) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *HpkePrivateKey) GetPublicKey() *HpkePublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *HpkePrivateKey) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

type HpkeKeyFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params *HpkeParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *HpkeKeyFormat) Reset() {
	*x = HpkeKeyFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_hpke_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HpkeKeyFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HpkeKeyFormat) ProtoMessage() {}

func (x *HpkeKeyFormat) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_hpke_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HpkeKeyFormat.ProtoReflect.Descriptor instead.
func (*HpkeKeyFormat) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_hpke_proto_rawDescGZIP(), []int{3}
}

func (x *HpkeKeyFormat) GetParams() *HpkeParams {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_third_party_tink_proto_hpke_proto protoreflect.FileDescriptor

var file_third_party_tink_proto_hpke_proto_rawDesc = []byte{
	0x0a, 0x21, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x69,
	0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x70, 0x6b, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x22, 0x9c, 0x01, 0x0a, 0x0a, 0x48, 0x70, 0x6b, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2d, 0x0a, 0x03, 0x6b, 0x65, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x48, 0x70, 0x6b, 0x65, 0x4b, 0x65, 0x6d,
	0x52, 0x03, 0x6b, 0x65, 0x6d, 0x12, 0x2d, 0x0a, 0x03, 0x6b, 0x64, 0x66, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x48, 0x70, 0x6b, 0x65, 0x4b, 0x64, 0x66, 0x52,
	0x03, 0x6b, 0x64, 0x66, 0x12, 0x30, 0x0a, 0x04, 0x61, 0x65, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x48, 0x70, 0x6b, 0x65, 0x41, 0x65, 0x61, 0x64,
	0x52, 0x04, 0x61, 0x65, 0x61, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x0d, 0x48, 0x70, 0x6b, 0x65, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x48, 0x70, 0x6b, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x8d, 0x01, 0x0a, 0x0e, 0x48, 0x70,
	0x6b, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e,
	0x48, 0x70, 0x6b, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x47, 0x0a, 0x0d, 0x48, 0x70, 0x6b,
	0x65, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e,
	0x48, 0x70, 0x6b, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x2a, 0x8c, 0x01, 0x0a, 0x07, 0x48, 0x70, 0x6b, 0x65, 0x4b, 0x65, 0x6d, 0x12, 0x0f,
	0x0a, 0x0b, 0x4b, 0x45, 0x4d, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x1c, 0x0a, 0x18, 0x44, 0x48, 0x4b, 0x45, 0x4d, 0x5f, 0x58, 0x32, 0x35, 0x35, 0x31, 0x39, 0x5f,
	0x48, 0x4b, 0x44, 0x46, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x01, 0x12, 0x1a, 0x0a,
	0x16, 0x44, 0x48, 0x4b, 0x45, 0x4d, 0x5f, 0x50, 0x32, 0x35, 0x36, 0x5f, 0x48, 0x4b, 0x44, 0x46,
	0x5f, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x48, 0x4b,
	0x45, 0x4d, 0x5f, 0x50, 0x33, 0x38, 0x34, 0x5f, 0x48, 0x4b, 0x44, 0x46, 0x5f, 0x53, 0x48, 0x41,
	0x33, 0x38, 0x34, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x48, 0x4b, 0x45, 0x4d, 0x5f, 0x50,
	0x35, 0x32, 0x31, 0x5f, 0x48, 0x4b, 0x44, 0x46, 0x5f, 0x53, 0x48, 0x41, 0x35, 0x31, 0x32, 0x10,
	0x04, 0x2a, 0x4d, 0x0a, 0x07, 0x48, 0x70, 0x6b, 0x65, 0x4b, 0x64, 0x66, 0x12, 0x0f, 0x0a, 0x0b,
	0x4b, 0x44, 0x46, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x48, 0x4b, 0x44, 0x46, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x48, 0x4b, 0x44, 0x46, 0x5f, 0x53, 0x48, 0x41, 0x33, 0x38, 0x34, 0x10, 0x02, 0x12,
	0x0f, 0x0a, 0x0b, 0x48, 0x4b, 0x44, 0x46, 0x5f, 0x53, 0x48, 0x41, 0x35, 0x31, 0x32, 0x10, 0x03,
	0x2a, 0x55, 0x0a, 0x08, 0x48, 0x70, 0x6b, 0x65, 0x41, 0x65, 0x61, 0x64, 0x12, 0x10, 0x0a, 0x0c,
	0x41, 0x45, 0x41, 0x44, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f,
	0x0a, 0x0b, 0x41, 0x45, 0x53, 0x5f, 0x31, 0x32, 0x38, 0x5f, 0x47, 0x43, 0x4d, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x41, 0x45, 0x53, 0x5f, 0x32, 0x35, 0x36, 0x5f, 0x47, 0x43, 0x4d, 0x10, 0x02,
	0x12, 0x15, 0x0a, 0x11, 0x43, 0x48, 0x41, 0x43, 0x48, 0x41, 0x32, 0x30, 0x5f, 0x50, 0x4f, 0x4c,
	0x59, 0x31, 0x33, 0x30, 0x35, 0x10, 0x03, 0x42, 0x49, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x69, 0x6e,
	0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x70, 0x6b, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_third_party_tink_proto_hpke_proto_rawDescOnce sync.Once
	file_third_party_tink_proto_hpke_proto_rawDescData = file_third_party_tink_proto_hpke_proto_rawDesc
)

func file_third_party_tink_proto_hpke_proto_rawDescGZIP() []byte {
	file_third_party_tink_proto_hpke_proto_rawDescOnce.Do(func() {
		file_third_party_tink_proto_hpke_proto_rawDescData = protoimpl.X.CompressGZIP(file_third_party_tink_proto_hpke_proto_rawDescData)
	})
	return file_third_party_tink_proto_hpke_proto_rawDescData
}

var file_third_party_tink_proto_hpke_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_third_party_tink_proto_hpke_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_third_party_tink_proto_hpke_proto_goTypes = []interface{}{
	(HpkeKem)(0),           // 0: google.crypto.tink.HpkeKem
	(HpkeKdf)(0),           // 1: google.crypto.tink.HpkeKdf
	(HpkeAead)(0),          // 2: google.crypto.tink.HpkeAead
	(*HpkeParams)(nil),     // 3: google.crypto.tink.HpkeParams
	(*HpkePublicKey)(nil),  // 4: google.crypto.tink.HpkePublicKey
	(*HpkePrivateKey)(nil), // 5: google.crypto.tink.HpkePrivateKey
	(*HpkeKeyFormat)(nil),  // 6: google.crypto.tink.HpkeKeyFormat
}
var file_third_party_tink_proto_hpke_proto_depIdxs = []int32{
	0, // 0: google.crypto.tink.HpkeParams.kem:type_name -> google.crypto.tink.HpkeKem
	1, // 1: google.crypto.tink.HpkeParams.kdf:type_name -> google.crypto.tink.HpkeKdf
	2, // 2: google.crypto.tink.HpkeParams.aead:type_name -> google.crypto.tink.HpkeAead
	3, // 3: google.crypto.tink.HpkePublicKey.params:type_name -> google.crypto.tink.HpkeParams
	4, // 4: google.crypto.tink.HpkePrivateKey.public_key:type_name -> google.crypto.tink.HpkePublicKey
	3, // 5: google.crypto.tink.HpkeKeyFormat.params:type_name -> google.crypto.tink.HpkeParams
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_third_party_tink_proto_hpke_proto_init() }
func file_third_party_tink_proto_hpke_proto_init() {
	if File_third_party_tink_proto_hpke_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_third_party_tink_proto_hpke_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HpkeParams); i {
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
		file_third_party_tink_proto_hpke_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HpkePublicKey); i {
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
		file_third_party_tink_proto_hpke_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HpkePrivateKey); i {
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
		file_third_party_tink_proto_hpke_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HpkeKeyFormat); i {
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
			RawDescriptor: file_third_party_tink_proto_hpke_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_third_party_tink_proto_hpke_proto_goTypes,
		DependencyIndexes: file_third_party_tink_proto_hpke_proto_depIdxs,
		EnumInfos:         file_third_party_tink_proto_hpke_proto_enumTypes,
		MessageInfos:      file_third_party_tink_proto_hpke_proto_msgTypes,
	}.Build()
	File_third_party_tink_proto_hpke_proto = out.File
	file_third_party_tink_proto_hpke_proto_rawDesc = nil
	file_third_party_tink_proto_hpke_proto_goTypes = nil
	file_third_party_tink_proto_hpke_proto_depIdxs = nil
}
