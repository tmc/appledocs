// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NSScrubberTextItemView */


/* debug [class_header]: Header for NSScrubberTextItemView */
// The class instance for the [ScrubberTextItemView] class.
var (
	ScrubberTextItemViewClass     _ScrubberTextItemViewClass
	ScrubberTextItemViewClassOnce sync.Once
)

func getScrubberTextItemViewClass() _ScrubberTextItemViewClass {
	ScrubberTextItemViewClassOnce.Do(func() {
		ScrubberTextItemViewClass = _ScrubberTextItemViewClass{objc.GetClass("NSScrubberTextItemView")}
	})
	return ScrubberTextItemViewClass
}

type _ScrubberTextItemViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScrubberTextItemView */
// An interface definition for the [ScrubberTextItemView] class.
type IScrubberTextItemView interface {
	IScrubberItemView
	
/* debug [class_interface_properties]: Properties for ScrubberTextItemView */
	// properties:
	TextField() ITextField
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScrubberTextItemView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScrubberTextItemView */
// Alloc allocates a new instance without initialization.
func (sc _ScrubberTextItemViewClass) Alloc() ScrubberTextItemView {
	rv := objc.Send[ScrubberTextItemView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScrubberTextItemViewClass) New() ScrubberTextItemView {
	rv := objc.Send[ScrubberTextItemView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberTextItemView) Init() ScrubberTextItemView {
	rv := objc.Send[ScrubberTextItemView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberTextItemView) Autorelease() ScrubberTextItemView {
	rv := objc.Send[ScrubberTextItemView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberTextItemView creates a new ScrubberTextItemView instance.
func NewScrubberTextItemView() ScrubberTextItemView {
	return getScrubberTextItemViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScrubberTextItemView */
// A concrete view subclass for displaying text for an item in a scrubber.
//
// Provide the text you want to display in the scrubber item to the property. If you want finer control over the appearance of the text, you can access the underlying text field using the property.


// A concrete view subclass for displaying text for an item in a scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberTextItemView
type ScrubberTextItemView struct {
	ScrubberItemView
}

// ScrubberTextItemViewFrom constructs a [ScrubberTextItemView] from an unsafe.Pointer.
//
// A concrete view subclass for displaying text for an item in a scrubber.
func ScrubberTextItemViewFrom(ptr unsafe.Pointer) ScrubberTextItemView {
	return ScrubberTextItemView{
		ScrubberItemView: ScrubberItemViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScrubberTextItemView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScrubberTextItemView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScrubberTextItemView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScrubberTextItemView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScrubberTextItemView */

// The text field that the scrubber item uses to display its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberTextItemView/textField
func (s_ ScrubberTextItemView) TextField() ITextField {
	rv := objc.Send[TextField](s_.ID, objc.Sel("textField"))
	return rv
}/* debug [instance_properties/getter]: textField */


// The text displayed for the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberTextItemView/title
func (s_ ScrubberTextItemView) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The text displayed for the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberTextItemView/title
func (s_ ScrubberTextItemView) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScrubberTextItemView */



