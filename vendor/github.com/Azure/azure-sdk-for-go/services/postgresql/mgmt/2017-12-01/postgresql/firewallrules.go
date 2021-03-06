package postgresql

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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// FirewallRulesClient is the the Microsoft Azure management API provides create, read, update, and delete
// functionality for Azure PostgreSQL resources including servers, databases, firewall rules, log files and
// configurations with new business model.
type FirewallRulesClient struct {
	BaseClient
}

// NewFirewallRulesClient creates an instance of the FirewallRulesClient client.
func NewFirewallRulesClient(subscriptionID string) FirewallRulesClient {
	return NewFirewallRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFirewallRulesClientWithBaseURI creates an instance of the FirewallRulesClient client.
func NewFirewallRulesClientWithBaseURI(baseURI string, subscriptionID string) FirewallRulesClient {
	return FirewallRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new firewall rule or updates an existing firewall rule.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// firewallRuleName - the name of the server firewall rule.
// parameters - the required parameters for creating or updating a firewall rule.
func (client FirewallRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters FirewallRule) (result FirewallRulesCreateOrUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.FirewallRuleProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.FirewallRuleProperties.StartIPAddress", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.FirewallRuleProperties.StartIPAddress", Name: validation.Pattern, Rule: `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`, Chain: nil}}},
					{Target: "parameters.FirewallRuleProperties.EndIPAddress", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.FirewallRuleProperties.EndIPAddress", Name: validation.Pattern, Rule: `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("postgresql.FirewallRulesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, firewallRuleName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresql.FirewallRulesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresql.FirewallRulesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client FirewallRulesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters FirewallRule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"firewallRuleName":  autorest.Encode("path", firewallRuleName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/firewallRules/{firewallRuleName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallRulesClient) CreateOrUpdateSender(req *http.Request) (future FirewallRulesCreateOrUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted))
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client FirewallRulesClient) CreateOrUpdateResponder(resp *http.Response) (result FirewallRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a server firewall rule.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// firewallRuleName - the name of the server firewall rule.
func (client FirewallRulesClient) Delete(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (result FirewallRulesDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, firewallRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresql.FirewallRulesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresql.FirewallRulesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client FirewallRulesClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"firewallRuleName":  autorest.Encode("path", firewallRuleName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/firewallRules/{firewallRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallRulesClient) DeleteSender(req *http.Request) (future FirewallRulesDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client FirewallRulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about a server firewall rule.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// firewallRuleName - the name of the server firewall rule.
func (client FirewallRulesClient) Get(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (result FirewallRule, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, firewallRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresql.FirewallRulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "postgresql.FirewallRulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresql.FirewallRulesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client FirewallRulesClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"firewallRuleName":  autorest.Encode("path", firewallRuleName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/firewallRules/{firewallRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallRulesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FirewallRulesClient) GetResponder(resp *http.Response) (result FirewallRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer list all the firewall rules in a given server.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
func (client FirewallRulesClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result FirewallRuleListResult, err error) {
	req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresql.FirewallRulesClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "postgresql.FirewallRulesClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresql.FirewallRulesClient", "ListByServer", resp, "Failure responding to request")
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client FirewallRulesClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/firewallRules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallRulesClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client FirewallRulesClient) ListByServerResponder(resp *http.Response) (result FirewallRuleListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
