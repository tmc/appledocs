
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SplitViewItem] class.
var SplitViewItemClass _SplitViewItemClass

func init() {
	SplitViewItemClass = _SplitViewItemClass{objc.GetClass("NSSplitViewItem")}
}

type _SplitViewItemClass struct {
	objc.Class
}

// An interface definition for the [SplitViewItem] class.
type ISplitViewItem interface {
	ID() objc.ID
}

type SplitViewItem struct {
	id objc.ID
}

func SplitViewItemFrom(ptr unsafe.Pointer) SplitViewItem {
	return SplitViewItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SplitViewItem) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SplitViewItemClass) Alloc() SplitViewItem {
	rv := objc.Send[SplitViewItem](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SplitViewItemClass) New() SplitViewItem {
	rv := objc.Send[SplitViewItem](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSplitViewItem creates and returns a new initialized instance.
func NewSplitViewItem() SplitViewItem {
	return SplitViewItemClass.New()
}

// Init initializes the instance.
func (s_ SplitViewItem) Init() SplitViewItem {
	rv := objc.Send[SplitViewItem](s_.ID(), selInit)
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitViewItem/init(inspectorWithViewController:)
func (sc _SplitViewItemClass) InspectorWithViewController(viewController unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.Class), objc.RegisterName("inspectorWithViewController:"), viewController)
	return rv
}
// A Boolean value that indicates whether full-height sidebars appear in the window after you set a style mask. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitViewItem/allowsFullHeightLayout
func (s_ SplitViewItem) AllowsFullHeightLayout() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("allowsFullHeightLayout"))
	return rv
}
// SetAllowsFullHeightLayout sets the value of the allowsFullHeightLayout property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitViewItem/allowsFullHeightLayout
func (s_ SplitViewItem) SetAllowsFullHeightLayout(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setAllowsFullHeightLayout:"), value)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitViewItem/bottomAlignedAccessoryViewControllers
func (s_ SplitViewItem) BottomAlignedAccessoryViewControllers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("bottomAlignedAccessoryViewControllers"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitViewItem/bottomAlignedAccessoryViewControllers
func (s_ SplitViewItem) SetBottomAlignedAccessoryViewControllers(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setBottomAlignedAccessoryViewControllers:"), value)
}
// The type of separator that the app displays between the title bar and content of a window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitViewItem/titlebarSeparatorStyle
func (s_ SplitViewItem) TitlebarSeparatorStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("titlebarSeparatorStyle"))
	return rv
}
// SetTitlebarSeparatorStyle sets the value of the titlebarSeparatorStyle property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitViewItem/titlebarSeparatorStyle
func (s_ SplitViewItem) SetTitlebarSeparatorStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setTitlebarSeparatorStyle:"), value)
}
// The following methods allow you to add accessory views to the top/bottom of this splitViewItem. See   for more details. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitViewItem/topAlignedAccessoryViewControllers
func (s_ SplitViewItem) TopAlignedAccessoryViewControllers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("topAlignedAccessoryViewControllers"))
	return rv
}
// SetTopAlignedAccessoryViewControllers sets the value of the topAlignedAccessoryViewControllers property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitViewItem/topAlignedAccessoryViewControllers
func (s_ SplitViewItem) SetTopAlignedAccessoryViewControllers(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setTopAlignedAccessoryViewControllers:"), value)
}
