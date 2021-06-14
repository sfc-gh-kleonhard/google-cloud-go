// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START assuredworkloads_v1beta1_generated_AssuredWorkloadsService_DeleteWorkload_sync]

package main

import (
	"context"

	assuredworkloads "cloud.google.com/go/assuredworkloads/apiv1beta1"
	assuredworkloadspb "google.golang.org/genproto/googleapis/cloud/assuredworkloads/v1beta1"
)

func main() {
	ctx := context.Background()
	c, err := assuredworkloads.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &assuredworkloadspb.DeleteWorkloadRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteWorkload(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END assuredworkloads_v1beta1_generated_AssuredWorkloadsService_DeleteWorkload_sync]