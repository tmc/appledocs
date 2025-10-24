// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMDLLightProbeIrradianceDataSource is the MDLLightProbeIrradianceDataSource protocol interface.
//
// Adopt this protocol to provide information for use in automatic placement of light probes around a scene.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.modelio/documentation/ModelIO/MDLLightProbeIrradianceDataSource
type PMDLLightProbeIrradianceDataSource interface {
	// Optional methods
	SphericalHarmonicsCoefficientsAtPosition(position unsafe.Pointer) foundation.Data
	HasSphericalHarmonicsCoefficientsAtPosition() bool
}

// MDLLightProbeIrradianceDataSource is a delegate implementation builder for the PMDLLightProbeIrradianceDataSource protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MDLLightProbeIrradianceDataSource struct {
	_SphericalHarmonicsCoefficientsAtPosition func(position unsafe.Pointer) foundation.Data
}

// SetSphericalHarmonicsCoefficientsAtPosition sets the handler for the SphericalHarmonicsCoefficientsAtPosition delegate method.
//
// Asks the data source to provide spherical harmonics coefficients that describe lighting conditions in all directions from the specified point in a scene.
func (d *MDLLightProbeIrradianceDataSource) SetSphericalHarmonicsCoefficientsAtPosition(f func(position unsafe.Pointer) foundation.Data) {
	d._SphericalHarmonicsCoefficientsAtPosition = f
}

// SphericalHarmonicsCoefficientsAtPosition implements the PMDLLightProbeIrradianceDataSource interface.
func (d *MDLLightProbeIrradianceDataSource) SphericalHarmonicsCoefficientsAtPosition(position unsafe.Pointer) foundation.Data {
	if d._SphericalHarmonicsCoefficientsAtPosition != nil {
		return d._SphericalHarmonicsCoefficientsAtPosition(position)
	}
	var zero foundation.Data
	return zero
}

// HasSphericalHarmonicsCoefficientsAtPosition returns true if a handler for SphericalHarmonicsCoefficientsAtPosition has been set.
func (d *MDLLightProbeIrradianceDataSource) HasSphericalHarmonicsCoefficientsAtPosition() bool {
	return d._SphericalHarmonicsCoefficientsAtPosition != nil
}
