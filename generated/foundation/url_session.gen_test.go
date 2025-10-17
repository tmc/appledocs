// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)


// ExampleNewURLSession demonstrates how to create a URLSession instance.
func ExampleNewURLSession() {
	_ = foundation.NewURLSession()
	// Output:
}

// ExampleNewSessionWithConfiguration demonstrates how to create a URLSession instance using NewSessionWithConfiguration.
// Creates a session with the specified session configuration.
func ExampleNewSessionWithConfiguration() {
	_ = foundation.NewSessionWithConfiguration(
		nil, // configuration unsafe.Pointer
	)
	// Output:
}

// ExampleNewSessionWithConfigurationDelegateDelegateQueue demonstrates how to create a URLSession instance using NewSessionWithConfigurationDelegateDelegateQueue.
// Creates a session with the specified session configuration, delegate, and operation queue.
func ExampleNewSessionWithConfigurationDelegateDelegateQueue() {
	_ = foundation.NewSessionWithConfigurationDelegateDelegateQueue(
		nil, // configuration unsafe.Pointer
		nil, // delegate unsafe.Pointer
		nil, // queue unsafe.Pointer
	)
	// Output:
}


