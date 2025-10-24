// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SCDisplay */


/* debug [class_header]: Header for SCDisplay */
// The class instance for the [Display] class.
var (
	DisplayClass     _DisplayClass
	DisplayClassOnce sync.Once
)

func getDisplayClass() _DisplayClass {
	DisplayClassOnce.Do(func() {
		DisplayClass = _DisplayClass{objc.GetClass("SCDisplay")}
	})
	return DisplayClass
}

type _DisplayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Display */
// An interface definition for the [Display] class.
type IDisplay interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Display */
	// properties:
	DisplayID() DirectDisplayID /* not a class type */
	Frame() corefoundation.CGRect
	Height() int
	Width() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Display */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Display */
// Alloc allocates a new instance without initialization.
func (dc _DisplayClass) Alloc() Display {
	rv := objc.Send[Display](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DisplayClass) New() Display {
	rv := objc.Send[Display](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ Display) Init() Display {
	rv := objc.Send[Display](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ Display) Autorelease() Display {
	rv := objc.Send[Display](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDisplay creates a new Display instance.
func NewDisplay() Display {
	return getDisplayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Display */
// An instance that represents a display device.
//
// A display object represents a physical display connected to a Mac. Query the display to retrieve its unique identifier and onscreen coordinates. Retrieve the available displays from an instance of . Select a display to capture and use it to create an instance of . Apply the filter to an instance of to limit its output to content matching your criteria.


// An instance that represents a display device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCDisplay
type Display struct {
	objectivec.Object
}

// DisplayFrom constructs a [Display] from an unsafe.Pointer.
//
// An instance that represents a display device.
func DisplayFrom(ptr unsafe.Pointer) Display {
	return Display{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Display *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Display */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Display */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Display */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Display */

// The Core Graphics display identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCDisplay/displayID
func (d_ Display) DisplayID() DirectDisplayID /* not a class type */ {
	rv := objc.Send[DirectDisplayID](d_.ID, objc.Sel("displayID"))
	return rv
}/* debug [instance_properties/getter]: displayID */


// The frame of the display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCDisplay/frame
func (d_ Display) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](d_.ID, objc.Sel("frame"))
	return rv
}/* debug [instance_properties/getter]: frame */


// The height of the display in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCDisplay/height
func (d_ Display) Height() int {
	rv := objc.Send[int](d_.ID, objc.Sel("height"))
	return rv
}/* debug [instance_properties/getter]: height */


// The width of the display in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCDisplay/width
func (d_ Display) Width() int {
	rv := objc.Send[int](d_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SCDisplay */



