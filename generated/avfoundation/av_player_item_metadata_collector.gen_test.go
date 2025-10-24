// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewPlayerItemMetadataCollector

// ExampleNewPlayerItemMetadataCollectorWithIdentifiersClassifyingLabels demonstrates how to create a PlayerItemMetadataCollector instance using NewPlayerItemMetadataCollectorWithIdentifiersClassifyingLabels.
// Creates a metadata collector to access a stream’s metadata groups matching the specified array of identifiers and classifying labels.
func ExampleNewPlayerItemMetadataCollectorWithIdentifiersClassifyingLabels() {
	_ = avfoundation.NewPlayerItemMetadataCollectorWithIdentifiersClassifyingLabels(
		[]avfoundation.string{}, // identifiers []string
		[]avfoundation.string{}, // classifyingLabels []string
	)
	// Output:
}
