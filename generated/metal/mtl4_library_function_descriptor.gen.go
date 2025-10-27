// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MTL4LibraryFunctionDescriptor] class.
var (
	MTL4LibraryFunctionDescriptorClass     _MTL4LibraryFunctionDescriptorClass
	MTL4LibraryFunctionDescriptorClassOnce sync.Once
)

func getMTL4LibraryFunctionDescriptorClass() _MTL4LibraryFunctionDescriptorClass {
	MTL4LibraryFunctionDescriptorClassOnce.Do(func() {
		MTL4LibraryFunctionDescriptorClass = _MTL4LibraryFunctionDescriptorClass{objc.GetClass("MTL4LibraryFunctionDescriptor")}
	})
	return MTL4LibraryFunctionDescriptorClass
}

type _MTL4LibraryFunctionDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [MTL4LibraryFunctionDescriptor] class.
type IMTL4LibraryFunctionDescriptor interface {
	IMTL4FunctionDescriptor
	

	// properties:
	Library() unsafe.Pointer
	SetLibrary(value unsafe.Pointer)
	Name() foundation.foundation.INSString
	SetName(value foundation.foundation.INSString)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MTL4LibraryFunctionDescriptorClass) Alloc() MTL4LibraryFunctionDescriptor {
	rv := objc.Send[MTL4LibraryFunctionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4LibraryFunctionDescriptorClass) New() MTL4LibraryFunctionDescriptor {
	rv := objc.Send[MTL4LibraryFunctionDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4LibraryFunctionDescriptor) Init() MTL4LibraryFunctionDescriptor {
	rv := objc.Send[MTL4LibraryFunctionDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4LibraryFunctionDescriptor) Autorelease() MTL4LibraryFunctionDescriptor {
	rv := objc.Send[MTL4LibraryFunctionDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4LibraryFunctionDescriptor creates a new MTL4LibraryFunctionDescriptor instance.
func NewMTL4LibraryFunctionDescriptor() MTL4LibraryFunctionDescriptor {
	return getMTL4LibraryFunctionDescriptorClass().New()
}





// Describes a shader function from a Metal library.


// Describes a shader function from a Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryFunctionDescriptor
type MTL4LibraryFunctionDescriptor struct {
	MTL4FunctionDescriptor
}

// MTL4LibraryFunctionDescriptorFrom constructs a [MTL4LibraryFunctionDescriptor] from an unsafe.Pointer.
//
// Describes a shader function from a Metal library.
func MTL4LibraryFunctionDescriptorFrom(ptr unsafe.Pointer) MTL4LibraryFunctionDescriptor {
	return MTL4LibraryFunctionDescriptor{
		MTL4FunctionDescriptor: MTL4FunctionDescriptorFrom(ptr),
	}
}

























// Returns a reference to the library containing the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryFunctionDescriptor/library
func (m_ MTL4LibraryFunctionDescriptor) Library() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("library"))
	return rv
}


// Returns a reference to the library containing the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryFunctionDescriptor/library
func (m_ MTL4LibraryFunctionDescriptor) SetLibrary(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLibrary:"), value)
}


// Assigns a name to the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryFunctionDescriptor/name
func (m_ MTL4LibraryFunctionDescriptor) Name() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}


// Assigns a name to the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryFunctionDescriptor/name
func (m_ MTL4LibraryFunctionDescriptor) SetName(value foundation.foundation.INSString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}








