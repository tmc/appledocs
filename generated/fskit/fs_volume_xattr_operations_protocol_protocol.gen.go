// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PFSVolumeXattrOperations is the FSVolumeXattrOperations protocol interface.
//
// Methods and properties implemented by volumes that natively or partially support extended attributes.
//
// Availability:
//   - macOS 15.4+
//
// See: doc://FSKit/documentation/FSKit/FSVolume/XattrOperations
type PFSVolumeXattrOperations interface {
	// Required methods
	GetXattrNamedOfItemReplyHandler(name IFSFileName, item IFSItem, reply unsafe.Pointer)
	ListXattrsOfItemReplyHandler(item IFSItem, reply unsafe.Pointer)
	SetXattrNamedToDataOnItemPolicyReplyHandler(name IFSFileName, value foundation.foundation.INSData, item IFSItem, policy FSSetXattrPolicy, reply unsafe.Pointer)
	// Optional methods
	SupportedXattrNamesForItem(item IFSItem) []FSFileName
	HasSupportedXattrNamesForItem() bool
}
