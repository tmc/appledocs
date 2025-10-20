// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// applicationProtocol is the application: protocol.
//
// Availability:
//   - macOS 10.7+
//
// Use this protocol when registering custom classes that conform to application:.
var applicationProtocol *objc.Protocol

func init() {
	applicationProtocol = objc.GetProtocol("application:")
}


