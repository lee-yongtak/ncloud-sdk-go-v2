/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type GetNetworkInterfaceListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 네트워크인터페이스번호리스트
NetworkInterfaceNoList []*string `json:"networkInterfaceNoList,omitempty"`

	// IP주소
Ip *string `json:"ip,omitempty"`

	// 네트워크인터페이스이름
NetworkInterfaceName *string `json:"networkInterfaceName,omitempty"`

	// 서버이름
ServerName *string `json:"serverName,omitempty"`

	// 서브넷이름
SubnetName *string `json:"subnetName,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`
}
