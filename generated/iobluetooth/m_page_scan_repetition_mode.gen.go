// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mPageScanRepetitionMode */


/* debug [class_header]: Header for mPageScanRepetitionMode */
// The class instance for the [mPageScanRepetitionMode] class.
var (
	MPageScanRepetitionModeClass     _mPageScanRepetitionModeClass
	MPageScanRepetitionModeClassOnce sync.Once
)

func getmPageScanRepetitionModeClass() _mPageScanRepetitionModeClass {
	MPageScanRepetitionModeClassOnce.Do(func() {
		MPageScanRepetitionModeClass = _mPageScanRepetitionModeClass{objc.GetClass("mPageScanRepetitionMode")}
	})
	return MPageScanRepetitionModeClass
}

type _mPageScanRepetitionModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mPageScanRepetitionMode */
// An interface definition for the [mPageScanRepetitionMode] class.
type ImPageScanRepetitionMode interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mPageScanRepetitionMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mPageScanRepetitionMode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mPageScanRepetitionMode */
// Alloc allocates a new instance without initialization.
func (mc _mPageScanRepetitionModeClass) Alloc() mPageScanRepetitionMode {
	rv := objc.Send[mPageScanRepetitionMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mPageScanRepetitionModeClass) New() mPageScanRepetitionMode {
	rv := objc.Send[mPageScanRepetitionMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mPageScanRepetitionMode) Init() mPageScanRepetitionMode {
	rv := objc.Send[mPageScanRepetitionMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mPageScanRepetitionMode) Autorelease() mPageScanRepetitionMode {
	rv := objc.Send[mPageScanRepetitionMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmPageScanRepetitionMode creates a new mPageScanRepetitionMode instance.
func NewmPageScanRepetitionMode() mPageScanRepetitionMode {
	return getmPageScanRepetitionModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mPageScanRepetitionMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mPageScanRepetitionMode
type mPageScanRepetitionMode struct {
	objectivec.Object
}

// mPageScanRepetitionModeFrom constructs a [mPageScanRepetitionMode] from an unsafe.Pointer.
func mPageScanRepetitionModeFrom(ptr unsafe.Pointer) mPageScanRepetitionMode {
	return mPageScanRepetitionMode{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mPageScanRepetitionMode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mPageScanRepetitionMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mPageScanRepetitionMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mPageScanRepetitionMode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mPageScanRepetitionMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mPageScanRepetitionMode */



