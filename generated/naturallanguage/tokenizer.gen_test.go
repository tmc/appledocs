// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage_test

import (
	"github.com/tmc/appledocs/generated/naturallanguage"
)

// Suppress unused import errors
var _ = naturallanguage.NewTokenizer

// ExampleNewTokenizerWithUnit demonstrates how to create a Tokenizer instance using NewTokenizerWithUnit.
// Creates a tokenizer with the specified unit.
func ExampleNewTokenizerWithUnit() {
	_ = naturallanguage.NewTokenizerWithUnit(
		naturallanguage.TokenUnit{}, // unit TokenUnit
	)
	// Output:
}

