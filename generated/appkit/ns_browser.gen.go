// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Browser] class.
var (
	BrowserClass     _BrowserClass
	BrowserClassOnce sync.Once
)

func getBrowserClass() _BrowserClass {
	BrowserClassOnce.Do(func() {
		BrowserClass = _BrowserClass{objc.GetClass("NSBrowser")}
	})
	return BrowserClass
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
// A browser displays information using a set of columns, which are indexed from left to right. Each successive column displays the next level down in the data hierarchy. This class uses the class to implement its user interface. Browsers have the following components: Columns Scroll views Matrices Browser cells To the user, browsers display data in columns and rows within each column. These components are arranged in the following component hierarchy:
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




