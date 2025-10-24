// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewFontAssetRequest

// ExampleNewFontAssetRequestWithFontDescriptorsOptions demonstrates how to create a FontAssetRequest instance using NewFontAssetRequestWithFontDescriptorsOptions.
func ExampleNewFontAssetRequestWithFontDescriptorsOptions() {
	_ = appkit.NewFontAssetRequestWithFontDescriptorsOptions(
		[]appkit.FontDescriptor{}, // fontDescriptors []FontDescriptor
		appkit.FontAssetRequestOptions{}, // options FontAssetRequestOptions
	)
	// Output:
}
