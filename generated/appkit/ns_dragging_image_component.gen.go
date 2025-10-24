// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DraggingImageComponent] class.
var (
	DraggingImageComponentClass     _DraggingImageComponentClass
	DraggingImageComponentClassOnce sync.Once
)

func getDraggingImageComponentClass() _DraggingImageComponentClass {
	DraggingImageComponentClassOnce.Do(func() {
		DraggingImageComponentClass = _DraggingImageComponentClass{objc.GetClass("NSDraggingImageComponent")}
	})
	return DraggingImageComponentClass
}

type _DraggingImageComponentClass struct {
	class objc.Class
}

// An interface definition for the [DraggingImageComponent] class.
type IDraggingImageComponent interface {
	objectivec.IObject
	// properties:
	Contents() objc.ID
	SetContents(value objc.ID)
	Frame() objc.IObject /* cross-framework: Rect */
	SetFrame(value objc.IObject /* cross-framework: Rect */)
	Key() objc.IObject /* cross-framework: DraggingImageComponentKey */
	SetKey(value objc.IObject /* cross-framework: DraggingImageComponentKey */)
	// methods:
}

// A single object in a dragging item.
//
// An array of instances are composited together to create the dragging image for an . instances can simply be considered as named images with a location used by an instance.


// A single object in a dragging item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent
type DraggingImageComponent struct {
	objectivec.Object
}

// DraggingImageComponentFrom constructs a [DraggingImageComponent] from an unsafe.Pointer.
//
// A single object in a dragging item.
func DraggingImageComponentFrom(ptr unsafe.Pointer) DraggingImageComponent {
	return DraggingImageComponent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DraggingImageComponentClass) Alloc() DraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DraggingImageComponentClass) New() DraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DraggingImageComponent) Init() DraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DraggingImageComponent) Autorelease() DraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDraggingImageComponent creates a new DraggingImageComponent instance.
func NewDraggingImageComponent() DraggingImageComponent {
	return getDraggingImageComponentClass().New()
}



// Initializes and returns a dragging image component with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/init(key:)
func NewDraggingImageComponentWithKey(key objc.IObject /* cross-framework: DraggingImageComponentKey */) DraggingImageComponent {
	instance := getDraggingImageComponentClass().Alloc()
	rv := objc.Send[DraggingImageComponent](instance.ID, objc.Sel("initWithKey:"), key)
	rv.Autorelease()
	return rv
}



// Creates and returns a dragging image component with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/draggingImageComponentWithKey:
func (dc _DraggingImageComponentClass) DraggingImageComponentWithKey(key objc.IObject /* cross-framework: DraggingImageComponentKey */) IDraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](objc.ID(dc.class), objc.Sel("draggingImageComponentWithKey:"), key)
	return rv
}


// An object providing the image contents of the component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/contents
func (d_ DraggingImageComponent) Contents() objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("contents"))
	return rv
}


// An object providing the image contents of the component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/contents
func (d_ DraggingImageComponent) SetContents(value objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContents:"), value)
}


// The coordinate space is the bounds of the parent dragging item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/frame
func (d_ DraggingImageComponent) Frame() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](d_.ID, objc.Sel("frame"))
	return rv
}


// The coordinate space is the bounds of the parent dragging item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/frame
func (d_ DraggingImageComponent) SetFrame(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFrame:"), value)
}


// The unique name of this image component instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/key
func (d_ DraggingImageComponent) Key() objc.IObject /* cross-framework: DraggingImageComponentKey */ {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("key"))
	return rv
}


// The unique name of this image component instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/key
func (d_ DraggingImageComponent) SetKey(value objc.IObject /* cross-framework: DraggingImageComponentKey */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setKey:"), value)
}


