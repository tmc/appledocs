// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import "github.com/ebitengine/purego/objc"

// valueForInputKeyProtocol is the valueForInputKey: protocol.
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.15)
//
// Use this protocol when registering custom classes that conform to valueForInputKey:.
var valueForInputKeyProtocol *objc.Protocol

func init() {
	valueForInputKeyProtocol = objc.GetProtocol("valueForInputKey:")
}

