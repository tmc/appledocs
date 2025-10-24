// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mIOService */


/* debug [class_header]: Header for mIOService */
// The class instance for the [mIOService] class.
var (
	MIOServiceClass     _mIOServiceClass
	MIOServiceClassOnce sync.Once
)

func getmIOServiceClass() _mIOServiceClass {
	MIOServiceClassOnce.Do(func() {
		MIOServiceClass = _mIOServiceClass{objc.GetClass("mIOService")}
	})
	return MIOServiceClass
}

type _mIOServiceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mIOService */
// An interface definition for the [mIOService] class.
type ImIOService interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mIOService */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mIOService */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mIOService */
// Alloc allocates a new instance without initialization.
func (mc _mIOServiceClass) Alloc() mIOService {
	rv := objc.Send[mIOService](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mIOServiceClass) New() mIOService {
	rv := objc.Send[mIOService](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIOService) Init() mIOService {
	rv := objc.Send[mIOService](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIOService) Autorelease() mIOService {
	rv := objc.Send[mIOService](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIOService creates a new mIOService instance.
func NewmIOService() mIOService {
	return getmIOServiceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mIOService */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothObject/mIOService
type mIOService struct {
	objectivec.Object
}

// mIOServiceFrom constructs a [mIOService] from an unsafe.Pointer.
func mIOServiceFrom(ptr unsafe.Pointer) mIOService {
	return mIOService{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mIOService *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mIOService */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mIOService */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mIOService */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mIOService */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mIOService */



