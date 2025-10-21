// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MESampleCursorChunk] class.
var (
	MESampleCursorChunkClass     _MESampleCursorChunkClass
	MESampleCursorChunkClassOnce sync.Once
)

func getMESampleCursorChunkClass() _MESampleCursorChunkClass {
	MESampleCursorChunkClassOnce.Do(func() {
		MESampleCursorChunkClass = _MESampleCursorChunkClass{objc.GetClass("MESampleCursorChunk")}
	})
	return MESampleCursorChunkClass
}

type _MESampleCursorChunkClass struct {
	class objc.Class
}

// An interface definition for the [MESampleCursorChunk] class.
type IMESampleCursorChunk interface {
	objectivec.IObject
}

// An object that provides information about the chunk of media at the location of a sample.
//
// The method returns an instance of this class.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MESampleCursorChunk
type MESampleCursorChunk struct {
	objectivec.Object
}

// MESampleCursorChunkFrom constructs a [MESampleCursorChunk] from an unsafe.Pointer.
//
// An object that provides information about the chunk of media at the location of a sample.
func MESampleCursorChunkFrom(ptr unsafe.Pointer) MESampleCursorChunk {
	return MESampleCursorChunk{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MESampleCursorChunkClass) Alloc() MESampleCursorChunk {
	rv := objc.Send[MESampleCursorChunk](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MESampleCursorChunkClass) New() MESampleCursorChunk {
	rv := objc.Send[MESampleCursorChunk](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MESampleCursorChunk) Init() MESampleCursorChunk {
	rv := objc.Send[MESampleCursorChunk](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MESampleCursorChunk) Autorelease() MESampleCursorChunk {
	rv := objc.Send[MESampleCursorChunk](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMESampleCursorChunk creates a new MESampleCursorChunk instance.
func NewMESampleCursorChunk() MESampleCursorChunk {
	return getMESampleCursorChunkClass().New()
}


// The byte source to use to read the data for the sample.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MESampleCursorChunk/byteSource
func (m_ MESampleCursorChunk) ByteSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("byteSource"))
	return rv
}



