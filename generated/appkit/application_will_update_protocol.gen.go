// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// applicationWillUpdateProtocol is the applicationWillUpdate: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to applicationWillUpdate:.
var applicationWillUpdateProtocol *objc.Protocol

func init() {
	applicationWillUpdateProtocol = objc.GetProtocol("applicationWillUpdate:")
}
