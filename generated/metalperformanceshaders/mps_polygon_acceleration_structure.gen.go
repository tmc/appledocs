// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PolygonAccelerationStructure] class.
var (
	PolygonAccelerationStructureClass     _PolygonAccelerationStructureClass
	PolygonAccelerationStructureClassOnce sync.Once
)

func getPolygonAccelerationStructureClass() _PolygonAccelerationStructureClass {
	PolygonAccelerationStructureClassOnce.Do(func() {
		PolygonAccelerationStructureClass = _PolygonAccelerationStructureClass{objc.GetClass("MPSPolygonAccelerationStructure")}
	})
	return PolygonAccelerationStructureClass
}

type _PolygonAccelerationStructureClass struct {
	class objc.Class
}

// An interface definition for the [PolygonAccelerationStructure] class.
type IPolygonAccelerationStructure interface {
	IAccelerationStructure
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure
type PolygonAccelerationStructure struct {
	AccelerationStructure
}

// PolygonAccelerationStructureFrom constructs a [PolygonAccelerationStructure] from an unsafe.Pointer.
func PolygonAccelerationStructureFrom(ptr unsafe.Pointer) PolygonAccelerationStructure {
	return PolygonAccelerationStructure{
		AccelerationStructure: AccelerationStructureFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PolygonAccelerationStructureClass) Alloc() PolygonAccelerationStructure {
	rv := objc.Send[PolygonAccelerationStructure](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PolygonAccelerationStructureClass) New() PolygonAccelerationStructure {
	rv := objc.Send[PolygonAccelerationStructure](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PolygonAccelerationStructure) Init() PolygonAccelerationStructure {
	rv := objc.Send[PolygonAccelerationStructure](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PolygonAccelerationStructure) Autorelease() PolygonAccelerationStructure {
	rv := objc.Send[PolygonAccelerationStructure](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPolygonAccelerationStructure creates a new PolygonAccelerationStructure instance.
func NewPolygonAccelerationStructure() PolygonAccelerationStructure {
	return getPolygonAccelerationStructureClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/vertexbufferoffset
func (p_ PolygonAccelerationStructure) VertexBufferOffset() int {
	rv := objc.Send[int](p_.ID, objc.Sel("vertexBufferOffset"))
	return rv
}


// SetVertexBufferOffset sets the value of the vertexBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/vertexbufferoffset
func (p_ PolygonAccelerationStructure) SetVertexBufferOffset(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVertexBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/polygoncount
func (p_ PolygonAccelerationStructure) PolygonCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("polygonCount"))
	return rv
}


// SetPolygonCount sets the value of the polygonCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/polygoncount
func (p_ PolygonAccelerationStructure) SetPolygonCount(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPolygonCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/vertexbuffer
func (p_ PolygonAccelerationStructure) VertexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("vertexBuffer"))
	return rv
}


// SetVertexBuffer sets the value of the vertexBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/vertexbuffer
func (p_ PolygonAccelerationStructure) SetVertexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVertexBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/indextype
func (p_ PolygonAccelerationStructure) IndexType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("indexType"))
	return rv
}


// SetIndexType sets the value of the indexType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/indextype
func (p_ PolygonAccelerationStructure) SetIndexType(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndexType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/indexbuffer
func (p_ PolygonAccelerationStructure) IndexBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("indexBuffer"))
	return rv
}


// SetIndexBuffer sets the value of the indexBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/indexbuffer
func (p_ PolygonAccelerationStructure) SetIndexBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndexBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/polygontype
func (p_ PolygonAccelerationStructure) PolygonType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("polygonType"))
	return rv
}


// SetPolygonType sets the value of the polygonType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/polygontype
func (p_ PolygonAccelerationStructure) SetPolygonType(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPolygonType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/vertexstride
func (p_ PolygonAccelerationStructure) VertexStride() int {
	rv := objc.Send[int](p_.ID, objc.Sel("vertexStride"))
	return rv
}


// SetVertexStride sets the value of the vertexStride property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/vertexstride
func (p_ PolygonAccelerationStructure) SetVertexStride(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVertexStride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/maskbufferoffset
func (p_ PolygonAccelerationStructure) MaskBufferOffset() int {
	rv := objc.Send[int](p_.ID, objc.Sel("maskBufferOffset"))
	return rv
}


// SetMaskBufferOffset sets the value of the maskBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/maskbufferoffset
func (p_ PolygonAccelerationStructure) SetMaskBufferOffset(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaskBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/maskbuffer
func (p_ PolygonAccelerationStructure) MaskBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("maskBuffer"))
	return rv
}


// SetMaskBuffer sets the value of the maskBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/maskbuffer
func (p_ PolygonAccelerationStructure) SetMaskBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaskBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/indexbufferoffset
func (p_ PolygonAccelerationStructure) IndexBufferOffset() int {
	rv := objc.Send[int](p_.ID, objc.Sel("indexBufferOffset"))
	return rv
}


// SetIndexBufferOffset sets the value of the indexBufferOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure/indexbufferoffset
func (p_ PolygonAccelerationStructure) SetIndexBufferOffset(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIndexBufferOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure/polygonBuffers
func (p_ PolygonAccelerationStructure) PolygonBuffers() []PolygonBuffer {
	rv := objc.Send[[]PolygonBuffer](p_.ID, objc.Sel("polygonBuffers"))
	return rv
}


// SetPolygonBuffers sets the value of the polygonBuffers property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure/polygonBuffers
func (p_ PolygonAccelerationStructure) SetPolygonBuffers(value []PolygonBuffer) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setPolygonBuffers:"), nsArray)
}



