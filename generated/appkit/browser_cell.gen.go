// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BrowserCell] class.
var (
	browserCellClass     _BrowserCellClass
	browserCellClassOnce sync.Once
)

func getBrowserCellClass() _BrowserCellClass {
	browserCellClassOnce.Do(func() {
		browserCellClass = _BrowserCellClass{objc.GetClass("NSBrowserCell")}
	})
	return browserCellClass
}

type _BrowserCellClass struct {
	class objc.Class
}

// An interface definition for the [BrowserCell] class.
type IBrowserCell interface {
	ICell
}

// The user interface of a browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell
type BrowserCell struct {
	Cell
}

// BrowserCellFrom constructs a [BrowserCell] from an unsafe.Pointer.
//
// The user interface of a browser.
func BrowserCellFrom(ptr unsafe.Pointer) BrowserCell {
	return BrowserCell{
		Cell: CellFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BrowserCellClass) Alloc() BrowserCell {
	rv := objc.Send[BrowserCell](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BrowserCellClass) New() BrowserCell {
	rv := objc.Send[BrowserCell](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BrowserCell) Init() BrowserCell {
	rv := objc.Send[BrowserCell](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BrowserCell) Autorelease() BrowserCell {
	rv := objc.Send[BrowserCell](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBrowserCell creates a new BrowserCell instance.
func NewBrowserCell() BrowserCell {
	return getBrowserCellClass().New()
}




