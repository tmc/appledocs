// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mReturnHighlightImage */


/* debug [class_header]: Header for mReturnHighlightImage */
// The class instance for the [mReturnHighlightImage] class.
var (
	MReturnHighlightImageClass     _mReturnHighlightImageClass
	MReturnHighlightImageClassOnce sync.Once
)

func getmReturnHighlightImageClass() _mReturnHighlightImageClass {
	MReturnHighlightImageClassOnce.Do(func() {
		MReturnHighlightImageClass = _mReturnHighlightImageClass{objc.GetClass("mReturnHighlightImage")}
	})
	return MReturnHighlightImageClass
}

type _mReturnHighlightImageClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mReturnHighlightImage */
// An interface definition for the [mReturnHighlightImage] class.
type ImReturnHighlightImage interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mReturnHighlightImage */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mReturnHighlightImage */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mReturnHighlightImage */
// Alloc allocates a new instance without initialization.
func (mc _mReturnHighlightImageClass) Alloc() mReturnHighlightImage {
	rv := objc.Send[mReturnHighlightImage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mReturnHighlightImageClass) New() mReturnHighlightImage {
	rv := objc.Send[mReturnHighlightImage](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mReturnHighlightImage) Init() mReturnHighlightImage {
	rv := objc.Send[mReturnHighlightImage](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mReturnHighlightImage) Autorelease() mReturnHighlightImage {
	rv := objc.Send[mReturnHighlightImage](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmReturnHighlightImage creates a new mReturnHighlightImage instance.
func NewmReturnHighlightImage() mReturnHighlightImage {
	return getmReturnHighlightImageClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mReturnHighlightImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/mReturnHighlightImage
type mReturnHighlightImage struct {
	objectivec.Object
}

// mReturnHighlightImageFrom constructs a [mReturnHighlightImage] from an unsafe.Pointer.
func mReturnHighlightImageFrom(ptr unsafe.Pointer) mReturnHighlightImage {
	return mReturnHighlightImage{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mReturnHighlightImage *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mReturnHighlightImage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mReturnHighlightImage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mReturnHighlightImage */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mReturnHighlightImage */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mReturnHighlightImage */



