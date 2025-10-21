// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DetectBarcodesRequest] class.
var (
	DetectBarcodesRequestClass     _DetectBarcodesRequestClass
	DetectBarcodesRequestClassOnce sync.Once
)

func getDetectBarcodesRequestClass() _DetectBarcodesRequestClass {
	DetectBarcodesRequestClassOnce.Do(func() {
		DetectBarcodesRequestClass = _DetectBarcodesRequestClass{objc.GetClass("VNDetectBarcodesRequest")}
	})
	return DetectBarcodesRequestClass
}

type _DetectBarcodesRequestClass struct {
	class objc.Class
}

// An interface definition for the [DetectBarcodesRequest] class.
type IDetectBarcodesRequest interface {
	IImageBasedRequest
}

// A request that detects barcodes in an image.
//
// This request returns an array of objects, one for each barcode it detects.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectBarcodesRequest
type DetectBarcodesRequest struct {
	ImageBasedRequest
}

// DetectBarcodesRequestFrom constructs a [DetectBarcodesRequest] from an unsafe.Pointer.
//
// A request that detects barcodes in an image.
func DetectBarcodesRequestFrom(ptr unsafe.Pointer) DetectBarcodesRequest {
	return DetectBarcodesRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DetectBarcodesRequestClass) Alloc() DetectBarcodesRequest {
	rv := objc.Send[DetectBarcodesRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DetectBarcodesRequestClass) New() DetectBarcodesRequest {
	rv := objc.Send[DetectBarcodesRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectBarcodesRequest) Init() DetectBarcodesRequest {
	rv := objc.Send[DetectBarcodesRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectBarcodesRequest) Autorelease() DetectBarcodesRequest {
	rv := objc.Send[DetectBarcodesRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectBarcodesRequest creates a new DetectBarcodesRequest instance.
func NewDetectBarcodesRequest() DetectBarcodesRequest {
	return getDetectBarcodesRequestClass().New()
}


// A Boolean value that indicates whether to coalesce multiple codes based on the symbology.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequest/coalescecompositesymbologies
func (d_ DetectBarcodesRequest) CoalesceCompositeSymbologies() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("coalesceCompositeSymbologies"))
	return rv
}


// SetCoalesceCompositeSymbologies sets the value of the coalesceCompositeSymbologies property.
// A Boolean value that indicates whether to coalesce multiple codes based on the symbology.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequest/coalescecompositesymbologies
func (d_ DetectBarcodesRequest) SetCoalesceCompositeSymbologies(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCoalesceCompositeSymbologies:"), value)
}

// The results of a barcode detection request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequest/results
func (d_ DetectBarcodesRequest) Results() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("results"))
	return rv
}


// SetResults sets the value of the results property.
// The results of a barcode detection request.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequest/results
func (d_ DetectBarcodesRequest) SetResults(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setResults:"), value)
}

// The barcode symbologies that the request detects in an image.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequest/symbologies
func (d_ DetectBarcodesRequest) Symbologies() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("symbologies"))
	return rv
}


// SetSymbologies sets the value of the symbologies property.
// The barcode symbologies that the request detects in an image.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequest/symbologies
func (d_ DetectBarcodesRequest) SetSymbologies(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSymbologies:"), value)
}

// A constant for specifying revision 1 of the barcode detection request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequestrevision1
func (d_ DetectBarcodesRequest) VNDetectBarcodesRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectBarcodesRequestRevision1"))
	return rv
}

// A constant for specifying revision 2 of the barcode detection request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequestrevision2
func (d_ DetectBarcodesRequest) VNDetectBarcodesRequestRevision2() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectBarcodesRequestRevision2"))
	return rv
}

// A constant for specifying revision 3 of the barcode detection request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequestrevision3
func (d_ DetectBarcodesRequest) VNDetectBarcodesRequestRevision3() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectBarcodesRequestRevision3"))
	return rv
}



