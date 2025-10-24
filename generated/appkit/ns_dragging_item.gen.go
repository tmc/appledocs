// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSDraggingItem */


/* debug [class_header]: Header for NSDraggingItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DraggingItem */
// An interface definition for the [DraggingItem] class.
type IDraggingItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DraggingItem */
	// properties:
	DraggingFrame() Rect /* not a class type */
	SetDraggingFrame(value Rect /* not a class type */)
	ImageComponents() []DraggingImageComponent
	ImageComponentsProvider() []objc.ID
	SetImageComponentsProvider(value []objc.ID)
	Item() objc.ID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DraggingItem */
	// methods:
	SetDraggingFrameContents(frame Rect /* not a class type */, contents objc.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DraggingItem */
// Alloc allocates a new instance without initialization.
func (dc _DraggingItemClass) Alloc() DraggingItem {
	rv := objc.Send[DraggingItem](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DraggingItem */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DraggingItem */

// Creates and returns a dragging item using the specified content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingItem/init(pasteboardWriter:)
func NewDraggingItemWithPasteboardWriter(pasteboardWriter unsafe.Pointer) DraggingItem {
	instance := getDraggingItemClass().Alloc()
	rv := objc.Send[DraggingItem](instance.ID, objc.Sel("initWithPasteboardWriter:"), pasteboardWriter)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDraggingItemWithPasteboardWriter */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DraggingItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DraggingItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DraggingItem */

// Sets the item’s dragging frame and contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingItem/setDraggingFrame(_:contents:)
func (d_ DraggingItem) SetDraggingFrameContents(frame Rect /* not a class type */, contents objc.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDraggingFrame:contents:"), frame, contents)
}/* debug [instance_methods/method]: SetDraggingFrameContents */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DraggingItem */

// The frame of the dragging item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingItem/draggingFrame
func (d_ DraggingItem) DraggingFrame() Rect /* not a class type */ {
	rv := objc.Send[Rect](d_.ID, objc.Sel("draggingFrame"))
	return rv
}/* debug [instance_properties/getter]: draggingFrame */


// The frame of the dragging item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingItem/draggingFrame
func (d_ DraggingItem) SetDraggingFrame(value Rect /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDraggingFrame:"), value)
}/* debug [instance_properties/setter]: draggingFrame */


// An array of dragging image components to use to create the drag image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingItem/imageComponents
func (d_ DraggingItem) ImageComponents() []DraggingImageComponent {
	rv := objc.Send[[]DraggingImageComponent](d_.ID, objc.Sel("imageComponents"))
	return rv
}/* debug [instance_properties/getter]: imageComponents */


// An array of blocks that provide the dragging image components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingItem/imageComponentsProvider
func (d_ DraggingItem) ImageComponentsProvider() []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("imageComponentsProvider"))
	return rv
}/* debug [instance_properties/getter]: imageComponentsProvider */


// An array of blocks that provide the dragging image components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingItem/imageComponentsProvider
func (d_ DraggingItem) SetImageComponentsProvider(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](d_.ID, objc.Sel("setImageComponentsProvider:"), nsArray)
}/* debug [instance_properties/setter]: imageComponentsProvider */


// The pasteboard reader or writer object dependent on the context where you use the dragging item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingItem/item
func (d_ DraggingItem) Item() objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("item"))
	return rv
}/* debug [instance_properties/getter]: item */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSDraggingItem */


