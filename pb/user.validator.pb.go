// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *UserEntity) Validate() error {
	return nil
}
func (this *GetUsersRequest) Validate() error {
	return nil
}
func (this *GetUsersResponse) Validate() error {
	for _, item := range this.Users {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Users", err)
			}
		}
	}
	return nil
}
func (this *GetUserRequest) Validate() error {
	return nil
}
func (this *GetUserResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
