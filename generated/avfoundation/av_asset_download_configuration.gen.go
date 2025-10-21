// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AssetDownloadConfiguration] class.
var (
	AssetDownloadConfigurationClass     _AssetDownloadConfigurationClass
	AssetDownloadConfigurationClassOnce sync.Once
)

func getAssetDownloadConfigurationClass() _AssetDownloadConfigurationClass {
	AssetDownloadConfigurationClassOnce.Do(func() {
		AssetDownloadConfigurationClass = _AssetDownloadConfigurationClass{objc.GetClass("AVAssetDownloadConfiguration")}
	})
	return AssetDownloadConfigurationClass
}

type _AssetDownloadConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [AssetDownloadConfiguration] class.
type IAssetDownloadConfiguration interface {
	objectivec.IObject
}

// An object that provides the configuration for a download task.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration
type AssetDownloadConfiguration struct {
	objectivec.Object
}

// AssetDownloadConfigurationFrom constructs a [AssetDownloadConfiguration] from an unsafe.Pointer.
//
// An object that provides the configuration for a download task.
func AssetDownloadConfigurationFrom(ptr unsafe.Pointer) AssetDownloadConfiguration {
	return AssetDownloadConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AssetDownloadConfigurationClass) Alloc() AssetDownloadConfiguration {
	rv := objc.Send[AssetDownloadConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AssetDownloadConfigurationClass) New() AssetDownloadConfiguration {
	rv := objc.Send[AssetDownloadConfiguration](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetDownloadConfiguration) Init() AssetDownloadConfiguration {
	rv := objc.Send[AssetDownloadConfiguration](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetDownloadConfiguration) Autorelease() AssetDownloadConfiguration {
	rv := objc.Send[AssetDownloadConfiguration](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetDownloadConfiguration creates a new AssetDownloadConfiguration instance.
func NewAssetDownloadConfiguration() AssetDownloadConfiguration {
	return getAssetDownloadConfigurationClass().New()
}




