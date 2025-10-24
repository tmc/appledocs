// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mLinkType */


/* debug [class_header]: Header for mLinkType */
// The class instance for the [mLinkType] class.
var (
	MLinkTypeClass     _mLinkTypeClass
	MLinkTypeClassOnce sync.Once
)

func getmLinkTypeClass() _mLinkTypeClass {
	MLinkTypeClassOnce.Do(func() {
		MLinkTypeClass = _mLinkTypeClass{objc.GetClass("mLinkType")}
	})
	return MLinkTypeClass
}

type _mLinkTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mLinkType */
// An interface definition for the [mLinkType] class.
type ImLinkType interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mLinkType */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mLinkType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mLinkType */
// Alloc allocates a new instance without initialization.
func (mc _mLinkTypeClass) Alloc() mLinkType {
	rv := objc.Send[mLinkType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mLinkTypeClass) New() mLinkType {
	rv := objc.Send[mLinkType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mLinkType) Init() mLinkType {
	rv := objc.Send[mLinkType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mLinkType) Autorelease() mLinkType {
	rv := objc.Send[mLinkType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmLinkType creates a new mLinkType instance.
func NewmLinkType() mLinkType {
	return getmLinkTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mLinkType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mLinkType
type mLinkType struct {
	objectivec.Object
}

// mLinkTypeFrom constructs a [mLinkType] from an unsafe.Pointer.
func mLinkTypeFrom(ptr unsafe.Pointer) mLinkType {
	return mLinkType{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mLinkType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mLinkType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mLinkType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mLinkType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mLinkType */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mLinkType */



