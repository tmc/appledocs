// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AssetWriterInputGroup] class.
var (
	AssetWriterInputGroupClass     _AssetWriterInputGroupClass
	AssetWriterInputGroupClassOnce sync.Once
)

func getAssetWriterInputGroupClass() _AssetWriterInputGroupClass {
	AssetWriterInputGroupClassOnce.Do(func() {
		AssetWriterInputGroupClass = _AssetWriterInputGroupClass{objc.GetClass("AVAssetWriterInputGroup")}
	})
	return AssetWriterInputGroupClass
}

type _AssetWriterInputGroupClass struct {
	class objc.Class
}

// An interface definition for the [AssetWriterInputGroup] class.
type IAssetWriterInputGroup interface {
	IMediaSelectionGroup
	// properties:
	DefaultInput() IAVAssetWriterInput
	SetDefaultInput(value IAVAssetWriterInput)
	Inputs() IAVAssetWriterInput
	SetInputs(value IAVAssetWriterInput)
	// methods:
}

// A group of inputs with tracks that are mutually exclusive to each other for playback or processing.
//
// Assets may contain multiple tracks of media that are mutually exclusive to each other when you play or process them. For example, an asset may contain multiple audio tracks for different spoken languages, but only one of them should play at a time. You use an input group to mark a collection of tracks as mutually exclusive to each other in the file the asset writer outputs.


// A group of inputs with tracks that are mutually exclusive to each other for playback or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputGroup
type AssetWriterInputGroup struct {
	MediaSelectionGroup
}

// AssetWriterInputGroupFrom constructs a [AssetWriterInputGroup] from an unsafe.Pointer.
//
// A group of inputs with tracks that are mutually exclusive to each other for playback or processing.
func AssetWriterInputGroupFrom(ptr unsafe.Pointer) AssetWriterInputGroup {
	return AssetWriterInputGroup{
		MediaSelectionGroup: MediaSelectionGroupFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AssetWriterInputGroupClass) Alloc() AssetWriterInputGroup {
	rv := objc.Send[AssetWriterInputGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AssetWriterInputGroupClass) New() AssetWriterInputGroup {
	rv := objc.Send[AssetWriterInputGroup](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetWriterInputGroup) Init() AssetWriterInputGroup {
	rv := objc.Send[AssetWriterInputGroup](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetWriterInputGroup) Autorelease() AssetWriterInputGroup {
	rv := objc.Send[AssetWriterInputGroup](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetWriterInputGroup creates a new AssetWriterInputGroup instance.
func NewAssetWriterInputGroup() AssetWriterInputGroup {
	return getAssetWriterInputGroupClass().New()
}



// The default input for the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputgroup/defaultinput
func (a_ AssetWriterInputGroup) DefaultInput() IAVAssetWriterInput {
	rv := objc.Send[AssetWriterInput](a_.ID, objc.Sel("defaultInput"))
	return rv
}


// The default input for the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputgroup/defaultinput
func (a_ AssetWriterInputGroup) SetDefaultInput(value IAVAssetWriterInput) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDefaultInput:"), value)
}


// The inputs with tracks that are mutually exclusive to each other for playback or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputgroup/inputs
func (a_ AssetWriterInputGroup) Inputs() IAVAssetWriterInput {
	rv := objc.Send[AssetWriterInput](a_.ID, objc.Sel("inputs"))
	return rv
}


// The inputs with tracks that are mutually exclusive to each other for playback or processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputgroup/inputs
func (a_ AssetWriterInputGroup) SetInputs(value IAVAssetWriterInput) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInputs:"), value)
}



