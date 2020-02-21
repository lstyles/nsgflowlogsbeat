package insights

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// TenantActivityLogsClient is the monitor Management Client
type TenantActivityLogsClient struct {
	BaseClient
}

// NewTenantActivityLogsClient creates an instance of the TenantActivityLogsClient client.
func NewTenantActivityLogsClient(subscriptionID string) TenantActivityLogsClient {
	return NewTenantActivityLogsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTenantActivityLogsClientWithBaseURI creates an instance of the TenantActivityLogsClient client.
func NewTenantActivityLogsClientWithBaseURI(baseURI string, subscriptionID string) TenantActivityLogsClient {
	return TenantActivityLogsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets the Activity Logs for the Tenant.<br>Everything that is applicable to the API to get the Activity Logs for
// the subscription is applicable to this API (the parameters, $filter, etc.).<br>One thing to point out here is that
// this API does *not* retrieve the logs at the individual subscription of the tenant but only surfaces the logs that
// were generated at the tenant level.
// Parameters:
// filter - reduces the set of data collected. <br>The **$filter** is very restricted and allows only the
// following patterns.<br>- List events for a resource group: $filter=eventTimestamp ge '<Start Time>' and
// eventTimestamp le '<End Time>' and eventChannels eq 'Admin, Operation' and resourceGroupName eq
// '<ResourceGroupName>'.<br>- List events for resource: $filter=eventTimestamp ge '<Start Time>' and
// eventTimestamp le '<End Time>' and eventChannels eq 'Admin, Operation' and resourceUri eq
// '<ResourceURI>'.<br>- List events for a subscription: $filter=eventTimestamp ge '<Start Time>' and
// eventTimestamp le '<End Time>' and eventChannels eq 'Admin, Operation'.<br>- List events for a resource
// provider: $filter=eventTimestamp ge '<Start Time>' and eventTimestamp le '<End Time>' and eventChannels eq
// 'Admin, Operation' and resourceProvider eq '<ResourceProviderName>'.<br>- List events for a correlation Id:
// api-version=2014-04-01&$filter=eventTimestamp ge '2014-07-16T04:36:37.6407898Z' and eventTimestamp le
// '2014-07-20T04:36:37.6407898Z' and eventChannels eq 'Admin, Operation' and correlationId eq
// '<CorrelationID>'.<br>**NOTE**: No other syntax is allowed.
// selectParameter - used to fetch events with only the given properties.<br>The **$select** argument is a
// comma separated list of property names to be returned. Possible values are: *authorization*, *claims*,
// *correlationId*, *description*, *eventDataId*, *eventName*, *eventTimestamp*, *httpRequest*, *level*,
// *operationId*, *operationName*, *properties*, *resourceGroupName*, *resourceProviderName*, *resourceId*,
// *status*, *submissionTimestamp*, *subStatus*, *subscriptionId*
func (client TenantActivityLogsClient) List(ctx context.Context, filter string, selectParameter string) (result EventDataCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantActivityLogsClient.List")
		defer func() {
			sc := -1
			if result.edc.Response.Response != nil {
				sc = result.edc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, filter, selectParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.TenantActivityLogsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.edc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.TenantActivityLogsClient", "List", resp, "Failure sending request")
		return
	}

	result.edc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.TenantActivityLogsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client TenantActivityLogsClient) ListPreparer(ctx context.Context, filter string, selectParameter string) (*http.Request, error) {
	const APIVersion = "2015-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/microsoft.insights/eventtypes/management/values"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client TenantActivityLogsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client TenantActivityLogsClient) ListResponder(resp *http.Response) (result EventDataCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client TenantActivityLogsClient) listNextResults(ctx context.Context, lastResults EventDataCollection) (result EventDataCollection, err error) {
	req, err := lastResults.eventDataCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "insights.TenantActivityLogsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "insights.TenantActivityLogsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.TenantActivityLogsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client TenantActivityLogsClient) ListComplete(ctx context.Context, filter string, selectParameter string) (result EventDataCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantActivityLogsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, filter, selectParameter)
	return
}
