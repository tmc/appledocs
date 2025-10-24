// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mReserved1 */


/* debug [class_header]: Header for mReserved1 */
// The class instance for the [mReserved1] class.
var (
	MReserved1Class     _mReserved1Class
	MReserved1ClassOnce sync.Once
)

func getmReserved1Class() _mReserved1Class {
	MReserved1ClassOnce.Do(func() {
		MReserved1Class = _mReserved1Class{objc.GetClass("mReserved1")}
	})
	return MReserved1Class
}

type _mReserved1Class struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mReserved1 */
// An interface definition for the [mReserved1] class.
type ImReserved1 interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mReserved1 */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mReserved1 */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mReserved1 */
// Alloc allocates a new instance without initialization.
func (mc _mReserved1Class) Alloc() mReserved1 {
	rv := objc.Send[mReserved1](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mReserved1Class) New() mReserved1 {
	rv := objc.Send[mReserved1](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mReserved1) Init() mReserved1 {
	rv := objc.Send[mReserved1](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mReserved1) Autorelease() mReserved1 {
	rv := objc.Send[mReserved1](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmReserved1 creates a new mReserved1 instance.
func NewmReserved1() mReserved1 {
	return getmReserved1Class().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mReserved1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mReserved1
type mReserved1 struct {
	objectivec.Object
}

// mReserved1From constructs a [mReserved1] from an unsafe.Pointer.
func mReserved1From(ptr unsafe.Pointer) mReserved1 {
	return mReserved1{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mReserved1 *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mReserved1 */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mReserved1 */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mReserved1 */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mReserved1 */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mReserved1 */



