// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DraggingItem] class.
var (
	DraggingItemClass     _DraggingItemClass
	DraggingItemClassOnce sync.Once
)

func getDraggingItemClass() _DraggingItemClass {
	DraggingItemClassOnce.Do(func() {
		DraggingItemClass = _DraggingItemClass{objc.GetClass("NSDraggingItem")}
	})
	return DraggingItemClass
}

type _DraggingItemClass struct {
	class objc.Class
}

// An interface definition for the [DraggingItem] class.
type IDraggingItem interface {
	objectivec.IObject
	// properties:
	DraggingFrame() objc.IObject /* cross-framework: Rect */
	SetDraggingFrame(value objc.IObject /* cross-framework: Rect */)
	ImageComponents() IDraggingImageComponent
	SetImageComponents(value IDraggingImageComponent)
	ImageComponentsProvider() IDraggingImageComponent
	SetImageComponentsProvider(value IDraggingImageComponent)
	Item() unsafe.Pointer
	SetItem(value unsafe.Pointer)
	// methods:
}

// A single dragged item within a dragging session.
//
// objects have extremely limited lifetimes. Don’t retain these items because changing outside of the prescribed lifetimes has no impact on the drag. When you call the method , the system immediately consumes the dragging items that pass to the method, and doesn’t retain them. Any further changes to the dragging item associated with the returned must occur with the enumeration method . When enumerating, the system creates instances right before giving them to the enumeration block. After returning from the block, the dragging item is no longer valid.


// A single dragged item within a dragging session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingItem
type DraggingItem struct {
	objectivec.Object
}

// DraggingItemFrom constructs a [DraggingItem] from an unsafe.Pointer.
//
// A single dragged item within a dragging session.
func DraggingItemFrom(ptr unsafe.Pointer) DraggingItem {
	return DraggingItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DraggingItemClass) Alloc() DraggingItem {
	rv := objc.Send[DraggingItem](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DraggingItemClass) New() DraggingItem {
	rv := objc.Send[DraggingItem](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DraggingItem) Init() DraggingItem {
	rv := objc.Send[DraggingItem](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DraggingItem) Autorelease() DraggingItem {
	rv := objc.Send[DraggingItem](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDraggingItem creates a new DraggingItem instance.
func NewDraggingItem() DraggingItem {
	return getDraggingItemClass().New()
}



// The frame of the dragging item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingitem/draggingframe
func (d_ DraggingItem) DraggingFrame() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[Rect](d_.ID, objc.Sel("draggingFrame"))
	return rv
}


// The frame of the dragging item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingitem/draggingframe
func (d_ DraggingItem) SetDraggingFrame(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDraggingFrame:"), value)
}


// An array of dragging image components to use to create the drag image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingitem/imagecomponents
func (d_ DraggingItem) ImageComponents() IDraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](d_.ID, objc.Sel("imageComponents"))
	return rv
}


// An array of dragging image components to use to create the drag image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingitem/imagecomponents
func (d_ DraggingItem) SetImageComponents(value IDraggingImageComponent) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setImageComponents:"), value)
}


// An array of blocks that provide the dragging image components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingitem/imagecomponentsprovider
func (d_ DraggingItem) ImageComponentsProvider() IDraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](d_.ID, objc.Sel("imageComponentsProvider"))
	return rv
}


// An array of blocks that provide the dragging image components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingitem/imagecomponentsprovider
func (d_ DraggingItem) SetImageComponentsProvider(value IDraggingImageComponent) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setImageComponentsProvider:"), value)
}


// The pasteboard reader or writer object dependent on the context where you use the dragging item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingitem/item
func (d_ DraggingItem) Item() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("item"))
	return rv
}


// The pasteboard reader or writer object dependent on the context where you use the dragging item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdraggingitem/item
func (d_ DraggingItem) SetItem(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setItem:"), value)
}



