// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Alert] class.
var alertClass = _AlertClass{objc.GetClass("NSAlert")}

type _AlertClass struct {
	class objc.Class
}

// A modal dialog or sheet attached to a document window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlert

type Alert struct {
	objectivec.Object
}

// AlertFrom constructs a [Alert] from an unsafe.Pointer.
//
// A modal dialog or sheet attached to a document window.
func AlertFrom(ptr unsafe.Pointer) Alert {
	return Alert{objectivec.Object{objc.ID(ptr)}}
}

// Runs the alert modally as a sheet attached to the specified window. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlert/beginSheetModal(for:completionHandler:)
func (a_ Alert) BeginSheetModalForWindowCompletionHandler(sheetWindow unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("beginSheetModalForWindow:completionHandler:"), sheetWindow, handler)
}


