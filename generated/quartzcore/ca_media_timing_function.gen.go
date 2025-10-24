// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CAMediaTimingFunction */


/* debug [class_header]: Header for CAMediaTimingFunction */
// The class instance for the [MediaTimingFunction] class.
var (
	MediaTimingFunctionClass     _MediaTimingFunctionClass
	MediaTimingFunctionClassOnce sync.Once
)

func getMediaTimingFunctionClass() _MediaTimingFunctionClass {
	MediaTimingFunctionClassOnce.Do(func() {
		MediaTimingFunctionClass = _MediaTimingFunctionClass{objc.GetClass("CAMediaTimingFunction")}
	})
	return MediaTimingFunctionClass
}

type _MediaTimingFunctionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaTimingFunction */
// An interface definition for the [MediaTimingFunction] class.
type IMediaTimingFunction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MediaTimingFunction */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaTimingFunction */
	// methods:
	GetControlPointAtIndexValues(idx uintptr /* not a class type */, ptr objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaTimingFunction */
// Alloc allocates a new instance without initialization.
func (mc _MediaTimingFunctionClass) Alloc() MediaTimingFunction {
	rv := objc.Send[MediaTimingFunction](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MediaTimingFunctionClass) New() MediaTimingFunction {
	rv := objc.Send[MediaTimingFunction](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaTimingFunction) Init() MediaTimingFunction {
	rv := objc.Send[MediaTimingFunction](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaTimingFunction) Autorelease() MediaTimingFunction {
	rv := objc.Send[MediaTimingFunction](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaTimingFunction creates a new MediaTimingFunction instance.
func NewMediaTimingFunction() MediaTimingFunction {
	return getMediaTimingFunctionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaTimingFunction */
// A function that defines the pacing of an animation as a timing curve.
//
// represents one segment of a function that defines the pacing of an animation as a timing curve. The function maps an input time normalized to the range to an output time also in the range . You can create a media timing function by supplying your own cubic Bézier curve control points using the method or by using one of the predefined timing functions.


// A function that defines the pacing of an animation as a timing curve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunction
type MediaTimingFunction struct {
	objectivec.Object
}

// MediaTimingFunctionFrom constructs a [MediaTimingFunction] from an unsafe.Pointer.
//
// A function that defines the pacing of an animation as a timing curve.
func MediaTimingFunctionFrom(ptr unsafe.Pointer) MediaTimingFunction {
	return MediaTimingFunction{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaTimingFunction */

// Returns an initialized timing function modeled as a cubic Bézier curve using the specified control points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunction/init(controlPoints:_:_:_:)
func NewMediaTimingFunctionWithControlPoints(c1x float32, c1y float32, c2x float32, c2y float32) MediaTimingFunction {
	instance := getMediaTimingFunctionClass().Alloc()
	rv := objc.Send[MediaTimingFunction](instance.ID, objc.Sel("initWithControlPoints::::"), c1x, c1y, c2x, c2y)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMediaTimingFunctionWithControlPoints */


// Creates and returns a new instance of configured with the predefined timing function specified by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunction/init(name:)
func NewMediaTimingFunctionWithName(name MediaTimingFunctionName /* typedef */) MediaTimingFunction {
	rv := objc.Send[MediaTimingFunction](objc.ID(getMediaTimingFunctionClass().class), objc.Sel("functionWithName:"), name)
	return rv
}/* debug [class_init_methods/constructor]: NewMediaTimingFunctionWithName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaTimingFunction */

// Creates and returns a new instance of timing function modeled as a cubic Bézier curve using the specified control points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunction/functionWithControlPoints::::
func (mc _MediaTimingFunctionClass) FunctionWithControlPoints(c1x float32, c1y float32, c2x float32, c2y float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("functionWithControlPoints::::"), c1x, c1y, c2x, c2y)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FunctionWithControlPoints) */


// Creates and returns a new instance of configured with the predefined timing function specified by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunction/init(name:)
func (mc _MediaTimingFunctionClass) FunctionWithName(name MediaTimingFunctionName /* typedef */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("functionWithName:"), name)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FunctionWithName) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaTimingFunction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaTimingFunction */

// Returns the control point for the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunction/getControlPoint(at:values:)
func (m_ MediaTimingFunction) GetControlPointAtIndexValues(idx uintptr /* not a class type */, ptr objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getControlPointAtIndex:values:"), idx, ptr)
}/* debug [instance_methods/method]: GetControlPointAtIndexValues */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaTimingFunction */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAMediaTimingFunction */


