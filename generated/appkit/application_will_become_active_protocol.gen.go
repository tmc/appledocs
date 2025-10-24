// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// applicationWillBecomeActiveProtocol is the applicationWillBecomeActive: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to applicationWillBecomeActive:.
var applicationWillBecomeActiveProtocol *objc.Protocol

func init() {
	applicationWillBecomeActiveProtocol = objc.GetProtocol("applicationWillBecomeActive:")
}

