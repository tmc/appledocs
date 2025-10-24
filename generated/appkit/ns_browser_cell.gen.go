// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	AlternateImage() IImage
	SetAlternateImage(value IImage)
	Image() IImage
	SetImage(value IImage)
	Leaf() bool
	SetLeaf(value bool)
	Loaded() bool
	SetLoaded(value bool)
	IsLeaf() bool
	SetIsLeaf(value bool)
	IsLoaded() bool
	SetIsLoaded(value bool)
	// methods:
	HighlightColorInView(controlView IView) IColor
	Reset()
	Set()
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell/init(imageCell:)
func NewBrowserCellImageCell(image IImage) BrowserCell {
	instance := getBrowserCellClass().Alloc()
	rv := objc.Send[BrowserCell](instance.ID, objc.Sel("initImageCell:"), image)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell/init(textCell:)
func NewBrowserCellTextCell(string_ objc.IObject /* cross-framework: NSString */) BrowserCell {
	instance := getBrowserCellClass().Alloc()
	rv := objc.Send[BrowserCell](instance.ID, objc.Sel("initTextCell:"), string_)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell/init(coder:)
func NewBrowserCellWithCoder(coder foundation.Coder) BrowserCell {
	instance := getBrowserCellClass().Alloc()
	rv := objc.Send[BrowserCell](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Returns the default image for branch cells in a browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell/branchImage
func (bc _BrowserCellClass) BranchImage() IImage {
	rv := objc.Send[Image](objc.ID(bc.class), objc.Sel("branchImage"))
	return rv
}

// Returns the default image for branch browser cells that are highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell/highlightedBranchImage
func (bc _BrowserCellClass) HighlightedBranchImage() IImage {
	rv := objc.Send[Image](objc.ID(bc.class), objc.Sel("highlightedBranchImage"))
	return rv
}

// Returns the highlight color that the receiver wants to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell/highlightColor(in:)
func (b_ BrowserCell) HighlightColorInView(controlView IView) IColor {
	rv := objc.Send[Color](b_.ID, objc.Sel("highlightColorInView:"), controlView)
	return rv
}


// Unhighlights the receiver and unsets its state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell/reset()
func (b_ BrowserCell) Reset() {
	objc.Send[objc.ID](b_.ID, objc.Sel("reset"))
}


// Highlights the receiver and sets its state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell/set()
func (b_ BrowserCell) Set() {
	objc.Send[objc.ID](b_.ID, objc.Sel("set"))
}


// The browser cell’s image for the highlighted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell/alternateImage
func (b_ BrowserCell) AlternateImage() IImage {
	rv := objc.Send[Image](b_.ID, objc.Sel("alternateImage"))
	return rv
}


// The browser cell’s image for the highlighted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell/alternateImage
func (b_ BrowserCell) SetAlternateImage(value IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAlternateImage:"), value)
}


// Returns the default image for branch cells in a browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell/branchImage
func (b_ BrowserCell) BranchImage() IImage {
	rv := objc.Send[Image](b_.ID, objc.Sel("branchImage"))
	return rv
}


// Returns the default image for branch browser cells that are highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell/highlightedBranchImage
func (b_ BrowserCell) HighlightedBranchImage() IImage {
	rv := objc.Send[Image](b_.ID, objc.Sel("highlightedBranchImage"))
	return rv
}


// The browser cell’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell/image
func (b_ BrowserCell) Image() IImage {
	rv := objc.Send[Image](b_.ID, objc.Sel("image"))
	return rv
}


// The browser cell’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell/image
func (b_ BrowserCell) SetImage(value IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImage:"), value)
}


// A Boolean that indicates whether the browser cell is a leaf or a branch cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell/isLeaf
func (b_ BrowserCell) Leaf() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("leaf"))
	return rv
}


// A Boolean that indicates whether the browser cell is a leaf or a branch cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell/isLeaf
func (b_ BrowserCell) SetLeaf(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLeaf:"), value)
}


// A Boolean that indicates whether the cell is ready to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell/isLoaded
func (b_ BrowserCell) Loaded() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("loaded"))
	return rv
}


// A Boolean that indicates whether the cell is ready to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowserCell/isLoaded
func (b_ BrowserCell) SetLoaded(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLoaded:"), value)
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


