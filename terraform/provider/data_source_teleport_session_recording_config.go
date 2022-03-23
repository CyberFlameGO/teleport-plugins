// Code generated by _gen/main.go DO NOT EDIT
/*
Copyright 2015-2022 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/gravitational/teleport-plugins/terraform/tfschema"
	apitypes "github.com/gravitational/teleport/api/types"
	"github.com/gravitational/trace"
)

// dataSourceTeleportSessionRecordingConfigType is the data source metadata type
type dataSourceTeleportSessionRecordingConfigType struct{}

// dataSourceTeleportSessionRecordingConfig is the resource
type dataSourceTeleportSessionRecordingConfig struct {
	p Provider
}

// GetSchema returns the data source schema
func (r dataSourceTeleportSessionRecordingConfigType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfschema.GenSchemaSessionRecordingConfigV2(ctx)
}

// NewDataSource creates the empty data source
func (r dataSourceTeleportSessionRecordingConfigType) NewDataSource(_ context.Context, p tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	return dataSourceTeleportSessionRecordingConfig{
		p: *(p.(*Provider)),
	}, nil
}

// Read reads teleport SessionRecordingConfig
func (r dataSourceTeleportSessionRecordingConfig) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	sessionRecordingConfigI, err := r.p.Client.GetSessionRecordingConfig(ctx)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading SessionRecordingConfig", trace.Wrap(err), "session_recording_config"))
		return
	}

    var state types.Object
	sessionRecordingConfig := sessionRecordingConfigI.(*apitypes.SessionRecordingConfigV2)
	diags := tfschema.CopySessionRecordingConfigV2ToTerraform(ctx, *sessionRecordingConfig, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}