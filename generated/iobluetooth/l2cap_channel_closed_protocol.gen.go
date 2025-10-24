// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import "github.com/ebitengine/purego/objc"

// l2capChannelClosedProtocol is the l2capChannelClosed: protocol.
//
// Availability:
//   - macOS +
//
// Use this protocol when registering custom classes that conform to l2capChannelClosed:.
var l2capChannelClosedProtocol *objc.Protocol

func init() {
	l2capChannelClosedProtocol = objc.GetProtocol("l2capChannelClosed:")
}

