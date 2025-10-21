// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension_test

import (
	"github.com/tmc/appledocs/generated/networkextension"
)

// Suppress unused import errors
var _ = networkextension.NewNEAppRule

// ExampleNewNEAppRuleWithSigningIdentifier demonstrates how to create a NEAppRule instance using NewNEAppRuleWithSigningIdentifier.
// Create an app rule that matches an app with a given signing identifier.
func ExampleNewNEAppRuleWithSigningIdentifier() {
	_ = networkextension.NewNEAppRuleWithSigningIdentifier(
		"signingIdentifier", // signingIdentifier string
	)
	// Output:
}
// ExampleNewNEAppRuleWithSigningIdentifierDesignatedRequirement demonstrates how to create a NEAppRule instance using NewNEAppRuleWithSigningIdentifierDesignatedRequirement.
// Create an app rule that matches an app with a given signing identifier and a given designated requirement.
func ExampleNewNEAppRuleWithSigningIdentifierDesignatedRequirement() {
	_ = networkextension.NewNEAppRuleWithSigningIdentifierDesignatedRequirement(
		"signingIdentifier", // signingIdentifier string
		"designatedRequirement", // designatedRequirement string
	)
	// Output:
}
