// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZVirtioConsolePortArray */

/* debug [class_header]: Header for VZVirtioConsolePortArray */
// The class instance for the [VZVirtioConsolePortArray] class.
var (
	VZVirtioConsolePortArrayClass     _VZVirtioConsolePortArrayClass
	VZVirtioConsolePortArrayClassOnce sync.Once
)

func getVZVirtioConsolePortArrayClass() _VZVirtioConsolePortArrayClass {
	VZVirtioConsolePortArrayClassOnce.Do(func() {
		VZVirtioConsolePortArrayClass = _VZVirtioConsolePortArrayClass{objc.GetClass("VZVirtioConsolePortArray")}
	})
	return VZVirtioConsolePortArrayClass
}

type _VZVirtioConsolePortArrayClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZVirtioConsolePortArray */
// An interface definition for the [VZVirtioConsolePortArray] class.
type IVZVirtioConsolePortArray interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZVirtioConsolePortArray */
	// properties:
	MaximumPortCount() uint32 /* not a class type */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZVirtioConsolePortArray */
	// methods:
	ObjectAtIndexedSubscript(portIndex uint) IVZVirtioConsolePort
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZVirtioConsolePortArray */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioConsolePortArrayClass) Alloc() VZVirtioConsolePortArray {
	rv := objc.Send[VZVirtioConsolePortArray](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZVirtioConsolePortArrayClass) New() VZVirtioConsolePortArray {
	rv := objc.Send[VZVirtioConsolePortArray](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioConsolePortArray) Init() VZVirtioConsolePortArray {
	rv := objc.Send[VZVirtioConsolePortArray](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioConsolePortArray) Autorelease() VZVirtioConsolePortArray {
	rv := objc.Send[VZVirtioConsolePortArray](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsolePortArray creates a new VZVirtioConsolePortArray instance.
func NewVZVirtioConsolePortArray() VZVirtioConsolePortArray {
	return getVZVirtioConsolePortArrayClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZVirtioConsolePortArray */
// A class that represents a collection of Virtio console ports.

// A class that represents a collection of Virtio console ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortArray
type VZVirtioConsolePortArray struct {
	objectivec.Object
}

// VZVirtioConsolePortArrayFrom constructs a [VZVirtioConsolePortArray] from an unsafe.Pointer.
//
// A class that represents a collection of Virtio console ports.
func VZVirtioConsolePortArrayFrom(ptr unsafe.Pointer) VZVirtioConsolePortArray {
	return VZVirtioConsolePortArray{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZVirtioConsolePortArray */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZVirtioConsolePortArray */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZVirtioConsolePortArray */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZVirtioConsolePortArray */

// Returns the Virtio console port at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortArray/subscript(_:)
func (v_ VZVirtioConsolePortArray) ObjectAtIndexedSubscript(portIndex uint) IVZVirtioConsolePort {
	rv := objc.Send[VZVirtioConsolePort](v_.ID, objc.Sel("objectAtIndexedSubscript:"), portIndex)
	return rv
} /* debug [instance_methods/method]: ObjectAtIndexedSubscript */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZVirtioConsolePortArray */

// An unsigned integer that represents the maximum number of ports allocated by this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortArray/maximumPortCount
func (v_ VZVirtioConsolePortArray) MaximumPortCount() uint32 /* not a class type */ {
	rv := objc.Send[uint32](v_.ID, objc.Sel("maximumPortCount"))
	return rv
} /* debug [instance_properties/getter]: maximumPortCount */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZVirtioConsolePortArray */
