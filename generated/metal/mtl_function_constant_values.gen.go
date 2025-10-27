// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [FunctionConstantValues] class.
var (
	FunctionConstantValuesClass     _FunctionConstantValuesClass
	FunctionConstantValuesClassOnce sync.Once
)

func getFunctionConstantValuesClass() _FunctionConstantValuesClass {
	FunctionConstantValuesClassOnce.Do(func() {
		FunctionConstantValuesClass = _FunctionConstantValuesClass{objc.GetClass("MTLFunctionConstantValues")}
	})
	return FunctionConstantValuesClass
}

type _FunctionConstantValuesClass struct {
	class objc.Class
}





// An interface definition for the [FunctionConstantValues] class.
type IFunctionConstantValues interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	Reset()
	SetConstantValueTypeAtIndex(value objectivec.IObject, type_ DataType, index uint)
	SetConstantValueTypeWithName(value objectivec.IObject, type_ DataType, name foundation.foundation.INSString)
	SetConstantValuesTypeWithRange(values objectivec.IObject, type_ DataType, range_ foundation.Range)


}





// Alloc allocates a new instance without initialization.
func (fc _FunctionConstantValuesClass) Alloc() FunctionConstantValues {
	rv := objc.Send[FunctionConstantValues](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FunctionConstantValuesClass) New() FunctionConstantValues {
	rv := objc.Send[FunctionConstantValues](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FunctionConstantValues) Init() FunctionConstantValues {
	rv := objc.Send[FunctionConstantValues](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FunctionConstantValues) Autorelease() FunctionConstantValues {
	rv := objc.Send[FunctionConstantValues](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFunctionConstantValues creates a new FunctionConstantValues instance.
func NewFunctionConstantValues() FunctionConstantValues {
	return getFunctionConstantValuesClass().New()
}





// A set of constant values that specialize a graphics or compute GPU function.
//
// An instance sets constant values for function constants. You declare function constants with the attribute in MSL (Metal Shading Language) source code. See the for more information. With an instance, you can set each constant value individually with an index or a name, or set multiple constant values with an index range. You can apply a single instance to multiple instances of any kind, such as a vertex function and a fragment function. When you create a specialized function, subsequent changes to its constant values have no effect. However, you can reset, add, or modify a constant value in your instance and reuse it to create another instance.


// A set of constant values that specialize a graphics or compute GPU function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstantValues
type FunctionConstantValues struct {
	objectivec.Object
}

// FunctionConstantValuesFrom constructs a [FunctionConstantValues] from an unsafe.Pointer.
//
// A set of constant values that specialize a graphics or compute GPU function.
func FunctionConstantValuesFrom(ptr unsafe.Pointer) FunctionConstantValues {
	return FunctionConstantValues{objectivec.Object{objc.ID(ptr)}}
}




















// Deletes all previously set constant values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstantValues/reset()
func (f_ FunctionConstantValues) Reset() {
	objc.Send[objc.ID](f_.ID, objc.Sel("reset"))
}


// Sets a value for a function constant at a specific index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstantValues/setConstantValue(_:type:index:)
func (f_ FunctionConstantValues) SetConstantValueTypeAtIndex(value objectivec.IObject, type_ DataType, index uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setConstantValue:type:atIndex:"), value, type_, index)
}


// Sets a value for a function constant with a specific name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstantValues/setConstantValue(_:type:withName:)
func (f_ FunctionConstantValues) SetConstantValueTypeWithName(value objectivec.IObject, type_ DataType, name foundation.foundation.INSString) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setConstantValue:type:withName:"), value, type_, name)
}


// Sets values for a group of function constants within a specific index range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionConstantValues/setConstantValues:type:withRange:
func (f_ FunctionConstantValues) SetConstantValuesTypeWithRange(values objectivec.IObject, type_ DataType, range_ foundation.Range) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setConstantValues:type:withRange:"), values, type_, range_)
}













