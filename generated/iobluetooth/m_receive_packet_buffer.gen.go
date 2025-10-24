// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mReceivePacketBuffer */


/* debug [class_header]: Header for mReceivePacketBuffer */
// The class instance for the [mReceivePacketBuffer] class.
var (
	MReceivePacketBufferClass     _mReceivePacketBufferClass
	MReceivePacketBufferClassOnce sync.Once
)

func getmReceivePacketBufferClass() _mReceivePacketBufferClass {
	MReceivePacketBufferClassOnce.Do(func() {
		MReceivePacketBufferClass = _mReceivePacketBufferClass{objc.GetClass("mReceivePacketBuffer")}
	})
	return MReceivePacketBufferClass
}

type _mReceivePacketBufferClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mReceivePacketBuffer */
// An interface definition for the [mReceivePacketBuffer] class.
type ImReceivePacketBuffer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mReceivePacketBuffer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mReceivePacketBuffer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mReceivePacketBuffer */
// Alloc allocates a new instance without initialization.
func (mc _mReceivePacketBufferClass) Alloc() mReceivePacketBuffer {
	rv := objc.Send[mReceivePacketBuffer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mReceivePacketBufferClass) New() mReceivePacketBuffer {
	rv := objc.Send[mReceivePacketBuffer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mReceivePacketBuffer) Init() mReceivePacketBuffer {
	rv := objc.Send[mReceivePacketBuffer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mReceivePacketBuffer) Autorelease() mReceivePacketBuffer {
	rv := objc.Send[mReceivePacketBuffer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmReceivePacketBuffer creates a new mReceivePacketBuffer instance.
func NewmReceivePacketBuffer() mReceivePacketBuffer {
	return getmReceivePacketBufferClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mReceivePacketBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/mReceivePacketBuffer
type mReceivePacketBuffer struct {
	objectivec.Object
}

// mReceivePacketBufferFrom constructs a [mReceivePacketBuffer] from an unsafe.Pointer.
func mReceivePacketBufferFrom(ptr unsafe.Pointer) mReceivePacketBuffer {
	return mReceivePacketBuffer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mReceivePacketBuffer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mReceivePacketBuffer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mReceivePacketBuffer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mReceivePacketBuffer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mReceivePacketBuffer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mReceivePacketBuffer */



