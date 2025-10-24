// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCollectionLayoutGroupCustomItem */


/* debug [class_header]: Header for NSCollectionLayoutGroupCustomItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CollectionLayoutGroupCustomItem */
// An interface definition for the [CollectionLayoutGroupCustomItem] class.
type ICollectionLayoutGroupCustomItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CollectionLayoutGroupCustomItem */
	// properties:
	Frame() Rect /* not a class type */
	ZIndex() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CollectionLayoutGroupCustomItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CollectionLayoutGroupCustomItem */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CollectionLayoutGroupCustomItem */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CollectionLayoutGroupCustomItem */

// Creates a custom item with the specified frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem/init(frame:)
func NewCollectionLayoutGroupCustomItemWithFrame(frame Rect /* not a class type */) CollectionLayoutGroupCustomItem {
	rv := objc.Send[CollectionLayoutGroupCustomItem](objc.ID(getCollectionLayoutGroupCustomItemClass().class), objc.Sel("customItemWithFrame:"), frame)
	return rv
}/* debug [class_init_methods/constructor]: NewCollectionLayoutGroupCustomItemWithFrame */


// Creates a custom item with the specified frame and vertical stacking order in relation to other items in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem/init(frame:zIndex:)
func NewCollectionLayoutGroupCustomItemWithFrameZIndex(frame Rect /* not a class type */, zIndex int) CollectionLayoutGroupCustomItem {
	rv := objc.Send[CollectionLayoutGroupCustomItem](objc.ID(getCollectionLayoutGroupCustomItemClass().class), objc.Sel("customItemWithFrame:zIndex:"), frame, zIndex)
	return rv
}/* debug [class_init_methods/constructor]: NewCollectionLayoutGroupCustomItemWithFrameZIndex */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CollectionLayoutGroupCustomItem */

// Creates a custom item with the specified frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem/init(frame:)
func (cc _CollectionLayoutGroupCustomItemClass) CustomItemWithFrame(frame Rect /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("customItemWithFrame:"), frame)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CustomItemWithFrame) */


// Creates a custom item with the specified frame and vertical stacking order in relation to other items in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem/init(frame:zIndex:)
func (cc _CollectionLayoutGroupCustomItemClass) CustomItemWithFrameZIndex(frame Rect /* not a class type */, zIndex int) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("customItemWithFrame:zIndex:"), frame, zIndex)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CustomItemWithFrameZIndex) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CollectionLayoutGroupCustomItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CollectionLayoutGroupCustomItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CollectionLayoutGroupCustomItem */

// The frame of the custom item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem/frame
func (c_ CollectionLayoutGroupCustomItem) Frame() Rect /* not a class type */ {
	rv := objc.Send[Rect](c_.ID, objc.Sel("frame"))
	return rv
}/* debug [instance_properties/getter]: frame */


// The vertical stacking order of the custom item in relation to other items in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItem/zIndex
func (c_ CollectionLayoutGroupCustomItem) ZIndex() int {
	rv := objc.Send[int](c_.ID, objc.Sel("zIndex"))
	return rv
}/* debug [instance_properties/getter]: zIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCollectionLayoutGroupCustomItem */


