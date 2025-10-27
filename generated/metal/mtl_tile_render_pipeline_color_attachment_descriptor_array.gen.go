// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [TileRenderPipelineColorAttachmentDescriptorArray] class.
var (
	TileRenderPipelineColorAttachmentDescriptorArrayClass     _TileRenderPipelineColorAttachmentDescriptorArrayClass
	TileRenderPipelineColorAttachmentDescriptorArrayClassOnce sync.Once
)

func getTileRenderPipelineColorAttachmentDescriptorArrayClass() _TileRenderPipelineColorAttachmentDescriptorArrayClass {
	TileRenderPipelineColorAttachmentDescriptorArrayClassOnce.Do(func() {
		TileRenderPipelineColorAttachmentDescriptorArrayClass = _TileRenderPipelineColorAttachmentDescriptorArrayClass{objc.GetClass("MTLTileRenderPipelineColorAttachmentDescriptorArray")}
	})
	return TileRenderPipelineColorAttachmentDescriptorArrayClass
}

type _TileRenderPipelineColorAttachmentDescriptorArrayClass struct {
	class objc.Class
}





// An interface definition for the [TileRenderPipelineColorAttachmentDescriptorArray] class.
type ITileRenderPipelineColorAttachmentDescriptorArray interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	SetObjectAtIndexedSubscript(attachment IMTLTileRenderPipelineColorAttachmentDescriptor, attachmentIndex uint)
	ObjectAtIndexedSubscript(attachmentIndex uint) ITileRenderPipelineColorAttachmentDescriptor


}





// Alloc allocates a new instance without initialization.
func (tc _TileRenderPipelineColorAttachmentDescriptorArrayClass) Alloc() TileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[TileRenderPipelineColorAttachmentDescriptorArray](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TileRenderPipelineColorAttachmentDescriptorArrayClass) New() TileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[TileRenderPipelineColorAttachmentDescriptorArray](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TileRenderPipelineColorAttachmentDescriptorArray) Init() TileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[TileRenderPipelineColorAttachmentDescriptorArray](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TileRenderPipelineColorAttachmentDescriptorArray) Autorelease() TileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[TileRenderPipelineColorAttachmentDescriptorArray](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTileRenderPipelineColorAttachmentDescriptorArray creates a new TileRenderPipelineColorAttachmentDescriptorArray instance.
func NewTileRenderPipelineColorAttachmentDescriptorArray() TileRenderPipelineColorAttachmentDescriptorArray {
	return getTileRenderPipelineColorAttachmentDescriptorArrayClass().New()
}





// An array of color attachment descriptors for the tile render pipeline.


// An array of color attachment descriptors for the tile render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineColorAttachmentDescriptorArray
type TileRenderPipelineColorAttachmentDescriptorArray struct {
	objectivec.Object
}

// TileRenderPipelineColorAttachmentDescriptorArrayFrom constructs a [TileRenderPipelineColorAttachmentDescriptorArray] from an unsafe.Pointer.
//
// An array of color attachment descriptors for the tile render pipeline.
func TileRenderPipelineColorAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) TileRenderPipelineColorAttachmentDescriptorArray {
	return TileRenderPipelineColorAttachmentDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}




















// Sets the render pipeline state for a specified color attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineColorAttachmentDescriptorArray/setObject:atIndexedSubscript:
func (t_ TileRenderPipelineColorAttachmentDescriptorArray) SetObjectAtIndexedSubscript(attachment IMTLTileRenderPipelineColorAttachmentDescriptor, attachmentIndex uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setObject:atIndexedSubscript:"), attachment, attachmentIndex)
}


// Returns the render pipeline state for the specified color attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineColorAttachmentDescriptorArray/subscript(_:)
func (t_ TileRenderPipelineColorAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) ITileRenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[TileRenderPipelineColorAttachmentDescriptor](t_.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}













