// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSStringDrawingContext */


/* debug [class_header]: Header for NSStringDrawingContext */
// The class instance for the [StringDrawingContext] class.
var (
	StringDrawingContextClass     _StringDrawingContextClass
	StringDrawingContextClassOnce sync.Once
)

func getStringDrawingContextClass() _StringDrawingContextClass {
	StringDrawingContextClassOnce.Do(func() {
		StringDrawingContextClass = _StringDrawingContextClass{objc.GetClass("NSStringDrawingContext")}
	})
	return StringDrawingContextClass
}

type _StringDrawingContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for StringDrawingContext */
// An interface definition for the [StringDrawingContext] class.
type IStringDrawingContext interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for StringDrawingContext */
	// properties:
	MinimumScaleFactor() float64
	SetMinimumScaleFactor(value float64)
	ActualScaleFactor() float64
	SetActualScaleFactor(value float64)
	TotalBounds() corefoundation.CGRect
	SetTotalBounds(value corefoundation.CGRect)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for StringDrawingContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for StringDrawingContext */
// Alloc allocates a new instance without initialization.
func (sc _StringDrawingContextClass) Alloc() StringDrawingContext {
	rv := objc.Send[StringDrawingContext](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StringDrawingContextClass) New() StringDrawingContext {
	rv := objc.Send[StringDrawingContext](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StringDrawingContext) Init() StringDrawingContext {
	rv := objc.Send[StringDrawingContext](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StringDrawingContext) Autorelease() StringDrawingContext {
	rv := objc.Send[StringDrawingContext](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStringDrawingContext creates a new StringDrawingContext instance.
func NewStringDrawingContext() StringDrawingContext {
	return getStringDrawingContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for StringDrawingContext */
// An object that manages metrics for drawing attributed strings.
//
// Prior to drawing, you can create an instance of this class and use it to specify the minimum scale factor and tracking adjustments for a string. After drawing, you can retrieve the actual values that were used during drawing. To use this class, allocate and initialize a new instance, set the minimum values, and pass your object to one of the corresponding methods that take the context object as a parameter. Upon completion of drawing, you can use the actual drawing values to make adjustments or record where the string was actually drawn.


// An object that manages metrics for drawing attributed strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStringDrawingContext
type StringDrawingContext struct {
	objectivec.Object
}

// StringDrawingContextFrom constructs a [StringDrawingContext] from an unsafe.Pointer.
//
// An object that manages metrics for drawing attributed strings.
func StringDrawingContextFrom(ptr unsafe.Pointer) StringDrawingContext {
	return StringDrawingContext{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for StringDrawingContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for StringDrawingContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for StringDrawingContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for StringDrawingContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for StringDrawingContext */

// The scale factor that determines the smallest font size to use during drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStringDrawingContext/minimumScaleFactor
func (s_ StringDrawingContext) MinimumScaleFactor() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minimumScaleFactor"))
	return rv
}/* debug [instance_properties/getter]: minimumScaleFactor */


// The scale factor that determines the smallest font size to use during drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStringDrawingContext/minimumScaleFactor
func (s_ StringDrawingContext) SetMinimumScaleFactor(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumScaleFactor:"), value)
}/* debug [instance_properties/setter]: minimumScaleFactor */


// The actual scale factor that the system applied to the font during drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstringdrawingcontext/actualscalefactor
func (s_ StringDrawingContext) ActualScaleFactor() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("actualScaleFactor"))
	return rv
}/* debug [instance_properties/getter]: actualScaleFactor */


// The actual scale factor that the system applied to the font during drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstringdrawingcontext/actualscalefactor
func (s_ StringDrawingContext) SetActualScaleFactor(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setActualScaleFactor:"), value)
}/* debug [instance_properties/setter]: actualScaleFactor */


// The most recent bounding rectangle that the system used to draw the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstringdrawingcontext/totalbounds
func (s_ StringDrawingContext) TotalBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s_.ID, objc.Sel("totalBounds"))
	return rv
}/* debug [instance_properties/getter]: totalBounds */


// The most recent bounding rectangle that the system used to draw the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstringdrawingcontext/totalbounds
func (s_ StringDrawingContext) SetTotalBounds(value corefoundation.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTotalBounds:"), value)
}/* debug [instance_properties/setter]: totalBounds */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSStringDrawingContext */



