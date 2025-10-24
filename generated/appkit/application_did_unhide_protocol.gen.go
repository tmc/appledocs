// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// applicationDidUnhideProtocol is the applicationDidUnhide: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to applicationDidUnhide:.
var applicationDidUnhideProtocol *objc.Protocol

func init() {
	applicationDidUnhideProtocol = objc.GetProtocol("applicationDidUnhide:")
}

