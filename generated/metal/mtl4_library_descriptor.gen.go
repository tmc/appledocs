// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MTL4LibraryDescriptor] class.
var (
	MTL4LibraryDescriptorClass     _MTL4LibraryDescriptorClass
	MTL4LibraryDescriptorClassOnce sync.Once
)

func getMTL4LibraryDescriptorClass() _MTL4LibraryDescriptorClass {
	MTL4LibraryDescriptorClassOnce.Do(func() {
		MTL4LibraryDescriptorClass = _MTL4LibraryDescriptorClass{objc.GetClass("MTL4LibraryDescriptor")}
	})
	return MTL4LibraryDescriptorClass
}

type _MTL4LibraryDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [MTL4LibraryDescriptor] class.
type IMTL4LibraryDescriptor interface {
	objectivec.IObject
	

	// properties:
	Name() foundation.foundation.INSString
	SetName(value foundation.foundation.INSString)
	Options() IMTLCompileOptions
	SetOptions(value IMTLCompileOptions)
	Source() foundation.foundation.INSString
	SetSource(value foundation.foundation.INSString)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MTL4LibraryDescriptorClass) Alloc() MTL4LibraryDescriptor {
	rv := objc.Send[MTL4LibraryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4LibraryDescriptorClass) New() MTL4LibraryDescriptor {
	rv := objc.Send[MTL4LibraryDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4LibraryDescriptor) Init() MTL4LibraryDescriptor {
	rv := objc.Send[MTL4LibraryDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4LibraryDescriptor) Autorelease() MTL4LibraryDescriptor {
	rv := objc.Send[MTL4LibraryDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4LibraryDescriptor creates a new MTL4LibraryDescriptor instance.
func NewMTL4LibraryDescriptor() MTL4LibraryDescriptor {
	return getMTL4LibraryDescriptorClass().New()
}





// Serves as the base descriptor for creating a Metal library.


// Serves as the base descriptor for creating a Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor
type MTL4LibraryDescriptor struct {
	objectivec.Object
}

// MTL4LibraryDescriptorFrom constructs a [MTL4LibraryDescriptor] from an unsafe.Pointer.
//
// Serves as the base descriptor for creating a Metal library.
func MTL4LibraryDescriptorFrom(ptr unsafe.Pointer) MTL4LibraryDescriptor {
	return MTL4LibraryDescriptor{objectivec.Object{objc.ID(ptr)}}
}

























// Assigns an optional name to the Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor/name
func (m_ MTL4LibraryDescriptor) Name() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}


// Assigns an optional name to the Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor/name
func (m_ MTL4LibraryDescriptor) SetName(value foundation.foundation.INSString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}


// Provides compile-time options for the Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor/options
func (m_ MTL4LibraryDescriptor) Options() IMTLCompileOptions {
	rv := objc.Send[CompileOptions](m_.ID, objc.Sel("options"))
	return rv
}


// Provides compile-time options for the Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor/options
func (m_ MTL4LibraryDescriptor) SetOptions(value IMTLCompileOptions) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptions:"), value)
}


// Assigns an optional string containing the source code of the shader language program to compile into a Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor/source
func (m_ MTL4LibraryDescriptor) Source() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("source"))
	return rv
}


// Assigns an optional string containing the source code of the shader language program to compile into a Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor/source
func (m_ MTL4LibraryDescriptor) SetSource(value foundation.foundation.INSString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSource:"), value)
}








