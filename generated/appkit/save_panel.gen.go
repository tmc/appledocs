// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SavePanel] class.
var (
	savePanelClass     _SavePanelClass
	savePanelClassOnce sync.Once
)

func getSavePanelClass() _SavePanelClass {
	savePanelClassOnce.Do(func() {
		savePanelClass = _SavePanelClass{objc.GetClass("NSSavePanel")}
	})
	return savePanelClass
}

type _SavePanelClass struct {
	class objc.Class
}

// An interface definition for the [SavePanel] class.
type ISavePanel interface {
	IPanel
	BeginWithCompletionHandler(handler unsafe.Pointer)
	BeginSheetModalForWindowCompletionHandler(window unsafe.Pointer, handler unsafe.Pointer)
	Cancel(sender objc.ID)
	Ok(sender objc.ID)
	RunModal() unsafe.Pointer
	ValidateVisibleColumns()
}

// A panel that prompts the user for information about where to save a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel

type SavePanel struct {
	Panel
}

// SavePanelFrom constructs a [SavePanel] from an unsafe.Pointer.
//
// A panel that prompts the user for information about where to save a file.
func SavePanelFrom(ptr unsafe.Pointer) SavePanel {
	return SavePanel{
		Panel: PanelFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _SavePanelClass) Alloc() SavePanel {
	rv := objc.Send[SavePanel](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SavePanelClass) New() SavePanel {
	rv := objc.Send[SavePanel](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SavePanel) Init() SavePanel {
	rv := objc.Send[SavePanel](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SavePanel) Autorelease() SavePanel {
	rv := objc.Send[SavePanel](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSavePanel creates a new SavePanel instance.
func NewSavePanel() SavePanel {
	return getSavePanelClass().New()
}


// Creates a new Save panel and initializes it with default information. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/savePanel
func (sc _SavePanelClass) SavePanel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("savePanel"))
	return rv
}
// Presents the panel as a modeless window. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/begin(completionHandler:)
func (s_ SavePanel) BeginWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginWithCompletionHandler:"), handler)
}
// Presents the panel as a sheet modal to the specified window. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/beginSheetModal(for:completionHandler:)
func (s_ SavePanel) BeginSheetModalForWindowCompletionHandler(window unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginSheetModalForWindow:completionHandler:"), window, handler)
}
// The action method that the panel calls when the user clicks the Cancel button. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/cancel(_:)
func (s_ SavePanel) Cancel(sender objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("cancel:"), sender)
}
// The action method that the panel calls when the user clicks the OK button. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/ok(_:)
func (s_ SavePanel) Ok(sender objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("ok:"), sender)
}
// Displays the panel and begins its event loop with the current working (or last-selected) directory as the default starting point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/runModal()
func (s_ SavePanel) RunModal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("runModal"))
	return rv
}
// Validates and reloads the browser columns visible in the panel. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/validateVisibleColumns()
func (s_ SavePanel) ValidateVisibleColumns() {
	objc.Send[objc.ID](s_.ID, objc.Sel("validateVisibleColumns"))
}


