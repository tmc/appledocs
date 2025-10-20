// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LogicalToPhysicalColorAttachmentMap] class.
var (
	LogicalToPhysicalColorAttachmentMapClass     _LogicalToPhysicalColorAttachmentMapClass
	LogicalToPhysicalColorAttachmentMapClassOnce sync.Once
)

func getLogicalToPhysicalColorAttachmentMapClass() _LogicalToPhysicalColorAttachmentMapClass {
	LogicalToPhysicalColorAttachmentMapClassOnce.Do(func() {
		LogicalToPhysicalColorAttachmentMapClass = _LogicalToPhysicalColorAttachmentMapClass{objc.GetClass("MTLLogicalToPhysicalColorAttachmentMap")}
	})
	return LogicalToPhysicalColorAttachmentMapClass
}

type _LogicalToPhysicalColorAttachmentMapClass struct {
	class objc.Class
}

// An interface definition for the [LogicalToPhysicalColorAttachmentMap] class.
type ILogicalToPhysicalColorAttachmentMap interface {
	objectivec.IObject
	GetPhysicalIndexForLogicalIndex(logicalIndex uint) uint
	Reset()
	SetPhysicalIndexForLogicalIndex(physicalIndex uint, logicalIndex uint)
}

// Allows you to easily specify color attachment remapping from logical to physical indices.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogicalToPhysicalColorAttachmentMap
type LogicalToPhysicalColorAttachmentMap struct {
	objectivec.Object
}

// LogicalToPhysicalColorAttachmentMapFrom constructs a [LogicalToPhysicalColorAttachmentMap] from an unsafe.Pointer.
//
// Allows you to easily specify color attachment remapping from logical to physical indices.
func LogicalToPhysicalColorAttachmentMapFrom(ptr unsafe.Pointer) LogicalToPhysicalColorAttachmentMap {
	return LogicalToPhysicalColorAttachmentMap{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LogicalToPhysicalColorAttachmentMapClass) Alloc() LogicalToPhysicalColorAttachmentMap {
	rv := objc.Send[LogicalToPhysicalColorAttachmentMap](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LogicalToPhysicalColorAttachmentMapClass) New() LogicalToPhysicalColorAttachmentMap {
	rv := objc.Send[LogicalToPhysicalColorAttachmentMap](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LogicalToPhysicalColorAttachmentMap) Init() LogicalToPhysicalColorAttachmentMap {
	rv := objc.Send[LogicalToPhysicalColorAttachmentMap](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LogicalToPhysicalColorAttachmentMap) Autorelease() LogicalToPhysicalColorAttachmentMap {
	rv := objc.Send[LogicalToPhysicalColorAttachmentMap](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLogicalToPhysicalColorAttachmentMap creates a new LogicalToPhysicalColorAttachmentMap instance.
func NewLogicalToPhysicalColorAttachmentMap() LogicalToPhysicalColorAttachmentMap {
	return getLogicalToPhysicalColorAttachmentMapClass().New()
}


// Queries the physical color attachment index corresponding to a logical index.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogicalToPhysicalColorAttachmentMap/getPhysicalIndexForLogicalIndex:
func (l_ LogicalToPhysicalColorAttachmentMap) GetPhysicalIndexForLogicalIndex(logicalIndex uint) uint {
	rv := objc.Send[uint](l_.ID, objc.Sel("getPhysicalIndexForLogicalIndex:"), logicalIndex)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogicalToPhysicalColorAttachmentMap/reset()
func (l_ LogicalToPhysicalColorAttachmentMap) Reset() {
	objc.Send[objc.ID](l_.ID, objc.Sel("reset"))
}

// Maps a physical color attachment index to a logical index.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogicalToPhysicalColorAttachmentMap/setPhysicalIndex:forLogicalIndex:
func (l_ LogicalToPhysicalColorAttachmentMap) SetPhysicalIndexForLogicalIndex(physicalIndex uint, logicalIndex uint) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setPhysicalIndex:forLogicalIndex:"), physicalIndex, logicalIndex)
}



