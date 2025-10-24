// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKLookAroundScene */


/* debug [class_header]: Header for MKLookAroundScene */
// The class instance for the [MKLookAroundScene] class.
var (
	MKLookAroundSceneClass     _MKLookAroundSceneClass
	MKLookAroundSceneClassOnce sync.Once
)

func getMKLookAroundSceneClass() _MKLookAroundSceneClass {
	MKLookAroundSceneClassOnce.Do(func() {
		MKLookAroundSceneClass = _MKLookAroundSceneClass{objc.GetClass("MKLookAroundScene")}
	})
	return MKLookAroundSceneClass
}

type _MKLookAroundSceneClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKLookAroundScene */
// An interface definition for the [MKLookAroundScene] class.
type IMKLookAroundScene interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKLookAroundScene */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKLookAroundScene */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKLookAroundScene */
// Alloc allocates a new instance without initialization.
func (mc _MKLookAroundSceneClass) Alloc() MKLookAroundScene {
	rv := objc.Send[MKLookAroundScene](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKLookAroundSceneClass) New() MKLookAroundScene {
	rv := objc.Send[MKLookAroundScene](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKLookAroundScene) Init() MKLookAroundScene {
	rv := objc.Send[MKLookAroundScene](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKLookAroundScene) Autorelease() MKLookAroundScene {
	rv := objc.Send[MKLookAroundScene](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKLookAroundScene creates a new MKLookAroundScene instance.
func NewMKLookAroundScene() MKLookAroundScene {
	return getMKLookAroundSceneClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKLookAroundScene */
// A utility class that encapsulates information the framework requires to retrieve and display a specific Look Around location’s imagery.


// A utility class that encapsulates information the framework requires to retrieve and display a specific Look Around location’s imagery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundScene
type MKLookAroundScene struct {
	objectivec.Object
}

// MKLookAroundSceneFrom constructs a [MKLookAroundScene] from an unsafe.Pointer.
//
// A utility class that encapsulates information the framework requires to retrieve and display a specific Look Around location’s imagery.
func MKLookAroundSceneFrom(ptr unsafe.Pointer) MKLookAroundScene {
	return MKLookAroundScene{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKLookAroundScene *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKLookAroundScene */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKLookAroundScene */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKLookAroundScene */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKLookAroundScene */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKLookAroundScene */



