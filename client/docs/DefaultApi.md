# {{classname}}

All URIs are relative to *http://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Geolocate**](DefaultApi.md#Geolocate) | **Get** /geolocate/{ip} | Get geolocation of an IP

# **Geolocate**
> map[string]interface{} Geolocate(ctx, ip)
Get geolocation of an IP

Queries ipgeolocation.io for a valid IP address' geolocation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ip** | **string**| IP Address | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

