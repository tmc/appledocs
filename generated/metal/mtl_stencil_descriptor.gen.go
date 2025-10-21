// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StencilDescriptor] class.
var (
	StencilDescriptorClass     _StencilDescriptorClass
	StencilDescriptorClassOnce sync.Once
)

func getStencilDescriptorClass() _StencilDescriptorClass {
	StencilDescriptorClassOnce.Do(func() {
		StencilDescriptorClass = _StencilDescriptorClass{objc.GetClass("MTLStencilDescriptor")}
	})
	return StencilDescriptorClass
}

type _StencilDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [StencilDescriptor] class.
type IStencilDescriptor interface {
	objectivec.IObject
}

// An object that defines the front-facing or back-facing stencil operations of a depth and stencil state object.
//
// A stencil test is a comparison between a masked reference value and a masked value stored in a stencil attachment. (A value is by performing a logical AND operation on it with the value.) The object defines how to update the contents of the stencil attachment, based on the results of the stencil test and the depth test. The property defines the stencil test. The , , and properties specify what to do to a stencil value stored in the stencil attachment for three different test outcomes: if the stencil test fails, if the stencil test passes and the depth test fails, or if both stencil and depth tests succeed, respectively. determines which stencil bits can be modified as the result of a stencil operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor
type StencilDescriptor struct {
	objectivec.Object
}

// StencilDescriptorFrom constructs a [StencilDescriptor] from an unsafe.Pointer.
//
// An object that defines the front-facing or back-facing stencil operations of a depth and stencil state object.
func StencilDescriptorFrom(ptr unsafe.Pointer) StencilDescriptor {
	return StencilDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StencilDescriptorClass) Alloc() StencilDescriptor {
	rv := objc.Send[StencilDescriptor](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StencilDescriptorClass) New() StencilDescriptor {
	rv := objc.Send[StencilDescriptor](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StencilDescriptor) Init() StencilDescriptor {
	rv := objc.Send[StencilDescriptor](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StencilDescriptor) Autorelease() StencilDescriptor {
	rv := objc.Send[StencilDescriptor](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStencilDescriptor creates a new StencilDescriptor instance.
func NewStencilDescriptor() StencilDescriptor {
	return getStencilDescriptorClass().New()
}


// The operation that is performed to update the values in the stencil attachment when the stencil test passes, but the depth test fails.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor/depthFailureOperation
func (s_ StencilDescriptor) DepthFailureOperation() StencilOperation {
	rv := objc.Send[StencilOperation](s_.ID, objc.Sel("depthFailureOperation"))
	return rv
}


// SetDepthFailureOperation sets the value of the depthFailureOperation property.
// The operation that is performed to update the values in the stencil attachment when the stencil test passes, but the depth test fails.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor/depthFailureOperation
func (s_ StencilDescriptor) SetDepthFailureOperation(value IStencilOperation) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDepthFailureOperation:"), value)
}

// The operation that is performed to update the values in the stencil attachment when both the stencil test and the depth test pass.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor/depthStencilPassOperation
func (s_ StencilDescriptor) DepthStencilPassOperation() StencilOperation {
	rv := objc.Send[StencilOperation](s_.ID, objc.Sel("depthStencilPassOperation"))
	return rv
}


// SetDepthStencilPassOperation sets the value of the depthStencilPassOperation property.
// The operation that is performed to update the values in the stencil attachment when both the stencil test and the depth test pass.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor/depthStencilPassOperation
func (s_ StencilDescriptor) SetDepthStencilPassOperation(value IStencilOperation) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDepthStencilPassOperation:"), value)
}

// A bitmask that determines from which bits that stencil comparison tests can read.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor/readMask
func (s_ StencilDescriptor) ReadMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("readMask"))
	return rv
}


// SetReadMask sets the value of the readMask property.
// A bitmask that determines from which bits that stencil comparison tests can read.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor/readMask
func (s_ StencilDescriptor) SetReadMask(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setReadMask:"), value)
}

// The comparison that is performed between the masked reference value and a masked value in the stencil attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor/stencilCompareFunction
func (s_ StencilDescriptor) StencilCompareFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stencilCompareFunction"))
	return rv
}


// SetStencilCompareFunction sets the value of the stencilCompareFunction property.
// The comparison that is performed between the masked reference value and a masked value in the stencil attachment.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor/stencilCompareFunction
func (s_ StencilDescriptor) SetStencilCompareFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStencilCompareFunction:"), value)
}

// The operation that is performed to update the values in the stencil attachment when the stencil test fails.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor/stencilFailureOperation
func (s_ StencilDescriptor) StencilFailureOperation() StencilOperation {
	rv := objc.Send[StencilOperation](s_.ID, objc.Sel("stencilFailureOperation"))
	return rv
}


// SetStencilFailureOperation sets the value of the stencilFailureOperation property.
// The operation that is performed to update the values in the stencil attachment when the stencil test fails.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor/stencilFailureOperation
func (s_ StencilDescriptor) SetStencilFailureOperation(value IStencilOperation) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStencilFailureOperation:"), value)
}

// A bitmask that determines to which bits that stencil operations can write.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor/writeMask
func (s_ StencilDescriptor) WriteMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("writeMask"))
	return rv
}


// SetWriteMask sets the value of the writeMask property.
// A bitmask that determines to which bits that stencil operations can write.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilDescriptor/writeMask
func (s_ StencilDescriptor) SetWriteMask(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setWriteMask:"), value)
}



