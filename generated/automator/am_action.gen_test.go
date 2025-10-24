// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator_test

import (
	"github.com/tmc/appledocs/generated/automator"
)

// Suppress unused import errors
var _ = automator.NewAMAction

// ExampleAMAction_Activated demonstrates using Activated on a AMAction instance.
// Allows the action to synchronize its information with settings in another app.
func ExampleAMAction_Activated() {
	obj := automator.NewAMAction()
	obj.Activated()
	// Output:
	}

// ExampleAMAction_Closed demonstrates using Closed on a AMAction instance.
// Invoked by Automator when the receiving action is removed from a workflow, allowing it to perform cleanup operations.
func ExampleAMAction_Closed() {
	obj := automator.NewAMAction()
	obj.Closed()
	// Output:
	}

// ExampleAMAction_Opened demonstrates using Opened on a AMAction instance.
// Allows the action to initialize its user interface.
func ExampleAMAction_Opened() {
	obj := automator.NewAMAction()
	obj.Opened()
	// Output:
	}

// ExampleAMAction_ParametersUpdated demonstrates using ParametersUpdated on a AMAction instance.
// Requests the action to update its user interface from its stored parameters, which have changed.
func ExampleAMAction_ParametersUpdated() {
	obj := automator.NewAMAction()
	obj.ParametersUpdated()
	// Output:
	}

// ExampleAMAction_Reset demonstrates using Reset on a AMAction instance.
// Resets the action to its initial state.
func ExampleAMAction_Reset() {
	obj := automator.NewAMAction()
	obj.Reset()
	// Output:
	}

// ExampleAMAction_Stop demonstrates using Stop on a AMAction instance.
// Stops the action from running.
func ExampleAMAction_Stop() {
	obj := automator.NewAMAction()
	obj.Stop()
	// Output:
	}

// ExampleAMAction_UpdateParameters demonstrates using UpdateParameters on a AMAction instance.
// Requests the action to update its stored set of parameters from the settings in the action’s user interface.
func ExampleAMAction_UpdateParameters() {
	obj := automator.NewAMAction()
	obj.UpdateParameters()
	// Output:
	}

// ExampleAMAction_WillFinishRunning demonstrates using WillFinishRunning on a AMAction instance.
// Provides an opportunity for an action to perform cleanup operations, such as closing windows and deallocating memory.
func ExampleAMAction_WillFinishRunning() {
	obj := automator.NewAMAction()
	obj.WillFinishRunning()
	// Output:
	}

