// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mAttributeDataElement */


/* debug [class_header]: Header for mAttributeDataElement */
// The class instance for the [mAttributeDataElement] class.
var (
	MAttributeDataElementClass     _mAttributeDataElementClass
	MAttributeDataElementClassOnce sync.Once
)

func getmAttributeDataElementClass() _mAttributeDataElementClass {
	MAttributeDataElementClassOnce.Do(func() {
		MAttributeDataElementClass = _mAttributeDataElementClass{objc.GetClass("mAttributeDataElement")}
	})
	return MAttributeDataElementClass
}

type _mAttributeDataElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mAttributeDataElement */
// An interface definition for the [mAttributeDataElement] class.
type ImAttributeDataElement interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mAttributeDataElement */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mAttributeDataElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mAttributeDataElement */
// Alloc allocates a new instance without initialization.
func (mc _mAttributeDataElementClass) Alloc() mAttributeDataElement {
	rv := objc.Send[mAttributeDataElement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mAttributeDataElementClass) New() mAttributeDataElement {
	rv := objc.Send[mAttributeDataElement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mAttributeDataElement) Init() mAttributeDataElement {
	rv := objc.Send[mAttributeDataElement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mAttributeDataElement) Autorelease() mAttributeDataElement {
	rv := objc.Send[mAttributeDataElement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmAttributeDataElement creates a new mAttributeDataElement instance.
func NewmAttributeDataElement() mAttributeDataElement {
	return getmAttributeDataElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mAttributeDataElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/mAttributeDataElement
type mAttributeDataElement struct {
	objectivec.Object
}

// mAttributeDataElementFrom constructs a [mAttributeDataElement] from an unsafe.Pointer.
func mAttributeDataElementFrom(ptr unsafe.Pointer) mAttributeDataElement {
	return mAttributeDataElement{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mAttributeDataElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mAttributeDataElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mAttributeDataElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mAttributeDataElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mAttributeDataElement */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mAttributeDataElement */



