// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mUUIDData */


/* debug [class_header]: Header for mUUIDData */
// The class instance for the [mUUIDData] class.
var (
	MUUIDDataClass     _mUUIDDataClass
	MUUIDDataClassOnce sync.Once
)

func getmUUIDDataClass() _mUUIDDataClass {
	MUUIDDataClassOnce.Do(func() {
		MUUIDDataClass = _mUUIDDataClass{objc.GetClass("mUUIDData")}
	})
	return MUUIDDataClass
}

type _mUUIDDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mUUIDData */
// An interface definition for the [mUUIDData] class.
type ImUUIDData interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mUUIDData */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mUUIDData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mUUIDData */
// Alloc allocates a new instance without initialization.
func (mc _mUUIDDataClass) Alloc() mUUIDData {
	rv := objc.Send[mUUIDData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mUUIDDataClass) New() mUUIDData {
	rv := objc.Send[mUUIDData](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mUUIDData) Init() mUUIDData {
	rv := objc.Send[mUUIDData](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mUUIDData) Autorelease() mUUIDData {
	rv := objc.Send[mUUIDData](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmUUIDData creates a new mUUIDData instance.
func NewmUUIDData() mUUIDData {
	return getmUUIDDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mUUIDData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/mUUIDData
type mUUIDData struct {
	objectivec.Object
}

// mUUIDDataFrom constructs a [mUUIDData] from an unsafe.Pointer.
func mUUIDDataFrom(ptr unsafe.Pointer) mUUIDData {
	return mUUIDData{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mUUIDData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mUUIDData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mUUIDData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mUUIDData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mUUIDData */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mUUIDData */



