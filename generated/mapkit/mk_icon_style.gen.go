// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKIconStyle */


/* debug [class_header]: Header for MKIconStyle */
// The class instance for the [MKIconStyle] class.
var (
	MKIconStyleClass     _MKIconStyleClass
	MKIconStyleClassOnce sync.Once
)

func getMKIconStyleClass() _MKIconStyleClass {
	MKIconStyleClassOnce.Do(func() {
		MKIconStyleClass = _MKIconStyleClass{objc.GetClass("MKIconStyle")}
	})
	return MKIconStyleClass
}

type _MKIconStyleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKIconStyle */
// An interface definition for the [MKIconStyle] class.
type IMKIconStyle interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKIconStyle */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKIconStyle */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKIconStyle */
// Alloc allocates a new instance without initialization.
func (mc _MKIconStyleClass) Alloc() MKIconStyle {
	rv := objc.Send[MKIconStyle](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKIconStyleClass) New() MKIconStyle {
	rv := objc.Send[MKIconStyle](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKIconStyle) Init() MKIconStyle {
	rv := objc.Send[MKIconStyle](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKIconStyle) Autorelease() MKIconStyle {
	rv := objc.Send[MKIconStyle](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKIconStyle creates a new MKIconStyle instance.
func NewMKIconStyle() MKIconStyle {
	return getMKIconStyleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKIconStyle */
// A class you use to customize the annotation view icon of a point of interest (POI) on the map.


// A class you use to customize the annotation view icon of a point of interest (POI) on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKIconStyle
type MKIconStyle struct {
	objectivec.Object
}

// MKIconStyleFrom constructs a [MKIconStyle] from an unsafe.Pointer.
//
// A class you use to customize the annotation view icon of a point of interest (POI) on the map.
func MKIconStyleFrom(ptr unsafe.Pointer) MKIconStyle {
	return MKIconStyle{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKIconStyle *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKIconStyle */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKIconStyle */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKIconStyle */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKIconStyle */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKIconStyle */


