// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import "github.com/ebitengine/purego/objc"

// setImageProtocol is the setImage: protocol.
//
// Availability:
//   - macOS 10.4+
//
// Use this protocol when registering custom classes that conform to setImage:.
var setImageProtocol *objc.Protocol

func init() {
	setImageProtocol = objc.GetProtocol("setImage:")
}

