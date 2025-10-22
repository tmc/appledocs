// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PolygonBuffer] class.
var (
	PolygonBufferClass     _PolygonBufferClass
	PolygonBufferClassOnce sync.Once
)

func getPolygonBufferClass() _PolygonBufferClass {
	PolygonBufferClassOnce.Do(func() {
		PolygonBufferClass = _PolygonBufferClass{objc.GetClass("MPSPolygonBuffer")}
	})
	return PolygonBufferClass
}

type _PolygonBufferClass struct {
	class objc.Class
}

// An interface definition for the [PolygonBuffer] class.
type IPolygonBuffer interface {
	objectivec.IObject
	IndexBufferOffset() uint
	SetIndexBufferOffset(value uint)
	IndexBuffer() unsafe.Pointer
	SetIndexBuffer(value unsafe.Pointer)
	MaskBuffer() unsafe.Pointer
	SetMaskBuffer(value unsafe.Pointer)
	MaskBufferOffset() int
	SetMaskBufferOffset(value int)
	PolygonCount() int
	SetPolygonCount(value int)
	VertexBuffer() unsafe.Pointer
	SetVertexBuffer(value unsafe.Pointer)
	VertexBufferOffset() int
	SetVertexBufferOffset(value int)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonBuffer
type PolygonBuffer struct {
	objectivec.Object
}

// PolygonBufferFrom constructs a [PolygonBuffer] from an unsafe.Pointer.
func PolygonBufferFrom(ptr unsafe.Pointer) PolygonBuffer {
	return PolygonBuffer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PolygonBufferClass) Alloc() PolygonBuffer {
	rv := objc.Send[PolygonBuffer](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PolygonBufferClass) New() PolygonBuffer {
	rv := objc.Send[PolygonBuffer](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PolygonBuffer) Init() PolygonBuffer {
	rv := objc.Send[PolygonBuffer](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PolygonBuffer) Autorelease() PolygonBuffer {
	rv := objc.Send[PolygonBuffer](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPolygonBuffer creates a new PolygonBuffer instance.
func NewPolygonBuffer() PolygonBuffer {
	return getPolygonBufferClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonBuffer/indexBufferOffset
func (p_ PolygonBuffer) IndexBufferOffset() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("indexBufferOffset"))
	return rv
}


// SetIndexBufferOffset sets the value of the indexBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonBuffer/indexBufferOffset
func (p_ PolygonBuffer) SetIndexBufferOffset(value uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndexBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/indexbuffer
func (p_ PolygonBuffer) IndexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("indexBuffer"))
	return rv
}


// SetIndexBuffer sets the value of the indexBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/indexbuffer
func (p_ PolygonBuffer) SetIndexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndexBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/maskbuffer
func (p_ PolygonBuffer) MaskBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("maskBuffer"))
	return rv
}


// SetMaskBuffer sets the value of the maskBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/maskbuffer
func (p_ PolygonBuffer) SetMaskBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaskBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/maskbufferoffset
func (p_ PolygonBuffer) MaskBufferOffset() int {
	rv := objc.Send[int](p_.ID, objc.Sel("maskBufferOffset"))
	return rv
}


// SetMaskBufferOffset sets the value of the maskBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/maskbufferoffset
func (p_ PolygonBuffer) SetMaskBufferOffset(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaskBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/polygoncount
func (p_ PolygonBuffer) PolygonCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("polygonCount"))
	return rv
}


// SetPolygonCount sets the value of the polygonCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/polygoncount
func (p_ PolygonBuffer) SetPolygonCount(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPolygonCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/vertexbuffer
func (p_ PolygonBuffer) VertexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("vertexBuffer"))
	return rv
}


// SetVertexBuffer sets the value of the vertexBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/vertexbuffer
func (p_ PolygonBuffer) SetVertexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVertexBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/vertexbufferoffset
func (p_ PolygonBuffer) VertexBufferOffset() int {
	rv := objc.Send[int](p_.ID, objc.Sel("vertexBufferOffset"))
	return rv
}


// SetVertexBufferOffset sets the value of the vertexBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/vertexbufferoffset
func (p_ PolygonBuffer) SetVertexBufferOffset(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVertexBufferOffset:"), value)
}



