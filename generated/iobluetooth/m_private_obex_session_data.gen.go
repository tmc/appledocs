// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mPrivateOBEXSessionData */


/* debug [class_header]: Header for mPrivateOBEXSessionData */
// The class instance for the [mPrivateOBEXSessionData] class.
var (
	MPrivateOBEXSessionDataClass     _mPrivateOBEXSessionDataClass
	MPrivateOBEXSessionDataClassOnce sync.Once
)

func getmPrivateOBEXSessionDataClass() _mPrivateOBEXSessionDataClass {
	MPrivateOBEXSessionDataClassOnce.Do(func() {
		MPrivateOBEXSessionDataClass = _mPrivateOBEXSessionDataClass{objc.GetClass("mPrivateOBEXSessionData")}
	})
	return MPrivateOBEXSessionDataClass
}

type _mPrivateOBEXSessionDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mPrivateOBEXSessionData */
// An interface definition for the [mPrivateOBEXSessionData] class.
type ImPrivateOBEXSessionData interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mPrivateOBEXSessionData */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mPrivateOBEXSessionData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mPrivateOBEXSessionData */
// Alloc allocates a new instance without initialization.
func (mc _mPrivateOBEXSessionDataClass) Alloc() mPrivateOBEXSessionData {
	rv := objc.Send[mPrivateOBEXSessionData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mPrivateOBEXSessionDataClass) New() mPrivateOBEXSessionData {
	rv := objc.Send[mPrivateOBEXSessionData](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mPrivateOBEXSessionData) Init() mPrivateOBEXSessionData {
	rv := objc.Send[mPrivateOBEXSessionData](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mPrivateOBEXSessionData) Autorelease() mPrivateOBEXSessionData {
	rv := objc.Send[mPrivateOBEXSessionData](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmPrivateOBEXSessionData creates a new mPrivateOBEXSessionData instance.
func NewmPrivateOBEXSessionData() mPrivateOBEXSessionData {
	return getmPrivateOBEXSessionDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mPrivateOBEXSessionData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/mPrivateOBEXSessionData
type mPrivateOBEXSessionData struct {
	objectivec.Object
}

// mPrivateOBEXSessionDataFrom constructs a [mPrivateOBEXSessionData] from an unsafe.Pointer.
func mPrivateOBEXSessionDataFrom(ptr unsafe.Pointer) mPrivateOBEXSessionData {
	return mPrivateOBEXSessionData{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mPrivateOBEXSessionData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mPrivateOBEXSessionData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mPrivateOBEXSessionData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mPrivateOBEXSessionData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mPrivateOBEXSessionData */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mPrivateOBEXSessionData */



