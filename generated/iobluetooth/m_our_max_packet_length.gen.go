// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mOurMaxPacketLength */


/* debug [class_header]: Header for mOurMaxPacketLength */
// The class instance for the [mOurMaxPacketLength] class.
var (
	MOurMaxPacketLengthClass     _mOurMaxPacketLengthClass
	MOurMaxPacketLengthClassOnce sync.Once
)

func getmOurMaxPacketLengthClass() _mOurMaxPacketLengthClass {
	MOurMaxPacketLengthClassOnce.Do(func() {
		MOurMaxPacketLengthClass = _mOurMaxPacketLengthClass{objc.GetClass("mOurMaxPacketLength")}
	})
	return MOurMaxPacketLengthClass
}

type _mOurMaxPacketLengthClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mOurMaxPacketLength */
// An interface definition for the [mOurMaxPacketLength] class.
type ImOurMaxPacketLength interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mOurMaxPacketLength */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mOurMaxPacketLength */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mOurMaxPacketLength */
// Alloc allocates a new instance without initialization.
func (mc _mOurMaxPacketLengthClass) Alloc() mOurMaxPacketLength {
	rv := objc.Send[mOurMaxPacketLength](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mOurMaxPacketLengthClass) New() mOurMaxPacketLength {
	rv := objc.Send[mOurMaxPacketLength](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOurMaxPacketLength) Init() mOurMaxPacketLength {
	rv := objc.Send[mOurMaxPacketLength](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOurMaxPacketLength) Autorelease() mOurMaxPacketLength {
	rv := objc.Send[mOurMaxPacketLength](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOurMaxPacketLength creates a new mOurMaxPacketLength instance.
func NewmOurMaxPacketLength() mOurMaxPacketLength {
	return getmOurMaxPacketLengthClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mOurMaxPacketLength */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/mOurMaxPacketLength
type mOurMaxPacketLength struct {
	objectivec.Object
}

// mOurMaxPacketLengthFrom constructs a [mOurMaxPacketLength] from an unsafe.Pointer.
func mOurMaxPacketLengthFrom(ptr unsafe.Pointer) mOurMaxPacketLength {
	return mOurMaxPacketLength{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mOurMaxPacketLength *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mOurMaxPacketLength */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mOurMaxPacketLength */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mOurMaxPacketLength */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mOurMaxPacketLength */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mOurMaxPacketLength */



