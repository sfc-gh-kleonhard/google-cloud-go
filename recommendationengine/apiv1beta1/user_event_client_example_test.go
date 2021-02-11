// Copyright 2021 Google LLC
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

package recommendationengine_test

import (
	"context"

	recommendationengine "cloud.google.com/go/recommendationengine/apiv1beta1"
	"google.golang.org/api/iterator"
	recommendationenginepb "google.golang.org/genproto/googleapis/cloud/recommendationengine/v1beta1"
)

func ExampleNewUserEventClient() {
	ctx := context.Background()
	c, err := recommendationengine.NewUserEventClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleUserEventClient_WriteUserEvent() {
	// import recommendationenginepb "google.golang.org/genproto/googleapis/cloud/recommendationengine/v1beta1"

	ctx := context.Background()
	c, err := recommendationengine.NewUserEventClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recommendationenginepb.WriteUserEventRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.WriteUserEvent(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleUserEventClient_CollectUserEvent() {
	// import recommendationenginepb "google.golang.org/genproto/googleapis/cloud/recommendationengine/v1beta1"

	ctx := context.Background()
	c, err := recommendationengine.NewUserEventClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recommendationenginepb.CollectUserEventRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CollectUserEvent(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleUserEventClient_ListUserEvents() {
	// import recommendationenginepb "google.golang.org/genproto/googleapis/cloud/recommendationengine/v1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := recommendationengine.NewUserEventClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recommendationenginepb.ListUserEventsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListUserEvents(ctx, req)
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

func ExampleUserEventClient_PurgeUserEvents() {
	// import recommendationenginepb "google.golang.org/genproto/googleapis/cloud/recommendationengine/v1beta1"

	ctx := context.Background()
	c, err := recommendationengine.NewUserEventClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recommendationenginepb.PurgeUserEventsRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.PurgeUserEvents(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleUserEventClient_ImportUserEvents() {
	// import recommendationenginepb "google.golang.org/genproto/googleapis/cloud/recommendationengine/v1beta1"

	ctx := context.Background()
	c, err := recommendationengine.NewUserEventClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recommendationenginepb.ImportUserEventsRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ImportUserEvents(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}