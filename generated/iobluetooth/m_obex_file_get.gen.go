// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mOBEXFileGet */


/* debug [class_header]: Header for mOBEXFileGet */
// The class instance for the [mOBEXFileGet] class.
var (
	MOBEXFileGetClass     _mOBEXFileGetClass
	MOBEXFileGetClassOnce sync.Once
)

func getmOBEXFileGetClass() _mOBEXFileGetClass {
	MOBEXFileGetClassOnce.Do(func() {
		MOBEXFileGetClass = _mOBEXFileGetClass{objc.GetClass("mOBEXFileGet")}
	})
	return MOBEXFileGetClass
}

type _mOBEXFileGetClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mOBEXFileGet */
// An interface definition for the [mOBEXFileGet] class.
type ImOBEXFileGet interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mOBEXFileGet */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mOBEXFileGet */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mOBEXFileGet */
// Alloc allocates a new instance without initialization.
func (mc _mOBEXFileGetClass) Alloc() mOBEXFileGet {
	rv := objc.Send[mOBEXFileGet](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mOBEXFileGetClass) New() mOBEXFileGet {
	rv := objc.Send[mOBEXFileGet](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOBEXFileGet) Init() mOBEXFileGet {
	rv := objc.Send[mOBEXFileGet](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOBEXFileGet) Autorelease() mOBEXFileGet {
	rv := objc.Send[mOBEXFileGet](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOBEXFileGet creates a new mOBEXFileGet instance.
func NewmOBEXFileGet() mOBEXFileGet {
	return getmOBEXFileGetClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mOBEXFileGet */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mOBEXFileGet
type mOBEXFileGet struct {
	objectivec.Object
}

// mOBEXFileGetFrom constructs a [mOBEXFileGet] from an unsafe.Pointer.
func mOBEXFileGetFrom(ptr unsafe.Pointer) mOBEXFileGet {
	return mOBEXFileGet{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mOBEXFileGet *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mOBEXFileGet */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mOBEXFileGet */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mOBEXFileGet */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mOBEXFileGet */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mOBEXFileGet */



