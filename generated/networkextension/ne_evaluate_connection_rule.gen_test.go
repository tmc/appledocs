// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension_test

import (
	"github.com/tmc/appledocs/generated/networkextension"
)

// Suppress unused import errors
var _ = networkextension.NewNEEvaluateConnectionRule

// ExampleNewNEEvaluateConnectionRuleWithMatchDomainsAndAction demonstrates how to create a NEEvaluateConnectionRule instance using NewNEEvaluateConnectionRuleWithMatchDomainsAndAction.
// Initialize an   instance with a list of destination host domains and an action.
func ExampleNewNEEvaluateConnectionRuleWithMatchDomainsAndAction() {
	_ = networkextension.NewNEEvaluateConnectionRuleWithMatchDomainsAndAction(
		[]networkextension.string{}, // domains []string
		networkextension.NEEvaluateConnectionRuleAction{}, // action NEEvaluateConnectionRuleAction
	)
	// Output:
}
