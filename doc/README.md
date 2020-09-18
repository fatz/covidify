# Documentation for Covidify

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *https://covidify.testing.mesosphe.re*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*PublicApi* | [**addVisit**](Apis/PublicApi.md#addvisit) | **POST** /visit | adds an Visit entry
*PublicApi* | [**checkVisit**](Apis/PublicApi.md#checkvisit) | **GET** /visit/{visitID} | Visit status check
*RestrictedApi* | [**addReportVisitor**](Apis/RestrictedApi.md#addreportvisitor) | **POST** /report/visitor | Report an infected visitor


<a name="documentation-for-models"></a>
## Documentation for Models

 - [ReportVisitor](.//Models/ReportVisitor.md)
 - [Visit](.//Models/Visit.md)
 - [VisitRisk](.//Models/VisitRisk.md)
 - [Visitor](.//Models/Visitor.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

All endpoints do not require authorization.
