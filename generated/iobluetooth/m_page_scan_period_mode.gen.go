// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mPageScanPeriodMode */


/* debug [class_header]: Header for mPageScanPeriodMode */
// The class instance for the [mPageScanPeriodMode] class.
var (
	MPageScanPeriodModeClass     _mPageScanPeriodModeClass
	MPageScanPeriodModeClassOnce sync.Once
)

func getmPageScanPeriodModeClass() _mPageScanPeriodModeClass {
	MPageScanPeriodModeClassOnce.Do(func() {
		MPageScanPeriodModeClass = _mPageScanPeriodModeClass{objc.GetClass("mPageScanPeriodMode")}
	})
	return MPageScanPeriodModeClass
}

type _mPageScanPeriodModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mPageScanPeriodMode */
// An interface definition for the [mPageScanPeriodMode] class.
type ImPageScanPeriodMode interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mPageScanPeriodMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mPageScanPeriodMode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mPageScanPeriodMode */
// Alloc allocates a new instance without initialization.
func (mc _mPageScanPeriodModeClass) Alloc() mPageScanPeriodMode {
	rv := objc.Send[mPageScanPeriodMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mPageScanPeriodModeClass) New() mPageScanPeriodMode {
	rv := objc.Send[mPageScanPeriodMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mPageScanPeriodMode) Init() mPageScanPeriodMode {
	rv := objc.Send[mPageScanPeriodMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mPageScanPeriodMode) Autorelease() mPageScanPeriodMode {
	rv := objc.Send[mPageScanPeriodMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmPageScanPeriodMode creates a new mPageScanPeriodMode instance.
func NewmPageScanPeriodMode() mPageScanPeriodMode {
	return getmPageScanPeriodModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mPageScanPeriodMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mPageScanPeriodMode
type mPageScanPeriodMode struct {
	objectivec.Object
}

// mPageScanPeriodModeFrom constructs a [mPageScanPeriodMode] from an unsafe.Pointer.
func mPageScanPeriodModeFrom(ptr unsafe.Pointer) mPageScanPeriodMode {
	return mPageScanPeriodMode{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mPageScanPeriodMode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mPageScanPeriodMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mPageScanPeriodMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mPageScanPeriodMode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mPageScanPeriodMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mPageScanPeriodMode */



