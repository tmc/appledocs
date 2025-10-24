// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import "github.com/ebitengine/purego/objc"

// deviceAttestationProtocol is the deviceAttestation: protocol.
//
// Availability:
//   - Mac Catalyst 16.1+ (Deprecated in 16.4)
//   - iOS 16.1+ (Deprecated in 16.4)
//   - iPadOS 16.1+ (Deprecated in 16.4)
//   - macOS 13.0+ (Deprecated in 13.3)
//   - tvOS 16.1+ (Deprecated in 16.4)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 9.1+ (Deprecated in 9.4)
//
// Use this protocol when registering custom classes that conform to deviceAttestation:.
var deviceAttestationProtocol *objc.Protocol

func init() {
	deviceAttestationProtocol = objc.GetProtocol("deviceAttestation:")
}

