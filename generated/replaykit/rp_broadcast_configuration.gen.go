// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RPBroadcastConfiguration] class.
var (
	RPBroadcastConfigurationClass     _RPBroadcastConfigurationClass
	RPBroadcastConfigurationClassOnce sync.Once
)

func getRPBroadcastConfigurationClass() _RPBroadcastConfigurationClass {
	RPBroadcastConfigurationClassOnce.Do(func() {
		RPBroadcastConfigurationClass = _RPBroadcastConfigurationClass{objc.GetClass("RPBroadcastConfiguration")}
	})
	return RPBroadcastConfigurationClass
}

type _RPBroadcastConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [RPBroadcastConfiguration] class.
type IRPBroadcastConfiguration interface {
	objectivec.IObject
}

// An object used to configure the movie clips produced during a live broadcast.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastConfiguration
type RPBroadcastConfiguration struct {
	objectivec.Object
}

// RPBroadcastConfigurationFrom constructs a [RPBroadcastConfiguration] from an unsafe.Pointer.
//
// An object used to configure the movie clips produced during a live broadcast.
func RPBroadcastConfigurationFrom(ptr unsafe.Pointer) RPBroadcastConfiguration {
	return RPBroadcastConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RPBroadcastConfigurationClass) Alloc() RPBroadcastConfiguration {
	rv := objc.Send[RPBroadcastConfiguration](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RPBroadcastConfigurationClass) New() RPBroadcastConfiguration {
	rv := objc.Send[RPBroadcastConfiguration](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RPBroadcastConfiguration) Init() RPBroadcastConfiguration {
	rv := objc.Send[RPBroadcastConfiguration](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RPBroadcastConfiguration) Autorelease() RPBroadcastConfiguration {
	rv := objc.Send[RPBroadcastConfiguration](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRPBroadcastConfiguration creates a new RPBroadcastConfiguration instance.
func NewRPBroadcastConfiguration() RPBroadcastConfiguration {
	return getRPBroadcastConfigurationClass().New()
}


// The duration of movie clips sent the to the movie clip handler extension.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastConfiguration/clipDuration
func (r_ RPBroadcastConfiguration) ClipDuration() TimeInterval {
	rv := objc.Send[TimeInterval](r_.ID, objc.Sel("clipDuration"))
	return rv
}


// SetClipDuration sets the value of the clipDuration property.
// The duration of movie clips sent the to the movie clip handler extension.

//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastConfiguration/clipDuration
func (r_ RPBroadcastConfiguration) SetClipDuration(value TimeInterval) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setClipDuration:"), value)
}

// The compression properties for encoding movie clips that are to be overwritten.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastConfiguration/videoCompressionProperties
func (r_ RPBroadcastConfiguration) VideoCompressionProperties() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("videoCompressionProperties"))
	return rv
}


// SetVideoCompressionProperties sets the value of the videoCompressionProperties property.
// The compression properties for encoding movie clips that are to be overwritten.

//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastConfiguration/videoCompressionProperties
func (r_ RPBroadcastConfiguration) SetVideoCompressionProperties(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setVideoCompressionProperties:"), value)
}



