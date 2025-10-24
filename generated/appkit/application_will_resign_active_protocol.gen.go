// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// applicationWillResignActiveProtocol is the applicationWillResignActive: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to applicationWillResignActive:.
var applicationWillResignActiveProtocol *objc.Protocol

func init() {
	applicationWillResignActiveProtocol = objc.GetProtocol("applicationWillResignActive:")
}

