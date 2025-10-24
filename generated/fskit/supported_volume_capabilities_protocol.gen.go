// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// supportedVolumeCapabilitiesProtocol is the supportedVolumeCapabilities protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to supportedVolumeCapabilities.
var supportedVolumeCapabilitiesProtocol *objc.Protocol

func init() {
	supportedVolumeCapabilitiesProtocol = objc.GetProtocol("supportedVolumeCapabilities")
}

