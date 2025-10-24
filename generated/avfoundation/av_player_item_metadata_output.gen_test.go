// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewPlayerItemMetadataOutput

// ExampleNewPlayerItemMetadataOutputWithIdentifiers demonstrates how to create a PlayerItemMetadataOutput instance using NewPlayerItemMetadataOutputWithIdentifiers.
// Creates an instance of AVPlayerItemMetadataOutput.
func ExampleNewPlayerItemMetadataOutputWithIdentifiers() {
	_ = avfoundation.NewPlayerItemMetadataOutputWithIdentifiers(
		[]avfoundation.string{}, // identifiers []string
	)
	// Output:
}
