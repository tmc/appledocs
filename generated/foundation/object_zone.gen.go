// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class objectZone */


/* debug [class_header]: Header for objectZone */
// The class instance for the [objectZone] class.
var (
	ObjectZoneClass     _objectZoneClass
	ObjectZoneClassOnce sync.Once
)

func getobjectZoneClass() _objectZoneClass {
	ObjectZoneClassOnce.Do(func() {
		ObjectZoneClass = _objectZoneClass{objc.GetClass("objectZone")}
	})
	return ObjectZoneClass
}

type _objectZoneClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for objectZone */
// An interface definition for the [objectZone] class.
type IobjectZone interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for objectZone */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for objectZone */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for objectZone */
// Alloc allocates a new instance without initialization.
func (oc _objectZoneClass) Alloc() objectZone {
	rv := objc.Send[objectZone](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _objectZoneClass) New() objectZone {
	rv := objc.Send[objectZone](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ objectZone) Init() objectZone {
	rv := objc.Send[objectZone](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ objectZone) Autorelease() objectZone {
	rv := objc.Send[objectZone](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewobjectZone creates a new objectZone instance.
func NewobjectZone() objectZone {
	return getobjectZoneClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for objectZone */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/objectZone-c.ivar
type objectZone struct {
	objectivec.Object
}

// objectZoneFrom constructs a [objectZone] from an unsafe.Pointer.
func objectZoneFrom(ptr unsafe.Pointer) objectZone {
	return objectZone{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for objectZone *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for objectZone */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for objectZone */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for objectZone */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for objectZone */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class objectZone */



