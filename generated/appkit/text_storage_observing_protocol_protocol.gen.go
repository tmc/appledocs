// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"
)

// PTextStorageObserving is the NSTextStorageObserving protocol interface.
//
// Optional methods that delegates implement to handle editing and transaction processing.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextStorageObserving
type PTextStorageObserving interface {
	// Required methods
	PerformEditingTransactionForTextStorageUsingBlock(textStorage ITextStorage, transaction unsafe.Pointer)/* debug [protocol_interface/required_method]: PerformEditingTransactionForTextStorageUsingBlock */
	ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange(textStorage ITextStorage, editMask TextStorageEditActions, newCharRange corefoundation.Range, delta int, invalidatedCharRange corefoundation.Range)/* debug [protocol_interface/required_method]: ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange */
}
