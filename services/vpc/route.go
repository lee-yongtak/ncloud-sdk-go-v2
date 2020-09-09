/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vpc/v2
 *
 * API version: 2020-08-27T11:47:46Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type Route struct {

	// 목적지CIDR블록
DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty"`

	// 목적지이름
TargetName *string `json:"targetName,omitempty"`

	// 라우트테이블번호
RouteTableNo *string `json:"routeTableNo,omitempty"`

	// 목적지유형
TargetType *CommonCode `json:"targetType,omitempty"`

	// 목적지번호
TargetNo *string `json:"targetNo,omitempty"`

	// Default여부
IsDefault *bool `json:"isDefault,omitempty"`
}
