// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [FunctionStitchingAttributeAlwaysInline] class.
type IFunctionStitchingAttributeAlwaysInline interface {
	objectivec.IObject
}

// An attribute to specify that Metal needs to inline all of the function calls when generating the stitched function.
//
// To inline functions in a call graph, instantiate an instance of this class and assign it as an attribute on the .
//
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

// Alloc allocates a new instance without initialization.
func (fc _FunctionStitchingAttributeAlwaysInlineClass) Alloc() FunctionStitchingAttributeAlwaysInline {
	rv := objc.Send[FunctionStitchingAttributeAlwaysInline](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A list of attributes to configure how the Metal device object generates the new stitched function.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinggraph/attributes
func (f_ FunctionStitchingAttributeAlwaysInline) Attributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("attributes"))
	return rv
}


// SetAttributes sets the value of the attributes property.
// A list of attributes to configure how the Metal device object generates the new stitched function.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinggraph/attributes
func (f_ FunctionStitchingAttributeAlwaysInline) SetAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAttributes:"), value)
}



