// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSScrubberItemView */


/* debug [class_header]: Header for NSScrubberItemView */
// The class instance for the [ScrubberItemView] class.
var (
	ScrubberItemViewClass     _ScrubberItemViewClass
	ScrubberItemViewClassOnce sync.Once
)

func getScrubberItemViewClass() _ScrubberItemViewClass {
	ScrubberItemViewClassOnce.Do(func() {
		ScrubberItemViewClass = _ScrubberItemViewClass{objc.GetClass("NSScrubberItemView")}
	})
	return ScrubberItemViewClass
}

type _ScrubberItemViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScrubberItemView */
// An interface definition for the [ScrubberItemView] class.
type IScrubberItemView interface {
	IScrubberArrangedView
	
/* debug [class_interface_properties]: Properties for ScrubberItemView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScrubberItemView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScrubberItemView */
// Alloc allocates a new instance without initialization.
func (sc _ScrubberItemViewClass) Alloc() ScrubberItemView {
	rv := objc.Send[ScrubberItemView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScrubberItemViewClass) New() ScrubberItemView {
	rv := objc.Send[ScrubberItemView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberItemView) Init() ScrubberItemView {
	rv := objc.Send[ScrubberItemView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberItemView) Autorelease() ScrubberItemView {
	rv := objc.Send[ScrubberItemView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberItemView creates a new ScrubberItemView instance.
func NewScrubberItemView() ScrubberItemView {
	return getScrubberItemViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScrubberItemView */
// An item at a specific index position in the scrubber.


// An item at a specific index position in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberItemView
type ScrubberItemView struct {
	ScrubberArrangedView
}

// ScrubberItemViewFrom constructs a [ScrubberItemView] from an unsafe.Pointer.
//
// An item at a specific index position in the scrubber.
func ScrubberItemViewFrom(ptr unsafe.Pointer) ScrubberItemView {
	return ScrubberItemView{
		ScrubberArrangedView: ScrubberArrangedViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScrubberItemView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScrubberItemView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScrubberItemView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScrubberItemView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScrubberItemView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScrubberItemView */



