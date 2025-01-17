//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsaas

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SaasSubscriptionLevelClient contains the methods for the SaasSubscriptionLevel group.
// Don't use this type directly, use NewSaasSubscriptionLevelClient() instead.
type SaasSubscriptionLevelClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSaasSubscriptionLevelClient creates a new instance of SaasSubscriptionLevelClient with the specified values.
func NewSaasSubscriptionLevelClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SaasSubscriptionLevelClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &SaasSubscriptionLevelClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Creates or updates a SaaS resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SaasSubscriptionLevelClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters SaasResourceCreation, options *SaasSubscriptionLevelBeginCreateOrUpdateOptions) (SaasSubscriptionLevelCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return SaasSubscriptionLevelCreateOrUpdatePollerResponse{}, err
	}
	result := SaasSubscriptionLevelCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SaasSubscriptionLevelClient.CreateOrUpdate", "location", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return SaasSubscriptionLevelCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &SaasSubscriptionLevelCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a SaaS resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SaasSubscriptionLevelClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters SaasResourceCreation, options *SaasSubscriptionLevelBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SaasSubscriptionLevelClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters SaasResourceCreation, options *SaasSubscriptionLevelBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SaaS/resources/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SaasSubscriptionLevelClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the specified SaaS.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SaasSubscriptionLevelClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, options *SaasSubscriptionLevelBeginDeleteOptions) (SaasSubscriptionLevelDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return SaasSubscriptionLevelDeletePollerResponse{}, err
	}
	result := SaasSubscriptionLevelDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SaasSubscriptionLevelClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return SaasSubscriptionLevelDeletePollerResponse{}, err
	}
	result.Poller = &SaasSubscriptionLevelDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified SaaS.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SaasSubscriptionLevelClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, options *SaasSubscriptionLevelBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SaasSubscriptionLevelClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *SaasSubscriptionLevelBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SaaS/resources/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SaasSubscriptionLevelClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets information about the specified Subscription Level SaaS.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SaasSubscriptionLevelClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *SaasSubscriptionLevelGetOptions) (SaasSubscriptionLevelGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return SaasSubscriptionLevelGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SaasSubscriptionLevelGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SaasSubscriptionLevelGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SaasSubscriptionLevelClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *SaasSubscriptionLevelGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SaaS/resources/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SaasSubscriptionLevelClient) getHandleResponse(resp *http.Response) (SaasSubscriptionLevelGetResponse, error) {
	result := SaasSubscriptionLevelGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SaasResource); err != nil {
		return SaasSubscriptionLevelGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SaasSubscriptionLevelClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListAccessToken - Gets the ISV access token for a specified Subscription Level SaaS.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SaasSubscriptionLevelClient) ListAccessToken(ctx context.Context, resourceGroupName string, resourceName string, options *SaasSubscriptionLevelListAccessTokenOptions) (SaasSubscriptionLevelListAccessTokenResponse, error) {
	req, err := client.listAccessTokenCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return SaasSubscriptionLevelListAccessTokenResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SaasSubscriptionLevelListAccessTokenResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SaasSubscriptionLevelListAccessTokenResponse{}, client.listAccessTokenHandleError(resp)
	}
	return client.listAccessTokenHandleResponse(resp)
}

// listAccessTokenCreateRequest creates the ListAccessToken request.
func (client *SaasSubscriptionLevelClient) listAccessTokenCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *SaasSubscriptionLevelListAccessTokenOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SaaS/resources/{resourceName}/listAccessToken"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAccessTokenHandleResponse handles the ListAccessToken response.
func (client *SaasSubscriptionLevelClient) listAccessTokenHandleResponse(resp *http.Response) (SaasSubscriptionLevelListAccessTokenResponse, error) {
	result := SaasSubscriptionLevelListAccessTokenResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessTokenResult); err != nil {
		return SaasSubscriptionLevelListAccessTokenResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listAccessTokenHandleError handles the ListAccessToken error response.
func (client *SaasSubscriptionLevelClient) listAccessTokenHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByAzureSubscription - Gets information about all the Subscription Level SaaS in a certain Azure subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SaasSubscriptionLevelClient) ListByAzureSubscription(options *SaasSubscriptionLevelListByAzureSubscriptionOptions) *SaasSubscriptionLevelListByAzureSubscriptionPager {
	return &SaasSubscriptionLevelListByAzureSubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByAzureSubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SaasSubscriptionLevelListByAzureSubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SaasResourceResponseWithContinuation.NextLink)
		},
	}
}

