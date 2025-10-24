// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mActionInProgress */


/* debug [class_header]: Header for mActionInProgress */
// The class instance for the [mActionInProgress] class.
var (
	MActionInProgressClass     _mActionInProgressClass
	MActionInProgressClassOnce sync.Once
)

func getmActionInProgressClass() _mActionInProgressClass {
	MActionInProgressClassOnce.Do(func() {
		MActionInProgressClass = _mActionInProgressClass{objc.GetClass("mActionInProgress")}
	})
	return MActionInProgressClass
}

type _mActionInProgressClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mActionInProgress */
// An interface definition for the [mActionInProgress] class.
type ImActionInProgress interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mActionInProgress */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mActionInProgress */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mActionInProgress */
// Alloc allocates a new instance without initialization.
func (mc _mActionInProgressClass) Alloc() mActionInProgress {
	rv := objc.Send[mActionInProgress](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mActionInProgressClass) New() mActionInProgress {
	rv := objc.Send[mActionInProgress](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mActionInProgress) Init() mActionInProgress {
	rv := objc.Send[mActionInProgress](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mActionInProgress) Autorelease() mActionInProgress {
	rv := objc.Send[mActionInProgress](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmActionInProgress creates a new mActionInProgress instance.
func NewmActionInProgress() mActionInProgress {
	return getmActionInProgressClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mActionInProgress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mActionInProgress
type mActionInProgress struct {
	objectivec.Object
}

// mActionInProgressFrom constructs a [mActionInProgress] from an unsafe.Pointer.
func mActionInProgressFrom(ptr unsafe.Pointer) mActionInProgress {
	return mActionInProgress{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mActionInProgress *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mActionInProgress */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mActionInProgress */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mActionInProgress */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mActionInProgress */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mActionInProgress */



