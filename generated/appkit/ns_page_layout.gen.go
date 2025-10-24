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
	// properties:
	AccessoryControllers() objc.IObject /* cross-framework: ViewController */
	SetAccessoryControllers(value objc.IObject /* cross-framework: ViewController */)
	PrintInfo() IPrintInfo
	SetPrintInfo(value IPrintInfo)
	// methods:
}

// A panel that queries the user for information such as paper type and orientation.
//
// A page layout panel is typically displayed in response to the user selecting the Page Setup menu item. You obtain an instance with the class method. The pane can then be run as a sheet using or modally using or . For design guidance, see .


// A panel that queries the user for information such as paper type and orientation.
//
// [Full Topic]
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



// An array of accessory view controllers belonging to the page layout panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagelayout/accessorycontrollers
func (p_ PageLayout) AccessoryControllers() objc.IObject /* cross-framework: ViewController */ {
	rv := objc.Send[ViewController](p_.ID, objc.Sel("accessoryControllers"))
	return rv
}


// An array of accessory view controllers belonging to the page layout panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagelayout/accessorycontrollers
func (p_ PageLayout) SetAccessoryControllers(value objc.IObject /* cross-framework: ViewController */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAccessoryControllers:"), value)
}


// The printing information object used when the page layout panel is run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagelayout/printinfo
func (p_ PageLayout) PrintInfo() IPrintInfo {
	rv := objc.Send[PrintInfo](p_.ID, objc.Sel("printInfo"))
	return rv
}


// The printing information object used when the page layout panel is run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagelayout/printinfo
func (p_ PageLayout) SetPrintInfo(value IPrintInfo) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPrintInfo:"), value)
}



