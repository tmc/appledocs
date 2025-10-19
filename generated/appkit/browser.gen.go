// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Browser] class.
var (
	browserClass     _BrowserClass
	browserClassOnce sync.Once
)

func getBrowserClass() _BrowserClass {
	browserClassOnce.Do(func() {
		browserClass = _BrowserClass{objc.GetClass("NSBrowser")}
	})
	return browserClass
}

type _BrowserClass struct {
	class objc.Class
}

// An interface definition for the [Browser] class.
type IBrowser interface {
	IControl
}

// An interface that displays a hierarchically organized list of data items that can be navigated and selected.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser
type Browser struct {
	Control
}

// BrowserFrom constructs a [Browser] from an unsafe.Pointer.
//
// An interface that displays a hierarchically organized list of data items that can be navigated and selected.
func BrowserFrom(ptr unsafe.Pointer) Browser {
	return Browser{
		Control: ControlFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BrowserClass) Alloc() Browser {
	rv := objc.Send[Browser](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BrowserClass) New() Browser {
	rv := objc.Send[Browser](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ Browser) Init() Browser {
	rv := objc.Send[Browser](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ Browser) Autorelease() Browser {
	rv := objc.Send[Browser](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBrowser creates a new Browser instance.
func NewBrowser() Browser {
	return getBrowserClass().New()
}




