// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
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
	PerformEditingTransactionForTextStorageUsingBlock(textStorage ITextStorage, transaction unsafe.Pointer)
	ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange(textStorage ITextStorage, editMask TextStorageEditActions, newCharRange corefoundation.Range, delta int, invalidatedCharRange corefoundation.Range)
}
