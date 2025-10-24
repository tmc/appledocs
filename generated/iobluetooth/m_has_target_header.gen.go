// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mHasTargetHeader */


/* debug [class_header]: Header for mHasTargetHeader */
// The class instance for the [mHasTargetHeader] class.
var (
	MHasTargetHeaderClass     _mHasTargetHeaderClass
	MHasTargetHeaderClassOnce sync.Once
)

func getmHasTargetHeaderClass() _mHasTargetHeaderClass {
	MHasTargetHeaderClassOnce.Do(func() {
		MHasTargetHeaderClass = _mHasTargetHeaderClass{objc.GetClass("mHasTargetHeader")}
	})
	return MHasTargetHeaderClass
}

type _mHasTargetHeaderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mHasTargetHeader */
// An interface definition for the [mHasTargetHeader] class.
type ImHasTargetHeader interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mHasTargetHeader */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mHasTargetHeader */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mHasTargetHeader */
// Alloc allocates a new instance without initialization.
func (mc _mHasTargetHeaderClass) Alloc() mHasTargetHeader {
	rv := objc.Send[mHasTargetHeader](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mHasTargetHeaderClass) New() mHasTargetHeader {
	rv := objc.Send[mHasTargetHeader](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mHasTargetHeader) Init() mHasTargetHeader {
	rv := objc.Send[mHasTargetHeader](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mHasTargetHeader) Autorelease() mHasTargetHeader {
	rv := objc.Send[mHasTargetHeader](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmHasTargetHeader creates a new mHasTargetHeader instance.
func NewmHasTargetHeader() mHasTargetHeader {
	return getmHasTargetHeaderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mHasTargetHeader */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mHasTargetHeader
type mHasTargetHeader struct {
	objectivec.Object
}

// mHasTargetHeaderFrom constructs a [mHasTargetHeader] from an unsafe.Pointer.
func mHasTargetHeaderFrom(ptr unsafe.Pointer) mHasTargetHeader {
	return mHasTargetHeader{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mHasTargetHeader *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mHasTargetHeader */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mHasTargetHeader */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mHasTargetHeader */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mHasTargetHeader */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mHasTargetHeader */



