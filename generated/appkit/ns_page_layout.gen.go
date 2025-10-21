// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PageLayout] class.
var (
	PageLayoutClass     _PageLayoutClass
	PageLayoutClassOnce sync.Once
)

func getPageLayoutClass() _PageLayoutClass {
	PageLayoutClassOnce.Do(func() {
		PageLayoutClass = _PageLayoutClass{objc.GetClass("NSPageLayout")}
	})
	return PageLayoutClass
}

type _PageLayoutClass struct {
	class objc.Class
}

// An interface definition for the [PageLayout] class.
type IPageLayout interface {
	objectivec.IObject
	AccessoryView() unsafe.Pointer
}

// A panel that queries the user for information such as paper type and orientation.
//
// A page layout panel is typically displayed in response to the user selecting the Page Setup menu item. You obtain an instance with the class method. The pane can then be run as a sheet using or modally using or . For design guidance, see .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout
type PageLayout struct {
	objectivec.Object
}

// PageLayoutFrom constructs a [PageLayout] from an unsafe.Pointer.
//
// A panel that queries the user for information such as paper type and orientation.
func PageLayoutFrom(ptr unsafe.Pointer) PageLayout {
	return PageLayout{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PageLayoutClass) Alloc() PageLayout {
	rv := objc.Send[PageLayout](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PageLayoutClass) New() PageLayout {
	rv := objc.Send[PageLayout](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PageLayout) Init() PageLayout {
	rv := objc.Send[PageLayout](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PageLayout) Autorelease() PageLayout {
	rv := objc.Send[PageLayout](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPageLayout creates a new PageLayout instance.
func NewPageLayout() PageLayout {
	return getPageLayoutClass().New()
}


// Returns a newly created page layout object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout/pageLayout
func (pc _PageLayoutClass) PageLayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("pageLayout"))
	return rv
}

// Returns the page layout panel’s accessory view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout/accessoryView
func (p_ PageLayout) AccessoryView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("accessoryView"))
	return rv
}

// An array of accessory view controllers belonging to the page layout panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout/accessoryControllers
func (p_ PageLayout) AccessoryControllers() []ViewController {
	rv := objc.Send[[]ViewController](p_.ID, objc.Sel("accessoryControllers"))
	return rv
}

// The printing information object used when the page layout panel is run.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagelayout/printinfo
func (p_ PageLayout) PrintInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("printInfo"))
	return rv
}


// SetPrintInfo sets the value of the printInfo property.
// The printing information object used when the page layout panel is run.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagelayout/printinfo
func (p_ PageLayout) SetPrintInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPrintInfo:"), value)
}



