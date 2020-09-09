/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vpc/v2
 *
 * API version: 2020-08-27T11:47:46Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type SetNatGatewayDescriptionRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// NATGateway인스턴스번호
NatGatewayInstanceNo *string `json:"natGatewayInstanceNo"`

	// NATGateway설명
NatGatewayDescription *string `json:"natGatewayDescription,omitempty"`
}
