// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKAnnotationView */


/* debug [class_header]: Header for MKAnnotationView */
// The class instance for the [MKAnnotationView] class.
var (
	MKAnnotationViewClass     _MKAnnotationViewClass
	MKAnnotationViewClassOnce sync.Once
)

func getMKAnnotationViewClass() _MKAnnotationViewClass {
	MKAnnotationViewClassOnce.Do(func() {
		MKAnnotationViewClass = _MKAnnotationViewClass{objc.GetClass("MKAnnotationView")}
	})
	return MKAnnotationViewClass
}

type _MKAnnotationViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKAnnotationView */
// An interface definition for the [MKAnnotationView] class.
type IMKAnnotationView interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKAnnotationView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKAnnotationView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKAnnotationView */
// Alloc allocates a new instance without initialization.
func (mc _MKAnnotationViewClass) Alloc() MKAnnotationView {
	rv := objc.Send[MKAnnotationView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKAnnotationViewClass) New() MKAnnotationView {
	rv := objc.Send[MKAnnotationView](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKAnnotationView) Init() MKAnnotationView {
	rv := objc.Send[MKAnnotationView](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKAnnotationView) Autorelease() MKAnnotationView {
	rv := objc.Send[MKAnnotationView](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKAnnotationView creates a new MKAnnotationView instance.
func NewMKAnnotationView() MKAnnotationView {
	return getMKAnnotationViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKAnnotationView */
// A parent class referenced by other MapKit classes.


// A parent class referenced by other MapKit classes. [Full Topic]
type MKAnnotationView struct {
	objectivec.Object
}

// MKAnnotationViewFrom constructs a [MKAnnotationView] from an unsafe.Pointer.
//
// A parent class referenced by other MapKit classes.
func MKAnnotationViewFrom(ptr unsafe.Pointer) MKAnnotationView {
	return MKAnnotationView{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKAnnotationView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKAnnotationView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKAnnotationView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKAnnotationView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKAnnotationView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKAnnotationView */



