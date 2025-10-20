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




