// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// applicationDidUpdateProtocol is the applicationDidUpdate: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to applicationDidUpdate:.
var applicationDidUpdateProtocol *objc.Protocol

func init() {
	applicationDidUpdateProtocol = objc.GetProtocol("applicationDidUpdate:")
}

