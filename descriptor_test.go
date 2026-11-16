package remotesubscriptions

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestManifestCopyPreservesSourceDataIdentity(t *testing.T) {
	first := ManifestJSON()
	var value struct {
		ID, Delivery, Version string
		DefaultEnabled        bool
		Database              struct {
			Tables, Settings                []string
			SchemaVersion, MigrationVersion string
			CompatibilityRange              struct{ MinimumSchema, MaximumSchema string }
		}
	}
	if err := json.Unmarshal(first, &value); err != nil {
		t.Fatal(err)
	}
	if value.ID != "remote-outbound-subscriptions" || value.Delivery != "in-process" || !value.DefaultEnabled || value.Version != "1" || value.Database.SchemaVersion != "3" || value.Database.MigrationVersion != "3" {
		t.Fatalf("descriptor identity changed: %+v", value)
	}
	if !reflect.DeepEqual(value.Database.Tables, []string{"remote_outbound_subscriptions", "remote_outbound_groups", "remote_outbound_group_connections", "remote_outbound_connections"}) || !reflect.DeepEqual(value.Database.Settings, []string{"subRemoteConversionPolicy", "subRemoteGroupAdaptation"}) {
		t.Fatal("durable table or settings identities changed")
	}
	if value.Database.CompatibilityRange.MinimumSchema != "1" || value.Database.CompatibilityRange.MaximumSchema != "3" {
		t.Fatal("compatibility changed")
	}
	first[0] = 0
	if !json.Valid(ManifestJSON()) {
		t.Fatal("caller mutated the embedded descriptor")
	}
}
