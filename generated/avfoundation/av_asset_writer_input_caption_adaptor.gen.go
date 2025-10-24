// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AssetWriterInputCaptionAdaptor] class.
var (
	AssetWriterInputCaptionAdaptorClass     _AssetWriterInputCaptionAdaptorClass
	AssetWriterInputCaptionAdaptorClassOnce sync.Once
)

func getAssetWriterInputCaptionAdaptorClass() _AssetWriterInputCaptionAdaptorClass {
	AssetWriterInputCaptionAdaptorClassOnce.Do(func() {
		AssetWriterInputCaptionAdaptorClass = _AssetWriterInputCaptionAdaptorClass{objc.GetClass("AVAssetWriterInputCaptionAdaptor")}
	})
	return AssetWriterInputCaptionAdaptorClass
}

type _AssetWriterInputCaptionAdaptorClass struct {
	class objc.Class
}





// An interface definition for the [AssetWriterInputCaptionAdaptor] class.
type IAssetWriterInputCaptionAdaptor interface {
	objectivec.IObject
	

	// properties:
	AssetWriterInput() IAVAssetWriterInput


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AssetWriterInputCaptionAdaptorClass) Alloc() AssetWriterInputCaptionAdaptor {
	rv := objc.Send[AssetWriterInputCaptionAdaptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetWriterInputCaptionAdaptorClass) New() AssetWriterInputCaptionAdaptor {
	rv := objc.Send[AssetWriterInputCaptionAdaptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetWriterInputCaptionAdaptor) Init() AssetWriterInputCaptionAdaptor {
	rv := objc.Send[AssetWriterInputCaptionAdaptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetWriterInputCaptionAdaptor) Autorelease() AssetWriterInputCaptionAdaptor {
	rv := objc.Send[AssetWriterInputCaptionAdaptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetWriterInputCaptionAdaptor creates a new AssetWriterInputCaptionAdaptor instance.
func NewAssetWriterInputCaptionAdaptor() AssetWriterInputCaptionAdaptor {
	return getAssetWriterInputCaptionAdaptorClass().New()
}





// An object that appends captions to an asset writer input.


// An object that appends captions to an asset writer input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputCaptionAdaptor
type AssetWriterInputCaptionAdaptor struct {
	objectivec.Object
}

// AssetWriterInputCaptionAdaptorFrom constructs a [AssetWriterInputCaptionAdaptor] from an unsafe.Pointer.
//
// An object that appends captions to an asset writer input.
func AssetWriterInputCaptionAdaptorFrom(ptr unsafe.Pointer) AssetWriterInputCaptionAdaptor {
	return AssetWriterInputCaptionAdaptor{objectivec.Object{objc.ID(ptr)}}
}






// Creates a new caption adaptor that writes to the specified asset writer input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputCaptionAdaptor/init(assetWriterInput:)
func NewAssetWriterInputCaptionAdaptorWithAssetWriterInput(input IAVAssetWriterInput) AssetWriterInputCaptionAdaptor {
	instance := getAssetWriterInputCaptionAdaptorClass().Alloc()
	rv := objc.Send[AssetWriterInputCaptionAdaptor](instance.ID, objc.Sel("initWithAssetWriterInput:"), input)
	rv.Autorelease()
	return rv
}







// A class method that creates a new caption adaptor that writes to the specified asset writer input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputCaptionAdaptor/assetWriterInputCaptionAdaptorWithAssetWriterInput:
func (ac _AssetWriterInputCaptionAdaptorClass) AssetWriterInputCaptionAdaptorWithAssetWriterInput(input IAVAssetWriterInput) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetWriterInputCaptionAdaptorWithAssetWriterInput:"), input)
	return rv
}

















// The associated asset writer input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputCaptionAdaptor/assetWriterInput
func (a_ AssetWriterInputCaptionAdaptor) AssetWriterInput() IAVAssetWriterInput {
	rv := objc.Send[AssetWriterInput](a_.ID, objc.Sel("assetWriterInput"))
	return rv
}







