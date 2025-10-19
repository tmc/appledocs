// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVAssetDownloadConfiguration] class.
var aVAssetDownloadConfigurationClass = _AVAssetDownloadConfigurationClass{objc.GetClass("AVAssetDownloadConfiguration")}

type _AVAssetDownloadConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [AVAssetDownloadConfiguration] class.
type IAVAssetDownloadConfiguration interface {
	objectivec.IObject
}

// An object that provides the configuration for a download task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration

type AVAssetDownloadConfiguration struct {
	objectivec.Object
}

// AVAssetDownloadConfigurationFrom constructs a [AVAssetDownloadConfiguration] from an unsafe.Pointer.
//
// An object that provides the configuration for a download task.
func AVAssetDownloadConfigurationFrom(ptr unsafe.Pointer) AVAssetDownloadConfiguration {
	return AVAssetDownloadConfiguration{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ac _AVAssetDownloadConfigurationClass) Alloc() AVAssetDownloadConfiguration {
	rv := objc.Send[AVAssetDownloadConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVAssetDownloadConfigurationClass) New() AVAssetDownloadConfiguration {
	rv := objc.Send[AVAssetDownloadConfiguration](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVAssetDownloadConfiguration) Init() AVAssetDownloadConfiguration {
	rv := objc.Send[AVAssetDownloadConfiguration](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVAssetDownloadConfiguration) Autorelease() AVAssetDownloadConfiguration {
	rv := objc.Send[AVAssetDownloadConfiguration](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetDownloadConfiguration creates a new AVAssetDownloadConfiguration instance.
func NewAVAssetDownloadConfiguration() AVAssetDownloadConfiguration {
	return aVAssetDownloadConfigurationClass.New()
}




