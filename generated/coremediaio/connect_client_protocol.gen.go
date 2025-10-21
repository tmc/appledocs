// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import "github.com/ebitengine/purego/objc"

// connectClientProtocol is the connectClient: protocol.
//
// Availability:
//   - Mac Catalyst 15.4+
//   - macOS 12.3+
//
// Use this protocol when registering custom classes that conform to connectClient:.
var connectClientProtocol *objc.Protocol

func init() {
	connectClientProtocol = objc.GetProtocol("connectClient:")
}
