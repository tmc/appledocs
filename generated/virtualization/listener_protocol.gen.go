// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import "github.com/ebitengine/purego/objc"

// listenerProtocol is the listener: protocol.
//
// Availability:
//   - macOS 11.0+
//
// Use this protocol when registering custom classes that conform to listener:.
var listenerProtocol *objc.Protocol

func init() {
	listenerProtocol = objc.GetProtocol("listener:")
}

