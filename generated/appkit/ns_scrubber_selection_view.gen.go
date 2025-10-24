// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSScrubberSelectionView */


/* debug [class_header]: Header for NSScrubberSelectionView */
// The class instance for the [ScrubberSelectionView] class.
var (
	ScrubberSelectionViewClass     _ScrubberSelectionViewClass
	ScrubberSelectionViewClassOnce sync.Once
)

func getScrubberSelectionViewClass() _ScrubberSelectionViewClass {
	ScrubberSelectionViewClassOnce.Do(func() {
		ScrubberSelectionViewClass = _ScrubberSelectionViewClass{objc.GetClass("NSScrubberSelectionView")}
	})
	return ScrubberSelectionViewClass
}

type _ScrubberSelectionViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScrubberSelectionView */
// An interface definition for the [ScrubberSelectionView] class.
type IScrubberSelectionView interface {
	IScrubberArrangedView
	
/* debug [class_interface_properties]: Properties for ScrubberSelectionView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScrubberSelectionView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScrubberSelectionView */
// Alloc allocates a new instance without initialization.
func (sc _ScrubberSelectionViewClass) Alloc() ScrubberSelectionView {
	rv := objc.Send[ScrubberSelectionView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScrubberSelectionViewClass) New() ScrubberSelectionView {
	rv := objc.Send[ScrubberSelectionView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberSelectionView) Init() ScrubberSelectionView {
	rv := objc.Send[ScrubberSelectionView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberSelectionView) Autorelease() ScrubberSelectionView {
	rv := objc.Send[ScrubberSelectionView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberSelectionView creates a new ScrubberSelectionView instance.
func NewScrubberSelectionView() ScrubberSelectionView {
	return getScrubberSelectionViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScrubberSelectionView */
// An abstract base class for specifying the appearance of a highlighted or selected item in a scrubber.
//
// Create a subclass to customize the selection or highlight appearance of an item in your scrubber control. You need to return an instance of your subclass from the method on .


// An abstract base class for specifying the appearance of a highlighted or selected item in a scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionView
type ScrubberSelectionView struct {
	ScrubberArrangedView
}

// ScrubberSelectionViewFrom constructs a [ScrubberSelectionView] from an unsafe.Pointer.
//
// An abstract base class for specifying the appearance of a highlighted or selected item in a scrubber.
func ScrubberSelectionViewFrom(ptr unsafe.Pointer) ScrubberSelectionView {
	return ScrubberSelectionView{
		ScrubberArrangedView: ScrubberArrangedViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScrubberSelectionView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScrubberSelectionView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScrubberSelectionView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScrubberSelectionView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScrubberSelectionView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScrubberSelectionView */



