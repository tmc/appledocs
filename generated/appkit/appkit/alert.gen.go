// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Alert] class.
var AlertClass objc.Class

func init() {
	AlertClass = objc.GetClass("NSAlert")
}

type Alert struct {
	objc.ID
}

func AlertFrom(ptr unsafe.Pointer) Alert {
	return Alert{
		ID: objc.ID(ptr),
	}
}


// Runs the alert modally as a sheet attached to the specified window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSAlert/beginSheetModal(for:completionHandler:)
func (a_ Alert) BeginSheetModalForWindowCompletionHandler(sheetWindow unsafe.Pointer, handler unsafe.Pointer) {
	sel := objc.RegisterName("beginSheetModalForWindow:completionHandler:")
	a_.ID.Send(sel, sheetWindow, handler)
}


