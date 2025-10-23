// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AssetWriterInputMetadataAdaptor] class.
var (
	AssetWriterInputMetadataAdaptorClass     _AssetWriterInputMetadataAdaptorClass
	AssetWriterInputMetadataAdaptorClassOnce sync.Once
)

func getAssetWriterInputMetadataAdaptorClass() _AssetWriterInputMetadataAdaptorClass {
	AssetWriterInputMetadataAdaptorClassOnce.Do(func() {
		AssetWriterInputMetadataAdaptorClass = _AssetWriterInputMetadataAdaptorClass{objc.GetClass("AVAssetWriterInputMetadataAdaptor")}
	})
	return AssetWriterInputMetadataAdaptorClass
}

type _AssetWriterInputMetadataAdaptorClass struct {
	class objc.Class
}

// An interface definition for the [AssetWriterInputMetadataAdaptor] class.
type IAssetWriterInputMetadataAdaptor interface {
	objectivec.IObject
	// properties:
	AssetWriterInput() IAVAssetWriterInput
	SetAssetWriterInput(value IAVAssetWriterInput)
	// methods:
}

// An object that appends timed metadata groups to an asset writer input.
//
// Use a metadata adaptor to append track-level metadata, packaged as instances of , to an asset writer input.


// An object that appends timed metadata groups to an asset writer input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputMetadataAdaptor
type AssetWriterInputMetadataAdaptor struct {
	objectivec.Object
}

// AssetWriterInputMetadataAdaptorFrom constructs a [AssetWriterInputMetadataAdaptor] from an unsafe.Pointer.
//
// An object that appends timed metadata groups to an asset writer input.
func AssetWriterInputMetadataAdaptorFrom(ptr unsafe.Pointer) AssetWriterInputMetadataAdaptor {
	return AssetWriterInputMetadataAdaptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AssetWriterInputMetadataAdaptorClass) Alloc() AssetWriterInputMetadataAdaptor {
	rv := objc.Send[AssetWriterInputMetadataAdaptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AssetWriterInputMetadataAdaptorClass) New() AssetWriterInputMetadataAdaptor {
	rv := objc.Send[AssetWriterInputMetadataAdaptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetWriterInputMetadataAdaptor) Init() AssetWriterInputMetadataAdaptor {
	rv := objc.Send[AssetWriterInputMetadataAdaptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetWriterInputMetadataAdaptor) Autorelease() AssetWriterInputMetadataAdaptor {
	rv := objc.Send[AssetWriterInputMetadataAdaptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetWriterInputMetadataAdaptor creates a new AssetWriterInputMetadataAdaptor instance.
func NewAssetWriterInputMetadataAdaptor() AssetWriterInputMetadataAdaptor {
	return getAssetWriterInputMetadataAdaptorClass().New()
}



// The input for the metadata adaptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputmetadataadaptor/assetwriterinput
func (a_ AssetWriterInputMetadataAdaptor) AssetWriterInput() IAVAssetWriterInput {
	rv := objc.Send[AssetWriterInput](a_.ID, objc.Sel("assetWriterInput"))
	return rv
}


// The input for the metadata adaptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinputmetadataadaptor/assetwriterinput
func (a_ AssetWriterInputMetadataAdaptor) SetAssetWriterInput(value IAVAssetWriterInput) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAssetWriterInput:"), value)
}



