// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class currentOffsetInBuffer */


/* debug [class_header]: Header for currentOffsetInBuffer */
// The class instance for the [currentOffsetInBuffer] class.
var (
	CurrentOffsetInBufferClass     _currentOffsetInBufferClass
	CurrentOffsetInBufferClassOnce sync.Once
)

func getcurrentOffsetInBufferClass() _currentOffsetInBufferClass {
	CurrentOffsetInBufferClassOnce.Do(func() {
		CurrentOffsetInBufferClass = _currentOffsetInBufferClass{objc.GetClass("currentOffsetInBuffer")}
	})
	return CurrentOffsetInBufferClass
}

type _currentOffsetInBufferClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for currentOffsetInBuffer */
// An interface definition for the [currentOffsetInBuffer] class.
type IcurrentOffsetInBuffer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for currentOffsetInBuffer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for currentOffsetInBuffer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for currentOffsetInBuffer */
// Alloc allocates a new instance without initialization.
func (cc _currentOffsetInBufferClass) Alloc() currentOffsetInBuffer {
	rv := objc.Send[currentOffsetInBuffer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _currentOffsetInBufferClass) New() currentOffsetInBuffer {
	rv := objc.Send[currentOffsetInBuffer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ currentOffsetInBuffer) Init() currentOffsetInBuffer {
	rv := objc.Send[currentOffsetInBuffer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ currentOffsetInBuffer) Autorelease() currentOffsetInBuffer {
	rv := objc.Send[currentOffsetInBuffer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewcurrentOffsetInBuffer creates a new currentOffsetInBuffer instance.
func NewcurrentOffsetInBuffer() currentOffsetInBuffer {
	return getcurrentOffsetInBufferClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for currentOffsetInBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/currentOffsetInBuffer
type currentOffsetInBuffer struct {
	objectivec.Object
}

// currentOffsetInBufferFrom constructs a [currentOffsetInBuffer] from an unsafe.Pointer.
func currentOffsetInBufferFrom(ptr unsafe.Pointer) currentOffsetInBuffer {
	return currentOffsetInBuffer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for currentOffsetInBuffer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for currentOffsetInBuffer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for currentOffsetInBuffer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for currentOffsetInBuffer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for currentOffsetInBuffer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class currentOffsetInBuffer */



