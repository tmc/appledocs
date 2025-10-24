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
	GetXattrNamedOfItemReplyHandler(name IFSFileName, item IFSItem, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: GetXattrNamedOfItemReplyHandler */
	ListXattrsOfItemReplyHandler(item IFSItem, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: ListXattrsOfItemReplyHandler */
	SetXattrNamedToDataOnItemPolicyReplyHandler(name IFSFileName, value objc.IObject /* cross-framework: NSData */, item IFSItem, policy FSSetXattrPolicy, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: SetXattrNamedToDataOnItemPolicyReplyHandler */
	// Optional methods
	SupportedXattrNamesForItem(item IFSItem) []FSFileName
	HasSupportedXattrNamesForItem() bool
}
