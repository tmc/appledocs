// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mPasskeyIndex */


/* debug [class_header]: Header for mPasskeyIndex */
// The class instance for the [mPasskeyIndex] class.
var (
	MPasskeyIndexClass     _mPasskeyIndexClass
	MPasskeyIndexClassOnce sync.Once
)

func getmPasskeyIndexClass() _mPasskeyIndexClass {
	MPasskeyIndexClassOnce.Do(func() {
		MPasskeyIndexClass = _mPasskeyIndexClass{objc.GetClass("mPasskeyIndex")}
	})
	return MPasskeyIndexClass
}

type _mPasskeyIndexClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mPasskeyIndex */
// An interface definition for the [mPasskeyIndex] class.
type ImPasskeyIndex interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mPasskeyIndex */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mPasskeyIndex */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mPasskeyIndex */
// Alloc allocates a new instance without initialization.
func (mc _mPasskeyIndexClass) Alloc() mPasskeyIndex {
	rv := objc.Send[mPasskeyIndex](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mPasskeyIndexClass) New() mPasskeyIndex {
	rv := objc.Send[mPasskeyIndex](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mPasskeyIndex) Init() mPasskeyIndex {
	rv := objc.Send[mPasskeyIndex](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mPasskeyIndex) Autorelease() mPasskeyIndex {
	rv := objc.Send[mPasskeyIndex](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmPasskeyIndex creates a new mPasskeyIndex instance.
func NewmPasskeyIndex() mPasskeyIndex {
	return getmPasskeyIndexClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mPasskeyIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/mPasskeyIndex
type mPasskeyIndex struct {
	objectivec.Object
}

// mPasskeyIndexFrom constructs a [mPasskeyIndex] from an unsafe.Pointer.
func mPasskeyIndexFrom(ptr unsafe.Pointer) mPasskeyIndex {
	return mPasskeyIndex{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mPasskeyIndex *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mPasskeyIndex */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mPasskeyIndex */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mPasskeyIndex */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mPasskeyIndex */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mPasskeyIndex */



