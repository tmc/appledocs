// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewRuleSystem

// ExampleNewRuleSystem demonstrates how to create a RuleSystem instance.
// Initializes a new, empty rule system.
func ExampleNewRuleSystem() {
	_ = gameplaykit.NewRuleSystem()
	// Output:
}
// ExampleRuleSystem_Evaluate demonstrates using Evaluate on a RuleSystem instance.
// Evaluates the rule system, executing the list of rules in its agenda.
func ExampleRuleSystem_Evaluate() {
	obj := gameplaykit.NewRuleSystem()
	obj.Evaluate()
	// Output:
	}

// ExampleRuleSystem_RemoveAllRules demonstrates using RemoveAllRules on a RuleSystem instance.
// Removes all rules from the system.
func ExampleRuleSystem_RemoveAllRules() {
	obj := gameplaykit.NewRuleSystem()
	obj.RemoveAllRules()
	// Output:
	}

// ExampleRuleSystem_Reset demonstrates using Reset on a RuleSystem instance.
// Returns the rule system to its original agenda and clears all facts.
func ExampleRuleSystem_Reset() {
	obj := gameplaykit.NewRuleSystem()
	obj.Reset()
	// Output:
	}

