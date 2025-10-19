// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Alert] class.
var (
	alertClass     _AlertClass
	alertClassOnce sync.Once
)

func getAlertClass() _AlertClass {
	alertClassOnce.Do(func() {
		alertClass = _AlertClass{objc.GetClass("NSAlert")}
	})
	return alertClass
}

type _AlertClass struct {
	class objc.Class
}

// An interface definition for the [Alert] class.
type IAlert interface {
	objectivec.IObject
	BeginSheetModalForWindowCompletionHandler(sheetWindow unsafe.Pointer, handler unsafe.Pointer)
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
// Alloc allocates a new instance without initialization.
func (ac _AlertClass) Alloc() Alert {
	rv := objc.Send[Alert](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AlertClass) New() Alert {
	rv := objc.Send[Alert](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Alert) Init() Alert {
	rv := objc.Send[Alert](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Alert) Autorelease() Alert {
	rv := objc.Send[Alert](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAlert creates a new Alert instance.
func NewAlert() Alert {
	return getAlertClass().New()
}


// Runs the alert modally as a sheet attached to the specified window. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlert/beginSheetModal(for:completionHandler:)
func (a_ Alert) BeginSheetModalForWindowCompletionHandler(sheetWindow unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("beginSheetModalForWindow:completionHandler:"), sheetWindow, handler)
}


