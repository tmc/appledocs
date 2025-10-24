// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNTextObservation */


/* debug [class_header]: Header for VNTextObservation */
// The class instance for the [TextObservation] class.
var (
	TextObservationClass     _TextObservationClass
	TextObservationClassOnce sync.Once
)

func getTextObservationClass() _TextObservationClass {
	TextObservationClassOnce.Do(func() {
		TextObservationClass = _TextObservationClass{objc.GetClass("VNTextObservation")}
	})
	return TextObservationClass
}

type _TextObservationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextObservation */
// An interface definition for the [TextObservation] class.
type ITextObservation interface {
	IRectangleObservation
	
/* debug [class_interface_properties]: Properties for TextObservation */
	// properties:
	CharacterBoxes() []RectangleObservation
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextObservation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextObservation */
// Alloc allocates a new instance without initialization.
func (tc _TextObservationClass) Alloc() TextObservation {
	rv := objc.Send[TextObservation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextObservationClass) New() TextObservation {
	rv := objc.Send[TextObservation](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextObservation) Init() TextObservation {
	rv := objc.Send[TextObservation](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextObservation) Autorelease() TextObservation {
	rv := objc.Send[TextObservation](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextObservation creates a new TextObservation instance.
func NewTextObservation() TextObservation {
	return getTextObservationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextObservation */
// Information about regions of text that an image-analysis request detects.
//
// This type of observation results from a . It expresses the location of each detected character by its bounding box.


// Information about regions of text that an image-analysis request detects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTextObservation
type TextObservation struct {
	RectangleObservation
}

// TextObservationFrom constructs a [TextObservation] from an unsafe.Pointer.
//
// Information about regions of text that an image-analysis request detects.
func TextObservationFrom(ptr unsafe.Pointer) TextObservation {
	return TextObservation{
		RectangleObservation: RectangleObservationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextObservation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextObservation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextObservation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextObservation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextObservation */

// An array of detected individual character bounding boxes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTextObservation/characterBoxes
func (t_ TextObservation) CharacterBoxes() []RectangleObservation {
	rv := objc.Send[[]RectangleObservation](t_.ID, objc.Sel("characterBoxes"))
	return rv
}/* debug [instance_properties/getter]: characterBoxes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNTextObservation */



