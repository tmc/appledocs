// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SavePanel] class.
var SavePanelClass objc.Class

func init() {
	SavePanelClass = objc.GetClass("NSSavePanel")
}

type SavePanel struct {
	objc.ID
}

func SavePanelFrom(ptr unsafe.Pointer) SavePanel {
	return SavePanel{
		ID: objc.ID(ptr),
	}
}


// Creates a new Save panel and initializes it with default information. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/savePanel
func (sc SavePanel) SavePanel() unsafe.Pointer {
	sel := objc.RegisterName("savePanel")
	ret := objc.ID(SavePanelClass).Send(sel)
	return unsafe.Pointer(ret)
}
// Presents the panel as a modeless window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/begin(completionHandler:)
func (s_ SavePanel) BeginWithCompletionHandler(handler unsafe.Pointer) {
	sel := objc.RegisterName("beginWithCompletionHandler:")
	s_.ID.Send(sel, handler)
}
// Presents the panel as a sheet modal to the specified window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/beginSheetModal(for:completionHandler:)
func (s_ SavePanel) BeginSheetModalForWindowCompletionHandler(window unsafe.Pointer, handler unsafe.Pointer) {
	sel := objc.RegisterName("beginSheetModalForWindow:completionHandler:")
	s_.ID.Send(sel, window, handler)
}
// The action method that the panel calls when the user clicks the Cancel button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/cancel(_:)
func (s_ SavePanel) Cancel(sender objc.ID) {
	sel := objc.RegisterName("cancel:")
	s_.ID.Send(sel, sender)
}
// The action method that the panel calls when the user clicks the OK button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/ok(_:)
func (s_ SavePanel) Ok(sender objc.ID) {
	sel := objc.RegisterName("ok:")
	s_.ID.Send(sel, sender)
}
// Displays the panel and begins its event loop with the current working (or last-selected) directory as the default starting point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/runModal()
func (s_ SavePanel) RunModal() unsafe.Pointer {
	sel := objc.RegisterName("runModal")
	ret := s_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Validates and reloads the browser columns visible in the panel. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/validateVisibleColumns()
func (s_ SavePanel) ValidateVisibleColumns() {
	sel := objc.RegisterName("validateVisibleColumns")
	s_.ID.Send(sel)
}

