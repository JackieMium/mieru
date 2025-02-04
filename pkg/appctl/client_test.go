// Copyright (C) 2021  mieru authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package appctl

import (
	"os"
	"testing"

	"github.com/enfein/mieru/pkg/appctl/appctlpb"
	"google.golang.org/protobuf/proto"
)

func TestApply2ClientConfig(t *testing.T) {
	beforeClientTest(t)

	// Apply config1, and then apply config2.
	configFile1 := "testdata/client_apply_config_1.json"
	if err := ApplyJSONClientConfig(configFile1); err != nil {
		t.Errorf("ApplyJSONClientConfig(%q) failed: %v", configFile1, err)
	}
	configFile2 := "testdata/client_apply_config_2.json"
	if err := ApplyJSONClientConfig(configFile2); err != nil {
		t.Errorf("ApplyJSONClientConfig(%q) failed: %v", configFile2, err)
	}
	merged, err := LoadClientConfig()
	if err != nil {
		t.Errorf("LoadClientConfig() failed: %v", err)
	}

	// Apply only config2. The client config should be the same.
	if err := deleteClientConfigFile(); err != nil {
		t.Fatalf("failed to delete client config file")
	}
	if err := StoreClientConfig(&appctlpb.ClientConfig{}); err != nil {
		t.Fatalf("failed to create empty client config file")
	}
	if err := ApplyJSONClientConfig(configFile2); err != nil {
		t.Errorf("ApplyJSONClientConfig(%q) failed: %v", configFile2, err)
	}
	want, err := LoadClientConfig()
	if err != nil {
		t.Errorf("LoadClientConfig() failed: %v", err)
	}
	if !proto.Equal(merged, want) {
		mergedJSON, _ := jsonMarshalOption.Marshal(merged)
		wantJSON, _ := jsonMarshalOption.Marshal(want)
		t.Errorf("client config doesn't equal:\ngot = %v\nwant = %v", string(mergedJSON), string(wantJSON))
	}

	afterClientTest(t)
}

func TestClientApplyReject(t *testing.T) {
	cases := []string{
		"testdata/client_reject_active_profile_mismatch.json",
		"testdata/client_reject_multiple_port_bindings.json",
		"testdata/client_reject_multiple_servers.json",
		"testdata/client_reject_no_active_profile.json",
		"testdata/client_reject_no_password.json",
		"testdata/client_reject_no_port_binding.json",
		"testdata/client_reject_no_port.json",
		"testdata/client_reject_no_profile_name.json",
		"testdata/client_reject_no_protocol.json",
		"testdata/client_reject_no_rpc_port.json",
		"testdata/client_reject_no_server_addr.json",
		"testdata/client_reject_no_socks5_port.json",
		"testdata/client_reject_no_user_name.json",
	}
	for _, c := range cases {
		t.Run(c, func(t *testing.T) {
			beforeClientTest(t)
			if err := ApplyJSONClientConfig(c); err == nil {
				t.Errorf("want error in ApplyJSONClientConfig(%q), got no error", c)
			}
			afterClientTest(t)
		})
	}
}

func TestClientDeleteProfile(t *testing.T) {
	beforeClientTest(t)

	configFile := "testdata/client_before_delete_profile.json"
	if err := ApplyJSONClientConfig(configFile); err != nil {
		t.Fatalf("ApplyJSONClientConfig(%q) failed: %v", configFile, err)
	}
	if err := DeleteClientConfigProfile("default"); err != nil {
		t.Errorf("DeleteClientConfigProfile(%q) failed: %v", "default", err)
	}
	if err := DeleteClientConfigProfile("this profile doesn't exist"); err != nil {
		t.Errorf("DeleteClientConfigProfile(%q) failed: %v", "this profile doesn't exist", err)
	}
	got, err := LoadClientConfig()
	if err != nil {
		t.Errorf("LoadClientConfig() failed: %v", err)
	}

	// Compare the result with client_after_delete_profile.json
	wantFile := "testdata/client_after_delete_profile.json"
	if err := deleteClientConfigFile(); err != nil {
		t.Fatalf("failed to delete client config file")
	}
	if err := StoreClientConfig(&appctlpb.ClientConfig{}); err != nil {
		t.Fatalf("failed to create empty client config file")
	}
	if err := ApplyJSONClientConfig(wantFile); err != nil {
		t.Fatalf("ApplyJSONClientConfig(%q) failed: %v", wantFile, err)
	}
	want, err := LoadClientConfig()
	if err != nil {
		t.Errorf("LoadClientConfig() failed: %v", err)
	}
	if !proto.Equal(got, want) {
		gotJSON, _ := jsonMarshalOption.Marshal(got)
		wantJSON, _ := jsonMarshalOption.Marshal(want)
		t.Errorf("client config doesn't equal:\ngot = %v\nwant = %v", string(gotJSON), string(wantJSON))
	}

	afterClientTest(t)
}

func TestClientDeleteProfileRejectActiveProfile(t *testing.T) {
	beforeClientTest(t)

	configFile := "testdata/client_before_delete_profile.json"
	if err := ApplyJSONClientConfig(configFile); err != nil {
		t.Fatalf("ApplyJSONClientConfig(%q) failed: %v", configFile, err)
	}
	if err := DeleteClientConfigProfile("new"); err == nil {
		t.Errorf("want error in DeleteClientConfigProfile(%q), got no error", "new")
	}

	afterClientTest(t)
}

func beforeClientTest(t *testing.T) {
	dir := os.TempDir()
	if dir == "" {
		t.Fatalf("failed to get system temporary directory for the test")
	}
	cachedClientConfigDir = dir
	if err := deleteClientConfigFile(); err != nil {
		t.Fatalf("failed to clean client config file before the test")
	}
	if err := StoreClientConfig(&appctlpb.ClientConfig{}); err != nil {
		t.Fatalf("failed to create empty client config file before the test")
	}
}

func afterClientTest(t *testing.T) {
	if err := deleteClientConfigFile(); err != nil {
		t.Fatalf("failed to clean client config file after the test")
	}
}
