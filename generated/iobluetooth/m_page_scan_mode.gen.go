// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mPageScanMode */


/* debug [class_header]: Header for mPageScanMode */
// The class instance for the [mPageScanMode] class.
var (
	MPageScanModeClass     _mPageScanModeClass
	MPageScanModeClassOnce sync.Once
)

func getmPageScanModeClass() _mPageScanModeClass {
	MPageScanModeClassOnce.Do(func() {
		MPageScanModeClass = _mPageScanModeClass{objc.GetClass("mPageScanMode")}
	})
	return MPageScanModeClass
}

type _mPageScanModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mPageScanMode */
// An interface definition for the [mPageScanMode] class.
type ImPageScanMode interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mPageScanMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mPageScanMode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mPageScanMode */
// Alloc allocates a new instance without initialization.
func (mc _mPageScanModeClass) Alloc() mPageScanMode {
	rv := objc.Send[mPageScanMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mPageScanModeClass) New() mPageScanMode {
	rv := objc.Send[mPageScanMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mPageScanMode) Init() mPageScanMode {
	rv := objc.Send[mPageScanMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mPageScanMode) Autorelease() mPageScanMode {
	rv := objc.Send[mPageScanMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmPageScanMode creates a new mPageScanMode instance.
func NewmPageScanMode() mPageScanMode {
	return getmPageScanModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mPageScanMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mPageScanMode
type mPageScanMode struct {
	objectivec.Object
}

// mPageScanModeFrom constructs a [mPageScanMode] from an unsafe.Pointer.
func mPageScanModeFrom(ptr unsafe.Pointer) mPageScanMode {
	return mPageScanMode{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mPageScanMode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mPageScanMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mPageScanMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mPageScanMode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mPageScanMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mPageScanMode */



