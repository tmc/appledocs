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
	

	// properties:
	IndexBuffer() Buffer get set /* not a class type */
	SetIndexBuffer(value Buffer get set /* not a class type */)
	IndexBufferOffset() objectivec.IObject
	SetIndexBufferOffset(value objectivec.IObject)
	MaskBuffer() Buffer get set /* not a class type */
	SetMaskBuffer(value Buffer get set /* not a class type */)
	MaskBufferOffset() objectivec.IObject
	SetMaskBufferOffset(value objectivec.IObject)
	PolygonCount() objectivec.IObject
	SetPolygonCount(value objectivec.IObject)
	VertexBuffer() Buffer get set /* not a class type */
	SetVertexBuffer(value Buffer get set /* not a class type */)
	VertexBufferOffset() objectivec.IObject
	SetVertexBufferOffset(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PolygonBufferClass) Alloc() PolygonBuffer {
	rv := objc.Send[PolygonBuffer](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonBuffer
type PolygonBuffer struct {
	objectivec.Object
}

// PolygonBufferFrom constructs a [PolygonBuffer] from an unsafe.Pointer.
func PolygonBufferFrom(ptr unsafe.Pointer) PolygonBuffer {
	return PolygonBuffer{objectivec.Object{objc.ID(ptr)}}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/3152583-initwithcoder
func NewPolygonBufferWithCoder(aDecoder foundation.Coder) PolygonBuffer {
	instance := getPolygonBufferClass().Alloc()
	rv := objc.Send[PolygonBuffer](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/3152586-polygonbuffer
func (pc _PolygonBufferClass) PolygonBuffer() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("polygonBuffer"))
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/3152580-indexbuffer
func (p_ PolygonBuffer) IndexBuffer() Buffer get set /* not a class type */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("indexBuffer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/3152580-indexbuffer
func (p_ PolygonBuffer) SetIndexBuffer(value Buffer get set /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndexBuffer:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/3152581-indexbufferoffset
func (p_ PolygonBuffer) IndexBufferOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("indexBufferOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/3152581-indexbufferoffset
func (p_ PolygonBuffer) SetIndexBufferOffset(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndexBufferOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/3152584-maskbuffer
func (p_ PolygonBuffer) MaskBuffer() Buffer get set /* not a class type */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("maskBuffer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/3152584-maskbuffer
func (p_ PolygonBuffer) SetMaskBuffer(value Buffer get set /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaskBuffer:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/3152585-maskbufferoffset
func (p_ PolygonBuffer) MaskBufferOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("maskBufferOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/3152585-maskbufferoffset
func (p_ PolygonBuffer) SetMaskBufferOffset(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaskBufferOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/3152587-polygoncount
func (p_ PolygonBuffer) PolygonCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("polygonCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/3152587-polygoncount
func (p_ PolygonBuffer) SetPolygonCount(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPolygonCount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/3152588-vertexbuffer
func (p_ PolygonBuffer) VertexBuffer() Buffer get set /* not a class type */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("vertexBuffer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/3152588-vertexbuffer
func (p_ PolygonBuffer) SetVertexBuffer(value Buffer get set /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVertexBuffer:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/3152589-vertexbufferoffset
func (p_ PolygonBuffer) VertexBufferOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("vertexBufferOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonbuffer/3152589-vertexbufferoffset
func (p_ PolygonBuffer) SetVertexBufferOffset(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVertexBufferOffset:"), value)
}







