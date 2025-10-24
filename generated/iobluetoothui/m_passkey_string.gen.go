// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mPasskeyString */


/* debug [class_header]: Header for mPasskeyString */
// The class instance for the [mPasskeyString] class.
var (
	MPasskeyStringClass     _mPasskeyStringClass
	MPasskeyStringClassOnce sync.Once
)

func getmPasskeyStringClass() _mPasskeyStringClass {
	MPasskeyStringClassOnce.Do(func() {
		MPasskeyStringClass = _mPasskeyStringClass{objc.GetClass("mPasskeyString")}
	})
	return MPasskeyStringClass
}

type _mPasskeyStringClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mPasskeyString */
// An interface definition for the [mPasskeyString] class.
type ImPasskeyString interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mPasskeyString */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mPasskeyString */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mPasskeyString */
// Alloc allocates a new instance without initialization.
func (mc _mPasskeyStringClass) Alloc() mPasskeyString {
	rv := objc.Send[mPasskeyString](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mPasskeyStringClass) New() mPasskeyString {
	rv := objc.Send[mPasskeyString](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mPasskeyString) Init() mPasskeyString {
	rv := objc.Send[mPasskeyString](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mPasskeyString) Autorelease() mPasskeyString {
	rv := objc.Send[mPasskeyString](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmPasskeyString creates a new mPasskeyString instance.
func NewmPasskeyString() mPasskeyString {
	return getmPasskeyStringClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mPasskeyString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/mPasskeyString
type mPasskeyString struct {
	objectivec.Object
}

// mPasskeyStringFrom constructs a [mPasskeyString] from an unsafe.Pointer.
func mPasskeyStringFrom(ptr unsafe.Pointer) mPasskeyString {
	return mPasskeyString{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mPasskeyString *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mPasskeyString */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mPasskeyString */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mPasskeyString */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mPasskeyString */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mPasskeyString */



