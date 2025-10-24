// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
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
	CloudShareForUserInterfaceItem(item objc.IObject) cloudkit.CKShare
}
