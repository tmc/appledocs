// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import "github.com/ebitengine/purego/objc"

// handsFreeProtocol is the handsFree: protocol.
//
// Availability:
//   - macOS 10.7+
//
// Use this protocol when registering custom classes that conform to handsFree:.
var handsFreeProtocol *objc.Protocol

func init() {
	handsFreeProtocol = objc.GetProtocol("handsFree:")
}
