// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AssetReader] class.
var (
	AssetReaderClass     _AssetReaderClass
	AssetReaderClassOnce sync.Once
)

func getAssetReaderClass() _AssetReaderClass {
	AssetReaderClassOnce.Do(func() {
		AssetReaderClass = _AssetReaderClass{objc.GetClass("AVAssetReader")}
	})
	return AssetReaderClass
}

type _AssetReaderClass struct {
	class objc.Class
}

// An interface definition for the [AssetReader] class.
type IAssetReader interface {
	objectivec.IObject
	// properties:
	Asset() IAVAsset
	SetAsset(value IAVAsset)
	Error() Error
	SetError(value Error)
	Outputs() AssetReaderOutput /* not a class type */
	SetOutputs(value AssetReaderOutput /* not a class type */)
	Status() unsafe.Pointer
	SetStatus(value unsafe.Pointer)
	TimeRange() TimeRange /* not a class type */
	SetTimeRange(value TimeRange /* not a class type */)
	// methods:
}

// An object that reads media data from an asset.
//
// Use an asset reader to read media data from instances of . The assets you read may represent file-based media like QuickTime movies or MPEG-4 files, or media that you compose from multiple sources using .


// An object that reads media data from an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader
type AssetReader struct {
	objectivec.Object
}

// AssetReaderFrom constructs a [AssetReader] from an unsafe.Pointer.
//
// An object that reads media data from an asset.
func AssetReaderFrom(ptr unsafe.Pointer) AssetReader {
	return AssetReader{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AssetReaderClass) Alloc() AssetReader {
	rv := objc.Send[AssetReader](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AssetReaderClass) New() AssetReader {
	rv := objc.Send[AssetReader](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetReader) Init() AssetReader {
	rv := objc.Send[AssetReader](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetReader) Autorelease() AssetReader {
	rv := objc.Send[AssetReader](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetReader creates a new AssetReader instance.
func NewAssetReader() AssetReader {
	return getAssetReaderClass().New()
}



// The asset from which to read media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/asset
func (a_ AssetReader) Asset() IAVAsset {
	rv := objc.Send[Asset](a_.ID, objc.Sel("asset"))
	return rv
}


// The asset from which to read media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/asset
func (a_ AssetReader) SetAsset(value IAVAsset) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAsset:"), value)
}


// An error that describes the reason for a failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/error
func (a_ AssetReader) Error() Error {
	rv := objc.Send[Error](a_.ID, objc.Sel("error"))
	return rv
}


// An error that describes the reason for a failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/error
func (a_ AssetReader) SetError(value Error) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setError:"), value)
}


// The outputs from which you read media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/outputs
func (a_ AssetReader) Outputs() AssetReaderOutput /* not a class type */ {
	rv := objc.Send[AssetReaderOutput](a_.ID, objc.Sel("outputs"))
	return rv
}


// The outputs from which you read media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/outputs
func (a_ AssetReader) SetOutputs(value AssetReaderOutput /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputs:"), value)
}


// The status of reading sample buffers from the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/status-swift.property
func (a_ AssetReader) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("status"))
	return rv
}


// The status of reading sample buffers from the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/status-swift.property
func (a_ AssetReader) SetStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStatus:"), value)
}


// The time range within the asset to read.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/timerange
func (a_ AssetReader) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](a_.ID, objc.Sel("timeRange"))
	return rv
}


// The time range within the asset to read.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreader/timerange
func (a_ AssetReader) SetTimeRange(value TimeRange /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTimeRange:"), value)
}



