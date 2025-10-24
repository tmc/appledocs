// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MESampleCursorChunk */


/* debug [class_header]: Header for MESampleCursorChunk */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MESampleCursorChunk */
// An interface definition for the [MESampleCursorChunk] class.
type IMESampleCursorChunk interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MESampleCursorChunk */
	// properties:
	ByteSource() IMEByteSource
	ChunkInfo() objc.IObject /* cross-framework: SampleCursorChunkInfo */
	ChunkStorageRange() objc.IObject /* cross-framework: SampleCursorStorageRange */
	SampleIndexWithinChunk() Index /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MESampleCursorChunk */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MESampleCursorChunk */
// Alloc allocates a new instance without initialization.
func (mc _MESampleCursorChunkClass) Alloc() MESampleCursorChunk {
	rv := objc.Send[MESampleCursorChunk](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MESampleCursorChunk */
// An object that provides information about the chunk of media at the location of a sample.
//
// The method returns an instance of this class.


// An object that provides information about the chunk of media at the location of a sample.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MESampleCursorChunk */

// Creates a new sample cursor chunk with byte source and chunk data that you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MESampleCursorChunk/init(byteSource:chunkStorageRange:chunkInfo:sampleIndexWithinChunk:)
func NewMESampleCursorChunkWithByteSourceChunkStorageRangeChunkInfoSampleIndexWithinChunk(byteSource IMEByteSource, chunkStorageRange objc.IObject /* cross-framework: SampleCursorStorageRange */, chunkInfo objc.IObject /* cross-framework: SampleCursorChunkInfo */, sampleIndexWithinChunk Index /* not a class type */) MESampleCursorChunk {
	instance := getMESampleCursorChunkClass().Alloc()
	rv := objc.Send[MESampleCursorChunk](instance.ID, objc.Sel("initWithByteSource:chunkStorageRange:chunkInfo:sampleIndexWithinChunk:"), byteSource, chunkStorageRange, chunkInfo, sampleIndexWithinChunk)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMESampleCursorChunkWithByteSourceChunkStorageRangeChunkInfoSampleIndexWithinChunk */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MESampleCursorChunk */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MESampleCursorChunk */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MESampleCursorChunk */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MESampleCursorChunk */

// The byte source to use to read the data for the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MESampleCursorChunk/byteSource
func (m_ MESampleCursorChunk) ByteSource() IMEByteSource {
	rv := objc.Send[MEByteSource](m_.ID, objc.Sel("byteSource"))
	return rv
}/* debug [instance_properties/getter]: byteSource */


// An object that provides details about the chunk in the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MESampleCursorChunk/chunkInfo
func (m_ MESampleCursorChunk) ChunkInfo() objc.IObject /* cross-framework: SampleCursorChunkInfo */ {
	rv := objc.Send[avfoundation.SampleCursorChunkInfo](m_.ID, objc.Sel("chunkInfo"))
	return rv
}/* debug [instance_properties/getter]: chunkInfo */


// The offset location and length of the sample’s chunk within the byte source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MESampleCursorChunk/chunkStorageRange
func (m_ MESampleCursorChunk) ChunkStorageRange() objc.IObject /* cross-framework: SampleCursorStorageRange */ {
	rv := objc.Send[avfoundation.SampleCursorStorageRange](m_.ID, objc.Sel("chunkStorageRange"))
	return rv
}/* debug [instance_properties/getter]: chunkStorageRange */


// The offset index of the sample within the chunk, in samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MESampleCursorChunk/sampleIndexWithinChunk
func (m_ MESampleCursorChunk) SampleIndexWithinChunk() Index /* not a class type */ {
	rv := objc.Send[Index](m_.ID, objc.Sel("sampleIndexWithinChunk"))
	return rv
}/* debug [instance_properties/getter]: sampleIndexWithinChunk */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MESampleCursorChunk */


