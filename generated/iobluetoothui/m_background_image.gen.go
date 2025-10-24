// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mBackgroundImage */


/* debug [class_header]: Header for mBackgroundImage */
// The class instance for the [mBackgroundImage] class.
var (
	MBackgroundImageClass     _mBackgroundImageClass
	MBackgroundImageClassOnce sync.Once
)

func getmBackgroundImageClass() _mBackgroundImageClass {
	MBackgroundImageClassOnce.Do(func() {
		MBackgroundImageClass = _mBackgroundImageClass{objc.GetClass("mBackgroundImage")}
	})
	return MBackgroundImageClass
}

type _mBackgroundImageClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mBackgroundImage */
// An interface definition for the [mBackgroundImage] class.
type ImBackgroundImage interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mBackgroundImage */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mBackgroundImage */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mBackgroundImage */
// Alloc allocates a new instance without initialization.
func (mc _mBackgroundImageClass) Alloc() mBackgroundImage {
	rv := objc.Send[mBackgroundImage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mBackgroundImageClass) New() mBackgroundImage {
	rv := objc.Send[mBackgroundImage](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mBackgroundImage) Init() mBackgroundImage {
	rv := objc.Send[mBackgroundImage](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mBackgroundImage) Autorelease() mBackgroundImage {
	rv := objc.Send[mBackgroundImage](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmBackgroundImage creates a new mBackgroundImage instance.
func NewmBackgroundImage() mBackgroundImage {
	return getmBackgroundImageClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mBackgroundImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/mBackgroundImage
type mBackgroundImage struct {
	objectivec.Object
}

// mBackgroundImageFrom constructs a [mBackgroundImage] from an unsafe.Pointer.
func mBackgroundImageFrom(ptr unsafe.Pointer) mBackgroundImage {
	return mBackgroundImage{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mBackgroundImage *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mBackgroundImage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mBackgroundImage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mBackgroundImage */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mBackgroundImage */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mBackgroundImage */



