// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CollectionLayoutGroupCustomItem] class.
var (
	CollectionLayoutGroupCustomItemClass     _CollectionLayoutGroupCustomItemClass
	CollectionLayoutGroupCustomItemClassOnce sync.Once
)

func getCollectionLayoutGroupCustomItemClass() _CollectionLayoutGroupCustomItemClass {
	CollectionLayoutGroupCustomItemClassOnce.Do(func() {
		CollectionLayoutGroupCustomItemClass = _CollectionLayoutGroupCustomItemClass{objc.GetClass("NSCollectionLayoutGroupCustomItem")}
	})
	return CollectionLayoutGroupCustomItemClass
}

type _CollectionLayoutGroupCustomItemClass struct {
	class objc.Class
}





// An interface definition for the [CollectionLayoutGroupCustomItem] class.
type ICollectionLayoutGroupCustomItem interface {
	objectivec.IObject
	

	// properties:
	Frame() corefoundation.CGRect
	ZIndex() int


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CollectionLayoutGroupCustomItemClass) Alloc() CollectionLayoutGroupCustomItem {
	rv := objc.Send[CollectionLayoutGroupCustomItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CollectionLayoutGroupCustomItemClass) New() CollectionLayoutGroupCustomItem {
	rv := objc.Send[CollectionLayoutGroupCustomItem](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionLayoutGroupCustomItem) Init() CollectionLayoutGroupCustomItem {
	rv := objc.Send[CollectionLayoutGroupCustomItem](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionLayoutGroupCustomItem) Autorelease() CollectionLayoutGroupCustomItem {
	rv := objc.Send[CollectionLayoutGroupCustomItem](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionLayoutGroupCustomItem creates a new CollectionLayoutGroupCustomItem instance.
func NewCollectionLayoutGroupCustomItem() CollectionLayoutGroupCustomItem {
	return getCollectionLayoutGroupCustomItemClass().New()
}





// An item used in a group with a custom layout arrangement.
//
// You use a custom item if you want to specify a layout with a custom arrangement, like a radial or diagonal layout. You use custom items within a group that’s created with . Instead of providing a layout size for the custom item, like you do when you create an , you provide a frame instead.


// An item used in a group with a custom layout arrangement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem
type CollectionLayoutGroupCustomItem struct {
	objectivec.Object
}

// CollectionLayoutGroupCustomItemFrom constructs a [CollectionLayoutGroupCustomItem] from an unsafe.Pointer.
//
// An item used in a group with a custom layout arrangement.
func CollectionLayoutGroupCustomItemFrom(ptr unsafe.Pointer) CollectionLayoutGroupCustomItem {
	return CollectionLayoutGroupCustomItem{objectivec.Object{objc.ID(ptr)}}
}






// Creates a custom item with the specified frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem/init(frame:)
func NewCollectionLayoutGroupCustomItemWithFrame(frame corefoundation.CGRect) CollectionLayoutGroupCustomItem {
	rv := objc.Send[CollectionLayoutGroupCustomItem](objc.ID(getCollectionLayoutGroupCustomItemClass().class), objc.Sel("customItemWithFrame:"), frame)
	return rv
}


// Creates a custom item with the specified frame and vertical stacking order in relation to other items in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem/init(frame:zIndex:)
func NewCollectionLayoutGroupCustomItemWithFrameZIndex(frame corefoundation.CGRect, zIndex int) CollectionLayoutGroupCustomItem {
	rv := objc.Send[CollectionLayoutGroupCustomItem](objc.ID(getCollectionLayoutGroupCustomItemClass().class), objc.Sel("customItemWithFrame:zIndex:"), frame, zIndex)
	return rv
}







// Creates a custom item with the specified frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem/init(frame:)
func (cc _CollectionLayoutGroupCustomItemClass) CustomItemWithFrame(frame corefoundation.CGRect) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("customItemWithFrame:"), frame)
	return rv
}


// Creates a custom item with the specified frame and vertical stacking order in relation to other items in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem/init(frame:zIndex:)
func (cc _CollectionLayoutGroupCustomItemClass) CustomItemWithFrameZIndex(frame corefoundation.CGRect, zIndex int) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("customItemWithFrame:zIndex:"), frame, zIndex)
	return rv
}

















// The frame of the custom item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem/frame
func (c_ CollectionLayoutGroupCustomItem) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("frame"))
	return rv
}


// The vertical stacking order of the custom item in relation to other items in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem/zIndex
func (c_ CollectionLayoutGroupCustomItem) ZIndex() int {
	rv := objc.Send[int](c_.ID, objc.Sel("zIndex"))
	return rv
}







