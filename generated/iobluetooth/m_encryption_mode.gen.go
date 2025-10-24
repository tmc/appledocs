// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mEncryptionMode */


/* debug [class_header]: Header for mEncryptionMode */
// The class instance for the [mEncryptionMode] class.
var (
	MEncryptionModeClass     _mEncryptionModeClass
	MEncryptionModeClassOnce sync.Once
)

func getmEncryptionModeClass() _mEncryptionModeClass {
	MEncryptionModeClassOnce.Do(func() {
		MEncryptionModeClass = _mEncryptionModeClass{objc.GetClass("mEncryptionMode")}
	})
	return MEncryptionModeClass
}

type _mEncryptionModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mEncryptionMode */
// An interface definition for the [mEncryptionMode] class.
type ImEncryptionMode interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mEncryptionMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mEncryptionMode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mEncryptionMode */
// Alloc allocates a new instance without initialization.
func (mc _mEncryptionModeClass) Alloc() mEncryptionMode {
	rv := objc.Send[mEncryptionMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mEncryptionModeClass) New() mEncryptionMode {
	rv := objc.Send[mEncryptionMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mEncryptionMode) Init() mEncryptionMode {
	rv := objc.Send[mEncryptionMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mEncryptionMode) Autorelease() mEncryptionMode {
	rv := objc.Send[mEncryptionMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmEncryptionMode creates a new mEncryptionMode instance.
func NewmEncryptionMode() mEncryptionMode {
	return getmEncryptionModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mEncryptionMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mEncryptionMode
type mEncryptionMode struct {
	objectivec.Object
}

// mEncryptionModeFrom constructs a [mEncryptionMode] from an unsafe.Pointer.
func mEncryptionModeFrom(ptr unsafe.Pointer) mEncryptionMode {
	return mEncryptionMode{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mEncryptionMode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mEncryptionMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mEncryptionMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mEncryptionMode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mEncryptionMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mEncryptionMode */



