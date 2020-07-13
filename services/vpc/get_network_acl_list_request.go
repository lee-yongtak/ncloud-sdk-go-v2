/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.apigw.ntruss.com/vpc/v2
 *
 * API version: 2020-07-06T00:53:31Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type GetNetworkAclListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 네트워크ACL이름
NetworkAclName *string `json:"networkAclName,omitempty"`

	// 네트워크ACL상태코드
NetworkAclStatusCode *string `json:"networkAclStatusCode,omitempty"`

	// 네트워크ACL번호리스트
NetworkAclNoList []*string `json:"networkAclNoList,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`

	// VPC번호
VpcNo *string `json:"vpcNo,omitempty"`
}
