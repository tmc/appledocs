// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (

	"github.com/tmc/appledocs/generated/objectivec"
)

// PPlugInRegistration is the CIPlugInRegistration protocol interface.
//
// The interface for loading Core Image image units.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.coreimage/documentation/CoreImage/CIPlugInRegistration
type PPlugInRegistration interface {
	// Required methods
	Load(host objectivec.IObject) bool/* debug [protocol_interface/required_method]: Load */
}
