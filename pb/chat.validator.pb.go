// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.proto

package proto

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/gogo/googleapis/google/api"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ChatEntity) Validate() error {
	return nil
}
func (this *GetChatsRequest) Validate() error {
	return nil
}
func (this *GetChatsResponse) Validate() error {
	for _, item := range this.Chats {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Chats", err)
			}
		}
	}
	return nil
}
