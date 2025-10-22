// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EDRMetadata] class.
var (
	EDRMetadataClass     _EDRMetadataClass
	EDRMetadataClassOnce sync.Once
)

func getEDRMetadataClass() _EDRMetadataClass {
	EDRMetadataClassOnce.Do(func() {
		EDRMetadataClass = _EDRMetadataClass{objc.GetClass("CAEDRMetadata")}
	})
	return EDRMetadataClass
}

type _EDRMetadataClass struct {
	class objc.Class
}

// An interface definition for the [EDRMetadata] class.
type IEDRMetadata interface {
	objectivec.IObject
	EdrMetadata() CAEDRMetadata
	SetEdrMetadata(value IEDRMetadata)
}

// Metadata describing how extended dynamic range (EDR) values should be tone mapped.
//
// If you need specific tone-mapping behavior, set the property of a to point to an instance of this class.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEDRMetadata
type EDRMetadata struct {
	objectivec.Object
}

// EDRMetadataFrom constructs a [EDRMetadata] from an unsafe.Pointer.
//
// Metadata describing how extended dynamic range (EDR) values should be tone mapped.
func EDRMetadataFrom(ptr unsafe.Pointer) EDRMetadata {
	return EDRMetadata{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EDRMetadataClass) Alloc() EDRMetadata {
	rv := objc.Send[EDRMetadata](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EDRMetadataClass) New() EDRMetadata {
	rv := objc.Send[EDRMetadata](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EDRMetadata) Init() EDRMetadata {
	rv := objc.Send[EDRMetadata](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EDRMetadata) Autorelease() EDRMetadata {
	rv := objc.Send[EDRMetadata](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEDRMetadata creates a new EDRMetadata instance.
func NewEDRMetadata() EDRMetadata {
	return getEDRMetadataClass().New()
}


// Creates EDR metadata for HDR10 content based on mastering display color information and content light levels.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEDRMetadata/hdr10(displayInfo:contentInfo:opticalOutputScale:)
func (ec _EDRMetadataClass) HDR10MetadataWithDisplayInfoContentInfoOpticalOutputScale(displayData foundation.IData, contentData foundation.IData, scale float32) EDRMetadata {
	rv := objc.Send[EDRMetadata](objc.ID(ec.class), objc.Sel("HDR10MetadataWithDisplayInfo:contentInfo:opticalOutputScale:"), displayData, contentData, scale)
	return rv
}

// Creates EDR metadata for HDR10 content based on the luminance characteristics of a mastering display.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEDRMetadata/hdr10(minLuminance:maxLuminance:opticalOutputScale:)
func (ec _EDRMetadataClass) HDR10MetadataWithMinLuminanceMaxLuminanceOpticalOutputScale(minNits float32, maxNits float32, scale float32) EDRMetadata {
	rv := objc.Send[EDRMetadata](objc.ID(ec.class), objc.Sel("HDR10MetadataWithMinLuminance:maxLuminance:opticalOutputScale:"), minNits, maxNits, scale)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEDRMetadata/hlg(ambientViewingEnvironment:)
func (ec _EDRMetadataClass) HLGMetadataWithAmbientViewingEnvironment(data foundation.IData) EDRMetadata {
	rv := objc.Send[EDRMetadata](objc.ID(ec.class), objc.Sel("HLGMetadataWithAmbientViewingEnvironment:"), data)
	return rv
}

// Extended dynamic range (EDR) metadata for the Hybrid Log-Gamma (HLG) transfer function.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEDRMetadata/hlg
func (ec _EDRMetadataClass) HLGMetadata() EDRMetadata {
	rv := objc.Send[CAEDRMetadata](objc.ID(ec.class), objc.Sel("HLGMetadata"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEDRMetadata/isAvailable
func (ec _EDRMetadataClass) Available() bool {
	rv := objc.Send[bool](objc.ID(ec.class), objc.Sel("available"))
	return rv
}
// Extended dynamic range (EDR) metadata for the Hybrid Log-Gamma (HLG) transfer function.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEDRMetadata/hlg
func (e_ EDRMetadata) HLGMetadata() CAEDRMetadata {
	rv := objc.Send[CAEDRMetadata](e_.ID, objc.Sel("HLGMetadata"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEDRMetadata/isAvailable
func (e_ EDRMetadata) Available() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("available"))
	return rv
}

// Metadata describing the tone mapping to apply to the extended dynamic range (EDR) values in the layer.
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/edrmetadata
func (e_ EDRMetadata) EdrMetadata() CAEDRMetadata {
	rv := objc.Send[CAEDRMetadata](e_.ID, objc.Sel("edrMetadata"))
	return rv
}


// SetEdrMetadata sets the value of the edrMetadata property.
// Metadata describing the tone mapping to apply to the extended dynamic range (EDR) values in the layer.

//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/edrmetadata
func (e_ EDRMetadata) SetEdrMetadata(value IEDRMetadata) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEdrMetadata:"), value)
}



