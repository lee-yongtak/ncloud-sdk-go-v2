/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vpc/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type NatGatewayInstance struct {

	// VPC번호
VpcNo *string `json:"vpcNo,omitempty"`

	// VPC이름
VpcName *string `json:"vpcName,omitempty"`

	// NATGateway인스턴스번호
NatGatewayInstanceNo *string `json:"natGatewayInstanceNo,omitempty"`

	// NATGateway이름
NatGatewayName *string `json:"natGatewayName,omitempty"`

	// 공인IP주소
PublicIp *string `json:"publicIp,omitempty"`

	// NATGateway인스턴스상태
NatGatewayInstanceStatus *CommonCode `json:"natGatewayInstanceStatus,omitempty"`

	// NATGateway인스턴스상태이름
NatGatewayInstanceStatusName *string `json:"natGatewayInstanceStatusName,omitempty"`

	// NATGateway인스턴스OP
NatGatewayInstanceOperation *CommonCode `json:"natGatewayInstanceOperation,omitempty"`

	// 생성일시
CreateDate *string `json:"createDate,omitempty"`

	// NATGateway설명
NatGatewayDescription *string `json:"natGatewayDescription,omitempty"`

	// ZONE코드
ZoneCode *string `json:"zoneCode,omitempty"`
}
