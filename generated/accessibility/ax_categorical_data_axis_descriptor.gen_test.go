// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility_test

import (
	"github.com/tmc/appledocs/generated/accessibility"
)

// Suppress unused import errors
var _ = accessibility.NewAXCategoricalDataAxisDescriptor

// ExampleNewAXCategoricalDataAxisDescriptorWithTitleCategoryOrder demonstrates how to create a AXCategoricalDataAxisDescriptor instance using NewAXCategoricalDataAxisDescriptorWithTitleCategoryOrder.
// Creates a categorical data axis with the specified title and an array of categories   in the specified order.
func ExampleNewAXCategoricalDataAxisDescriptorWithTitleCategoryOrder() {
	_ = accessibility.NewAXCategoricalDataAxisDescriptorWithTitleCategoryOrder(
		"title", // title string
		[]accessibility.string{}, // categoryOrder []string
	)
	// Output:
}
