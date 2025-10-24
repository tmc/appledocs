// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mMaxPacketLength */


/* debug [class_header]: Header for mMaxPacketLength */
// The class instance for the [mMaxPacketLength] class.
var (
	MMaxPacketLengthClass     _mMaxPacketLengthClass
	MMaxPacketLengthClassOnce sync.Once
)

func getmMaxPacketLengthClass() _mMaxPacketLengthClass {
	MMaxPacketLengthClassOnce.Do(func() {
		MMaxPacketLengthClass = _mMaxPacketLengthClass{objc.GetClass("mMaxPacketLength")}
	})
	return MMaxPacketLengthClass
}

type _mMaxPacketLengthClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mMaxPacketLength */
// An interface definition for the [mMaxPacketLength] class.
type ImMaxPacketLength interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mMaxPacketLength */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mMaxPacketLength */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mMaxPacketLength */
// Alloc allocates a new instance without initialization.
func (mc _mMaxPacketLengthClass) Alloc() mMaxPacketLength {
	rv := objc.Send[mMaxPacketLength](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mMaxPacketLengthClass) New() mMaxPacketLength {
	rv := objc.Send[mMaxPacketLength](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mMaxPacketLength) Init() mMaxPacketLength {
	rv := objc.Send[mMaxPacketLength](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mMaxPacketLength) Autorelease() mMaxPacketLength {
	rv := objc.Send[mMaxPacketLength](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmMaxPacketLength creates a new mMaxPacketLength instance.
func NewmMaxPacketLength() mMaxPacketLength {
	return getmMaxPacketLengthClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mMaxPacketLength */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mMaxPacketLength
type mMaxPacketLength struct {
	objectivec.Object
}

// mMaxPacketLengthFrom constructs a [mMaxPacketLength] from an unsafe.Pointer.
func mMaxPacketLengthFrom(ptr unsafe.Pointer) mMaxPacketLength {
	return mMaxPacketLength{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mMaxPacketLength *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mMaxPacketLength */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mMaxPacketLength */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mMaxPacketLength */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mMaxPacketLength */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mMaxPacketLength */



