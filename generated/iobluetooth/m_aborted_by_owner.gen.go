// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mAbortedByOwner */


/* debug [class_header]: Header for mAbortedByOwner */
// The class instance for the [mAbortedByOwner] class.
var (
	MAbortedByOwnerClass     _mAbortedByOwnerClass
	MAbortedByOwnerClassOnce sync.Once
)

func getmAbortedByOwnerClass() _mAbortedByOwnerClass {
	MAbortedByOwnerClassOnce.Do(func() {
		MAbortedByOwnerClass = _mAbortedByOwnerClass{objc.GetClass("mAbortedByOwner")}
	})
	return MAbortedByOwnerClass
}

type _mAbortedByOwnerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mAbortedByOwner */
// An interface definition for the [mAbortedByOwner] class.
type ImAbortedByOwner interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mAbortedByOwner */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mAbortedByOwner */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mAbortedByOwner */
// Alloc allocates a new instance without initialization.
func (mc _mAbortedByOwnerClass) Alloc() mAbortedByOwner {
	rv := objc.Send[mAbortedByOwner](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mAbortedByOwnerClass) New() mAbortedByOwner {
	rv := objc.Send[mAbortedByOwner](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mAbortedByOwner) Init() mAbortedByOwner {
	rv := objc.Send[mAbortedByOwner](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mAbortedByOwner) Autorelease() mAbortedByOwner {
	rv := objc.Send[mAbortedByOwner](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmAbortedByOwner creates a new mAbortedByOwner instance.
func NewmAbortedByOwner() mAbortedByOwner {
	return getmAbortedByOwnerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mAbortedByOwner */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mAbortedByOwner
type mAbortedByOwner struct {
	objectivec.Object
}

// mAbortedByOwnerFrom constructs a [mAbortedByOwner] from an unsafe.Pointer.
func mAbortedByOwnerFrom(ptr unsafe.Pointer) mAbortedByOwner {
	return mAbortedByOwner{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mAbortedByOwner *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mAbortedByOwner */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mAbortedByOwner */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mAbortedByOwner */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mAbortedByOwner */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mAbortedByOwner */



