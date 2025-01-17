//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerlockbox

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

// GetClient contains the methods for the Get group.
// Don't use this type directly, use NewGetClient() instead.
type GetClient struct {
	ep string
	pl runtime.Pipeline
}

// NewGetClient creates a new instance of GetClient with the specified values.
func NewGetClient(credential azcore.TokenCredential, options *arm.ClientOptions) *GetClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &GetClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// TenantOptedIn - Get Customer Lockbox request
// If the operation fails it returns the *ErrorResponse error type.
func (client *GetClient) TenantOptedIn(ctx context.Context, tenantID string, options *GetTenantOptedInOptions) (GetTenantOptedInResponse, error) {
	req, err := client.tenantOptedInCreateRequest(ctx, tenantID, options)
	if err != nil {
		return GetTenantOptedInResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GetTenantOptedInResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GetTenantOptedInResponse{}, client.tenantOptedInHandleError(resp)
	}
	return client.tenantOptedInHandleResponse(resp)
}

// tenantOptedInCreateRequest creates the TenantOptedIn request.
func (client *GetClient) tenantOptedInCreateRequest(ctx context.Context, tenantID string, options *GetTenantOptedInOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.CustomerLockbox/tenantOptedIn/{tenantId}"
	if tenantID == "" {
		return nil, errors.New("parameter tenantID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tenantId}", url.PathEscape(tenantID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-02-28-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// tenantOptedInHandleResponse handles the TenantOptedIn response.
func (client *GetClient) tenantOptedInHandleResponse(resp *http.Response) (GetTenantOptedInResponse, error) {
	result := GetTenantOptedInResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantOptInResponse); err != nil {
		return GetTenantOptedInResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// tenantOptedInHandleError handles the TenantOptedIn error response.
func (client *GetClient) tenantOptedInHandleError(resp *http.Response) error {
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
