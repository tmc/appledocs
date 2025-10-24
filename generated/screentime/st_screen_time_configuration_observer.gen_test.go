// Code generated from Apple documentation for ScreenTime. DO NOT EDIT.

package screentime_test

import (
	"github.com/tmc/appledocs/generated/screentime"
)

// Suppress unused import errors
var _ = screentime.NewSTScreenTimeConfigurationObserver

// ExampleSTScreenTimeConfigurationObserver_StartObserving demonstrates using StartObserving on a STScreenTimeConfigurationObserver instance.
// Starts observing changes to the current configuration.
func ExampleSTScreenTimeConfigurationObserver_StartObserving() {
	obj := screentime.NewSTScreenTimeConfigurationObserver()
	obj.StartObserving()
	// Output:
	}