// listByAzureSubscriptionCreateRequest creates the ListByAzureSubscription request.
func (client *SaasSubscriptionLevelClient) listByAzureSubscriptionCreateRequest(ctx context.Context, options *SaasSubscriptionLevelListByAzureSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.SaaS/resources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAzureSubscriptionHandleResponse handles the ListByAzureSubscription response.
func (client *SaasSubscriptionLevelClient) listByAzureSubscriptionHandleResponse(resp *http.Response) (SaasSubscriptionLevelListByAzureSubscriptionResponse, error) {
	result := SaasSubscriptionLevelListByAzureSubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SaasResourceResponseWithContinuation); err != nil {
		return SaasSubscriptionLevelListByAzureSubscriptionResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByAzureSubscriptionHandleError handles the ListByAzureSubscription error response.
func (client *SaasSubscriptionLevelClient) listByAzureSubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Gets information about all the Subscription Level SaaS in a certain resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SaasSubscriptionLevelClient) ListByResourceGroup(resourceGroupName string, options *SaasSubscriptionLevelListByResourceGroupOptions) *SaasSubscriptionLevelListByResourceGroupPager {
	return &SaasSubscriptionLevelListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp SaasSubscriptionLevelListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SaasResourceResponseWithContinuation.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SaasSubscriptionLevelClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SaasSubscriptionLevelListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SaaS/resources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SaasSubscriptionLevelClient) listByResourceGroupHandleResponse(resp *http.Response) (SaasSubscriptionLevelListByResourceGroupResponse, error) {
	result := SaasSubscriptionLevelListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SaasResourceResponseWithContinuation); err != nil {
		return SaasSubscriptionLevelListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *SaasSubscriptionLevelClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginMoveResources - Move a specified Subscription Level SaaS.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SaasSubscriptionLevelClient) BeginMoveResources(ctx context.Context, resourceGroupName string, moveResourceParameter MoveResource, options *SaasSubscriptionLevelBeginMoveResourcesOptions) (SaasSubscriptionLevelMoveResourcesPollerResponse, error) {
	resp, err := client.moveResources(ctx, resourceGroupName, moveResourceParameter, options)
	if err != nil {
		return SaasSubscriptionLevelMoveResourcesPollerResponse{}, err
	}
	result := SaasSubscriptionLevelMoveResourcesPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SaasSubscriptionLevelClient.MoveResources", "location", resp, client.pl, client.moveResourcesHandleError)
	if err != nil {
		return SaasSubscriptionLevelMoveResourcesPollerResponse{}, err
	}
	result.Poller = &SaasSubscriptionLevelMoveResourcesPoller{
		pt: pt,
	}
	return result, nil
}

// MoveResources - Move a specified Subscription Level SaaS.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SaasSubscriptionLevelClient) moveResources(ctx context.Context, resourceGroupName string, moveResourceParameter MoveResource, options *SaasSubscriptionLevelBeginMoveResourcesOptions) (*http.Response, error) {
	req, err := client.moveResourcesCreateRequest(ctx, resourceGroupName, moveResourceParameter, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.moveResourcesHandleError(resp)
	}
	return resp, nil
}

// moveResourcesCreateRequest creates the MoveResources request.
func (client *SaasSubscriptionLevelClient) moveResourcesCreateRequest(ctx context.Context, resourceGroupName string, moveResourceParameter MoveResource, options *SaasSubscriptionLevelBeginMoveResourcesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/moveResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, moveResourceParameter)
}

