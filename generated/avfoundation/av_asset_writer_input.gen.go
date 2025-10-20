// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AssetWriterInput] class.
var (
	AssetWriterInputClass     _AssetWriterInputClass
	AssetWriterInputClassOnce sync.Once
)

func getAssetWriterInputClass() _AssetWriterInputClass {
	AssetWriterInputClassOnce.Do(func() {
		AssetWriterInputClass = _AssetWriterInputClass{objc.GetClass("AVAssetWriterInput")}
	})
	return AssetWriterInputClass
}

type _AssetWriterInputClass struct {
	class objc.Class
}

// An interface definition for the [AssetWriterInput] class.
type IAssetWriterInput interface {
	objectivec.IObject
	AppendSampleBuffer(sampleBuffer unsafe.Pointer) bool
}

// An object that appends media samples to a track in an asset writer’s output file.
//
// Create an asset writer input to write a single track of media, and optional track-level metadata, to the output file. To write multiple concurrent tracks with ideal interleaving of media data, observe the value of the property of each input. You can use an asset writer input to create tracks in a QuickTime movie file that aren’t self-contained, and instead reference sample data that exists in another file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput
type AssetWriterInput struct {
	objectivec.Object
}

// AssetWriterInputFrom constructs a [AssetWriterInput] from an unsafe.Pointer.
//
// An object that appends media samples to a track in an asset writer’s output file.
func AssetWriterInputFrom(ptr unsafe.Pointer) AssetWriterInput {
	return AssetWriterInput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AssetWriterInputClass) Alloc() AssetWriterInput {
	rv := objc.Send[AssetWriterInput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AssetWriterInputClass) New() AssetWriterInput {
	rv := objc.Send[AssetWriterInput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetWriterInput) Init() AssetWriterInput {
	rv := objc.Send[AssetWriterInput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetWriterInput) Autorelease() AssetWriterInput {
	rv := objc.Send[AssetWriterInput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetWriterInput creates a new AssetWriterInput instance.
func NewAssetWriterInput() AssetWriterInput {
	return getAssetWriterInputClass().New()
}


// Creates an input to append sample buffers of the specified type to the output file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/init(mediaType:outputSettings:)
func NewAssetWriterInputWithMediaTypeOutputSettings(mediaType unsafe.Pointer, outputSettings unsafe.Pointer) AssetWriterInput {
	instance := getAssetWriterInputClass().Alloc()
	rv := objc.Send[AssetWriterInput](instance.ID, objc.Sel("initWithMediaType:outputSettings:"), mediaType, outputSettings)
	rv.Autorelease()
	return rv
}

// Creates an input that appends sample buffers of the specified type and format hint to the output file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/init(mediaType:outputSettings:sourceFormatHint:)
func NewAssetWriterInputWithMediaTypeOutputSettingsSourceFormatHint(mediaType unsafe.Pointer, outputSettings unsafe.Pointer, sourceFormatHint unsafe.Pointer) AssetWriterInput {
	instance := getAssetWriterInputClass().Alloc()
	rv := objc.Send[AssetWriterInput](instance.ID, objc.Sel("initWithMediaType:outputSettings:sourceFormatHint:"), mediaType, outputSettings, sourceFormatHint)
	rv.Autorelease()
	return rv
}


// Appends a sample buffer to an input to write to the output file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/append(_:)
func (a_ AssetWriterInput) AppendSampleBuffer(sampleBuffer unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("appendSampleBuffer:"), sampleBuffer)
	return rv
}


