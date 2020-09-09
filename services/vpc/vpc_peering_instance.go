/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vpc/v2
 *
 * API version: 2020-08-27T11:47:46Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type VpcPeeringInstance struct {

	// VPCPeering인스턴스번호
VpcPeeringInstanceNo *string `json:"vpcPeeringInstanceNo,omitempty"`

	// VPCPeering이름
VpcPeeringName *string `json:"vpcPeeringName,omitempty"`

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 생성일시
CreateDate *string `json:"createDate,omitempty"`

	// 마지막수정일시
LastModifiyDate *string `json:"lastModifiyDate,omitempty"`

	// VPCPeering인스턴스상태
VpcPeeringInstanceStatus *CommonCode `json:"vpcPeeringInstanceStatus,omitempty"`

	// VPCPeering인스턴스상태이름
VpcPeeringInstanceStatusName *string `json:"vpcPeeringInstanceStatusName,omitempty"`

	// VPCPeering인스턴스OP
VpcPeeringInstanceOperation *CommonCode `json:"vpcPeeringInstanceOperation,omitempty"`

	// 요청VPC번호
SourceVpcNo *string `json:"sourceVpcNo,omitempty"`

	// 요청VPC이름
SourceVpcName *string `json:"sourceVpcName,omitempty"`

	// 요청VPC IPv4 CIDR블록
SourceVpcIpv4CidrBlock *string `json:"sourceVpcIpv4CidrBlock,omitempty"`

	// 요청VPC소유자ID
SourceVpcLoginId *string `json:"sourceVpcLoginId,omitempty"`

	// 수락VPC번호
TargetVpcNo *string `json:"targetVpcNo,omitempty"`

	// 수락VPC이름
TargetVpcName *string `json:"targetVpcName,omitempty"`

	// 수락VPC IPv4 CIDR블록
TargetVpcIpv4CidrBlock *string `json:"targetVpcIpv4CidrBlock,omitempty"`

	// 수락VPC소유자ID
TargetVpcLoginId *string `json:"targetVpcLoginId,omitempty"`

	// VPCPeering설명
VpcPeeringDescription *string `json:"vpcPeeringDescription,omitempty"`

	// 역방향VPCPeering존재여부
HasReverseVpcPeering *bool `json:"hasReverseVpcPeering,omitempty"`

	// 계정간의VPCPeering여부
IsBetweenAccounts *bool `json:"isBetweenAccounts,omitempty"`

	// 역방향VPCPeering인스턴스번호
ReverseVpcPeeringInstanceNo *string `json:"reverseVpcPeeringInstanceNo,omitempty"`
}
