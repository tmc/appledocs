// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// speechSynthesizerProtocol is the speechSynthesizer: protocol.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 14.0)
//
// Use this protocol when registering custom classes that conform to speechSynthesizer:.
var speechSynthesizerProtocol *objc.Protocol

func init() {
	speechSynthesizerProtocol = objc.GetProtocol("speechSynthesizer:")
}
