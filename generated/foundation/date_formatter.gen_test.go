// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewDateFormatter


// ExampleNewDateFormatterWithDateFormatAllowNaturalLanguage demonstrates how to create a DateFormatter instance using NewDateFormatterWithDateFormatAllowNaturalLanguage.
// Initializes and returns an   instance that uses the OS X 10.0 formatting behavior and the given date format string in its conversions.
func ExampleNewDateFormatterWithDateFormatAllowNaturalLanguage() {
	_ = foundation.NewDateFormatterWithDateFormatAllowNaturalLanguage(
		"format", // format string
		false, // flag bool
	)
	// Output:
}


