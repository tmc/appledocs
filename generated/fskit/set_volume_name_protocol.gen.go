// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// setVolumeNameProtocol is the setVolumeName: protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to setVolumeName:.
var setVolumeNameProtocol *objc.Protocol

func init() {
	setVolumeNameProtocol = objc.GetProtocol("setVolumeName:")
}

