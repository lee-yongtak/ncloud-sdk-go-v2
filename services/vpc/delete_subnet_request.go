/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vpc/v2
 *
 * API version: 2020-08-27T11:47:46Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type DeleteSubnetRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 서브넷번호
SubnetNo *string `json:"subnetNo"`
}
