// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import "github.com/ebitengine/purego/objc"

// isReadyForMoreMediaDataProtocol is the isReadyForMoreMediaData protocol.
//
// Availability:
//   - macOS 15.0+
//
// Use this protocol when registering custom classes that conform to isReadyForMoreMediaData.
var isReadyForMoreMediaDataProtocol *objc.Protocol

func init() {
	isReadyForMoreMediaDataProtocol = objc.GetProtocol("isReadyForMoreMediaData")
}

