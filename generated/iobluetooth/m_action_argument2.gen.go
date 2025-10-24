// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mActionArgument2 */


/* debug [class_header]: Header for mActionArgument2 */
// The class instance for the [mActionArgument2] class.
var (
	MActionArgument2Class     _mActionArgument2Class
	MActionArgument2ClassOnce sync.Once
)

func getmActionArgument2Class() _mActionArgument2Class {
	MActionArgument2ClassOnce.Do(func() {
		MActionArgument2Class = _mActionArgument2Class{objc.GetClass("mActionArgument2")}
	})
	return MActionArgument2Class
}

type _mActionArgument2Class struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mActionArgument2 */
// An interface definition for the [mActionArgument2] class.
type ImActionArgument2 interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mActionArgument2 */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mActionArgument2 */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mActionArgument2 */
// Alloc allocates a new instance without initialization.
func (mc _mActionArgument2Class) Alloc() mActionArgument2 {
	rv := objc.Send[mActionArgument2](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mActionArgument2Class) New() mActionArgument2 {
	rv := objc.Send[mActionArgument2](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mActionArgument2) Init() mActionArgument2 {
	rv := objc.Send[mActionArgument2](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mActionArgument2) Autorelease() mActionArgument2 {
	rv := objc.Send[mActionArgument2](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmActionArgument2 creates a new mActionArgument2 instance.
func NewmActionArgument2() mActionArgument2 {
	return getmActionArgument2Class().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mActionArgument2 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mActionArgument2
type mActionArgument2 struct {
	objectivec.Object
}

// mActionArgument2From constructs a [mActionArgument2] from an unsafe.Pointer.
func mActionArgument2From(ptr unsafe.Pointer) mActionArgument2 {
	return mActionArgument2{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mActionArgument2 *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mActionArgument2 */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mActionArgument2 */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mActionArgument2 */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mActionArgument2 */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mActionArgument2 */



