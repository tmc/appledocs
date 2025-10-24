// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/cloudkit"
)

// PCloudSharingValidation is the NSCloudSharingValidation protocol interface.
//
// A protocol that a Cloud-sharing toolbar item uses to get validation of an item.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSCloudSharingValidation
type PCloudSharingValidation interface {
	// Required methods
	CloudShareForUserInterfaceItem(item unsafe.Pointer) cloudkit.CKShare/* debug [protocol_interface/required_method]: CloudShareForUserInterfaceItem */
}
