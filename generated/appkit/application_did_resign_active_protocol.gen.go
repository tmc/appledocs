// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// applicationDidResignActiveProtocol is the applicationDidResignActive: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to applicationDidResignActive:.
var applicationDidResignActiveProtocol *objc.Protocol

func init() {
	applicationDidResignActiveProtocol = objc.GetProtocol("applicationDidResignActive:")
}
