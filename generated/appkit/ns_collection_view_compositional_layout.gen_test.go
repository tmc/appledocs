// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewCollectionViewCompositionalLayout

// ExampleNewCollectionViewCompositionalLayoutWithSectionProvider demonstrates how to create a CollectionViewCompositionalLayout instance using NewCollectionViewCompositionalLayoutWithSectionProvider.
// Creates a compositional layout object with a section provider to supply the layout’s sections.
func ExampleNewCollectionViewCompositionalLayoutWithSectionProvider() {
	_ = appkit.NewCollectionViewCompositionalLayoutWithSectionProvider(
		appkit.CollectionViewCompositionalLayoutSectionProvider /* not a class type */{}, // sectionProvider CollectionViewCompositionalLayoutSectionProvider /* not a class type */
	)
	// Output:
}
