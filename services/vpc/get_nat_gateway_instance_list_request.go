/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vpc/v2
 *
 * API version: 2020-08-27T11:47:46Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type GetNatGatewayInstanceListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// ZONE코드
ZoneCode *string `json:"zoneCode,omitempty"`

	// NATGateway인스턴스번호리스트
NatGatewayInstanceNoList []*string `json:"natGatewayInstanceNoList,omitempty"`

	// 공인IP주소
PublicIp *string `json:"publicIp,omitempty"`

	// VPC이름
VpcName *string `json:"vpcName,omitempty"`

	// natGatewayName
NatGatewayName *string `json:"natGatewayName,omitempty"`

	// NATGateway인스턴스상태코드
NatGatewayInstanceStatusCode *string `json:"natGatewayInstanceStatusCode,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`
}
