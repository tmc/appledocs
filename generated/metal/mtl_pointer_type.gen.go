// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [PointerType] class.
var (
	PointerTypeClass     _PointerTypeClass
	PointerTypeClassOnce sync.Once
)

func getPointerTypeClass() _PointerTypeClass {
	PointerTypeClassOnce.Do(func() {
		PointerTypeClass = _PointerTypeClass{objc.GetClass("MTLPointerType")}
	})
	return PointerTypeClass
}

type _PointerTypeClass struct {
	class objc.Class
}





// An interface definition for the [PointerType] class.
type IPointerType interface {
	IType
	

	// properties:
	Access() BindingAccess
	Alignment() uint
	DataSize() uint
	ElementIsArgumentBuffer() bool
	ElementType() DataType


	

	// methods:
	ElementArrayType() IArrayType
	ElementStructType() IStructType


}





// Alloc allocates a new instance without initialization.
func (pc _PointerTypeClass) Alloc() PointerType {
	rv := objc.Send[PointerType](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PointerTypeClass) New() PointerType {
	rv := objc.Send[PointerType](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PointerType) Init() PointerType {
	rv := objc.Send[PointerType](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PointerType) Autorelease() PointerType {
	rv := objc.Send[PointerType](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPointerType creates a new PointerType instance.
func NewPointerType() PointerType {
	return getPointerTypeClass().New()
}





// A description of a pointer.


// A description of a pointer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPointerType
type PointerType struct {
	Type
}

// PointerTypeFrom constructs a [PointerType] from an unsafe.Pointer.
//
// A description of a pointer.
func PointerTypeFrom(ptr unsafe.Pointer) PointerType {
	return PointerType{
		Type: TypeFrom(ptr),
	}
}




















// Provides a description of the underlying array when the pointer points to an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPointerType/elementArrayType()
func (p_ PointerType) ElementArrayType() IArrayType {
	rv := objc.Send[ArrayType](p_.ID, objc.Sel("elementArrayType"))
	return rv
}


// Provides a description of the underlying struct when the pointer points to a struct.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPointerType/elementStructType()
func (p_ PointerType) ElementStructType() IStructType {
	rv := objc.Send[StructType](p_.ID, objc.Sel("elementStructType"))
	return rv
}







// The function’s read/write access to the element data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPointerType/access
func (p_ PointerType) Access() BindingAccess {
	rv := objc.Send[BindingAccess](p_.ID, objc.Sel("access"))
	return rv
}


// The required byte alignment in memory for the element data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPointerType/alignment
func (p_ PointerType) Alignment() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("alignment"))
	return rv
}


// The size, in bytes, of the element data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPointerType/dataSize
func (p_ PointerType) DataSize() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("dataSize"))
	return rv
}


// A Boolean value that indicates whether the element is an argument buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPointerType/elementIsArgumentBuffer
func (p_ PointerType) ElementIsArgumentBuffer() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("elementIsArgumentBuffer"))
	return rv
}


// The data type of the element data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPointerType/elementType
func (p_ PointerType) ElementType() DataType {
	rv := objc.Send[DataType](p_.ID, objc.Sel("elementType"))
	return rv
}








