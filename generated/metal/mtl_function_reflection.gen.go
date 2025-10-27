// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [FunctionReflection] class.
var (
	FunctionReflectionClass     _FunctionReflectionClass
	FunctionReflectionClassOnce sync.Once
)

func getFunctionReflectionClass() _FunctionReflectionClass {
	FunctionReflectionClassOnce.Do(func() {
		FunctionReflectionClass = _FunctionReflectionClass{objc.GetClass("MTLFunctionReflection")}
	})
	return FunctionReflectionClass
}

type _FunctionReflectionClass struct {
	class objc.Class
}





// An interface definition for the [FunctionReflection] class.
type IFunctionReflection interface {
	objectivec.IObject
	

	// properties:
	Bindings() []objc.ID
	UserAnnotation() foundation.foundation.INSString


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (fc _FunctionReflectionClass) Alloc() FunctionReflection {
	rv := objc.Send[FunctionReflection](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FunctionReflectionClass) New() FunctionReflection {
	rv := objc.Send[FunctionReflection](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FunctionReflection) Init() FunctionReflection {
	rv := objc.Send[FunctionReflection](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FunctionReflection) Autorelease() FunctionReflection {
	rv := objc.Send[FunctionReflection](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFunctionReflection creates a new FunctionReflection instance.
func NewFunctionReflection() FunctionReflection {
	return getFunctionReflectionClass().New()
}





// Represents a reflection object containing information about a function in a Metal library.


// Represents a reflection object containing information about a function in a Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionReflection
type FunctionReflection struct {
	objectivec.Object
}

// FunctionReflectionFrom constructs a [FunctionReflection] from an unsafe.Pointer.
//
// Represents a reflection object containing information about a function in a Metal library.
func FunctionReflectionFrom(ptr unsafe.Pointer) FunctionReflection {
	return FunctionReflection{objectivec.Object{objc.ID(ptr)}}
}

























// Provides a list of inputs and outputs of the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionReflection/bindings
func (f_ FunctionReflection) Bindings() []objc.ID {
	rv := objc.Send[[]objc.ID](f_.ID, objc.Sel("bindings"))
	return rv
}


// The string passed to the user annotation attribute for this function. Null if no user annotation is present for this function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionReflection/userAnnotation
func (f_ FunctionReflection) UserAnnotation() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("userAnnotation"))
	return rv
}








