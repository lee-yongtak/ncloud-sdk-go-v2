/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.apigw.ntruss.com/vpc/v2
 *
 * API version: 2020-07-06T00:53:31Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type CreateVpcPeeringInstanceRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// VPCPeering설명
VpcPeeringDescription *string `json:"vpcPeeringDescription,omitempty"`

	// 요청VPC번호
SourceVpcNo *string `json:"sourceVpcNo"`

	// 수락VPC소유자ID
TargetVpcLoginId *string `json:"targetVpcLoginId,omitempty"`

	// 수락VPC이름
TargetVpcName *string `json:"targetVpcName,omitempty"`

	// 수락VPC번호
TargetVpcNo *string `json:"targetVpcNo"`

	// VPCPeering이름
VpcPeeringName *string `json:"vpcPeeringName,omitempty"`
}
