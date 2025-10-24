// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension_test

import (
	"github.com/tmc/appledocs/generated/networkextension"
)

// Suppress unused import errors
var _ = networkextension.NewNEFilterSettings

// ExampleNewNEFilterSettingsWithRulesDefaultAction demonstrates how to create a NEFilterSettings instance using NewNEFilterSettingsWithRulesDefaultAction.
// Creates a new settings instance from an array of rules and a default action.
func ExampleNewNEFilterSettingsWithRulesDefaultAction() {
	_ = networkextension.NewNEFilterSettingsWithRulesDefaultAction(
		[]networkextension.NEFilterRule{}, // rules []NEFilterRule
		networkextension.NEFilterAction{}, // defaultAction NEFilterAction
	)
	// Output:
}
