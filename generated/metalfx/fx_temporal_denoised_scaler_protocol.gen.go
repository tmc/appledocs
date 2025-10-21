// Code generated from Apple documentation for MetalFX. DO NOT EDIT.

package metalfx

import "github.com/ebitengine/purego/objc"

// FXTemporalDenoisedScalerProtocol is the MTLFXTemporalDenoisedScaler protocol.
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 26.0+
//   - tvOS 18.0+
//
// Use this protocol when registering custom classes that conform to MTLFXTemporalDenoisedScaler.
var FXTemporalDenoisedScalerProtocol *objc.Protocol

func init() {
	FXTemporalDenoisedScalerProtocol = objc.GetProtocol("MTLFXTemporalDenoisedScaler")
}
