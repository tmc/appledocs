
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Alert] class.
var AlertClass _AlertClass

func init() {
	AlertClass = _AlertClass{objc.GetClass("NSAlert")}
}

type _AlertClass struct {
	objc.Class
}

// An interface definition for the [Alert] class.
type IAlert interface {
	ID() objc.ID
	BeginSheetModalForWindowCompletionHandler(sheetWindow unsafe.Pointer, handler unsafe.Pointer)
}

type Alert struct {
	id objc.ID
}

func AlertFrom(ptr unsafe.Pointer) Alert {
	return Alert{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ Alert) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _AlertClass) Alloc() Alert {
	rv := objc.Send[Alert](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _AlertClass) New() Alert {
	rv := objc.Send[Alert](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewAlert creates and returns a new initialized instance.
func NewAlert() Alert {
	return AlertClass.New()
}

// Init initializes the instance.
func (a_ Alert) Init() Alert {
	rv := objc.Send[Alert](a_.ID(), selInit)
	return rv
}
// Runs the alert modally as a sheet attached to the specified window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSAlert/beginSheetModal(for:completionHandler:)
func (a_ Alert) BeginSheetModalForWindowCompletionHandler(sheetWindow unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID(), objc.RegisterName("beginSheetModalForWindow:completionHandler:"), sheetWindow, handler)
}
