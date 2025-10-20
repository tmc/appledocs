// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// shouldPerformSegueWithIdentifierProtocol is the shouldPerformSegueWithIdentifier: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to shouldPerformSegueWithIdentifier:.
var shouldPerformSegueWithIdentifierProtocol *objc.Protocol

func init() {
	shouldPerformSegueWithIdentifierProtocol = objc.GetProtocol("shouldPerformSegueWithIdentifier:")
}
