//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// WorkspaceManagedSQLServerEncryptionProtectorClient contains the methods for the WorkspaceManagedSQLServerEncryptionProtector group.
// Don't use this type directly, use NewWorkspaceManagedSQLServerEncryptionProtectorClient() instead.
type WorkspaceManagedSQLServerEncryptionProtectorClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewWorkspaceManagedSQLServerEncryptionProtectorClient creates a new instance of WorkspaceManagedSQLServerEncryptionProtectorClient with the specified values.
func NewWorkspaceManagedSQLServerEncryptionProtectorClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WorkspaceManagedSQLServerEncryptionProtectorClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &WorkspaceManagedSQLServerEncryptionProtectorClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Updates workspace managed sql server's encryption protector.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, encryptionProtectorName EncryptionProtectorName, parameters EncryptionProtector, options *WorkspaceManagedSQLServerEncryptionProtectorBeginCreateOrUpdateOptions) (WorkspaceManagedSQLServerEncryptionProtectorCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, encryptionProtectorName, parameters, options)
	if err != nil {
		return WorkspaceManagedSQLServerEncryptionProtectorCreateOrUpdatePollerResponse{}, err
	}
	result := WorkspaceManagedSQLServerEncryptionProtectorCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WorkspaceManagedSQLServerEncryptionProtectorClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return WorkspaceManagedSQLServerEncryptionProtectorCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &WorkspaceManagedSQLServerEncryptionProtectorCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Updates workspace managed sql server's encryption protector.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, encryptionProtectorName EncryptionProtectorName, parameters EncryptionProtector, options *WorkspaceManagedSQLServerEncryptionProtectorBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, encryptionProtectorName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, encryptionProtectorName EncryptionProtectorName, parameters EncryptionProtector, options *WorkspaceManagedSQLServerEncryptionProtectorBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/encryptionProtector/{encryptionProtectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Get - Get workspace managed sql server's encryption protector.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, encryptionProtectorName EncryptionProtectorName, options *WorkspaceManagedSQLServerEncryptionProtectorGetOptions) (WorkspaceManagedSQLServerEncryptionProtectorGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, encryptionProtectorName, options)
	if err != nil {
		return WorkspaceManagedSQLServerEncryptionProtectorGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceManagedSQLServerEncryptionProtectorGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceManagedSQLServerEncryptionProtectorGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, encryptionProtectorName EncryptionProtectorName, options *WorkspaceManagedSQLServerEncryptionProtectorGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/encryptionProtector/{encryptionProtectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) getHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerEncryptionProtectorGetResponse, error) {
	result := WorkspaceManagedSQLServerEncryptionProtectorGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EncryptionProtector); err != nil {
		return WorkspaceManagedSQLServerEncryptionProtectorGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) getHandleError(resp *http.Response) error {
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

// List - Get list of encryption protectors for workspace managed sql server.
// If the operation fails it returns a generic error.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) List(resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerEncryptionProtectorListOptions) *WorkspaceManagedSQLServerEncryptionProtectorListPager {
	return &WorkspaceManagedSQLServerEncryptionProtectorListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
		},
		advancer: func(ctx context.Context, resp WorkspaceManagedSQLServerEncryptionProtectorListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.EncryptionProtectorListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerEncryptionProtectorListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/encryptionProtector"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) listHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerEncryptionProtectorListResponse, error) {
	result := WorkspaceManagedSQLServerEncryptionProtectorListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EncryptionProtectorListResult); err != nil {
		return WorkspaceManagedSQLServerEncryptionProtectorListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginRevalidate - Revalidates workspace managed sql server's existing encryption protector.
// If the operation fails it returns a generic error.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) BeginRevalidate(ctx context.Context, resourceGroupName string, workspaceName string, encryptionProtectorName EncryptionProtectorName, options *WorkspaceManagedSQLServerEncryptionProtectorBeginRevalidateOptions) (WorkspaceManagedSQLServerEncryptionProtectorRevalidatePollerResponse, error) {
	resp, err := client.revalidate(ctx, resourceGroupName, workspaceName, encryptionProtectorName, options)
	if err != nil {
		return WorkspaceManagedSQLServerEncryptionProtectorRevalidatePollerResponse{}, err
	}
	result := WorkspaceManagedSQLServerEncryptionProtectorRevalidatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WorkspaceManagedSQLServerEncryptionProtectorClient.Revalidate", "", resp, client.pl, client.revalidateHandleError)
	if err != nil {
		return WorkspaceManagedSQLServerEncryptionProtectorRevalidatePollerResponse{}, err
	}
	result.Poller = &WorkspaceManagedSQLServerEncryptionProtectorRevalidatePoller{
		pt: pt,
	}
	return result, nil
}

// Revalidate - Revalidates workspace managed sql server's existing encryption protector.
// If the operation fails it returns a generic error.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) revalidate(ctx context.Context, resourceGroupName string, workspaceName string, encryptionProtectorName EncryptionProtectorName, options *WorkspaceManagedSQLServerEncryptionProtectorBeginRevalidateOptions) (*http.Response, error) {
	req, err := client.revalidateCreateRequest(ctx, resourceGroupName, workspaceName, encryptionProtectorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.revalidateHandleError(resp)
	}
	return resp, nil
}

// revalidateCreateRequest creates the Revalidate request.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) revalidateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, encryptionProtectorName EncryptionProtectorName, options *WorkspaceManagedSQLServerEncryptionProtectorBeginRevalidateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/encryptionProtector/{encryptionProtectorName}/revalidate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// revalidateHandleError handles the Revalidate error response.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) revalidateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
