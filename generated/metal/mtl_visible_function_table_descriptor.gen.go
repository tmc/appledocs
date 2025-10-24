// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLVisibleFunctionTableDescriptor */


/* debug [class_header]: Header for MTLVisibleFunctionTableDescriptor */
// The class instance for the [VisibleFunctionTableDescriptor] class.
var (
	VisibleFunctionTableDescriptorClass     _VisibleFunctionTableDescriptorClass
	VisibleFunctionTableDescriptorClassOnce sync.Once
)

func getVisibleFunctionTableDescriptorClass() _VisibleFunctionTableDescriptorClass {
	VisibleFunctionTableDescriptorClassOnce.Do(func() {
		VisibleFunctionTableDescriptorClass = _VisibleFunctionTableDescriptorClass{objc.GetClass("MTLVisibleFunctionTableDescriptor")}
	})
	return VisibleFunctionTableDescriptorClass
}

type _VisibleFunctionTableDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VisibleFunctionTableDescriptor */
// An interface definition for the [VisibleFunctionTableDescriptor] class.
type IVisibleFunctionTableDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VisibleFunctionTableDescriptor */
	// properties:
	FunctionCount() uint
	SetFunctionCount(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VisibleFunctionTableDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VisibleFunctionTableDescriptor */
// Alloc allocates a new instance without initialization.
func (vc _VisibleFunctionTableDescriptorClass) Alloc() VisibleFunctionTableDescriptor {
	rv := objc.Send[VisibleFunctionTableDescriptor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VisibleFunctionTableDescriptorClass) New() VisibleFunctionTableDescriptor {
	rv := objc.Send[VisibleFunctionTableDescriptor](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VisibleFunctionTableDescriptor) Init() VisibleFunctionTableDescriptor {
	rv := objc.Send[VisibleFunctionTableDescriptor](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VisibleFunctionTableDescriptor) Autorelease() VisibleFunctionTableDescriptor {
	rv := objc.Send[VisibleFunctionTableDescriptor](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVisibleFunctionTableDescriptor creates a new VisibleFunctionTableDescriptor instance.
func NewVisibleFunctionTableDescriptor() VisibleFunctionTableDescriptor {
	return getVisibleFunctionTableDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VisibleFunctionTableDescriptor */
// A specification of how to create a visible function table.


// A specification of how to create a visible function table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTableDescriptor
type VisibleFunctionTableDescriptor struct {
	objectivec.Object
}

// VisibleFunctionTableDescriptorFrom constructs a [VisibleFunctionTableDescriptor] from an unsafe.Pointer.
//
// A specification of how to create a visible function table.
func VisibleFunctionTableDescriptorFrom(ptr unsafe.Pointer) VisibleFunctionTableDescriptor {
	return VisibleFunctionTableDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VisibleFunctionTableDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VisibleFunctionTableDescriptor */

// Creates a default visible function table descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTableDescriptor/visibleFunctionTableDescriptor
func (vc _VisibleFunctionTableDescriptorClass) VisibleFunctionTableDescriptor() IVisibleFunctionTableDescriptor {
	rv := objc.Send[VisibleFunctionTableDescriptor](objc.ID(vc.class), objc.Sel("visibleFunctionTableDescriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VisibleFunctionTableDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VisibleFunctionTableDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VisibleFunctionTableDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VisibleFunctionTableDescriptor */

// The number of entries in the function table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTableDescriptor/functionCount
func (v_ VisibleFunctionTableDescriptor) FunctionCount() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("functionCount"))
	return rv
}/* debug [instance_properties/getter]: functionCount */


// The number of entries in the function table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTableDescriptor/functionCount
func (v_ VisibleFunctionTableDescriptor) SetFunctionCount(value uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFunctionCount:"), value)
}/* debug [instance_properties/setter]: functionCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLVisibleFunctionTableDescriptor */





