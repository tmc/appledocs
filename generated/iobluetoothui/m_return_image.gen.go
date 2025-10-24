// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mReturnImage */


/* debug [class_header]: Header for mReturnImage */
// The class instance for the [mReturnImage] class.
var (
	MReturnImageClass     _mReturnImageClass
	MReturnImageClassOnce sync.Once
)

func getmReturnImageClass() _mReturnImageClass {
	MReturnImageClassOnce.Do(func() {
		MReturnImageClass = _mReturnImageClass{objc.GetClass("mReturnImage")}
	})
	return MReturnImageClass
}

type _mReturnImageClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mReturnImage */
// An interface definition for the [mReturnImage] class.
type ImReturnImage interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mReturnImage */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mReturnImage */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mReturnImage */
// Alloc allocates a new instance without initialization.
func (mc _mReturnImageClass) Alloc() mReturnImage {
	rv := objc.Send[mReturnImage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mReturnImageClass) New() mReturnImage {
	rv := objc.Send[mReturnImage](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mReturnImage) Init() mReturnImage {
	rv := objc.Send[mReturnImage](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mReturnImage) Autorelease() mReturnImage {
	rv := objc.Send[mReturnImage](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmReturnImage creates a new mReturnImage instance.
func NewmReturnImage() mReturnImage {
	return getmReturnImageClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mReturnImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/mReturnImage
type mReturnImage struct {
	objectivec.Object
}

// mReturnImageFrom constructs a [mReturnImage] from an unsafe.Pointer.
func mReturnImageFrom(ptr unsafe.Pointer) mReturnImage {
	return mReturnImage{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mReturnImage *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mReturnImage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mReturnImage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mReturnImage */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mReturnImage */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mReturnImage */



