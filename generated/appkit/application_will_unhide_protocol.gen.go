// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// applicationWillUnhideProtocol is the applicationWillUnhide: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to applicationWillUnhide:.
var applicationWillUnhideProtocol *objc.Protocol

func init() {
	applicationWillUnhideProtocol = objc.GetProtocol("applicationWillUnhide:")
}

