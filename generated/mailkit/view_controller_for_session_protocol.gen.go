// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import "github.com/ebitengine/purego/objc"

// viewControllerForSessionProtocol is the viewControllerForSession: protocol.
//
// Availability:
//   - macOS 12.0+
//
// Use this protocol when registering custom classes that conform to viewControllerForSession:.
var viewControllerForSessionProtocol *objc.Protocol

func init() {
	viewControllerForSessionProtocol = objc.GetProtocol("viewControllerForSession:")
}

