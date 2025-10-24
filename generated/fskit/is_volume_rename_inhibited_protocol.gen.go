// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// isVolumeRenameInhibitedProtocol is the isVolumeRenameInhibited protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to isVolumeRenameInhibited.
var isVolumeRenameInhibitedProtocol *objc.Protocol

func init() {
	isVolumeRenameInhibitedProtocol = objc.GetProtocol("isVolumeRenameInhibited")
}

