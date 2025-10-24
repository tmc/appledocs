// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// applicationDidBecomeActiveProtocol is the applicationDidBecomeActive: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to applicationDidBecomeActive:.
var applicationDidBecomeActiveProtocol *objc.Protocol

func init() {
	applicationDidBecomeActiveProtocol = objc.GetProtocol("applicationDidBecomeActive:")
}

