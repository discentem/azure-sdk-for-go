//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeviceupdate

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

// PrivateEndpointConnectionProxiesClient contains the methods for the PrivateEndpointConnectionProxies group.
// Don't use this type directly, use NewPrivateEndpointConnectionProxiesClient() instead.
type PrivateEndpointConnectionProxiesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPrivateEndpointConnectionProxiesClient creates a new instance of PrivateEndpointConnectionProxiesClient with the specified values.
func NewPrivateEndpointConnectionProxiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateEndpointConnectionProxiesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &PrivateEndpointConnectionProxiesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - (INTERNAL - DO NOT USE) Creates or updates the specified private endpoint connection proxy resource associated with the device
// update account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PrivateEndpointConnectionProxiesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionProxyID string, privateEndpointConnectionProxy PrivateEndpointConnectionProxy, options *PrivateEndpointConnectionProxiesBeginCreateOrUpdateOptions) (PrivateEndpointConnectionProxiesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, accountName, privateEndpointConnectionProxyID, privateEndpointConnectionProxy, options)
	if err != nil {
		return PrivateEndpointConnectionProxiesCreateOrUpdatePollerResponse{}, err
	}
	result := PrivateEndpointConnectionProxiesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PrivateEndpointConnectionProxiesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return PrivateEndpointConnectionProxiesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &PrivateEndpointConnectionProxiesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - (INTERNAL - DO NOT USE) Creates or updates the specified private endpoint connection proxy resource associated with the device update
