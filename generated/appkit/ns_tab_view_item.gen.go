// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [TabViewItem] class.
var (
	TabViewItemClass     _TabViewItemClass
	TabViewItemClassOnce sync.Once
)

func getTabViewItemClass() _TabViewItemClass {
	TabViewItemClassOnce.Do(func() {
		TabViewItemClass = _TabViewItemClass{objc.GetClass("NSTabViewItem")}
	})
	return TabViewItemClass
}

type _TabViewItemClass struct {
	class objc.Class
}

// An interface definition for the [TabViewItem] class.
type ITabViewItem interface {
	objectivec.IObject
}

// An item in a tab view.
//
// An is a convenient way for presenting information in multiple pages. A tab view is usually distinguished by a row of tabs that give the visual appearance of folder tabs. When the user clicks a tab, the tab view displays a view page provided by your application. A tab view keeps a zero-based array of tab view items, one for each tab in the view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem
type TabViewItem struct {
	objectivec.Object
}

// TabViewItemFrom constructs a [TabViewItem] from an unsafe.Pointer.
//
// An item in a tab view.
func TabViewItemFrom(ptr unsafe.Pointer) TabViewItem {
	return TabViewItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TabViewItemClass) Alloc() TabViewItem {
	rv := objc.Send[TabViewItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TabViewItemClass) New() TabViewItem {
	rv := objc.Send[TabViewItem](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TabViewItem) Init() TabViewItem {
	rv := objc.Send[TabViewItem](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TabViewItem) Autorelease() TabViewItem {
	rv := objc.Send[TabViewItem](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTabViewItem creates a new TabViewItem instance.
func NewTabViewItem() TabViewItem {
	return getTabViewItemClass().New()
}


// Sets the background color for content in the view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/color
func (t_ TabViewItem) Color() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("color"))
	return rv
}


// SetColor sets the value of the color property.
// Sets the background color for content in the view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/color
func (t_ TabViewItem) SetColor(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setColor:"), value)
}

// Returns the current display state of the tab associated with the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/tabState
func (t_ TabViewItem) TabState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("tabState"))
	return rv
}

// Returns the parent tab view for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/tabView
func (t_ TabViewItem) TabView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("tabView"))
	return rv
}



