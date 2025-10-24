// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PMTRStorage is the MTRStorage protocol interface.
//
// Availability:
//   - Mac Catalyst 16.4+
//   - iOS 16.4+
//   - iPadOS 16.4+
//   - macOS 13.3+
//   - tvOS 16.4+
//   - visionOS 1.0+
//   - watchOS 9.4+
//
// See: doc://com.apple.matter/documentation/Matter/MTRStorage
type PMTRStorage interface {
	// Required methods
	RemoveStorageDataForKey(key objc.IObject /* cross-framework: NSString */) bool/* debug [protocol_interface/required_method]: RemoveStorageDataForKey */
	SetStorageDataForKey(value objc.IObject /* cross-framework: NSData */, key objc.IObject /* cross-framework: NSString */) bool/* debug [protocol_interface/required_method]: SetStorageDataForKey */
	StorageDataForKey(key objc.IObject /* cross-framework: NSString */) foundation.Data/* debug [protocol_interface/required_method]: StorageDataForKey */
}
