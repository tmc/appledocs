// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BlitPassDescriptor] class.
var (
	BlitPassDescriptorClass     _BlitPassDescriptorClass
	BlitPassDescriptorClassOnce sync.Once
)

func getBlitPassDescriptorClass() _BlitPassDescriptorClass {
	BlitPassDescriptorClassOnce.Do(func() {
		BlitPassDescriptorClass = _BlitPassDescriptorClass{objc.GetClass("MTLBlitPassDescriptor")}
	})
	return BlitPassDescriptorClass
}

type _BlitPassDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [BlitPassDescriptor] class.
type IBlitPassDescriptor interface {
	objectivec.IObject
}

// A configuration you create to customize a blit command encoder, which affects the runtime behavior of the blit pass you encode with it.
//
// You can customize an encoder for a blit pass by creating and configuring an instance and passing it to .
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitPassDescriptor
type BlitPassDescriptor struct {
	objectivec.Object
}

// BlitPassDescriptorFrom constructs a [BlitPassDescriptor] from an unsafe.Pointer.
//
// A configuration you create to customize a blit command encoder, which affects the runtime behavior of the blit pass you encode with it.
func BlitPassDescriptorFrom(ptr unsafe.Pointer) BlitPassDescriptor {
	return BlitPassDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BlitPassDescriptorClass) Alloc() BlitPassDescriptor {
	rv := objc.Send[BlitPassDescriptor](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BlitPassDescriptorClass) New() BlitPassDescriptor {
	rv := objc.Send[BlitPassDescriptor](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BlitPassDescriptor) Init() BlitPassDescriptor {
	rv := objc.Send[BlitPassDescriptor](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BlitPassDescriptor) Autorelease() BlitPassDescriptor {
	rv := objc.Send[BlitPassDescriptor](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBlitPassDescriptor creates a new BlitPassDescriptor instance.
func NewBlitPassDescriptor() BlitPassDescriptor {
	return getBlitPassDescriptorClass().New()
}


// An array of counter sample buffer attachments that you configure for a blit pass.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitPassDescriptor/sampleBufferAttachments
func (b_ BlitPassDescriptor) SampleBufferAttachments() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("sampleBufferAttachments"))
	return rv
}



