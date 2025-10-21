// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import "github.com/ebitengine/purego/objc"

// attributesProtocol is the attributes protocol.
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.15)
//
// Use this protocol when registering custom classes that conform to attributes.
var attributesProtocol *objc.Protocol

func init() {
	attributesProtocol = objc.GetProtocol("attributes")
}
