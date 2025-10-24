// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mTheirMaxPacketLength */


/* debug [class_header]: Header for mTheirMaxPacketLength */
// The class instance for the [mTheirMaxPacketLength] class.
var (
	MTheirMaxPacketLengthClass     _mTheirMaxPacketLengthClass
	MTheirMaxPacketLengthClassOnce sync.Once
)

func getmTheirMaxPacketLengthClass() _mTheirMaxPacketLengthClass {
	MTheirMaxPacketLengthClassOnce.Do(func() {
		MTheirMaxPacketLengthClass = _mTheirMaxPacketLengthClass{objc.GetClass("mTheirMaxPacketLength")}
	})
	return MTheirMaxPacketLengthClass
}

type _mTheirMaxPacketLengthClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mTheirMaxPacketLength */
// An interface definition for the [mTheirMaxPacketLength] class.
type ImTheirMaxPacketLength interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mTheirMaxPacketLength */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mTheirMaxPacketLength */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mTheirMaxPacketLength */
// Alloc allocates a new instance without initialization.
func (mc _mTheirMaxPacketLengthClass) Alloc() mTheirMaxPacketLength {
	rv := objc.Send[mTheirMaxPacketLength](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mTheirMaxPacketLengthClass) New() mTheirMaxPacketLength {
	rv := objc.Send[mTheirMaxPacketLength](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mTheirMaxPacketLength) Init() mTheirMaxPacketLength {
	rv := objc.Send[mTheirMaxPacketLength](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mTheirMaxPacketLength) Autorelease() mTheirMaxPacketLength {
	rv := objc.Send[mTheirMaxPacketLength](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmTheirMaxPacketLength creates a new mTheirMaxPacketLength instance.
func NewmTheirMaxPacketLength() mTheirMaxPacketLength {
	return getmTheirMaxPacketLengthClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mTheirMaxPacketLength */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/mTheirMaxPacketLength
type mTheirMaxPacketLength struct {
	objectivec.Object
}

// mTheirMaxPacketLengthFrom constructs a [mTheirMaxPacketLength] from an unsafe.Pointer.
func mTheirMaxPacketLengthFrom(ptr unsafe.Pointer) mTheirMaxPacketLength {
	return mTheirMaxPacketLength{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mTheirMaxPacketLength *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mTheirMaxPacketLength */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mTheirMaxPacketLength */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mTheirMaxPacketLength */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mTheirMaxPacketLength */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mTheirMaxPacketLength */





