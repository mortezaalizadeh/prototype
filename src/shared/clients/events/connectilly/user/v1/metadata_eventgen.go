// Code generated by connectillyctl, DO NOT EDIT.

package v1

import (
	_ "embed"
)

const TopicName = "user.v1.event"

//go:embed schema/schema.proto
var protobuf string

func GetProtobufSchema() string {
	return protobuf
}