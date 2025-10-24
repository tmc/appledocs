// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSScrubberProportionalLayout */


/* debug [class_header]: Header for NSScrubberProportionalLayout */
// The class instance for the [ScrubberProportionalLayout] class.
var (
	ScrubberProportionalLayoutClass     _ScrubberProportionalLayoutClass
	ScrubberProportionalLayoutClassOnce sync.Once
)

func getScrubberProportionalLayoutClass() _ScrubberProportionalLayoutClass {
	ScrubberProportionalLayoutClassOnce.Do(func() {
		ScrubberProportionalLayoutClass = _ScrubberProportionalLayoutClass{objc.GetClass("NSScrubberProportionalLayout")}
	})
	return ScrubberProportionalLayoutClass
}

type _ScrubberProportionalLayoutClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScrubberProportionalLayout */
// An interface definition for the [ScrubberProportionalLayout] class.
type IScrubberProportionalLayout interface {
	IScrubberLayout
	
/* debug [class_interface_properties]: Properties for ScrubberProportionalLayout */
	// properties:
	NumberOfVisibleItems() int
	SetNumberOfVisibleItems(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScrubberProportionalLayout */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScrubberProportionalLayout */
// Alloc allocates a new instance without initialization.
func (sc _ScrubberProportionalLayoutClass) Alloc() ScrubberProportionalLayout {
	rv := objc.Send[ScrubberProportionalLayout](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScrubberProportionalLayoutClass) New() ScrubberProportionalLayout {
	rv := objc.Send[ScrubberProportionalLayout](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberProportionalLayout) Init() ScrubberProportionalLayout {
	rv := objc.Send[ScrubberProportionalLayout](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberProportionalLayout) Autorelease() ScrubberProportionalLayout {
	rv := objc.Send[ScrubberProportionalLayout](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberProportionalLayout creates a new ScrubberProportionalLayout instance.
func NewScrubberProportionalLayout() ScrubberProportionalLayout {
	return getScrubberProportionalLayoutClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScrubberProportionalLayout */
// A concrete layout object that sizes each item to some fraction of the scrubber’s visible size.


// A concrete layout object that sizes each item to some fraction of the scrubber’s visible size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberProportionalLayout
type ScrubberProportionalLayout struct {
	ScrubberLayout
}

// ScrubberProportionalLayoutFrom constructs a [ScrubberProportionalLayout] from an unsafe.Pointer.
//
// A concrete layout object that sizes each item to some fraction of the scrubber’s visible size.
func ScrubberProportionalLayoutFrom(ptr unsafe.Pointer) ScrubberProportionalLayout {
	return ScrubberProportionalLayout{
		ScrubberLayout: ScrubberLayoutFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScrubberProportionalLayout */

// Initializes and returns a newly allocated proprotional layout object from a storyboard or nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberProportionalLayout/init(coder:)
func NewScrubberProportionalLayoutWithCoder(coder foundation.Coder) ScrubberProportionalLayout {
	instance := getScrubberProportionalLayoutClass().Alloc()
	rv := objc.Send[ScrubberProportionalLayout](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewScrubberProportionalLayoutWithCoder */


// Initializes and returns a newly allocated proportional layout, configured to display the given number of items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberProportionalLayout/init(numberOfVisibleItems:)
func NewScrubberProportionalLayoutWithNumberOfVisibleItems(numberOfVisibleItems int) ScrubberProportionalLayout {
	instance := getScrubberProportionalLayoutClass().Alloc()
	rv := objc.Send[ScrubberProportionalLayout](instance.ID, objc.Sel("initWithNumberOfVisibleItems:"), numberOfVisibleItems)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewScrubberProportionalLayoutWithNumberOfVisibleItems */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScrubberProportionalLayout */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScrubberProportionalLayout */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScrubberProportionalLayout */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScrubberProportionalLayout */

// The number of items visible in the scrubber at once.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberProportionalLayout/numberOfVisibleItems
func (s_ ScrubberProportionalLayout) NumberOfVisibleItems() int {
	rv := objc.Send[int](s_.ID, objc.Sel("numberOfVisibleItems"))
	return rv
}/* debug [instance_properties/getter]: numberOfVisibleItems */


// The number of items visible in the scrubber at once.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberProportionalLayout/numberOfVisibleItems
func (s_ ScrubberProportionalLayout) SetNumberOfVisibleItems(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNumberOfVisibleItems:"), value)
}/* debug [instance_properties/setter]: numberOfVisibleItems */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScrubberProportionalLayout */


