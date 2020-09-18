# PublicApi

All URIs are relative to *https://covidify.testing.mesosphe.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addVisit**](PublicApi.md#addVisit) | **POST** /visit | adds an Visit entry
[**checkVisit**](PublicApi.md#checkVisit) | **GET** /visit/{visitID} | Visit status check


<a name="addVisit"></a>
# **addVisit**
> addVisit(visit)

adds an Visit entry

    Adds an visitor to the Database

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **visit** | [**Visit**](..//Models/Visit.md)| Inventory item to add | [optional]

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="checkVisit"></a>
# **checkVisit**
> VisitRisk checkVisit(visitID)

Visit status check

    Returns the infection risk for a Visit

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **visitID** | [**UUID**](..//Models/.md)| ID of Visit to return | [default to null]

### Return type

[**VisitRisk**](..//Models/VisitRisk.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

