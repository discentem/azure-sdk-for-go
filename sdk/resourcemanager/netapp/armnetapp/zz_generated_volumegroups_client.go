//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VolumeGroupsClient contains the methods for the VolumeGroups group.
// Don't use this type directly, use NewVolumeGroupsClient() instead.
type VolumeGroupsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewVolumeGroupsClient creates a new instance of VolumeGroupsClient with the specified values.
func NewVolumeGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VolumeGroupsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &VolumeGroupsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreate - Create a volume group along with specified volumes
// If the operation fails it returns a generic error.
func (client *VolumeGroupsClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, volumeGroupName string, body VolumeGroupDetails, options *VolumeGroupsBeginCreateOptions) (VolumeGroupsCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, accountName, volumeGroupName, body, options)
	if err != nil {
		return VolumeGroupsCreatePollerResponse{}, err
	}
	result := VolumeGroupsCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VolumeGroupsClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return VolumeGroupsCreatePollerResponse{}, err
	}
	result.Poller = &VolumeGroupsCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Create a volume group along with specified volumes
// If the operation fails it returns a generic error.
func (client *VolumeGroupsClient) create(ctx context.Context, resourceGroupName string, accountName string, volumeGroupName string, body VolumeGroupDetails, options *VolumeGroupsBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, volumeGroupName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *VolumeGroupsClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, volumeGroupName string, body VolumeGroupDetails, options *VolumeGroupsBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/volumeGroups/{volumeGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if volumeGroupName == "" {
		return nil, errors.New("parameter volumeGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeGroupName}", url.PathEscape(volumeGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// createHandleError handles the Create error response.
func (client *VolumeGroupsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Delete the specified volume group only if there are no volumes under volume group.
// If the operation fails it returns a generic error.
func (client *VolumeGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, volumeGroupName string, options *VolumeGroupsBeginDeleteOptions) (VolumeGroupsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, volumeGroupName, options)
	if err != nil {
		return VolumeGroupsDeletePollerResponse{}, err
	}
	result := VolumeGroupsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VolumeGroupsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return VolumeGroupsDeletePollerResponse{}, err
	}
	result.Poller = &VolumeGroupsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete the specified volume group only if there are no volumes under volume group.
// If the operation fails it returns a generic error.
func (client *VolumeGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, volumeGroupName string, options *VolumeGroupsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, volumeGroupName, options)
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
func (client *VolumeGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, volumeGroupName string, options *VolumeGroupsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/volumeGroups/{volumeGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if volumeGroupName == "" {
		return nil, errors.New("parameter volumeGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeGroupName}", url.PathEscape(volumeGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VolumeGroupsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Get details of the specified volume group
// If the operation fails it returns a generic error.
func (client *VolumeGroupsClient) Get(ctx context.Context, resourceGroupName string, accountName string, volumeGroupName string, options *VolumeGroupsGetOptions) (VolumeGroupsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, volumeGroupName, options)
	if err != nil {
		return VolumeGroupsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VolumeGroupsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VolumeGroupsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VolumeGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, volumeGroupName string, options *VolumeGroupsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/volumeGroups/{volumeGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if volumeGroupName == "" {
		return nil, errors.New("parameter volumeGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeGroupName}", url.PathEscape(volumeGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VolumeGroupsClient) getHandleResponse(resp *http.Response) (VolumeGroupsGetResponse, error) {
	result := VolumeGroupsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VolumeGroupDetails); err != nil {
		return VolumeGroupsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VolumeGroupsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByNetAppAccount - List all volume groups for given account
// If the operation fails it returns a generic error.
func (client *VolumeGroupsClient) ListByNetAppAccount(ctx context.Context, resourceGroupName string, accountName string, options *VolumeGroupsListByNetAppAccountOptions) (VolumeGroupsListByNetAppAccountResponse, error) {
	req, err := client.listByNetAppAccountCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return VolumeGroupsListByNetAppAccountResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VolumeGroupsListByNetAppAccountResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VolumeGroupsListByNetAppAccountResponse{}, client.listByNetAppAccountHandleError(resp)
	}
	return client.listByNetAppAccountHandleResponse(resp)
}

// listByNetAppAccountCreateRequest creates the ListByNetAppAccount request.
func (client *VolumeGroupsClient) listByNetAppAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *VolumeGroupsListByNetAppAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/volumeGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByNetAppAccountHandleResponse handles the ListByNetAppAccount response.
func (client *VolumeGroupsClient) listByNetAppAccountHandleResponse(resp *http.Response) (VolumeGroupsListByNetAppAccountResponse, error) {
	result := VolumeGroupsListByNetAppAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VolumeGroupList); err != nil {
		return VolumeGroupsListByNetAppAccountResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByNetAppAccountHandleError handles the ListByNetAppAccount error response.
func (client *VolumeGroupsClient) listByNetAppAccountHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
