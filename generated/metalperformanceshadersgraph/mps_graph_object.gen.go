// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSGraphObject */


/* debug [class_header]: Header for MPSGraphObject */
// The class instance for the [GraphObject] class.
var (
	GraphObjectClass     _GraphObjectClass
	GraphObjectClassOnce sync.Once
)

func getGraphObjectClass() _GraphObjectClass {
	GraphObjectClassOnce.Do(func() {
		GraphObjectClass = _GraphObjectClass{objc.GetClass("MPSGraphObject")}
	})
	return GraphObjectClass
}

type _GraphObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphObject */
// An interface definition for the [GraphObject] class.
type IGraphObject interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GraphObject */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphObject */
// Alloc allocates a new instance without initialization.
func (gc _GraphObjectClass) Alloc() GraphObject {
	rv := objc.Send[GraphObject](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphObjectClass) New() GraphObject {
	rv := objc.Send[GraphObject](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphObject) Init() GraphObject {
	rv := objc.Send[GraphObject](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphObject) Autorelease() GraphObject {
	rv := objc.Send[GraphObject](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphObject creates a new GraphObject instance.
func NewGraphObject() GraphObject {
	return getGraphObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphObject */
// The common base class for all Metal Performance Shaders Graph objects.
//
// Only the child classes should be used.


// The common base class for all Metal Performance Shaders Graph objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphObject
type GraphObject struct {
	objectivec.Object
}

// GraphObjectFrom constructs a [GraphObject] from an unsafe.Pointer.
//
// The common base class for all Metal Performance Shaders Graph objects.
func GraphObjectFrom(ptr unsafe.Pointer) GraphObject {
	return GraphObject{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphObject */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphObject */



