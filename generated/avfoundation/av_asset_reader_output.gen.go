// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AssetReaderOutput] class.
var (
	AssetReaderOutputClass     _AssetReaderOutputClass
	AssetReaderOutputClassOnce sync.Once
)

func getAssetReaderOutputClass() _AssetReaderOutputClass {
	AssetReaderOutputClassOnce.Do(func() {
		AssetReaderOutputClass = _AssetReaderOutputClass{objc.GetClass("AVAssetReaderOutput")}
	})
	return AssetReaderOutputClass
}

type _AssetReaderOutputClass struct {
	class objc.Class
}

// An interface definition for the [AssetReaderOutput] class.
type IAssetReaderOutput interface {
	objectivec.IObject
}

// An abstract class that defines the interface to read media samples from an asset reader.
//
// You add concrete output instances, such as or , to an asset reader to perform specific tasks.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutput
type AssetReaderOutput struct {
	objectivec.Object
}

// AssetReaderOutputFrom constructs a [AssetReaderOutput] from an unsafe.Pointer.
//
// An abstract class that defines the interface to read media samples from an asset reader.
func AssetReaderOutputFrom(ptr unsafe.Pointer) AssetReaderOutput {
	return AssetReaderOutput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AssetReaderOutputClass) Alloc() AssetReaderOutput {
	rv := objc.Send[AssetReaderOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AssetReaderOutputClass) New() AssetReaderOutput {
	rv := objc.Send[AssetReaderOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetReaderOutput) Init() AssetReaderOutput {
	rv := objc.Send[AssetReaderOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetReaderOutput) Autorelease() AssetReaderOutput {
	rv := objc.Send[AssetReaderOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetReaderOutput creates a new AssetReaderOutput instance.
func NewAssetReaderOutput() AssetReaderOutput {
	return getAssetReaderOutputClass().New()
}


// A Boolean value that indicates whether the output vends copied sample data.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/alwayscopiessampledata
func (a_ AssetReaderOutput) AlwaysCopiesSampleData() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("alwaysCopiesSampleData"))
	return rv
}


// SetAlwaysCopiesSampleData sets the value of the alwaysCopiesSampleData property.
// A Boolean value that indicates whether the output vends copied sample data.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/alwayscopiessampledata
func (a_ AssetReaderOutput) SetAlwaysCopiesSampleData(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAlwaysCopiesSampleData:"), value)
}

// The media type of samples that the output reads.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/mediatype
func (a_ AssetReaderOutput) MediaType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("mediaType"))
	return rv
}


// SetMediaType sets the value of the mediaType property.
// The media type of samples that the output reads.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/mediatype
func (a_ AssetReaderOutput) SetMediaType(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMediaType:"), value)
}

// A Boolean value that indicates whether the output supports reconfiguring the time ranges it reads.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/supportsrandomaccess
func (a_ AssetReaderOutput) SupportsRandomAccess() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("supportsRandomAccess"))
	return rv
}


// SetSupportsRandomAccess sets the value of the supportsRandomAccess property.
// A Boolean value that indicates whether the output supports reconfiguring the time ranges it reads.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/supportsrandomaccess
func (a_ AssetReaderOutput) SetSupportsRandomAccess(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSupportsRandomAccess:"), value)
}



