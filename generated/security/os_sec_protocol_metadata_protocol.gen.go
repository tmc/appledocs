// Code generated from Apple documentation for Security. DO NOT EDIT.

package security

import "github.com/ebitengine/purego/objc"

// OS_sec_protocol_metadataProtocol is the OS_sec_protocol_metadata protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// Use this protocol when registering custom classes that conform to OS_sec_protocol_metadata.
var OS_sec_protocol_metadataProtocol *objc.Protocol

func init() {
	OS_sec_protocol_metadataProtocol = objc.GetProtocol("OS_sec_protocol_metadata")
}
