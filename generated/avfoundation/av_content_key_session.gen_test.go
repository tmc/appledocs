// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewContentKeySession

// ExampleNewContentKeySessionWithKeySystem demonstrates how to create a ContentKeySession instance using NewContentKeySessionWithKeySystem.
// Creates a content key session to manage a collection of content decryption keys.
func ExampleNewContentKeySessionWithKeySystem() {
	_ = avfoundation.NewContentKeySessionWithKeySystem(
		avfoundation.ContentKeySystem{}, // keySystem ContentKeySystem
	)
	// Output:
}
// ExampleContentKeySession_Expire demonstrates using Expire on a ContentKeySession instance.
// Tells the delegate that the session expired as the result of normal, intentional processes.
func ExampleContentKeySession_Expire() {
	obj := avfoundation.NewContentKeySession()
	obj.Expire()
	// Output:
	}

