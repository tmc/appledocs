// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class outBuffer */


/* debug [class_header]: Header for outBuffer */
// The class instance for the [outBuffer] class.
var (
	OutBufferClass     _outBufferClass
	OutBufferClassOnce sync.Once
)

func getoutBufferClass() _outBufferClass {
	OutBufferClassOnce.Do(func() {
		OutBufferClass = _outBufferClass{objc.GetClass("outBuffer")}
	})
	return OutBufferClass
}

type _outBufferClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for outBuffer */
// An interface definition for the [outBuffer] class.
type IoutBuffer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for outBuffer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for outBuffer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for outBuffer */
// Alloc allocates a new instance without initialization.
func (oc _outBufferClass) Alloc() outBuffer {
	rv := objc.Send[outBuffer](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _outBufferClass) New() outBuffer {
	rv := objc.Send[outBuffer](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ outBuffer) Init() outBuffer {
	rv := objc.Send[outBuffer](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ outBuffer) Autorelease() outBuffer {
	rv := objc.Send[outBuffer](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewoutBuffer creates a new outBuffer instance.
func NewoutBuffer() outBuffer {
	return getoutBufferClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for outBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/outBuffer
type outBuffer struct {
	objectivec.Object
}

// outBufferFrom constructs a [outBuffer] from an unsafe.Pointer.
func outBufferFrom(ptr unsafe.Pointer) outBuffer {
	return outBuffer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for outBuffer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for outBuffer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for outBuffer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for outBuffer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for outBuffer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class outBuffer */



