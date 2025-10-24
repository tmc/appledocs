// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewCollectionLayoutAnchor

// ExampleNewCollectionLayoutAnchorWithEdges demonstrates how to create a CollectionLayoutAnchor instance using NewCollectionLayoutAnchorWithEdges.
// Creates an anchor with the specified edges to attach to.
func ExampleNewCollectionLayoutAnchorWithEdges() {
	_ = appkit.NewCollectionLayoutAnchorWithEdges(
		appkit.DirectionalRectEdge{}, // edges DirectionalRectEdge
	)
	// Output:
}
