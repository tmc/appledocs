// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mOBEXSession */


/* debug [class_header]: Header for mOBEXSession */
// The class instance for the [mOBEXSession] class.
var (
	MOBEXSessionClass     _mOBEXSessionClass
	MOBEXSessionClassOnce sync.Once
)

func getmOBEXSessionClass() _mOBEXSessionClass {
	MOBEXSessionClassOnce.Do(func() {
		MOBEXSessionClass = _mOBEXSessionClass{objc.GetClass("mOBEXSession")}
	})
	return MOBEXSessionClass
}

type _mOBEXSessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mOBEXSession */
// An interface definition for the [mOBEXSession] class.
type ImOBEXSession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mOBEXSession */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mOBEXSession */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mOBEXSession */
// Alloc allocates a new instance without initialization.
func (mc _mOBEXSessionClass) Alloc() mOBEXSession {
	rv := objc.Send[mOBEXSession](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mOBEXSessionClass) New() mOBEXSession {
	rv := objc.Send[mOBEXSession](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOBEXSession) Init() mOBEXSession {
	rv := objc.Send[mOBEXSession](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOBEXSession) Autorelease() mOBEXSession {
	rv := objc.Send[mOBEXSession](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOBEXSession creates a new mOBEXSession instance.
func NewmOBEXSession() mOBEXSession {
	return getmOBEXSessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mOBEXSession */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mOBEXSession
type mOBEXSession struct {
	objectivec.Object
}

// mOBEXSessionFrom constructs a [mOBEXSession] from an unsafe.Pointer.
func mOBEXSessionFrom(ptr unsafe.Pointer) mOBEXSession {
	return mOBEXSession{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mOBEXSession *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mOBEXSession */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mOBEXSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mOBEXSession */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mOBEXSession */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mOBEXSession */



