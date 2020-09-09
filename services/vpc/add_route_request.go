/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vpc/v2
 *
 * API version: 2020-08-27T11:47:46Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type AddRouteRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 라우트리스트
RouteList []*RouteParameter `json:"routeList"`

	// 라우트테이블번호
RouteTableNo *string `json:"routeTableNo"`

	// VPC번호
VpcNo *string `json:"vpcNo"`
}
