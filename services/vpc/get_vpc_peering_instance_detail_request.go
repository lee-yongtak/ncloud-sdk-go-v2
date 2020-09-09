/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vpc/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type GetVpcPeeringInstanceDetailRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// VPCPeering인스턴스번호
VpcPeeringInstanceNo *string `json:"vpcPeeringInstanceNo"`
}
