// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CLPlacemark */


/* debug [class_header]: Header for CLPlacemark */
// The class instance for the [Placemark] class.
var (
	PlacemarkClass     _PlacemarkClass
	PlacemarkClassOnce sync.Once
)

func getPlacemarkClass() _PlacemarkClass {
	PlacemarkClassOnce.Do(func() {
		PlacemarkClass = _PlacemarkClass{objc.GetClass("CLPlacemark")}
	})
	return PlacemarkClass
}

type _PlacemarkClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Placemark */
// An interface definition for the [Placemark] class.
type IPlacemark interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Placemark */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Placemark */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Placemark */
// Alloc allocates a new instance without initialization.
func (pc _PlacemarkClass) Alloc() Placemark {
	rv := objc.Send[Placemark](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlacemarkClass) New() Placemark {
	rv := objc.Send[Placemark](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Placemark) Init() Placemark {
	rv := objc.Send[Placemark](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Placemark) Autorelease() Placemark {
	rv := objc.Send[Placemark](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlacemark creates a new Placemark instance.
func NewPlacemark() Placemark {
	return getPlacemarkClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Placemark */
// A parent class referenced by other MapKit classes.


// A parent class referenced by other MapKit classes. [Full Topic]
type Placemark struct {
	objectivec.Object
}

// PlacemarkFrom constructs a [Placemark] from an unsafe.Pointer.
//
// A parent class referenced by other MapKit classes.
func PlacemarkFrom(ptr unsafe.Pointer) Placemark {
	return Placemark{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Placemark *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Placemark */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Placemark */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Placemark */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Placemark */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CLPlacemark */



