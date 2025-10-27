// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MTL4SpecializedFunctionDescriptor] class.
var (
	MTL4SpecializedFunctionDescriptorClass     _MTL4SpecializedFunctionDescriptorClass
	MTL4SpecializedFunctionDescriptorClassOnce sync.Once
)

func getMTL4SpecializedFunctionDescriptorClass() _MTL4SpecializedFunctionDescriptorClass {
	MTL4SpecializedFunctionDescriptorClassOnce.Do(func() {
		MTL4SpecializedFunctionDescriptorClass = _MTL4SpecializedFunctionDescriptorClass{objc.GetClass("MTL4SpecializedFunctionDescriptor")}
	})
	return MTL4SpecializedFunctionDescriptorClass
}

type _MTL4SpecializedFunctionDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [MTL4SpecializedFunctionDescriptor] class.
type IMTL4SpecializedFunctionDescriptor interface {
	IMTL4FunctionDescriptor
	

	// properties:
	ConstantValues() IMTLFunctionConstantValues
	SetConstantValues(value IMTLFunctionConstantValues)
	FunctionDescriptor() IMTL4FunctionDescriptor
	SetFunctionDescriptor(value IMTL4FunctionDescriptor)
	SpecializedName() foundation.foundation.INSString
	SetSpecializedName(value foundation.foundation.INSString)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MTL4SpecializedFunctionDescriptorClass) Alloc() MTL4SpecializedFunctionDescriptor {
	rv := objc.Send[MTL4SpecializedFunctionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4SpecializedFunctionDescriptorClass) New() MTL4SpecializedFunctionDescriptor {
	rv := objc.Send[MTL4SpecializedFunctionDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4SpecializedFunctionDescriptor) Init() MTL4SpecializedFunctionDescriptor {
	rv := objc.Send[MTL4SpecializedFunctionDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4SpecializedFunctionDescriptor) Autorelease() MTL4SpecializedFunctionDescriptor {
	rv := objc.Send[MTL4SpecializedFunctionDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4SpecializedFunctionDescriptor creates a new MTL4SpecializedFunctionDescriptor instance.
func NewMTL4SpecializedFunctionDescriptor() MTL4SpecializedFunctionDescriptor {
	return getMTL4SpecializedFunctionDescriptorClass().New()
}





// Groups together properties to configure and create a specialized function by passing it to a factory method.
//
// You can pass an instance of this class to any methods that accept a parameter to provide extra configuration, such as function constants or a name.


// Groups together properties to configure and create a specialized function by passing it to a factory method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4SpecializedFunctionDescriptor
type MTL4SpecializedFunctionDescriptor struct {
	MTL4FunctionDescriptor
}

// MTL4SpecializedFunctionDescriptorFrom constructs a [MTL4SpecializedFunctionDescriptor] from an unsafe.Pointer.
//
// Groups together properties to configure and create a specialized function by passing it to a factory method.
func MTL4SpecializedFunctionDescriptorFrom(ptr unsafe.Pointer) MTL4SpecializedFunctionDescriptor {
	return MTL4SpecializedFunctionDescriptor{
		MTL4FunctionDescriptor: MTL4FunctionDescriptorFrom(ptr),
	}
}

























// Configures optional function constant values to associate with the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4SpecializedFunctionDescriptor/constantValues
func (m_ MTL4SpecializedFunctionDescriptor) ConstantValues() IMTLFunctionConstantValues {
	rv := objc.Send[FunctionConstantValues](m_.ID, objc.Sel("constantValues"))
	return rv
}


// Configures optional function constant values to associate with the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4SpecializedFunctionDescriptor/constantValues
func (m_ MTL4SpecializedFunctionDescriptor) SetConstantValues(value IMTLFunctionConstantValues) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConstantValues:"), value)
}


// Provides a descriptor that corresponds to a base function that the specialization applies to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4SpecializedFunctionDescriptor/functionDescriptor
func (m_ MTL4SpecializedFunctionDescriptor) FunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](m_.ID, objc.Sel("functionDescriptor"))
	return rv
}


// Provides a descriptor that corresponds to a base function that the specialization applies to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4SpecializedFunctionDescriptor/functionDescriptor
func (m_ MTL4SpecializedFunctionDescriptor) SetFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFunctionDescriptor:"), value)
}


// Assigns an optional name to the specialized function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4SpecializedFunctionDescriptor/specializedName
func (m_ MTL4SpecializedFunctionDescriptor) SpecializedName() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("specializedName"))
	return rv
}


// Assigns an optional name to the specialized function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4SpecializedFunctionDescriptor/specializedName
func (m_ MTL4SpecializedFunctionDescriptor) SetSpecializedName(value foundation.foundation.INSString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSpecializedName:"), value)
}








