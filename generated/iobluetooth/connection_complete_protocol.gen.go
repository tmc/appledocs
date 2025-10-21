// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import "github.com/ebitengine/purego/objc"

// connectionCompleteProtocol is the connectionComplete: protocol.
//
// Use this protocol when registering custom classes that conform to connectionComplete:.
var connectionCompleteProtocol *objc.Protocol

func init() {
	connectionCompleteProtocol = objc.GetProtocol("connectionComplete:")
}
