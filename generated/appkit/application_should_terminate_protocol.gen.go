// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// applicationShouldTerminateProtocol is the applicationShouldTerminate: protocol.
//
// Availability:
//   - macOS +
//
// Use this protocol when registering custom classes that conform to applicationShouldTerminate:.
var applicationShouldTerminateProtocol *objc.Protocol

func init() {
	applicationShouldTerminateProtocol = objc.GetProtocol("applicationShouldTerminate:")
}

