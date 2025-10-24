// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class map */


/* debug [class_header]: Header for map */
// The class instance for the [Map] class.
var (
	MapClass     _MapClass
	MapClassOnce sync.Once
)

func getMapClass() _MapClass {
	MapClassOnce.Do(func() {
		MapClass = _MapClass{objc.GetClass("map")}
	})
	return MapClass
}

type _MapClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Map */
// An interface definition for the [Map] class.
type IMap interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Map */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Map */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Map */
// Alloc allocates a new instance without initialization.
func (mc _MapClass) Alloc() Map {
	rv := objc.Send[Map](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MapClass) New() Map {
	rv := objc.Send[Map](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Map) Init() Map {
	rv := objc.Send[Map](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Map) Autorelease() Map {
	rv := objc.Send[Map](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMap creates a new Map instance.
func NewMap() Map {
	return getMapClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Map */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/map
type Map struct {
	objectivec.Object
}

// MapFrom constructs a [Map] from an unsafe.Pointer.
func MapFrom(ptr unsafe.Pointer) Map {
	return Map{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Map *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Map */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Map */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Map */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Map */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class map */



