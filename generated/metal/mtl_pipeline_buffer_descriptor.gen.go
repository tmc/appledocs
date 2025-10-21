// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PipelineBufferDescriptor] class.
var (
	PipelineBufferDescriptorClass     _PipelineBufferDescriptorClass
	PipelineBufferDescriptorClassOnce sync.Once
)

func getPipelineBufferDescriptorClass() _PipelineBufferDescriptorClass {
	PipelineBufferDescriptorClassOnce.Do(func() {
		PipelineBufferDescriptorClass = _PipelineBufferDescriptorClass{objc.GetClass("MTLPipelineBufferDescriptor")}
	})
	return PipelineBufferDescriptorClass
}

type _PipelineBufferDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [PipelineBufferDescriptor] class.
type IPipelineBufferDescriptor interface {
	objectivec.IObject
}

// The mutability options for a buffer that a render or compute pipeline uses.
//
// Metal can perform additional optimizations if you guarantee that neither the CPU nor the GPU modify a buffer’s contents before starting a pass. Use immutable buffers as much as possible to take advantage of Metal optimizations. To declare that a buffer is immutable, set the property of their associated object to .
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineBufferDescriptor
type PipelineBufferDescriptor struct {
	objectivec.Object
}

// PipelineBufferDescriptorFrom constructs a [PipelineBufferDescriptor] from an unsafe.Pointer.
//
// The mutability options for a buffer that a render or compute pipeline uses.
func PipelineBufferDescriptorFrom(ptr unsafe.Pointer) PipelineBufferDescriptor {
	return PipelineBufferDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PipelineBufferDescriptorClass) Alloc() PipelineBufferDescriptor {
	rv := objc.Send[PipelineBufferDescriptor](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PipelineBufferDescriptorClass) New() PipelineBufferDescriptor {
	rv := objc.Send[PipelineBufferDescriptor](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PipelineBufferDescriptor) Init() PipelineBufferDescriptor {
	rv := objc.Send[PipelineBufferDescriptor](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PipelineBufferDescriptor) Autorelease() PipelineBufferDescriptor {
	rv := objc.Send[PipelineBufferDescriptor](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPipelineBufferDescriptor creates a new PipelineBufferDescriptor instance.
func NewPipelineBufferDescriptor() PipelineBufferDescriptor {
	return getPipelineBufferDescriptorClass().New()
}


// A mutability option that determines whether you can update a buffer’s contents before related commands use the buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineBufferDescriptor/mutability
func (p_ PipelineBufferDescriptor) Mutability() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("mutability"))
	return rv
}


// SetMutability sets the value of the mutability property.
// A mutability option that determines whether you can update a buffer’s contents before related commands use the buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineBufferDescriptor/mutability
func (p_ PipelineBufferDescriptor) SetMutability(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMutability:"), value)
}



