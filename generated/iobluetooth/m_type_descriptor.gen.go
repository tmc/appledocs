// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mTypeDescriptor */


/* debug [class_header]: Header for mTypeDescriptor */
// The class instance for the [mTypeDescriptor] class.
var (
	MTypeDescriptorClass     _mTypeDescriptorClass
	MTypeDescriptorClassOnce sync.Once
)

func getmTypeDescriptorClass() _mTypeDescriptorClass {
	MTypeDescriptorClassOnce.Do(func() {
		MTypeDescriptorClass = _mTypeDescriptorClass{objc.GetClass("mTypeDescriptor")}
	})
	return MTypeDescriptorClass
}

type _mTypeDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mTypeDescriptor */
// An interface definition for the [mTypeDescriptor] class.
type ImTypeDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mTypeDescriptor */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mTypeDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mTypeDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _mTypeDescriptorClass) Alloc() mTypeDescriptor {
	rv := objc.Send[mTypeDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mTypeDescriptorClass) New() mTypeDescriptor {
	rv := objc.Send[mTypeDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mTypeDescriptor) Init() mTypeDescriptor {
	rv := objc.Send[mTypeDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mTypeDescriptor) Autorelease() mTypeDescriptor {
	rv := objc.Send[mTypeDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmTypeDescriptor creates a new mTypeDescriptor instance.
func NewmTypeDescriptor() mTypeDescriptor {
	return getmTypeDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mTypeDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/mTypeDescriptor
type mTypeDescriptor struct {
	objectivec.Object
}

// mTypeDescriptorFrom constructs a [mTypeDescriptor] from an unsafe.Pointer.
func mTypeDescriptorFrom(ptr unsafe.Pointer) mTypeDescriptor {
	return mTypeDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mTypeDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mTypeDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mTypeDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mTypeDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mTypeDescriptor */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mTypeDescriptor */



