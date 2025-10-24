// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLFunctionStitchingAttributeAlwaysInline */


/* debug [class_header]: Header for MTLFunctionStitchingAttributeAlwaysInline */
// The class instance for the [FunctionStitchingAttributeAlwaysInline] class.
var (
	FunctionStitchingAttributeAlwaysInlineClass     _FunctionStitchingAttributeAlwaysInlineClass
	FunctionStitchingAttributeAlwaysInlineClassOnce sync.Once
)

func getFunctionStitchingAttributeAlwaysInlineClass() _FunctionStitchingAttributeAlwaysInlineClass {
	FunctionStitchingAttributeAlwaysInlineClassOnce.Do(func() {
		FunctionStitchingAttributeAlwaysInlineClass = _FunctionStitchingAttributeAlwaysInlineClass{objc.GetClass("MTLFunctionStitchingAttributeAlwaysInline")}
	})
	return FunctionStitchingAttributeAlwaysInlineClass
}

type _FunctionStitchingAttributeAlwaysInlineClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FunctionStitchingAttributeAlwaysInline */
// An interface definition for the [FunctionStitchingAttributeAlwaysInline] class.
type IFunctionStitchingAttributeAlwaysInline interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FunctionStitchingAttributeAlwaysInline */
	// properties:
	Attributes() FunctionStitchingAttribute /* not a class type */
	SetAttributes(value FunctionStitchingAttribute /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FunctionStitchingAttributeAlwaysInline */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FunctionStitchingAttributeAlwaysInline */
// Alloc allocates a new instance without initialization.
func (fc _FunctionStitchingAttributeAlwaysInlineClass) Alloc() FunctionStitchingAttributeAlwaysInline {
	rv := objc.Send[FunctionStitchingAttributeAlwaysInline](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FunctionStitchingAttributeAlwaysInlineClass) New() FunctionStitchingAttributeAlwaysInline {
	rv := objc.Send[FunctionStitchingAttributeAlwaysInline](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FunctionStitchingAttributeAlwaysInline) Init() FunctionStitchingAttributeAlwaysInline {
	rv := objc.Send[FunctionStitchingAttributeAlwaysInline](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FunctionStitchingAttributeAlwaysInline) Autorelease() FunctionStitchingAttributeAlwaysInline {
	rv := objc.Send[FunctionStitchingAttributeAlwaysInline](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFunctionStitchingAttributeAlwaysInline creates a new FunctionStitchingAttributeAlwaysInline instance.
func NewFunctionStitchingAttributeAlwaysInline() FunctionStitchingAttributeAlwaysInline {
	return getFunctionStitchingAttributeAlwaysInlineClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FunctionStitchingAttributeAlwaysInline */
// An attribute to specify that Metal needs to inline all of the function calls when generating the stitched function.
//
// To inline functions in a call graph, instantiate an instance of this class and assign it as an attribute on the .


// An attribute to specify that Metal needs to inline all of the function calls when generating the stitched function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingAttributeAlwaysInline
type FunctionStitchingAttributeAlwaysInline struct {
	objectivec.Object
}

// FunctionStitchingAttributeAlwaysInlineFrom constructs a [FunctionStitchingAttributeAlwaysInline] from an unsafe.Pointer.
//
// An attribute to specify that Metal needs to inline all of the function calls when generating the stitched function.
func FunctionStitchingAttributeAlwaysInlineFrom(ptr unsafe.Pointer) FunctionStitchingAttributeAlwaysInline {
	return FunctionStitchingAttributeAlwaysInline{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FunctionStitchingAttributeAlwaysInline *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FunctionStitchingAttributeAlwaysInline */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FunctionStitchingAttributeAlwaysInline */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FunctionStitchingAttributeAlwaysInline */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FunctionStitchingAttributeAlwaysInline */

// A list of attributes to configure how the Metal device object generates the new stitched function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinggraph/attributes
func (f_ FunctionStitchingAttributeAlwaysInline) Attributes() FunctionStitchingAttribute /* not a class type */ {
	rv := objc.Send[FunctionStitchingAttribute](f_.ID, objc.Sel("attributes"))
	return rv
}/* debug [instance_properties/getter]: attributes */


// A list of attributes to configure how the Metal device object generates the new stitched function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinggraph/attributes
func (f_ FunctionStitchingAttributeAlwaysInline) SetAttributes(value FunctionStitchingAttribute /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAttributes:"), value)
}/* debug [instance_properties/setter]: attributes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLFunctionStitchingAttributeAlwaysInline */



