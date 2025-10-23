// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewOrthography

// ExampleNewOrthographyWithCoder demonstrates how to create a Orthography instance using NewOrthographyWithCoder.
func ExampleNewOrthographyWithCoder() {
	_ = foundation.NewOrthographyWithCoder(
		foundation.Coder{}, // coder Coder
	)
	// Output:
}
// ExampleNewOrthographyWithDominantScriptLanguageMap demonstrates how to create a Orthography instance using NewOrthographyWithDominantScriptLanguageMap.
// Creates an orthography object with the specified dominant script and language map.
func ExampleNewOrthographyWithDominantScriptLanguageMap() {
	_ = foundation.NewOrthographyWithDominantScriptLanguageMap(
		"script", // script string
		foundation.IDictionary{}, // map IDictionary
	)
	// Output:
}
