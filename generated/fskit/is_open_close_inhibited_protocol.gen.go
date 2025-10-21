// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// isOpenCloseInhibitedProtocol is the isOpenCloseInhibited protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to isOpenCloseInhibited.
var isOpenCloseInhibitedProtocol *objc.Protocol

func init() {
	isOpenCloseInhibitedProtocol = objc.GetProtocol("isOpenCloseInhibited")
}
