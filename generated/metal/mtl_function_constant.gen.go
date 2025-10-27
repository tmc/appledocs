// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [FunctionConstant] class.
var (
	FunctionConstantClass     _FunctionConstantClass
	FunctionConstantClassOnce sync.Once
)

func getFunctionConstantClass() _FunctionConstantClass {
	FunctionConstantClassOnce.Do(func() {
		FunctionConstantClass = _FunctionConstantClass{objc.GetClass("MTLFunctionConstant")}
	})
	return FunctionConstantClass
}

type _FunctionConstantClass struct {
	class objc.Class
}





// An interface definition for the [FunctionConstant] class.
type IFunctionConstant interface {
	objectivec.IObject
	

	// properties:
	Index() uint
	Name() foundation.foundation.INSString
	Required() bool
	Type() DataType


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (fc _FunctionConstantClass) Alloc() FunctionConstant {
	rv := objc.Send[FunctionConstant](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FunctionConstantClass) New() FunctionConstant {
	rv := objc.Send[FunctionConstant](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FunctionConstant) Init() FunctionConstant {
	rv := objc.Send[FunctionConstant](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FunctionConstant) Autorelease() FunctionConstant {
	rv := objc.Send[FunctionConstant](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFunctionConstant creates a new FunctionConstant instance.
func NewFunctionConstant() FunctionConstant {
	return getFunctionConstantClass().New()
}





// A constant that specializes the behavior of a shader.
//
// Don’t create an instance directly. Instead, the list of function constants for a function by querying the property of an instance. An instance should only be obtained from a nonspecialized function created with the method. You only need an instance if you don’t have sufficient information to create an instance used to create a specialized function with the or method.


// A constant that specializes the behavior of a shader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstant
type FunctionConstant struct {
	objectivec.Object
}

// FunctionConstantFrom constructs a [FunctionConstant] from an unsafe.Pointer.
//
// A constant that specializes the behavior of a shader.
func FunctionConstantFrom(ptr unsafe.Pointer) FunctionConstant {
	return FunctionConstant{objectivec.Object{objc.ID(ptr)}}
}

























// The index of the function constant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstant/index
func (f_ FunctionConstant) Index() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("index"))
	return rv
}


// The name of the function constant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstant/name
func (f_ FunctionConstant) Name() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("name"))
	return rv
}


// A Boolean value indicating whether the function constant must be provided to specialize the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstant/required
func (f_ FunctionConstant) Required() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("required"))
	return rv
}


// The data type of the function constant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstant/type
func (f_ FunctionConstant) Type() DataType {
	rv := objc.Send[DataType](f_.ID, objc.Sel("type"))
	return rv
}








