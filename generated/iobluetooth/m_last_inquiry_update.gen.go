// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mLastInquiryUpdate */


/* debug [class_header]: Header for mLastInquiryUpdate */
// The class instance for the [mLastInquiryUpdate] class.
var (
	MLastInquiryUpdateClass     _mLastInquiryUpdateClass
	MLastInquiryUpdateClassOnce sync.Once
)

func getmLastInquiryUpdateClass() _mLastInquiryUpdateClass {
	MLastInquiryUpdateClassOnce.Do(func() {
		MLastInquiryUpdateClass = _mLastInquiryUpdateClass{objc.GetClass("mLastInquiryUpdate")}
	})
	return MLastInquiryUpdateClass
}

type _mLastInquiryUpdateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mLastInquiryUpdate */
// An interface definition for the [mLastInquiryUpdate] class.
type ImLastInquiryUpdate interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mLastInquiryUpdate */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mLastInquiryUpdate */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mLastInquiryUpdate */
// Alloc allocates a new instance without initialization.
func (mc _mLastInquiryUpdateClass) Alloc() mLastInquiryUpdate {
	rv := objc.Send[mLastInquiryUpdate](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mLastInquiryUpdateClass) New() mLastInquiryUpdate {
	rv := objc.Send[mLastInquiryUpdate](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mLastInquiryUpdate) Init() mLastInquiryUpdate {
	rv := objc.Send[mLastInquiryUpdate](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mLastInquiryUpdate) Autorelease() mLastInquiryUpdate {
	rv := objc.Send[mLastInquiryUpdate](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmLastInquiryUpdate creates a new mLastInquiryUpdate instance.
func NewmLastInquiryUpdate() mLastInquiryUpdate {
	return getmLastInquiryUpdateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mLastInquiryUpdate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mLastInquiryUpdate
type mLastInquiryUpdate struct {
	objectivec.Object
}

// mLastInquiryUpdateFrom constructs a [mLastInquiryUpdate] from an unsafe.Pointer.
func mLastInquiryUpdateFrom(ptr unsafe.Pointer) mLastInquiryUpdate {
	return mLastInquiryUpdate{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mLastInquiryUpdate *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mLastInquiryUpdate */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mLastInquiryUpdate */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mLastInquiryUpdate */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mLastInquiryUpdate */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mLastInquiryUpdate */



