// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [DepthStencilDescriptor] class.
var (
	DepthStencilDescriptorClass     _DepthStencilDescriptorClass
	DepthStencilDescriptorClassOnce sync.Once
)

func getDepthStencilDescriptorClass() _DepthStencilDescriptorClass {
	DepthStencilDescriptorClassOnce.Do(func() {
		DepthStencilDescriptorClass = _DepthStencilDescriptorClass{objc.GetClass("MTLDepthStencilDescriptor")}
	})
	return DepthStencilDescriptorClass
}

type _DepthStencilDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [DepthStencilDescriptor] class.
type IDepthStencilDescriptor interface {
	objectivec.IObject
}

// An instance that configures new instances.
//
// An instance is used to define a specific configuration of the depth and stencil stages of a rendering pipeline. To create an instance, use standard allocation and initialization techniques. To enable writing the depth value to a depth attachment, set the depthWriteEnabled property to . The depthCompareFunction property specifies how the depth test is performed. If a fragment’s depth value fails the depth test, the fragment is discarded. is a commonly used value for , because fragment values that are farther away from the viewer than the pixel depth value (a previously written fragment) fail the depth test and are considered occluded by the earlier depth value. The and properties define two independent stencil descriptors: one for front-facing primitives and the other for back-facing primitives, respectively. Both properties can be set to the same MTLStencilDescriptor instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDepthStencilDescriptor
type DepthStencilDescriptor struct {
	objectivec.Object
}

// DepthStencilDescriptorFrom constructs a [DepthStencilDescriptor] from an unsafe.Pointer.
//
// An instance that configures new instances.
func DepthStencilDescriptorFrom(ptr unsafe.Pointer) DepthStencilDescriptor {
	return DepthStencilDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DepthStencilDescriptorClass) Alloc() DepthStencilDescriptor {
	rv := objc.Send[DepthStencilDescriptor](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DepthStencilDescriptorClass) New() DepthStencilDescriptor {
	rv := objc.Send[DepthStencilDescriptor](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DepthStencilDescriptor) Init() DepthStencilDescriptor {
	rv := objc.Send[DepthStencilDescriptor](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DepthStencilDescriptor) Autorelease() DepthStencilDescriptor {
	rv := objc.Send[DepthStencilDescriptor](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDepthStencilDescriptor creates a new DepthStencilDescriptor instance.
func NewDepthStencilDescriptor() DepthStencilDescriptor {
	return getDepthStencilDescriptorClass().New()
}


// The stencil descriptor for back-facing primitives.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDepthStencilDescriptor/backFaceStencil
func (d_ DepthStencilDescriptor) BackFaceStencil() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("backFaceStencil"))
	return rv
}


// SetBackFaceStencil sets the value of the backFaceStencil property.
// The stencil descriptor for back-facing primitives.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDepthStencilDescriptor/backFaceStencil
func (d_ DepthStencilDescriptor) SetBackFaceStencil(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBackFaceStencil:"), value)
}

// The comparison that is performed between a fragment’s depth value and the depth value in the attachment, which determines whether to discard the fragment.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDepthStencilDescriptor/depthCompareFunction
func (d_ DepthStencilDescriptor) DepthCompareFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("depthCompareFunction"))
	return rv
}


// SetDepthCompareFunction sets the value of the depthCompareFunction property.
// The comparison that is performed between a fragment’s depth value and the depth value in the attachment, which determines whether to discard the fragment.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDepthStencilDescriptor/depthCompareFunction
func (d_ DepthStencilDescriptor) SetDepthCompareFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDepthCompareFunction:"), value)
}

// The stencil descriptor for front-facing primitives.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDepthStencilDescriptor/frontFaceStencil
func (d_ DepthStencilDescriptor) FrontFaceStencil() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("frontFaceStencil"))
	return rv
}


// SetFrontFaceStencil sets the value of the frontFaceStencil property.
// The stencil descriptor for front-facing primitives.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDepthStencilDescriptor/frontFaceStencil
func (d_ DepthStencilDescriptor) SetFrontFaceStencil(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFrontFaceStencil:"), value)
}

// A Boolean value that indicates whether depth values can be written to the depth attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDepthStencilDescriptor/isDepthWriteEnabled
func (d_ DepthStencilDescriptor) DepthWriteEnabled() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("depthWriteEnabled"))
	return rv
}


// SetDepthWriteEnabled sets the value of the depthWriteEnabled property.
// A Boolean value that indicates whether depth values can be written to the depth attachment.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDepthStencilDescriptor/isDepthWriteEnabled
func (d_ DepthStencilDescriptor) SetDepthWriteEnabled(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDepthWriteEnabled:"), value)
}

// A string that identifies this object.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDepthStencilDescriptor/label
func (d_ DepthStencilDescriptor) Label() string {
	rv := objc.Send[string](d_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// A string that identifies this object.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDepthStencilDescriptor/label
func (d_ DepthStencilDescriptor) SetLabel(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLabel:"), objc.String(value))
}



