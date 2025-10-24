// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mReserved2 */


/* debug [class_header]: Header for mReserved2 */
// The class instance for the [mReserved2] class.
var (
	MReserved2Class     _mReserved2Class
	MReserved2ClassOnce sync.Once
)

func getmReserved2Class() _mReserved2Class {
	MReserved2ClassOnce.Do(func() {
		MReserved2Class = _mReserved2Class{objc.GetClass("mReserved2")}
	})
	return MReserved2Class
}

type _mReserved2Class struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mReserved2 */
// An interface definition for the [mReserved2] class.
type ImReserved2 interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mReserved2 */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mReserved2 */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mReserved2 */
// Alloc allocates a new instance without initialization.
func (mc _mReserved2Class) Alloc() mReserved2 {
	rv := objc.Send[mReserved2](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mReserved2Class) New() mReserved2 {
	rv := objc.Send[mReserved2](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mReserved2) Init() mReserved2 {
	rv := objc.Send[mReserved2](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mReserved2) Autorelease() mReserved2 {
	rv := objc.Send[mReserved2](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmReserved2 creates a new mReserved2 instance.
func NewmReserved2() mReserved2 {
	return getmReserved2Class().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mReserved2 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mReserved2
type mReserved2 struct {
	objectivec.Object
}

// mReserved2From constructs a [mReserved2] from an unsafe.Pointer.
func mReserved2From(ptr unsafe.Pointer) mReserved2 {
	return mReserved2{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mReserved2 *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mReserved2 */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mReserved2 */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mReserved2 */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mReserved2 */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mReserved2 */



