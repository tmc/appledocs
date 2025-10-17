// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)


// ExampleNewErrorWithDomainCodeUserInfo demonstrates how to create a Error instance using NewErrorWithDomainCodeUserInfo.
// Returns an   object initialized for a given domain and code with a given   dictionary.
func ExampleNewErrorWithDomainCodeUserInfo() {
	_ = foundation.NewErrorWithDomainCodeUserInfo(
		nil, // domain unsafe.Pointer
		0, // code int
		nil, // dict unsafe.Pointer
	)
	// Output:
}


