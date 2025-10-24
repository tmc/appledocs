// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSScrubberLayoutAttributes */


/* debug [class_header]: Header for NSScrubberLayoutAttributes */
// The class instance for the [ScrubberLayoutAttributes] class.
var (
	ScrubberLayoutAttributesClass     _ScrubberLayoutAttributesClass
	ScrubberLayoutAttributesClassOnce sync.Once
)

func getScrubberLayoutAttributesClass() _ScrubberLayoutAttributesClass {
	ScrubberLayoutAttributesClassOnce.Do(func() {
		ScrubberLayoutAttributesClass = _ScrubberLayoutAttributesClass{objc.GetClass("NSScrubberLayoutAttributes")}
	})
	return ScrubberLayoutAttributesClass
}

type _ScrubberLayoutAttributesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScrubberLayoutAttributes */
// An interface definition for the [ScrubberLayoutAttributes] class.
type IScrubberLayoutAttributes interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ScrubberLayoutAttributes */
	// properties:
	Alpha() float64
	SetAlpha(value float64)
	Frame() Rect /* not a class type */
	SetFrame(value Rect /* not a class type */)
	ItemIndex() int
	SetItemIndex(value int)
	Hash() int
	SetHash(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScrubberLayoutAttributes */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScrubberLayoutAttributes */
// Alloc allocates a new instance without initialization.
func (sc _ScrubberLayoutAttributesClass) Alloc() ScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScrubberLayoutAttributesClass) New() ScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberLayoutAttributes) Init() ScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberLayoutAttributes) Autorelease() ScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberLayoutAttributes creates a new ScrubberLayoutAttributes instance.
func NewScrubberLayoutAttributes() ScrubberLayoutAttributes {
	return getScrubberLayoutAttributesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScrubberLayoutAttributes */
// The layout of a scrubber item.
//
// A layout attributes object is the model for the layout of a single item in a scrubber control. If you require model attributes in addition to those provided by this class, create a subclass and add appropriate attributes. Subclasses must implement , and the protocol.


// The layout of a scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes
type ScrubberLayoutAttributes struct {
	objectivec.Object
}

// ScrubberLayoutAttributesFrom constructs a [ScrubberLayoutAttributes] from an unsafe.Pointer.
//
// The layout of a scrubber item.
func ScrubberLayoutAttributesFrom(ptr unsafe.Pointer) ScrubberLayoutAttributes {
	return ScrubberLayoutAttributes{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScrubberLayoutAttributes */

// Creates a new layout attributes object for the specified scrubber item index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/init(forItemAt:)
func NewScrubberLayoutAttributesForItemAtIndex(index int) ScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](objc.ID(getScrubberLayoutAttributesClass().class), objc.Sel("layoutAttributesForItemAtIndex:"), index)
	return rv
}/* debug [class_init_methods/constructor]: NewScrubberLayoutAttributesForItemAtIndex */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScrubberLayoutAttributes */

// Creates a new layout attributes object for the specified scrubber item index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/init(forItemAt:)
func (sc _ScrubberLayoutAttributesClass) LayoutAttributesForItemAtIndex(index int) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("layoutAttributesForItemAtIndex:"), index)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayoutAttributesForItemAtIndex) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScrubberLayoutAttributes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScrubberLayoutAttributes */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScrubberLayoutAttributes */

// The item’s alpha value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/alpha
func (s_ ScrubberLayoutAttributes) Alpha() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("alpha"))
	return rv
}/* debug [instance_properties/getter]: alpha */


// The item’s alpha value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/alpha
func (s_ ScrubberLayoutAttributes) SetAlpha(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAlpha:"), value)
}/* debug [instance_properties/setter]: alpha */


// The frame of the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/frame
func (s_ ScrubberLayoutAttributes) Frame() Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("frame"))
	return rv
}/* debug [instance_properties/getter]: frame */


// The frame of the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/frame
func (s_ ScrubberLayoutAttributes) SetFrame(value Rect /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFrame:"), value)
}/* debug [instance_properties/setter]: frame */


// The index of the scrubber item that is represented by the item’s layout attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/itemIndex
func (s_ ScrubberLayoutAttributes) ItemIndex() int {
	rv := objc.Send[int](s_.ID, objc.Sel("itemIndex"))
	return rv
}/* debug [instance_properties/getter]: itemIndex */


// The index of the scrubber item that is represented by the item’s layout attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/itemIndex
func (s_ ScrubberLayoutAttributes) SetItemIndex(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setItemIndex:"), value)
}/* debug [instance_properties/setter]: itemIndex */


// Returns an integer that can be used as a table address in a hash table structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash
func (s_ ScrubberLayoutAttributes) Hash() int {
	rv := objc.Send[int](s_.ID, objc.Sel("hash"))
	return rv
}/* debug [instance_properties/getter]: hash */


// Returns an integer that can be used as a table address in a hash table structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash
func (s_ ScrubberLayoutAttributes) SetHash(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHash:"), value)
}/* debug [instance_properties/setter]: hash */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScrubberLayoutAttributes */


