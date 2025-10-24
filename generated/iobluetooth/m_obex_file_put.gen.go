// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mOBEXFilePut */


/* debug [class_header]: Header for mOBEXFilePut */
// The class instance for the [mOBEXFilePut] class.
var (
	MOBEXFilePutClass     _mOBEXFilePutClass
	MOBEXFilePutClassOnce sync.Once
)

func getmOBEXFilePutClass() _mOBEXFilePutClass {
	MOBEXFilePutClassOnce.Do(func() {
		MOBEXFilePutClass = _mOBEXFilePutClass{objc.GetClass("mOBEXFilePut")}
	})
	return MOBEXFilePutClass
}

type _mOBEXFilePutClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mOBEXFilePut */
// An interface definition for the [mOBEXFilePut] class.
type ImOBEXFilePut interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mOBEXFilePut */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mOBEXFilePut */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mOBEXFilePut */
// Alloc allocates a new instance without initialization.
func (mc _mOBEXFilePutClass) Alloc() mOBEXFilePut {
	rv := objc.Send[mOBEXFilePut](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mOBEXFilePutClass) New() mOBEXFilePut {
	rv := objc.Send[mOBEXFilePut](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOBEXFilePut) Init() mOBEXFilePut {
	rv := objc.Send[mOBEXFilePut](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOBEXFilePut) Autorelease() mOBEXFilePut {
	rv := objc.Send[mOBEXFilePut](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOBEXFilePut creates a new mOBEXFilePut instance.
func NewmOBEXFilePut() mOBEXFilePut {
	return getmOBEXFilePutClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mOBEXFilePut */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mOBEXFilePut
type mOBEXFilePut struct {
	objectivec.Object
}

// mOBEXFilePutFrom constructs a [mOBEXFilePut] from an unsafe.Pointer.
func mOBEXFilePutFrom(ptr unsafe.Pointer) mOBEXFilePut {
	return mOBEXFilePut{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mOBEXFilePut *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mOBEXFilePut */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mOBEXFilePut */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mOBEXFilePut */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mOBEXFilePut */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mOBEXFilePut */



