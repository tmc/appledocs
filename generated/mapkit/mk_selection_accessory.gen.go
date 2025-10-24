// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKSelectionAccessory */


/* debug [class_header]: Header for MKSelectionAccessory */
// The class instance for the [MKSelectionAccessory] class.
var (
	MKSelectionAccessoryClass     _MKSelectionAccessoryClass
	MKSelectionAccessoryClassOnce sync.Once
)

func getMKSelectionAccessoryClass() _MKSelectionAccessoryClass {
	MKSelectionAccessoryClassOnce.Do(func() {
		MKSelectionAccessoryClass = _MKSelectionAccessoryClass{objc.GetClass("MKSelectionAccessory")}
	})
	return MKSelectionAccessoryClass
}

type _MKSelectionAccessoryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKSelectionAccessory */
// An interface definition for the [MKSelectionAccessory] class.
type IMKSelectionAccessory interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKSelectionAccessory */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKSelectionAccessory */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKSelectionAccessory */
// Alloc allocates a new instance without initialization.
func (mc _MKSelectionAccessoryClass) Alloc() MKSelectionAccessory {
	rv := objc.Send[MKSelectionAccessory](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKSelectionAccessoryClass) New() MKSelectionAccessory {
	rv := objc.Send[MKSelectionAccessory](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKSelectionAccessory) Init() MKSelectionAccessory {
	rv := objc.Send[MKSelectionAccessory](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKSelectionAccessory) Autorelease() MKSelectionAccessory {
	rv := objc.Send[MKSelectionAccessory](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKSelectionAccessory creates a new MKSelectionAccessory instance.
func NewMKSelectionAccessory() MKSelectionAccessory {
	return getMKSelectionAccessoryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKSelectionAccessory */
// The type of accessory to display for a selected annotation.
//
// Implement in your map view delegate to specify a selection accessory for annotation content.


// The type of accessory to display for a selected annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKSelectionAccessory
type MKSelectionAccessory struct {
	objectivec.Object
}

// MKSelectionAccessoryFrom constructs a [MKSelectionAccessory] from an unsafe.Pointer.
//
// The type of accessory to display for a selected annotation.
func MKSelectionAccessoryFrom(ptr unsafe.Pointer) MKSelectionAccessory {
	return MKSelectionAccessory{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKSelectionAccessory *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKSelectionAccessory */

// Detailed information about a place
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKSelectionAccessory/mapItemDetail(_:)
func (mc _MKSelectionAccessoryClass) MapItemDetailWithPresentationStyle(presentationStyle IMKMapItemDetailSelectionAccessoryPresentationStyle) MKSelectionAccessory {
	rv := objc.Send[MKSelectionAccessory](objc.ID(mc.class), objc.Sel("mapItemDetailWithPresentationStyle:"), presentationStyle)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MapItemDetailWithPresentationStyle) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKSelectionAccessory */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKSelectionAccessory */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKSelectionAccessory */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKSelectionAccessory */



