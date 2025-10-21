// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import "github.com/ebitengine/purego/objc"

// rfcommChannelClosedProtocol is the rfcommChannelClosed: protocol.
//
// Use this protocol when registering custom classes that conform to rfcommChannelClosed:.
var rfcommChannelClosedProtocol *objc.Protocol

func init() {
	rfcommChannelClosedProtocol = objc.GetProtocol("rfcommChannelClosed:")
}
