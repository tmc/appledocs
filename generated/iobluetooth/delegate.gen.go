// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class delegate */


/* debug [class_header]: Header for delegate */
// The class instance for the [delegate] class.
var (
	DelegateClass     _delegateClass
	DelegateClassOnce sync.Once
)

func getdelegateClass() _delegateClass {
	DelegateClassOnce.Do(func() {
		DelegateClass = _delegateClass{objc.GetClass("delegate")}
	})
	return DelegateClass
}

type _delegateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for delegate */
// An interface definition for the [delegate] class.
type Idelegate interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for delegate */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for delegate */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for delegate */
// Alloc allocates a new instance without initialization.
func (dc _delegateClass) Alloc() delegate {
	rv := objc.Send[delegate](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _delegateClass) New() delegate {
	rv := objc.Send[delegate](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ delegate) Init() delegate {
	rv := objc.Send[delegate](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ delegate) Autorelease() delegate {
	rv := objc.Send[delegate](d_.ID, objc.Sel("autorelease"))
	return rv
}

// Newdelegate creates a new delegate instance.
func Newdelegate() delegate {
	return getdelegateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for delegate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/delegate-c.ivar
type delegate struct {
	objectivec.Object
}

// delegateFrom constructs a [delegate] from an unsafe.Pointer.
func delegateFrom(ptr unsafe.Pointer) delegate {
	return delegate{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for delegate *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for delegate */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for delegate */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for delegate */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for delegate */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class delegate */



