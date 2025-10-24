// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// inputClientBecomeActiveProtocol is the inputClientBecomeActive: protocol.
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Use this protocol when registering custom classes that conform to inputClientBecomeActive:.
var inputClientBecomeActiveProtocol *objc.Protocol

func init() {
	inputClientBecomeActiveProtocol = objc.GetProtocol("inputClientBecomeActive:")
}
