// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets_test

import (
	"github.com/tmc/appledocs/generated/backgroundassets"
)

// Suppress unused import errors
var _ = backgroundassets.NewBAAssetPackManifest

// ExampleBAAssetPackManifest_AllDownloads demonstrates using AllDownloads on a BAAssetPackManifest instance.
// Creates download objects for every asset pack in this manifest.
func ExampleBAAssetPackManifest_AllDownloads() {
	obj := backgroundassets.NewBAAssetPackManifest()
	_ = obj.AllDownloads()
	// Output:
	}

