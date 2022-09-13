// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package budgets_test

import (
	"context"

	budgets "cloud.google.com/go/billing/budgets/apiv1beta1"
	"google.golang.org/api/iterator"
	budgetspb "google.golang.org/genproto/googleapis/cloud/billing/budgets/v1beta1"
)

func ExampleNewBudgetClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := budgets.NewBudgetClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleBudgetClient_CreateBudget() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := budgets.NewBudgetClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &budgetspb.CreateBudgetRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/billing/budgets/v1beta1#CreateBudgetRequest.
	}
	resp, err := c.CreateBudget(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleBudgetClient_UpdateBudget() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := budgets.NewBudgetClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &budgetspb.UpdateBudgetRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/billing/budgets/v1beta1#UpdateBudgetRequest.
	}
	resp, err := c.UpdateBudget(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleBudgetClient_GetBudget() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := budgets.NewBudgetClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &budgetspb.GetBudgetRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/billing/budgets/v1beta1#GetBudgetRequest.
	}
	resp, err := c.GetBudget(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleBudgetClient_ListBudgets() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := budgets.NewBudgetClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &budgetspb.ListBudgetsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/billing/budgets/v1beta1#ListBudgetsRequest.
	}
	it := c.ListBudgets(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleBudgetClient_DeleteBudget() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := budgets.NewBudgetClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &budgetspb.DeleteBudgetRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/billing/budgets/v1beta1#DeleteBudgetRequest.
	}
	err = c.DeleteBudget(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
