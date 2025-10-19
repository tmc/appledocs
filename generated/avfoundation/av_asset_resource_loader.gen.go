// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVAssetResourceLoader] class.
var aVAssetResourceLoaderClass = _AVAssetResourceLoaderClass{objc.GetClass("AVAssetResourceLoader")}

type _AVAssetResourceLoaderClass struct {
	class objc.Class
}

// An interface definition for the [AVAssetResourceLoader] class.
type IAVAssetResourceLoader interface {
	objectivec.IObject
}

// An object that mediates resource requests from a URL asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader

type AVAssetResourceLoader struct {
	objectivec.Object
}

// AVAssetResourceLoaderFrom constructs a [AVAssetResourceLoader] from an unsafe.Pointer.
//
// An object that mediates resource requests from a URL asset.
func AVAssetResourceLoaderFrom(ptr unsafe.Pointer) AVAssetResourceLoader {
	return AVAssetResourceLoader{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ac _AVAssetResourceLoaderClass) Alloc() AVAssetResourceLoader {
	rv := objc.Send[AVAssetResourceLoader](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVAssetResourceLoaderClass) New() AVAssetResourceLoader {
	rv := objc.Send[AVAssetResourceLoader](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVAssetResourceLoader) Init() AVAssetResourceLoader {
	rv := objc.Send[AVAssetResourceLoader](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVAssetResourceLoader) Autorelease() AVAssetResourceLoader {
	rv := objc.Send[AVAssetResourceLoader](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetResourceLoader creates a new AVAssetResourceLoader instance.
func NewAVAssetResourceLoader() AVAssetResourceLoader {
	return aVAssetResourceLoaderClass.New()
}




