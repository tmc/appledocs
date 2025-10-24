// Code generated from Apple documentation for MediaAccessibility. DO NOT EDIT.

package mediaaccessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MAFlashingLightsProcessorResult */


/* debug [class_header]: Header for MAFlashingLightsProcessorResult */
// The class instance for the [MAFlashingLightsProcessorResult] class.
var (
	MAFlashingLightsProcessorResultClass     _MAFlashingLightsProcessorResultClass
	MAFlashingLightsProcessorResultClassOnce sync.Once
)

func getMAFlashingLightsProcessorResultClass() _MAFlashingLightsProcessorResultClass {
	MAFlashingLightsProcessorResultClassOnce.Do(func() {
		MAFlashingLightsProcessorResultClass = _MAFlashingLightsProcessorResultClass{objc.GetClass("MAFlashingLightsProcessorResult")}
	})
	return MAFlashingLightsProcessorResultClass
}

type _MAFlashingLightsProcessorResultClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MAFlashingLightsProcessorResult */
// An interface definition for the [MAFlashingLightsProcessorResult] class.
type IMAFlashingLightsProcessorResult interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MAFlashingLightsProcessorResult */
	// properties:
	IntensityLevel() float32
	MitigationLevel() float32
	SurfaceProcessed() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MAFlashingLightsProcessorResult */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MAFlashingLightsProcessorResult */
// Alloc allocates a new instance without initialization.
func (mc _MAFlashingLightsProcessorResultClass) Alloc() MAFlashingLightsProcessorResult {
	rv := objc.Send[MAFlashingLightsProcessorResult](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MAFlashingLightsProcessorResultClass) New() MAFlashingLightsProcessorResult {
	rv := objc.Send[MAFlashingLightsProcessorResult](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MAFlashingLightsProcessorResult) Init() MAFlashingLightsProcessorResult {
	rv := objc.Send[MAFlashingLightsProcessorResult](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MAFlashingLightsProcessorResult) Autorelease() MAFlashingLightsProcessorResult {
	rv := objc.Send[MAFlashingLightsProcessorResult](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMAFlashingLightsProcessorResult creates a new MAFlashingLightsProcessorResult instance.
func NewMAFlashingLightsProcessorResult() MAFlashingLightsProcessorResult {
	return getMAFlashingLightsProcessorResultClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MAFlashingLightsProcessorResult */
// An object that reports the result of the flashing lights processor.
//
// An object is the result of calling . This object indicates whether the method successfully processed the input surface, the intensity of flashing lights in the input surface, and the amount of mitigation in the output surface.


// An object that reports the result of the flashing lights processor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAFlashingLightsProcessorResult
type MAFlashingLightsProcessorResult struct {
	objectivec.Object
}

// MAFlashingLightsProcessorResultFrom constructs a [MAFlashingLightsProcessorResult] from an unsafe.Pointer.
//
// An object that reports the result of the flashing lights processor.
func MAFlashingLightsProcessorResultFrom(ptr unsafe.Pointer) MAFlashingLightsProcessorResult {
	return MAFlashingLightsProcessorResult{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MAFlashingLightsProcessorResult *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MAFlashingLightsProcessorResult */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MAFlashingLightsProcessorResult */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MAFlashingLightsProcessorResult */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MAFlashingLightsProcessorResult */

// The intensity of flashing lights in the input surface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAFlashingLightsProcessorResult/intensityLevel
func (m_ MAFlashingLightsProcessorResult) IntensityLevel() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("intensityLevel"))
	return rv
}/* debug [instance_properties/getter]: intensityLevel */


// The amount of mitigation in the output surface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAFlashingLightsProcessorResult/mitigationLevel
func (m_ MAFlashingLightsProcessorResult) MitigationLevel() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("mitigationLevel"))
	return rv
}/* debug [instance_properties/getter]: mitigationLevel */


// A Boolean value that indicates whether the flashing lights processor successfully processed the input surface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAFlashingLightsProcessorResult/surfaceProcessed
func (m_ MAFlashingLightsProcessorResult) SurfaceProcessed() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("surfaceProcessed"))
	return rv
}/* debug [instance_properties/getter]: surfaceProcessed */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MAFlashingLightsProcessorResult */



