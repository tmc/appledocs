// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVAssetReader] class.
var (
	aVAssetReaderClass     _AVAssetReaderClass
	aVAssetReaderClassOnce sync.Once
)

func getAVAssetReaderClass() _AVAssetReaderClass {
	aVAssetReaderClassOnce.Do(func() {
		aVAssetReaderClass = _AVAssetReaderClass{objc.GetClass("AVAssetReader")}
	})
	return aVAssetReaderClass
}

type _AVAssetReaderClass struct {
	class objc.Class
}

// An interface definition for the [AVAssetReader] class.
type IAVAssetReader interface {
	objectivec.IObject
}

// An object that reads media data from an asset.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader
type AVAssetReader struct {
	objectivec.Object
}

// AVAssetReaderFrom constructs a [AVAssetReader] from an unsafe.Pointer.
//
// An object that reads media data from an asset.
func AVAssetReaderFrom(ptr unsafe.Pointer) AVAssetReader {
	return AVAssetReader{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVAssetReaderClass) Alloc() AVAssetReader {
	rv := objc.Send[AVAssetReader](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVAssetReaderClass) New() AVAssetReader {
	rv := objc.Send[AVAssetReader](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVAssetReader) Init() AVAssetReader {
	rv := objc.Send[AVAssetReader](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVAssetReader) Autorelease() AVAssetReader {
	rv := objc.Send[AVAssetReader](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetReader creates a new AVAssetReader instance.
func NewAVAssetReader() AVAssetReader {
	return getAVAssetReaderClass().New()
}




