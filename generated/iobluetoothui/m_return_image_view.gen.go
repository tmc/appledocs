// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mReturnImageView */


/* debug [class_header]: Header for mReturnImageView */
// The class instance for the [mReturnImageView] class.
var (
	MReturnImageViewClass     _mReturnImageViewClass
	MReturnImageViewClassOnce sync.Once
)

func getmReturnImageViewClass() _mReturnImageViewClass {
	MReturnImageViewClassOnce.Do(func() {
		MReturnImageViewClass = _mReturnImageViewClass{objc.GetClass("mReturnImageView")}
	})
	return MReturnImageViewClass
}

type _mReturnImageViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mReturnImageView */
// An interface definition for the [mReturnImageView] class.
type ImReturnImageView interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mReturnImageView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mReturnImageView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mReturnImageView */
// Alloc allocates a new instance without initialization.
func (mc _mReturnImageViewClass) Alloc() mReturnImageView {
	rv := objc.Send[mReturnImageView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mReturnImageViewClass) New() mReturnImageView {
	rv := objc.Send[mReturnImageView](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mReturnImageView) Init() mReturnImageView {
	rv := objc.Send[mReturnImageView](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mReturnImageView) Autorelease() mReturnImageView {
	rv := objc.Send[mReturnImageView](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmReturnImageView creates a new mReturnImageView instance.
func NewmReturnImageView() mReturnImageView {
	return getmReturnImageViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mReturnImageView */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/mReturnImageView
type mReturnImageView struct {
	objectivec.Object
}

// mReturnImageViewFrom constructs a [mReturnImageView] from an unsafe.Pointer.
func mReturnImageViewFrom(ptr unsafe.Pointer) mReturnImageView {
	return mReturnImageView{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mReturnImageView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mReturnImageView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mReturnImageView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mReturnImageView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mReturnImageView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mReturnImageView */



