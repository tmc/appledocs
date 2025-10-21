// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import "github.com/ebitengine/purego/objc"

// compositionURLProtocol is the compositionURL protocol.
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.15)
//
// Use this protocol when registering custom classes that conform to compositionURL.
var compositionURLProtocol *objc.Protocol

func init() {
	compositionURLProtocol = objc.GetProtocol("compositionURL")
}
