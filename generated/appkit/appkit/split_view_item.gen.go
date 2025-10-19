// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SplitViewItem] class.
var (
	splitViewItemClass     _SplitViewItemClass
	splitViewItemClassOnce sync.Once
)

func getSplitViewItemClass() _SplitViewItemClass {
	splitViewItemClassOnce.Do(func() {
		splitViewItemClass = _SplitViewItemClass{objc.GetClass("NSSplitViewItem")}
	})
	return splitViewItemClass
}

type _SplitViewItemClass struct {
	class objc.Class
}

// An interface definition for the [SplitViewItem] class.
type ISplitViewItem interface {
	objectivec.IObject
}

// An item in a split view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem

type SplitViewItem struct {
	objectivec.Object
}

// SplitViewItemFrom constructs a [SplitViewItem] from an unsafe.Pointer.
//
// An item in a split view controller.
func SplitViewItemFrom(ptr unsafe.Pointer) SplitViewItem {
	return SplitViewItem{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (sc _SplitViewItemClass) Alloc() SplitViewItem {
	rv := objc.Send[SplitViewItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SplitViewItemClass) New() SplitViewItem {
	rv := objc.Send[SplitViewItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SplitViewItem) Init() SplitViewItem {
	rv := objc.Send[SplitViewItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SplitViewItem) Autorelease() SplitViewItem {
	rv := objc.Send[SplitViewItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSplitViewItem creates a new SplitViewItem instance.
func NewSplitViewItem() SplitViewItem {
	return getSplitViewItemClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(inspectorWithViewController:)
func NewSplitViewItemInspectorWithViewController(viewController unsafe.Pointer) SplitViewItem {
	rv := objc.Send[SplitViewItem](objc.ID(getSplitViewItemClass().class), objc.Sel("inspectorWithViewController:"), viewController)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(inspectorWithViewController:)
func (sc _SplitViewItemClass) InspectorWithViewController(viewController unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("inspectorWithViewController:"), viewController)
	return rv
}

