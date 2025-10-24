// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mDeviceForService */


/* debug [class_header]: Header for mDeviceForService */
// The class instance for the [mDeviceForService] class.
var (
	MDeviceForServiceClass     _mDeviceForServiceClass
	MDeviceForServiceClassOnce sync.Once
)

func getmDeviceForServiceClass() _mDeviceForServiceClass {
	MDeviceForServiceClassOnce.Do(func() {
		MDeviceForServiceClass = _mDeviceForServiceClass{objc.GetClass("mDeviceForService")}
	})
	return MDeviceForServiceClass
}

type _mDeviceForServiceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mDeviceForService */
// An interface definition for the [mDeviceForService] class.
type ImDeviceForService interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mDeviceForService */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mDeviceForService */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mDeviceForService */
// Alloc allocates a new instance without initialization.
func (mc _mDeviceForServiceClass) Alloc() mDeviceForService {
	rv := objc.Send[mDeviceForService](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mDeviceForServiceClass) New() mDeviceForService {
	rv := objc.Send[mDeviceForService](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mDeviceForService) Init() mDeviceForService {
	rv := objc.Send[mDeviceForService](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mDeviceForService) Autorelease() mDeviceForService {
	rv := objc.Send[mDeviceForService](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmDeviceForService creates a new mDeviceForService instance.
func NewmDeviceForService() mDeviceForService {
	return getmDeviceForServiceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mDeviceForService */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/mDeviceForService
type mDeviceForService struct {
	objectivec.Object
}

// mDeviceForServiceFrom constructs a [mDeviceForService] from an unsafe.Pointer.
func mDeviceForServiceFrom(ptr unsafe.Pointer) mDeviceForService {
	return mDeviceForService{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mDeviceForService *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mDeviceForService */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mDeviceForService */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mDeviceForService */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mDeviceForService */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mDeviceForService */



