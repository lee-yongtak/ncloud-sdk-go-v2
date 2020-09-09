/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vserver/v2
 *
 * API version: 2020-09-04T05:39:38Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type AccessControlGroup struct {

	// ACG번호
AccessControlGroupNo *string `json:"accessControlGroupNo,omitempty"`

	// ACG이름
AccessControlGroupName *string `json:"accessControlGroupName,omitempty"`

	// Default여부
IsDefault *bool `json:"isDefault,omitempty"`

	// VPC번호
VpcNo *string `json:"vpcNo,omitempty"`

	// ACG상태
AccessControlGroupStatus *CommonCode `json:"accessControlGroupStatus,omitempty"`

	// ACG설명
AccessControlGroupDescription *string `json:"accessControlGroupDescription,omitempty"`
}
