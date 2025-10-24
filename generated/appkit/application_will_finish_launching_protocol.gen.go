// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// applicationWillFinishLaunchingProtocol is the applicationWillFinishLaunching: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to applicationWillFinishLaunching:.
var applicationWillFinishLaunchingProtocol *objc.Protocol

func init() {
	applicationWillFinishLaunchingProtocol = objc.GetProtocol("applicationWillFinishLaunching:")
}

