// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"unsafe"
)

// PFileProviderTestingOperation is the NSFileProviderTestingOperation protocol interface.
//
// An operation that the system can schedule.
//
// Availability:
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 11.3+
//   - visionOS 1.0+
//
// See: doc://com.apple.fileprovider/documentation/FileProvider/NSFileProviderTestingOperation
type PFileProviderTestingOperation interface {
	// Required methods
	AsChildrenEnumeration() unsafe.Pointer/* debug [protocol_interface/required_method]: AsChildrenEnumeration */
	AsCollisionResolution() unsafe.Pointer/* debug [protocol_interface/required_method]: AsCollisionResolution */
	AsContentFetch() unsafe.Pointer/* debug [protocol_interface/required_method]: AsContentFetch */
	AsCreation() unsafe.Pointer/* debug [protocol_interface/required_method]: AsCreation */
	AsDeletion() unsafe.Pointer/* debug [protocol_interface/required_method]: AsDeletion */
	AsIngestion() unsafe.Pointer/* debug [protocol_interface/required_method]: AsIngestion */
	AsLookup() unsafe.Pointer/* debug [protocol_interface/required_method]: AsLookup */
	AsModification() unsafe.Pointer/* debug [protocol_interface/required_method]: AsModification */
}
