// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [PipelineBufferDescriptorArray] class.
var (
	PipelineBufferDescriptorArrayClass     _PipelineBufferDescriptorArrayClass
	PipelineBufferDescriptorArrayClassOnce sync.Once
)

func getPipelineBufferDescriptorArrayClass() _PipelineBufferDescriptorArrayClass {
	PipelineBufferDescriptorArrayClassOnce.Do(func() {
		PipelineBufferDescriptorArrayClass = _PipelineBufferDescriptorArrayClass{objc.GetClass("MTLPipelineBufferDescriptorArray")}
	})
	return PipelineBufferDescriptorArrayClass
}

type _PipelineBufferDescriptorArrayClass struct {
	class objc.Class
}





// An interface definition for the [PipelineBufferDescriptorArray] class.
type IPipelineBufferDescriptorArray interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	SetObjectAtIndexedSubscript(buffer IMTLPipelineBufferDescriptor, bufferIndex uint)
	ObjectAtIndexedSubscript(bufferIndex uint) IPipelineBufferDescriptor


}





// Alloc allocates a new instance without initialization.
func (pc _PipelineBufferDescriptorArrayClass) Alloc() PipelineBufferDescriptorArray {
	rv := objc.Send[PipelineBufferDescriptorArray](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PipelineBufferDescriptorArrayClass) New() PipelineBufferDescriptorArray {
	rv := objc.Send[PipelineBufferDescriptorArray](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PipelineBufferDescriptorArray) Init() PipelineBufferDescriptorArray {
	rv := objc.Send[PipelineBufferDescriptorArray](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PipelineBufferDescriptorArray) Autorelease() PipelineBufferDescriptorArray {
	rv := objc.Send[PipelineBufferDescriptorArray](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPipelineBufferDescriptorArray creates a new PipelineBufferDescriptorArray instance.
func NewPipelineBufferDescriptorArray() PipelineBufferDescriptorArray {
	return getPipelineBufferDescriptorArrayClass().New()
}





// An array of pipeline buffer descriptors.


// An array of pipeline buffer descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineBufferDescriptorArray
type PipelineBufferDescriptorArray struct {
	objectivec.Object
}

// PipelineBufferDescriptorArrayFrom constructs a [PipelineBufferDescriptorArray] from an unsafe.Pointer.
//
// An array of pipeline buffer descriptors.
func PipelineBufferDescriptorArrayFrom(ptr unsafe.Pointer) PipelineBufferDescriptorArray {
	return PipelineBufferDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}




















// Sets a pipeline buffer descriptor at the specified array index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineBufferDescriptorArray/setObject:atIndexedSubscript:
func (p_ PipelineBufferDescriptorArray) SetObjectAtIndexedSubscript(buffer IMTLPipelineBufferDescriptor, bufferIndex uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setObject:atIndexedSubscript:"), buffer, bufferIndex)
}


// Returns the pipeline buffer descriptor at the specified array index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineBufferDescriptorArray/subscript(_:)
func (p_ PipelineBufferDescriptorArray) ObjectAtIndexedSubscript(bufferIndex uint) IPipelineBufferDescriptor {
	rv := objc.Send[PipelineBufferDescriptor](p_.ID, objc.Sel("objectAtIndexedSubscript:"), bufferIndex)
	return rv
}













