# \V2Api

All URIs are relative to *https://ncloud.apigw.ntruss.com/clouddb/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCloudDBInstance**](V2Api.md#CreateCloudDBInstance) | **Post** /createCloudDBInstance | 
[**DeleteCloudDBServerInstance**](V2Api.md#DeleteCloudDBServerInstance) | **Post** /deleteCloudDBServerInstance | 
[**GetCloudDBConfigGroupList**](V2Api.md#GetCloudDBConfigGroupList) | **Post** /getCloudDBConfigGroupList | 
[**GetCloudDBImageProductList**](V2Api.md#GetCloudDBImageProductList) | **Post** /getCloudDBImageProductList | 
[**GetCloudDBInstanceList**](V2Api.md#GetCloudDBInstanceList) | **Post** /getCloudDBInstanceList | 
[**GetCloudDBProductList**](V2Api.md#GetCloudDBProductList) | **Post** /getCloudDBProductList | 
[**RebootCloudDBServerInstance**](V2Api.md#RebootCloudDBServerInstance) | **Post** /rebootCloudDBServerInstance | 


# **CreateCloudDBInstance**
> CreateCloudDbInstanceResponse CreateCloudDBInstance(ctx, createCloudDBInstanceRequest)


CloudDB인스턴스생성

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **createCloudDBInstanceRequest** | [**CreateCloudDbInstanceRequest**](CreateCloudDbInstanceRequest.md)| createCloudDBInstanceRequest | 

### Return type

[**CreateCloudDbInstanceResponse**](createCloudDBInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudDBServerInstance**
> DeleteCloudDbServerInstanceResponse DeleteCloudDBServerInstance(ctx, deleteCloudDBServerInstanceRequest)


CloudDB서버인스턴스삭제

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **deleteCloudDBServerInstanceRequest** | [**DeleteCloudDbServerInstanceRequest**](DeleteCloudDbServerInstanceRequest.md)| deleteCloudDBServerInstanceRequest | 

### Return type

[**DeleteCloudDbServerInstanceResponse**](deleteCloudDBServerInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudDBConfigGroupList**
> GetCloudDbConfigGroupListResponse GetCloudDBConfigGroupList(ctx, getCloudDBConfigGroupListRequest)


CloudDB설정그룹리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getCloudDBConfigGroupListRequest** | [**GetCloudDbConfigGroupListRequest**](GetCloudDbConfigGroupListRequest.md)| getCloudDBConfigGroupListRequest | 

### Return type

[**GetCloudDbConfigGroupListResponse**](getCloudDBConfigGroupListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudDBImageProductList**
> GetCloudDbImageProductListResponse GetCloudDBImageProductList(ctx, getCloudDBImageProductListRequest)


CloudDB이미지상품리스트

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getCloudDBImageProductListRequest** | [**GetCloudDbImageProductListRequest**](GetCloudDbImageProductListRequest.md)| getCloudDBImageProductListRequest | 

### Return type

[**GetCloudDbImageProductListResponse**](getCloudDBImageProductListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudDBInstanceList**
> GetCloudDbInstanceListResponse GetCloudDBInstanceList(ctx, getCloudDBInstanceListRequest)


CloudDB인스턴스리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getCloudDBInstanceListRequest** | [**GetCloudDbInstanceListRequest**](GetCloudDbInstanceListRequest.md)| getCloudDBInstanceListRequest | 

### Return type

[**GetCloudDbInstanceListResponse**](getCloudDBInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudDBProductList**
> GetCloudDbProductListResponse GetCloudDBProductList(ctx, getCloudDBProductListRequest)


CloudDB상품리스트조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **getCloudDBProductListRequest** | [**GetCloudDbProductListRequest**](GetCloudDbProductListRequest.md)| getCloudDBProductListRequest | 

### Return type

[**GetCloudDbProductListResponse**](getCloudDBProductListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RebootCloudDBServerInstance**
> RebootCloudDbServerInstanceResponse RebootCloudDBServerInstance(ctx, rebootCloudDBServerInstanceRequest)


CloudDB서버인스턴스재부팅

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **rebootCloudDBServerInstanceRequest** | [**RebootCloudDbServerInstanceRequest**](RebootCloudDbServerInstanceRequest.md)| rebootCloudDBServerInstanceRequest | 

### Return type

[**RebootCloudDbServerInstanceResponse**](rebootCloudDBServerInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
