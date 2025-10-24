// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PCKRecordKeyValueSetting is the CKRecordKeyValueSetting protocol interface.
//
// A protocol for managing the key-value pairs of a CloudKit record.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 3.0+
//
// See: doc://com.apple.cloudkit/documentation/CloudKit/CKRecordKeyValueSetting
type PCKRecordKeyValueSetting interface {
	// Required methods
	AllKeys() []string/* debug [protocol_interface/required_method]: AllKeys */
	ChangedKeys() []string/* debug [protocol_interface/required_method]: ChangedKeys */
	ObjectForKey(key objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: ObjectForKey */
	SetObjectForKey(object unsafe.Pointer, key objectivec.IObject)/* debug [protocol_interface/required_method]: SetObjectForKey */
	SetObjectForKeyedSubscript(object unsafe.Pointer, key objectivec.IObject)/* debug [protocol_interface/required_method]: SetObjectForKeyedSubscript */
	ObjectForKeyedSubscript(key objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: ObjectForKeyedSubscript */
}
