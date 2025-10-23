// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BrowserCell] class.
var (
	BrowserCellClass     _BrowserCellClass
	BrowserCellClassOnce sync.Once
)

func getBrowserCellClass() _BrowserCellClass {
	BrowserCellClassOnce.Do(func() {
		BrowserCellClass = _BrowserCellClass{objc.GetClass("NSBrowserCell")}
	})
	return BrowserCellClass
}

type _BrowserCellClass struct {
	class objc.Class
}

// An interface definition for the [BrowserCell] class.
type IBrowserCell interface {
	ICell
	AlternateImage() Image
	SetAlternateImage(value IImage)
	Image() Image
	SetImage(value IImage)
	IsLeaf() bool
	SetIsLeaf(value bool)
	IsLoaded() bool
	SetIsLoaded(value bool)
}

// The user interface of a browser.
//
// The class is the subclass of used by default to display data in the columns of an object. (Each column contains an object filled with objects.)


// The user interface of a browser.
//
// [Full Topic]
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



// The browser cell’s image for the highlighted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/alternateimage
func (b_ BrowserCell) AlternateImage() Image {
	rv := objc.Send[Image](b_.ID, objc.Sel("alternateImage"))
	return rv
}


// The browser cell’s image for the highlighted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/alternateimage
func (b_ BrowserCell) SetAlternateImage(value IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAlternateImage:"), value)
}


// The browser cell’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/image
func (b_ BrowserCell) Image() Image {
	rv := objc.Send[Image](b_.ID, objc.Sel("image"))
	return rv
}


// The browser cell’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/image
func (b_ BrowserCell) SetImage(value IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImage:"), value)
}


// A Boolean that indicates whether the browser cell is a leaf or a branch cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/isleaf
func (b_ BrowserCell) IsLeaf() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isLeaf"))
	return rv
}


// A Boolean that indicates whether the browser cell is a leaf or a branch cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/isleaf
func (b_ BrowserCell) SetIsLeaf(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsLeaf:"), value)
}


// A Boolean that indicates whether the cell is ready to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/isloaded
func (b_ BrowserCell) IsLoaded() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isLoaded"))
	return rv
}


// A Boolean that indicates whether the cell is ready to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/isloaded
func (b_ BrowserCell) SetIsLoaded(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsLoaded:"), value)
}



