// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mActionArgument3 */


/* debug [class_header]: Header for mActionArgument3 */
// The class instance for the [mActionArgument3] class.
var (
	MActionArgument3Class     _mActionArgument3Class
	MActionArgument3ClassOnce sync.Once
)

func getmActionArgument3Class() _mActionArgument3Class {
	MActionArgument3ClassOnce.Do(func() {
		MActionArgument3Class = _mActionArgument3Class{objc.GetClass("mActionArgument3")}
	})
	return MActionArgument3Class
}

type _mActionArgument3Class struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mActionArgument3 */
// An interface definition for the [mActionArgument3] class.
type ImActionArgument3 interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mActionArgument3 */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mActionArgument3 */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mActionArgument3 */
// Alloc allocates a new instance without initialization.
func (mc _mActionArgument3Class) Alloc() mActionArgument3 {
	rv := objc.Send[mActionArgument3](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mActionArgument3Class) New() mActionArgument3 {
	rv := objc.Send[mActionArgument3](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mActionArgument3) Init() mActionArgument3 {
	rv := objc.Send[mActionArgument3](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mActionArgument3) Autorelease() mActionArgument3 {
	rv := objc.Send[mActionArgument3](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmActionArgument3 creates a new mActionArgument3 instance.
func NewmActionArgument3() mActionArgument3 {
	return getmActionArgument3Class().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mActionArgument3 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mActionArgument3
type mActionArgument3 struct {
	objectivec.Object
}

// mActionArgument3From constructs a [mActionArgument3] from an unsafe.Pointer.
func mActionArgument3From(ptr unsafe.Pointer) mActionArgument3 {
	return mActionArgument3{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mActionArgument3 *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mActionArgument3 */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mActionArgument3 */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mActionArgument3 */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mActionArgument3 */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mActionArgument3 */



