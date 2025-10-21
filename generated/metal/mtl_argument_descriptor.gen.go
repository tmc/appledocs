// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ArgumentDescriptor] class.
var (
	ArgumentDescriptorClass     _ArgumentDescriptorClass
	ArgumentDescriptorClassOnce sync.Once
)

func getArgumentDescriptorClass() _ArgumentDescriptorClass {
	ArgumentDescriptorClassOnce.Do(func() {
		ArgumentDescriptorClass = _ArgumentDescriptorClass{objc.GetClass("MTLArgumentDescriptor")}
	})
	return ArgumentDescriptorClass
}

type _ArgumentDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [ArgumentDescriptor] class.
type IArgumentDescriptor interface {
	objectivec.IObject
}

// A representation of an argument within an argument buffer.
//
// This descriptor can represent arguments within flat structures only. It can represent arrays of allowed argument buffer data types, but it cannot represent arguments within nested structures. Argument buffers with simple, flat structures can be represented by an array of instances. You can then use this array to create an instance by calling the method. Argument buffers with complex, nested structures must define their structure in Metal shading language code, which can then be directly assigned to a specific buffer index of a function. You can then use this buffer index to call the method and create an instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor
type ArgumentDescriptor struct {
	objectivec.Object
}

// ArgumentDescriptorFrom constructs a [ArgumentDescriptor] from an unsafe.Pointer.
//
// A representation of an argument within an argument buffer.
func ArgumentDescriptorFrom(ptr unsafe.Pointer) ArgumentDescriptor {
	return ArgumentDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _ArgumentDescriptorClass) Alloc() ArgumentDescriptor {
	rv := objc.Send[ArgumentDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ArgumentDescriptorClass) New() ArgumentDescriptor {
	rv := objc.Send[ArgumentDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ArgumentDescriptor) Init() ArgumentDescriptor {
	rv := objc.Send[ArgumentDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ArgumentDescriptor) Autorelease() ArgumentDescriptor {
	rv := objc.Send[ArgumentDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewArgumentDescriptor creates a new ArgumentDescriptor instance.
func NewArgumentDescriptor() ArgumentDescriptor {
	return getArgumentDescriptorClass().New()
}


// Creates an empty argument descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/argumentDescriptor
func (ac _ArgumentDescriptorClass) ArgumentDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("argumentDescriptor"))
	return rv
}

// The length of an array argument.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/arrayLength
func (a_ ArgumentDescriptor) ArrayLength() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("arrayLength"))
	return rv
}


// SetArrayLength sets the value of the arrayLength property.
// The length of an array argument.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/arrayLength
func (a_ ArgumentDescriptor) SetArrayLength(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setArrayLength:"), value)
}



