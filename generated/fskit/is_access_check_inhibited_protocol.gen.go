// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// isAccessCheckInhibitedProtocol is the isAccessCheckInhibited protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to isAccessCheckInhibited.
var isAccessCheckInhibitedProtocol *objc.Protocol

func init() {
	isAccessCheckInhibitedProtocol = objc.GetProtocol("isAccessCheckInhibited")
}

