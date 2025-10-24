// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import "github.com/ebitengine/purego/objc"

// setValueProtocol is the setValue: protocol.
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.15)
//
// Use this protocol when registering custom classes that conform to setValue:.
var setValueProtocol *objc.Protocol

func init() {
	setValueProtocol = objc.GetProtocol("setValue:")
}

