//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armreservations

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// CalculateExchangeClient contains the methods for the CalculateExchange group.
// Don't use this type directly, use NewCalculateExchangeClient() instead.
type CalculateExchangeClient struct {
	ep string
	pl runtime.Pipeline
}

// NewCalculateExchangeClient creates a new instance of CalculateExchangeClient with the specified values.
func NewCalculateExchangeClient(credential azcore.TokenCredential, options *arm.ClientOptions) *CalculateExchangeClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &CalculateExchangeClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginPost - Calculates price for exchanging Reservations if there are no policy errors.
// If the operation fails it returns the *Error error type.
func (client *CalculateExchangeClient) BeginPost(ctx context.Context, body CalculateExchangeRequest, options *CalculateExchangeBeginPostOptions) (CalculateExchangePostPollerResponse, error) {
	resp, err := client.post(ctx, body, options)
	if err != nil {
		return CalculateExchangePostPollerResponse{}, err
	}
	result := CalculateExchangePostPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CalculateExchangeClient.Post", "azure-async-operation", resp, client.pl, client.postHandleError)
	if err != nil {
		return CalculateExchangePostPollerResponse{}, err
	}
	result.Poller = &CalculateExchangePostPoller{
		pt: pt,
	}
	return result, nil
}

// Post - Calculates price for exchanging Reservations if there are no policy errors.
// If the operation fails it returns the *Error error type.
func (client *CalculateExchangeClient) post(ctx context.Context, body CalculateExchangeRequest, options *CalculateExchangeBeginPostOptions) (*http.Response, error) {
	req, err := client.postCreateRequest(ctx, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.postHandleError(resp)
	}
	return resp, nil
}

// postCreateRequest creates the Post request.
func (client *CalculateExchangeClient) postCreateRequest(ctx context.Context, body CalculateExchangeRequest, options *CalculateExchangeBeginPostOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/calculateExchange"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// postHandleError handles the Post error response.
func (client *CalculateExchangeClient) postHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
