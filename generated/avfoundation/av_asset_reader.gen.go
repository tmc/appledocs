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
	CanAddOutput(output unsafe.Pointer) bool
	StartReading() bool
}

// An object that reads media data from an asset.
//
// Use an asset reader to read media data from instances of . The assets you read may represent file-based media like QuickTime movies or MPEG-4 files, or media that you compose from multiple sources using .
//
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


// Determines whether you can add the output to the asset reader.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/canAdd(_:)
func (a_ AssetReader) CanAddOutput(output unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canAddOutput:"), output)
	return rv
}

// Prepares the asset reader to start reading sample buffers from the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/startReading()
func (a_ AssetReader) StartReading() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("startReading"))
	return rv
}



