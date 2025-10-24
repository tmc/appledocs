// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mSizeDescriptor */


/* debug [class_header]: Header for mSizeDescriptor */
// The class instance for the [mSizeDescriptor] class.
var (
	MSizeDescriptorClass     _mSizeDescriptorClass
	MSizeDescriptorClassOnce sync.Once
)

func getmSizeDescriptorClass() _mSizeDescriptorClass {
	MSizeDescriptorClassOnce.Do(func() {
		MSizeDescriptorClass = _mSizeDescriptorClass{objc.GetClass("mSizeDescriptor")}
	})
	return MSizeDescriptorClass
}

type _mSizeDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mSizeDescriptor */
// An interface definition for the [mSizeDescriptor] class.
type ImSizeDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mSizeDescriptor */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mSizeDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mSizeDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _mSizeDescriptorClass) Alloc() mSizeDescriptor {
	rv := objc.Send[mSizeDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mSizeDescriptorClass) New() mSizeDescriptor {
	rv := objc.Send[mSizeDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mSizeDescriptor) Init() mSizeDescriptor {
	rv := objc.Send[mSizeDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mSizeDescriptor) Autorelease() mSizeDescriptor {
	rv := objc.Send[mSizeDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmSizeDescriptor creates a new mSizeDescriptor instance.
func NewmSizeDescriptor() mSizeDescriptor {
	return getmSizeDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mSizeDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/mSizeDescriptor
type mSizeDescriptor struct {
	objectivec.Object
}

// mSizeDescriptorFrom constructs a [mSizeDescriptor] from an unsafe.Pointer.
func mSizeDescriptorFrom(ptr unsafe.Pointer) mSizeDescriptor {
	return mSizeDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mSizeDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mSizeDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mSizeDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mSizeDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mSizeDescriptor */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mSizeDescriptor */



