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



