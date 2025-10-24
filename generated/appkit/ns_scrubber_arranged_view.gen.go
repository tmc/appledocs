// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSScrubberArrangedView */


/* debug [class_header]: Header for NSScrubberArrangedView */
// The class instance for the [ScrubberArrangedView] class.
var (
	ScrubberArrangedViewClass     _ScrubberArrangedViewClass
	ScrubberArrangedViewClassOnce sync.Once
)

func getScrubberArrangedViewClass() _ScrubberArrangedViewClass {
	ScrubberArrangedViewClassOnce.Do(func() {
		ScrubberArrangedViewClass = _ScrubberArrangedViewClass{objc.GetClass("NSScrubberArrangedView")}
	})
	return ScrubberArrangedViewClass
}

type _ScrubberArrangedViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScrubberArrangedView */
// An interface definition for the [ScrubberArrangedView] class.
type IScrubberArrangedView interface {
	IView
	
/* debug [class_interface_properties]: Properties for ScrubberArrangedView */
	// properties:
	Highlighted() bool
	SetHighlighted(value bool)
	Selected() bool
	SetSelected(value bool)
	IsHighlighted() bool
	SetIsHighlighted(value bool)
	IsSelected() bool
	SetIsSelected(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScrubberArrangedView */
	// methods:
	ApplyLayoutAttributes(layoutAttributes IScrubberLayoutAttributes)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScrubberArrangedView */
// Alloc allocates a new instance without initialization.
func (sc _ScrubberArrangedViewClass) Alloc() ScrubberArrangedView {
	rv := objc.Send[ScrubberArrangedView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScrubberArrangedViewClass) New() ScrubberArrangedView {
	rv := objc.Send[ScrubberArrangedView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberArrangedView) Init() ScrubberArrangedView {
	rv := objc.Send[ScrubberArrangedView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberArrangedView) Autorelease() ScrubberArrangedView {
	rv := objc.Send[ScrubberArrangedView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberArrangedView creates a new ScrubberArrangedView instance.
func NewScrubberArrangedView() ScrubberArrangedView {
	return getScrubberArrangedViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScrubberArrangedView */
// An abstract base class for the views whose layout is managed by a scrubber.


// An abstract base class for the views whose layout is managed by a scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView
type ScrubberArrangedView struct {
	View
}

// ScrubberArrangedViewFrom constructs a [ScrubberArrangedView] from an unsafe.Pointer.
//
// An abstract base class for the views whose layout is managed by a scrubber.
func ScrubberArrangedViewFrom(ptr unsafe.Pointer) ScrubberArrangedView {
	return ScrubberArrangedView{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScrubberArrangedView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScrubberArrangedView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScrubberArrangedView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScrubberArrangedView */

// Updates the layout of the arranged view to respect the provided layout attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/apply(_:)
func (s_ ScrubberArrangedView) ApplyLayoutAttributes(layoutAttributes IScrubberLayoutAttributes) {
	objc.Send[objc.ID](s_.ID, objc.Sel("applyLayoutAttributes:"), layoutAttributes)
}/* debug [instance_methods/method]: ApplyLayoutAttributes */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScrubberArrangedView */

// A Boolean value that specifies whether the view is currently highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/isHighlighted
func (s_ ScrubberArrangedView) Highlighted() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("highlighted"))
	return rv
}/* debug [instance_properties/getter]: highlighted */


// A Boolean value that specifies whether the view is currently highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/isHighlighted
func (s_ ScrubberArrangedView) SetHighlighted(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHighlighted:"), value)
}/* debug [instance_properties/setter]: highlighted */


// A Boolean value that specifies whether the current view is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/isSelected
func (s_ ScrubberArrangedView) Selected() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("selected"))
	return rv
}/* debug [instance_properties/getter]: selected */


// A Boolean value that specifies whether the current view is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/isSelected
func (s_ ScrubberArrangedView) SetSelected(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelected:"), value)
}/* debug [instance_properties/setter]: selected */


// A Boolean value that specifies whether the view is currently highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberarrangedview/ishighlighted
func (s_ ScrubberArrangedView) IsHighlighted() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isHighlighted"))
	return rv
}/* debug [instance_properties/getter]: isHighlighted */


// A Boolean value that specifies whether the view is currently highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberarrangedview/ishighlighted
func (s_ ScrubberArrangedView) SetIsHighlighted(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsHighlighted:"), value)
}/* debug [instance_properties/setter]: isHighlighted */


// A Boolean value that specifies whether the current view is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberarrangedview/isselected
func (s_ ScrubberArrangedView) IsSelected() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSelected"))
	return rv
}/* debug [instance_properties/getter]: isSelected */


// A Boolean value that specifies whether the current view is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberarrangedview/isselected
func (s_ ScrubberArrangedView) SetIsSelected(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsSelected:"), value)
}/* debug [instance_properties/setter]: isSelected */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScrubberArrangedView */



