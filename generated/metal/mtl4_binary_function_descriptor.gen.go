// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MTL4BinaryFunctionDescriptor] class.
var (
	MTL4BinaryFunctionDescriptorClass     _MTL4BinaryFunctionDescriptorClass
	MTL4BinaryFunctionDescriptorClassOnce sync.Once
)

func getMTL4BinaryFunctionDescriptorClass() _MTL4BinaryFunctionDescriptorClass {
	MTL4BinaryFunctionDescriptorClassOnce.Do(func() {
		MTL4BinaryFunctionDescriptorClass = _MTL4BinaryFunctionDescriptorClass{objc.GetClass("MTL4BinaryFunctionDescriptor")}
	})
	return MTL4BinaryFunctionDescriptorClass
}

type _MTL4BinaryFunctionDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [MTL4BinaryFunctionDescriptor] class.
type IMTL4BinaryFunctionDescriptor interface {
	objectivec.IObject
	

	// properties:
	FunctionDescriptor() IMTL4FunctionDescriptor
	SetFunctionDescriptor(value IMTL4FunctionDescriptor)
	Name() foundation.foundation.INSString
	SetName(value foundation.foundation.INSString)
	Options() MTL4BinaryFunctionOptions
	SetOptions(value MTL4BinaryFunctionOptions)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MTL4BinaryFunctionDescriptorClass) Alloc() MTL4BinaryFunctionDescriptor {
	rv := objc.Send[MTL4BinaryFunctionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4BinaryFunctionDescriptorClass) New() MTL4BinaryFunctionDescriptor {
	rv := objc.Send[MTL4BinaryFunctionDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4BinaryFunctionDescriptor) Init() MTL4BinaryFunctionDescriptor {
	rv := objc.Send[MTL4BinaryFunctionDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4BinaryFunctionDescriptor) Autorelease() MTL4BinaryFunctionDescriptor {
	rv := objc.Send[MTL4BinaryFunctionDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4BinaryFunctionDescriptor creates a new MTL4BinaryFunctionDescriptor instance.
func NewMTL4BinaryFunctionDescriptor() MTL4BinaryFunctionDescriptor {
	return getMTL4BinaryFunctionDescriptorClass().New()
}





// Base interface for other function-derived interfaces.


// Base interface for other function-derived interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4BinaryFunctionDescriptor
type MTL4BinaryFunctionDescriptor struct {
	objectivec.Object
}

// MTL4BinaryFunctionDescriptorFrom constructs a [MTL4BinaryFunctionDescriptor] from an unsafe.Pointer.
//
// Base interface for other function-derived interfaces.
func MTL4BinaryFunctionDescriptorFrom(ptr unsafe.Pointer) MTL4BinaryFunctionDescriptor {
	return MTL4BinaryFunctionDescriptor{objectivec.Object{objc.ID(ptr)}}
}

























// Provides the function descriptor corresponding to the function to compile into a binary function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4BinaryFunctionDescriptor/functionDescriptor
func (m_ MTL4BinaryFunctionDescriptor) FunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](m_.ID, objc.Sel("functionDescriptor"))
	return rv
}


// Provides the function descriptor corresponding to the function to compile into a binary function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4BinaryFunctionDescriptor/functionDescriptor
func (m_ MTL4BinaryFunctionDescriptor) SetFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFunctionDescriptor:"), value)
}


// Associates a string that uniquely identifies a binary function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4BinaryFunctionDescriptor/name
func (m_ MTL4BinaryFunctionDescriptor) Name() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}


// Associates a string that uniquely identifies a binary function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4BinaryFunctionDescriptor/name
func (m_ MTL4BinaryFunctionDescriptor) SetName(value foundation.foundation.INSString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}


// Configure the options to use at binary function creation time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4BinaryFunctionDescriptor/options
func (m_ MTL4BinaryFunctionDescriptor) Options() MTL4BinaryFunctionOptions {
	rv := objc.Send[MTL4BinaryFunctionOptions](m_.ID, objc.Sel("options"))
	return rv
}


// Configure the options to use at binary function creation time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4BinaryFunctionDescriptor/options
func (m_ MTL4BinaryFunctionDescriptor) SetOptions(value MTL4BinaryFunctionOptions) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptions:"), value)
}








