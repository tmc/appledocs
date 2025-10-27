// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [FunctionDescriptor] class.
var (
	FunctionDescriptorClass     _FunctionDescriptorClass
	FunctionDescriptorClassOnce sync.Once
)

func getFunctionDescriptorClass() _FunctionDescriptorClass {
	FunctionDescriptorClassOnce.Do(func() {
		FunctionDescriptorClass = _FunctionDescriptorClass{objc.GetClass("MTLFunctionDescriptor")}
	})
	return FunctionDescriptorClass
}

type _FunctionDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [FunctionDescriptor] class.
type IFunctionDescriptor interface {
	objectivec.IObject
	

	// properties:
	BinaryArchives() []objc.ID
	SetBinaryArchives(value []objc.ID)
	ConstantValues() IMTLFunctionConstantValues
	SetConstantValues(value IMTLFunctionConstantValues)
	Name() foundation.foundation.INSString
	SetName(value foundation.foundation.INSString)
	Options() FunctionOptions
	SetOptions(value FunctionOptions)
	SpecializedName() foundation.foundation.INSString
	SetSpecializedName(value foundation.foundation.INSString)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (fc _FunctionDescriptorClass) Alloc() FunctionDescriptor {
	rv := objc.Send[FunctionDescriptor](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FunctionDescriptorClass) New() FunctionDescriptor {
	rv := objc.Send[FunctionDescriptor](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FunctionDescriptor) Init() FunctionDescriptor {
	rv := objc.Send[FunctionDescriptor](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FunctionDescriptor) Autorelease() FunctionDescriptor {
	rv := objc.Send[FunctionDescriptor](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFunctionDescriptor creates a new FunctionDescriptor instance.
func NewFunctionDescriptor() FunctionDescriptor {
	return getFunctionDescriptorClass().New()
}





// A description of a function object to create.


// A description of a function object to create.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor
type FunctionDescriptor struct {
	objectivec.Object
}

// FunctionDescriptorFrom constructs a [FunctionDescriptor] from an unsafe.Pointer.
//
// A description of a function object to create.
func FunctionDescriptorFrom(ptr unsafe.Pointer) FunctionDescriptor {
	return FunctionDescriptor{objectivec.Object{objc.ID(ptr)}}
}










// Creates a default function descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/functionDescriptor
func (fc _FunctionDescriptorClass) FunctionDescriptor() IFunctionDescriptor {
	rv := objc.Send[FunctionDescriptor](objc.ID(fc.class), objc.Sel("functionDescriptor"))
	return rv
}

















// The binary archives to search for a previously-compiled version of this function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/binaryArchives
func (f_ FunctionDescriptor) BinaryArchives() []objc.ID {
	rv := objc.Send[[]objc.ID](f_.ID, objc.Sel("binaryArchives"))
	return rv
}


// The binary archives to search for a previously-compiled version of this function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/binaryArchives
func (f_ FunctionDescriptor) SetBinaryArchives(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](f_.ID, objc.Sel("setBinaryArchives:"), nsArray)
}


// The set of constant values assigned to the function constants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/constantValues
func (f_ FunctionDescriptor) ConstantValues() IMTLFunctionConstantValues {
	rv := objc.Send[FunctionConstantValues](f_.ID, objc.Sel("constantValues"))
	return rv
}


// The set of constant values assigned to the function constants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/constantValues
func (f_ FunctionDescriptor) SetConstantValues(value IMTLFunctionConstantValues) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setConstantValues:"), value)
}


// The name of the function to fetch from the library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/name
func (f_ FunctionDescriptor) Name() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("name"))
	return rv
}


// The name of the function to fetch from the library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/name
func (f_ FunctionDescriptor) SetName(value foundation.foundation.INSString) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setName:"), value)
}


// Flags specifying how Metal should create the new function object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/options
func (f_ FunctionDescriptor) Options() FunctionOptions {
	rv := objc.Send[FunctionOptions](f_.ID, objc.Sel("options"))
	return rv
}


// Flags specifying how Metal should create the new function object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/options
func (f_ FunctionDescriptor) SetOptions(value FunctionOptions) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOptions:"), value)
}


// A new name for the created function object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/specializedName
func (f_ FunctionDescriptor) SpecializedName() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("specializedName"))
	return rv
}


// A new name for the created function object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/specializedName
func (f_ FunctionDescriptor) SetSpecializedName(value foundation.foundation.INSString) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSpecializedName:"), value)
}








