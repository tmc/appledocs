// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mIsServer */


/* debug [class_header]: Header for mIsServer */
// The class instance for the [mIsServer] class.
var (
	MIsServerClass     _mIsServerClass
	MIsServerClassOnce sync.Once
)

func getmIsServerClass() _mIsServerClass {
	MIsServerClassOnce.Do(func() {
		MIsServerClass = _mIsServerClass{objc.GetClass("mIsServer")}
	})
	return MIsServerClass
}

type _mIsServerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mIsServer */
// An interface definition for the [mIsServer] class.
type ImIsServer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mIsServer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mIsServer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mIsServer */
// Alloc allocates a new instance without initialization.
func (mc _mIsServerClass) Alloc() mIsServer {
	rv := objc.Send[mIsServer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mIsServerClass) New() mIsServer {
	rv := objc.Send[mIsServer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIsServer) Init() mIsServer {
	rv := objc.Send[mIsServer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIsServer) Autorelease() mIsServer {
	rv := objc.Send[mIsServer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIsServer creates a new mIsServer instance.
func NewmIsServer() mIsServer {
	return getmIsServerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mIsServer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/mIsServer
type mIsServer struct {
	objectivec.Object
}

// mIsServerFrom constructs a [mIsServer] from an unsafe.Pointer.
func mIsServerFrom(ptr unsafe.Pointer) mIsServer {
	return mIsServer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mIsServer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mIsServer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mIsServer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mIsServer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mIsServer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mIsServer */