// account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PrivateEndpointConnectionProxiesClient) createOrUpdate(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionProxyID string, privateEndpointConnectionProxy PrivateEndpointConnectionProxy, options *PrivateEndpointConnectionProxiesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, privateEndpointConnectionProxyID, privateEndpointConnectionProxy, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrivateEndpointConnectionProxiesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionProxyID string, privateEndpointConnectionProxy PrivateEndpointConnectionProxy, options *PrivateEndpointConnectionProxiesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/privateEndpointConnectionProxies/{privateEndpointConnectionProxyId}"
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
	if privateEndpointConnectionProxyID == "" {
		return nil, errors.New("parameter privateEndpointConnectionProxyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionProxyId}", url.PathEscape(privateEndpointConnectionProxyID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, privateEndpointConnectionProxy)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *PrivateEndpointConnectionProxiesClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - (INTERNAL - DO NOT USE) Deletes the specified private endpoint connection proxy associated with the device update account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PrivateEndpointConnectionProxiesClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionProxyID string, options *PrivateEndpointConnectionProxiesBeginDeleteOptions) (PrivateEndpointConnectionProxiesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, privateEndpointConnectionProxyID, options)
	if err != nil {
		return PrivateEndpointConnectionProxiesDeletePollerResponse{}, err
	}
	result := PrivateEndpointConnectionProxiesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PrivateEndpointConnectionProxiesClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return PrivateEndpointConnectionProxiesDeletePollerResponse{}, err
	}
	result.Poller = &PrivateEndpointConnectionProxiesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - (INTERNAL - DO NOT USE) Deletes the specified private endpoint connection proxy associated with the device update account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PrivateEndpointConnectionProxiesClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionProxyID string, options *PrivateEndpointConnectionProxiesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, privateEndpointConnectionProxyID, options)
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
func (client *PrivateEndpointConnectionProxiesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionProxyID string, options *PrivateEndpointConnectionProxiesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/privateEndpointConnectionProxies/{privateEndpointConnectionProxyId}"
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
	if privateEndpointConnectionProxyID == "" {
		return nil, errors.New("parameter privateEndpointConnectionProxyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionProxyId}", url.PathEscape(privateEndpointConnectionProxyID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *PrivateEndpointConnectionProxiesClient) deleteHandleError(resp *http.Response) error {
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

// Get - (INTERNAL - DO NOT USE) Get the specified private endpoint connection proxy associated with the device update account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PrivateEndpointConnectionProxiesClient) Get(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionProxyID string, options *PrivateEndpointConnectionProxiesGetOptions) (PrivateEndpointConnectionProxiesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, privateEndpointConnectionProxyID, options)
	if err != nil {
		return PrivateEndpointConnectionProxiesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionProxiesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateEndpointConnectionProxiesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateEndpointConnectionProxiesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionProxyID string, options *PrivateEndpointConnectionProxiesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/privateEndpointConnectionProxies/{privateEndpointConnectionProxyId}"
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
	if privateEndpointConnectionProxyID == "" {
		return nil, errors.New("parameter privateEndpointConnectionProxyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionProxyId}", url.PathEscape(privateEndpointConnectionProxyID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateEndpointConnectionProxiesClient) getHandleResponse(resp *http.Response) (PrivateEndpointConnectionProxiesGetResponse, error) {
	result := PrivateEndpointConnectionProxiesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionProxy); err != nil {
		return PrivateEndpointConnectionProxiesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PrivateEndpointConnectionProxiesClient) getHandleError(resp *http.Response) error {
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

// ListByAccount - (INTERNAL - DO NOT USE) List all private endpoint connection proxies in a device update account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PrivateEndpointConnectionProxiesClient) ListByAccount(ctx context.Context, resourceGroupName string, accountName string, options *PrivateEndpointConnectionProxiesListByAccountOptions) (PrivateEndpointConnectionProxiesListByAccountResponse, error) {
	req, err := client.listByAccountCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return PrivateEndpointConnectionProxiesListByAccountResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionProxiesListByAccountResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateEndpointConnectionProxiesListByAccountResponse{}, client.listByAccountHandleError(resp)
	}
	return client.listByAccountHandleResponse(resp)
}

// listByAccountCreateRequest creates the ListByAccount request.
func (client *PrivateEndpointConnectionProxiesClient) listByAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *PrivateEndpointConnectionProxiesListByAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/privateEndpointConnectionProxies"
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
	reqQP.Set("api-version", "2020-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAccountHandleResponse handles the ListByAccount response.
func (client *PrivateEndpointConnectionProxiesClient) listByAccountHandleResponse(resp *http.Response) (PrivateEndpointConnectionProxiesListByAccountResponse, error) {
	result := PrivateEndpointConnectionProxiesListByAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionProxyListResult); err != nil {
		return PrivateEndpointConnectionProxiesListByAccountResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByAccountHandleError handles the ListByAccount error response.
func (client *PrivateEndpointConnectionProxiesClient) listByAccountHandleError(resp *http.Response) error {
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

// Validate - (INTERNAL - DO NOT USE) Validates a private endpoint connection proxy object.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PrivateEndpointConnectionProxiesClient) Validate(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionProxyID string, privateEndpointConnectionProxy PrivateEndpointConnectionProxy, options *PrivateEndpointConnectionProxiesValidateOptions) (PrivateEndpointConnectionProxiesValidateResponse, error) {
	req, err := client.validateCreateRequest(ctx, resourceGroupName, accountName, privateEndpointConnectionProxyID, privateEndpointConnectionProxy, options)
	if err != nil {
		return PrivateEndpointConnectionProxiesValidateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionProxiesValidateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateEndpointConnectionProxiesValidateResponse{}, client.validateHandleError(resp)
	}
	return PrivateEndpointConnectionProxiesValidateResponse{RawResponse: resp}, nil
}

// validateCreateRequest creates the Validate request.
func (client *PrivateEndpointConnectionProxiesClient) validateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionProxyID string, privateEndpointConnectionProxy PrivateEndpointConnectionProxy, options *PrivateEndpointConnectionProxiesValidateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/privateEndpointConnectionProxies/{privateEndpointConnectionProxyId}/validate"
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
	if privateEndpointConnectionProxyID == "" {
		return nil, errors.New("parameter privateEndpointConnectionProxyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionProxyId}", url.PathEscape(privateEndpointConnectionProxyID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, privateEndpointConnectionProxy)
}

// validateHandleError handles the Validate error response.
func (client *PrivateEndpointConnectionProxiesClient) validateHandleError(resp *http.Response) error {
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
