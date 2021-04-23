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

// [START area120tables_generated_area120_tables_apiv1alpha1_Client_BatchCreateRows]

package main

import (
	"context"

	tables "cloud.google.com/go/area120/tables/apiv1alpha1"
	tablespb "google.golang.org/genproto/googleapis/area120/tables/v1alpha1"
)

func main() {
	// import tablespb "google.golang.org/genproto/googleapis/area120/tables/v1alpha1"

	ctx := context.Background()
	c, err := tables.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &tablespb.BatchCreateRowsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.BatchCreateRows(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END area120tables_generated_area120_tables_apiv1alpha1_Client_BatchCreateRows]