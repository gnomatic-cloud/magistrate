// Code generated by protoc-gen-go-binary. DO NOT EDIT.
// source: pbcatalog/v1alpha1/service_endpoints.proto

package catalogv1alpha1

import (
	"google.golang.org/protobuf/proto"
)

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *ServiceEndpoints) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *ServiceEndpoints) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *Endpoint) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *Endpoint) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}