// moveResourcesHandleError handles the MoveResources error response.
func (client *SaasSubscriptionLevelClient) moveResourcesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Updates a SaaS Subscription Level resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SaasSubscriptionLevelClient) BeginUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters SaasResourceCreation, options *SaasSubscriptionLevelBeginUpdateOptions) (SaasSubscriptionLevelUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return SaasSubscriptionLevelUpdatePollerResponse{}, err
	}
	result := SaasSubscriptionLevelUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SaasSubscriptionLevelClient.Update", "location", resp, client.pl, client.updateHandleError)
	if err != nil {
		return SaasSubscriptionLevelUpdatePollerResponse{}, err
	}
	result.Poller = &SaasSubscriptionLevelUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates a SaaS Subscription Level resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SaasSubscriptionLevelClient) update(ctx context.Context, resourceGroupName string, resourceName string, parameters SaasResourceCreation, options *SaasSubscriptionLevelBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *SaasSubscriptionLevelClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters SaasResourceCreation, options *SaasSubscriptionLevelBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SaaS/resources/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleError handles the Update error response.
func (client *SaasSubscriptionLevelClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdateToUnsubscribed - Unsubscribe from a specified Subscription Level SaaS.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SaasSubscriptionLevelClient) BeginUpdateToUnsubscribed(ctx context.Context, resourceGroupName string, resourceName string, parameters DeleteOptions, options *SaasSubscriptionLevelBeginUpdateToUnsubscribedOptions) (SaasSubscriptionLevelUpdateToUnsubscribedPollerResponse, error) {
	resp, err := client.updateToUnsubscribed(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return SaasSubscriptionLevelUpdateToUnsubscribedPollerResponse{}, err
	}
	result := SaasSubscriptionLevelUpdateToUnsubscribedPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SaasSubscriptionLevelClient.UpdateToUnsubscribed", "location", resp, client.pl, client.updateToUnsubscribedHandleError)
	if err != nil {
		return SaasSubscriptionLevelUpdateToUnsubscribedPollerResponse{}, err
	}
	result.Poller = &SaasSubscriptionLevelUpdateToUnsubscribedPoller{
		pt: pt,
	}
	return result, nil
}

// UpdateToUnsubscribed - Unsubscribe from a specified Subscription Level SaaS.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SaasSubscriptionLevelClient) updateToUnsubscribed(ctx context.Context, resourceGroupName string, resourceName string, parameters DeleteOptions, options *SaasSubscriptionLevelBeginUpdateToUnsubscribedOptions) (*http.Response, error) {
	req, err := client.updateToUnsubscribedCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.updateToUnsubscribedHandleError(resp)
	}
	return resp, nil
}

// updateToUnsubscribedCreateRequest creates the UpdateToUnsubscribed request.
func (client *SaasSubscriptionLevelClient) updateToUnsubscribedCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters DeleteOptions, options *SaasSubscriptionLevelBeginUpdateToUnsubscribedOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SaaS/resources/{resourceName}/unsubscribe"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateToUnsubscribedHandleError handles the UpdateToUnsubscribed error response.
func (client *SaasSubscriptionLevelClient) updateToUnsubscribedHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ValidateMoveResources - Validate whether a specified Subscription Level SaaS can be moved.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SaasSubscriptionLevelClient) ValidateMoveResources(ctx context.Context, resourceGroupName string, moveResourceParameter MoveResource, options *SaasSubscriptionLevelValidateMoveResourcesOptions) (SaasSubscriptionLevelValidateMoveResourcesResponse, error) {
	req, err := client.validateMoveResourcesCreateRequest(ctx, resourceGroupName, moveResourceParameter, options)
	if err != nil {
		return SaasSubscriptionLevelValidateMoveResourcesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SaasSubscriptionLevelValidateMoveResourcesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SaasSubscriptionLevelValidateMoveResourcesResponse{}, client.validateMoveResourcesHandleError(resp)
	}
	return SaasSubscriptionLevelValidateMoveResourcesResponse{RawResponse: resp}, nil
}

// validateMoveResourcesCreateRequest creates the ValidateMoveResources request.
func (client *SaasSubscriptionLevelClient) validateMoveResourcesCreateRequest(ctx context.Context, resourceGroupName string, moveResourceParameter MoveResource, options *SaasSubscriptionLevelValidateMoveResourcesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/validateMoveResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, moveResourceParameter)
}

// validateMoveResourcesHandleError handles the ValidateMoveResources error response.
func (client *SaasSubscriptionLevelClient) validateMoveResourcesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
