// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// windowWillMiniaturizeProtocol is the windowWillMiniaturize: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to windowWillMiniaturize:.
var windowWillMiniaturizeProtocol *objc.Protocol

func init() {
	windowWillMiniaturizeProtocol = objc.GetProtocol("windowWillMiniaturize:")
}
