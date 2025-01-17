/*
Copyright 2015-2021 Gravitational, Inc.

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

package test

import (
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/trace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func (s *TerraformSuite) TestApp() {
	checkDestroyed := func(state *terraform.State) error {
		_, err := s.client.GetApp(s.Context(), "test")
		if trace.IsNotFound(err) {
			return nil
		}

		return err
	}

	name := "teleport_app.test"

	resource.Test(s.T(), resource.TestCase{
		ProtoV6ProviderFactories: s.terraformProviders,
		CheckDestroy:             checkDestroyed,
		Steps: []resource.TestStep{
			{
				Config: s.getFixture("app_0_create.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "kind", "app"),
					resource.TestCheckResourceAttr(name, "spec.uri", "localhost:3000"),
				),
			},
			{
				Config:   s.getFixture("app_0_create.tf"),
				PlanOnly: true,
			},
			{
				Config: s.getFixture("app_1_update.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "kind", "app"),
					resource.TestCheckResourceAttr(name, "spec.uri", "localhost:3000"),
				),
			},
			{
				Config:   s.getFixture("app_1_update.tf"),
				PlanOnly: true,
			},
		},
	})
}

func (s *TerraformSuite) TestImportApp() {
	r := "teleport_app"
	id := "test_import"
	name := r + "." + id

	app := &types.AppV3{
		Metadata: types.Metadata{
			Name: id,
		},
		Spec: types.AppSpecV3{
			URI: "localhost:3000/test",
		},
	}
	err := app.CheckAndSetDefaults()
	require.NoError(s.T(), err)

	err = s.client.CreateApp(s.Context(), app)
	require.NoError(s.T(), err)

	resource.Test(s.T(), resource.TestCase{
		ProtoV6ProviderFactories: s.terraformProviders,
		Steps: []resource.TestStep{
			{
				Config:        s.terraformConfig + "\n" + `resource "` + r + `" "` + id + `" { }`,
				ResourceName:  name,
				ImportState:   true,
				ImportStateId: id,
				ImportStateCheck: func(state []*terraform.InstanceState) error {
					require.Equal(s.T(), state[0].Attributes["kind"], "app")
					require.Equal(s.T(), state[0].Attributes["spec.uri"], "localhost:3000/test")

					return nil
				},
			},
		},
	})
}
