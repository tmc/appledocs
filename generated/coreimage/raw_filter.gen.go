// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RAWFilter] class.
var rAWFilterClass = _RAWFilterClass{objc.GetClass("CIRAWFilter")}

type _RAWFilterClass struct {
	class objc.Class
}

// An interface definition for the [RAWFilter] class.
type IRAWFilter interface {
	IFilter
}

// A filter subclass that produces an image by manipulating RAW image sensor data from a digital camera or scanner. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter

type RAWFilter struct {
	Filter
}

// RAWFilterFrom constructs a [RAWFilter] from an unsafe.Pointer.
//
// A filter subclass that produces an image by manipulating RAW image sensor data from a digital camera or scanner.
func RAWFilterFrom(ptr unsafe.Pointer) RAWFilter {
	return RAWFilter{
		Filter: FilterFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (rc _RAWFilterClass) Alloc() RAWFilter {
	rv := objc.Send[RAWFilter](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (rc _RAWFilterClass) New() RAWFilter {
	rv := objc.Send[RAWFilter](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RAWFilter) Init() RAWFilter {
	rv := objc.Send[RAWFilter](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RAWFilter) Autorelease() RAWFilter {
	rv := objc.Send[RAWFilter](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRAWFilter creates a new RAWFilter instance.
func NewRAWFilter() RAWFilter {
	return rAWFilterClass.New()
}


// Creates a RAW filter from the image at the URL location that you specify. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(imageURL:)
func NewRAWFilterWithImageURL(url unsafe.Pointer) RAWFilter {
	rv := objc.Send[RAWFilter](objc.ID(rAWFilterClass.class), objc.Sel("filterWithImageURL:"), url)
	rv.Autorelease()
	return rv
}
// Creates a RAW filter from the pixel buffer and its properties that you specify. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(cvPixelBuffer:properties:)
func NewRAWFilterWithCVPixelBufferProperties(buffer unsafe.Pointer, properties unsafe.Pointer) RAWFilter {
	rv := objc.Send[RAWFilter](objc.ID(rAWFilterClass.class), objc.Sel("filterWithCVPixelBuffer:properties:"), buffer, properties)
	rv.Autorelease()
	return rv
}
// Creates a RAW filter from the image data and type hint that you specify. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(imageData:identifierHint:)
func NewRAWFilterWithImageDataIdentifierHint(data unsafe.Pointer, identifierHint string) RAWFilter {
	rv := objc.Send[RAWFilter](objc.ID(rAWFilterClass.class), objc.Sel("filterWithImageData:identifierHint:"), data, identifierHint)
	rv.Autorelease()
	return rv
}


// Creates a RAW filter from the pixel buffer and its properties that you specify. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(cvPixelBuffer:properties:)
func (rc _RAWFilterClass) FilterWithCVPixelBufferProperties(buffer unsafe.Pointer, properties unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("filterWithCVPixelBuffer:properties:"), buffer, properties)
	return rv
}
// Creates a RAW filter from the image data and type hint that you specify. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(imageData:identifierHint:)
func (rc _RAWFilterClass) FilterWithImageDataIdentifierHint(data unsafe.Pointer, identifierHint string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("filterWithImageData:identifierHint:"), data, identifierHint)
	return rv
}
// Creates a RAW filter from the image at the URL location that you specify. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(imageURL:)
func (rc _RAWFilterClass) FilterWithImageURL(url unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("filterWithImageURL:"), url)
	return rv
}

