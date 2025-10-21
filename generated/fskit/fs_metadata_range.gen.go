// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [FSMetadataRange] class.
var (
	FSMetadataRangeClass     _FSMetadataRangeClass
	FSMetadataRangeClassOnce sync.Once
)

func getFSMetadataRangeClass() _FSMetadataRangeClass {
	FSMetadataRangeClassOnce.Do(func() {
		FSMetadataRangeClass = _FSMetadataRangeClass{objc.GetClass("FSMetadataRange")}
	})
	return FSMetadataRangeClass
}

type _FSMetadataRangeClass struct {
	class objc.Class
}

// An interface definition for the [FSMetadataRange] class.
type IFSMetadataRange interface {
	objectivec.IObject
}

// A range that describes contiguous metadata segments on disk.
//
// This type represents a range that begins at and ends at . Each segment in the range represents a single block in the resource’s buffer cache. For example, given an with the following properties: The range represents eight segments: from 0 to 511, then from 512 to 1023, and so on until a final segment of 3584 to 4095. Ensure that each metadata segment represents a range that’s already present in the resource’s buffer cache. Similarly, ensure that each segment’s offset and length matches the offset and length of the corresponding block in the buffer cache.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSMetadataRange
type FSMetadataRange struct {
	objectivec.Object
}

// FSMetadataRangeFrom constructs a [FSMetadataRange] from an unsafe.Pointer.
//
// A range that describes contiguous metadata segments on disk.
func FSMetadataRangeFrom(ptr unsafe.Pointer) FSMetadataRange {
	return FSMetadataRange{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FSMetadataRangeClass) Alloc() FSMetadataRange {
	rv := objc.Send[FSMetadataRange](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSMetadataRangeClass) New() FSMetadataRange {
	rv := objc.Send[FSMetadataRange](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSMetadataRange) Init() FSMetadataRange {
	rv := objc.Send[FSMetadataRange](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSMetadataRange) Autorelease() FSMetadataRange {
	rv := objc.Send[FSMetadataRange](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSMetadataRange creates a new FSMetadataRange instance.
func NewFSMetadataRange() FSMetadataRange {
	return getFSMetadataRangeClass().New()
}




// Initializes a metadata range with the given properties.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSMetadataRange/init(offset:segmentLength:segmentCount:)
func NewFSMetadataRangeWithOffsetSegmentLengthSegmentCount(startOffset unsafe.Pointer, segmentLength uint64, segmentCount uint64) FSMetadataRange {
	instance := getFSMetadataRangeClass().Alloc()
	rv := objc.Send[FSMetadataRange](instance.ID, objc.Sel("initWithOffset:segmentLength:segmentCount:"), startOffset, segmentLength, segmentCount)
	rv.Autorelease()
	return rv
}


// Creates a metadata range with the given properties.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSMetadataRange/rangeWithOffset:segmentLength:segmentCount:
func (fc _FSMetadataRangeClass) RangeWithOffsetSegmentLengthSegmentCount(startOffset unsafe.Pointer, segmentLength uint64, segmentCount uint64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("rangeWithOffset:segmentLength:segmentCount:"), startOffset, segmentLength, segmentCount)
	return rv
}

// The number of segments in the range.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSMetadataRange/segmentCount
func (f_ FSMetadataRange) SegmentCount() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("segmentCount"))
	return rv
}

// The segment length in bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSMetadataRange/segmentLength
func (f_ FSMetadataRange) SegmentLength() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("segmentLength"))
	return rv
}

// The start offset of the range in bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSMetadataRange/startOffset
func (f_ FSMetadataRange) StartOffset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("startOffset"))
	return rv
}


