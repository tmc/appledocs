// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKRouteStep */


/* debug [class_header]: Header for MKRouteStep */
// The class instance for the [MKRouteStep] class.
var (
	MKRouteStepClass     _MKRouteStepClass
	MKRouteStepClassOnce sync.Once
)

func getMKRouteStepClass() _MKRouteStepClass {
	MKRouteStepClassOnce.Do(func() {
		MKRouteStepClass = _MKRouteStepClass{objc.GetClass("MKRouteStep")}
	})
	return MKRouteStepClass
}

type _MKRouteStepClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKRouteStep */
// An interface definition for the [MKRouteStep] class.
type IMKRouteStep interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKRouteStep */
	// properties:
	Distance() LocationDistance /* not a class type */
	Instructions() objc.IObject /* cross-framework: NSString */
	Notice() objc.IObject /* cross-framework: NSString */
	Polyline() IMKPolyline
	TransportType() MKDirectionsTransportType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKRouteStep */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKRouteStep */
// Alloc allocates a new instance without initialization.
func (mc _MKRouteStepClass) Alloc() MKRouteStep {
	rv := objc.Send[MKRouteStep](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKRouteStepClass) New() MKRouteStep {
	rv := objc.Send[MKRouteStep](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKRouteStep) Init() MKRouteStep {
	rv := objc.Send[MKRouteStep](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKRouteStep) Autorelease() MKRouteStep {
	rv := objc.Send[MKRouteStep](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKRouteStep creates a new MKRouteStep instance.
func NewMKRouteStep() MKRouteStep {
	return getMKRouteStepClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKRouteStep */
// One portion of an overall route.
//
// Each object corresponds to a single instruction that the person needs to follow when navigating between two points. For example, a step might involve following a single road until continuing along the route requires a turn. You don’t create instances of this class directly. An object contains the objects associated with a route. For more information about requesting directions, see .


// One portion of an overall route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKRoute/Step
type MKRouteStep struct {
	objectivec.Object
}

// MKRouteStepFrom constructs a [MKRouteStep] from an unsafe.Pointer.
//
// One portion of an overall route.
func MKRouteStepFrom(ptr unsafe.Pointer) MKRouteStep {
	return MKRouteStep{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKRouteStep *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKRouteStep */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKRouteStep */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKRouteStep */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKRouteStep */

// The step distance, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKRoute/Step/distance
func (m_ MKRouteStep) Distance() LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](m_.ID, objc.Sel("distance"))
	return rv
}/* debug [instance_properties/getter]: distance */


// The written instructions for following the path that the step represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKRoute/Step/instructions
func (m_ MKRouteStep) Instructions() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("instructions"))
	return rv
}/* debug [instance_properties/getter]: instructions */


// Additional notices that apply to the step.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKRoute/Step/notice
func (m_ MKRouteStep) Notice() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("notice"))
	return rv
}/* debug [instance_properties/getter]: notice */


// The detailed step geometry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKRoute/Step/polyline
func (m_ MKRouteStep) Polyline() IMKPolyline {
	rv := objc.Send[MKPolyline](m_.ID, objc.Sel("polyline"))
	return rv
}/* debug [instance_properties/getter]: polyline */


// The transport type of the step.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKRoute/Step/transportType
func (m_ MKRouteStep) TransportType() MKDirectionsTransportType {
	rv := objc.Send[MKDirectionsTransportType](m_.ID, objc.Sel("transportType"))
	return rv
}/* debug [instance_properties/getter]: transportType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKRouteStep */



