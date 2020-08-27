/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vpc/v2
 *
 * API version: 2020-08-04T12:59:37Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type Subnet struct {

	// 서브넷번호
SubnetNo *string `json:"subnetNo,omitempty"`

	// VPC번호
VpcNo *string `json:"vpcNo,omitempty"`

	// ZONE코드
ZoneCode *string `json:"zoneCode,omitempty"`

	// 서브넷이름
SubnetName *string `json:"subnetName,omitempty"`

	// 서브넷
Subnet *string `json:"subnet,omitempty"`

	// 서브넷상태
SubnetStatus *CommonCode `json:"subnetStatus,omitempty"`

	// 생성일시
CreateDate *string `json:"createDate,omitempty"`

	// 서브넷유형
SubnetType *CommonCode `json:"subnetType,omitempty"`

	// 용도유형
UsageType *CommonCode `json:"usageType,omitempty"`

	// 네트워크ACL번호
NetworkAclNo *string `json:"networkAclNo,omitempty"`
}
