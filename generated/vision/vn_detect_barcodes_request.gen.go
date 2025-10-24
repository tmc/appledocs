// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNDetectBarcodesRequest */


/* debug [class_header]: Header for VNDetectBarcodesRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DetectBarcodesRequest */
// An interface definition for the [DetectBarcodesRequest] class.
type IDetectBarcodesRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for DetectBarcodesRequest */
	// properties:
	CoalesceCompositeSymbologies() bool
	SetCoalesceCompositeSymbologies(value bool)
	Results() []BarcodeObservation
	Symbologies() []string
	SetSymbologies(value []string)
	VNDetectBarcodesRequestRevision1() int
	VNDetectBarcodesRequestRevision2() int
	VNDetectBarcodesRequestRevision3() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DetectBarcodesRequest */
	// methods:
	SupportedSymbologiesAndReturnError(error_ objectivec.IObject) []string
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DetectBarcodesRequest */
// Alloc allocates a new instance without initialization.
func (dc _DetectBarcodesRequestClass) Alloc() DetectBarcodesRequest {
	rv := objc.Send[DetectBarcodesRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DetectBarcodesRequest */
// A request that detects barcodes in an image.
//
// This request returns an array of objects, one for each barcode it detects.


// A request that detects barcodes in an image.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DetectBarcodesRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DetectBarcodesRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DetectBarcodesRequest */

// The array of barcode symbologies that the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectBarcodesRequest/supportedSymbologies
func (dc _DetectBarcodesRequestClass) SupportedSymbologies() []string {
	rv := objc.Send[[]string](objc.ID(dc.class), objc.Sel("supportedSymbologies"))
	return rv
}/* debug [class_properties_class/property]: supportedSymbologies */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DetectBarcodesRequest */

// Returns the barcode symbologies that the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectBarcodesRequest/supportedSymbologies()
func (d_ DetectBarcodesRequest) SupportedSymbologiesAndReturnError(error_ objectivec.IObject) []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("supportedSymbologiesAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: SupportedSymbologiesAndReturnError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DetectBarcodesRequest */

// A Boolean value that indicates whether to coalesce multiple codes based on the symbology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectBarcodesRequest/coalesceCompositeSymbologies
func (d_ DetectBarcodesRequest) CoalesceCompositeSymbologies() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("coalesceCompositeSymbologies"))
	return rv
}/* debug [instance_properties/getter]: coalesceCompositeSymbologies */


// A Boolean value that indicates whether to coalesce multiple codes based on the symbology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectBarcodesRequest/coalesceCompositeSymbologies
func (d_ DetectBarcodesRequest) SetCoalesceCompositeSymbologies(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCoalesceCompositeSymbologies:"), value)
}/* debug [instance_properties/setter]: coalesceCompositeSymbologies */


// The results of a barcode detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectBarcodesRequest/results
func (d_ DetectBarcodesRequest) Results() []BarcodeObservation {
	rv := objc.Send[[]BarcodeObservation](d_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// The array of barcode symbologies that the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectBarcodesRequest/supportedSymbologies
func (d_ DetectBarcodesRequest) SupportedSymbologies() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("supportedSymbologies"))
	return rv
}/* debug [instance_properties/getter]: supportedSymbologies */


// The barcode symbologies that the request detects in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectBarcodesRequest/symbologies
func (d_ DetectBarcodesRequest) Symbologies() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("symbologies"))
	return rv
}/* debug [instance_properties/getter]: symbologies */


// The barcode symbologies that the request detects in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectBarcodesRequest/symbologies
func (d_ DetectBarcodesRequest) SetSymbologies(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](d_.ID, objc.Sel("setSymbologies:"), nsArray)
}/* debug [instance_properties/setter]: symbologies */


// A constant for specifying revision 1 of the barcode detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequestrevision1
func (d_ DetectBarcodesRequest) VNDetectBarcodesRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectBarcodesRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNDetectBarcodesRequestRevision1 */


// A constant for specifying revision 2 of the barcode detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequestrevision2
func (d_ DetectBarcodesRequest) VNDetectBarcodesRequestRevision2() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectBarcodesRequestRevision2"))
	return rv
}/* debug [instance_properties/getter]: VNDetectBarcodesRequestRevision2 */


// A constant for specifying revision 3 of the barcode detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectbarcodesrequestrevision3
func (d_ DetectBarcodesRequest) VNDetectBarcodesRequestRevision3() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectBarcodesRequestRevision3"))
	return rv
}/* debug [instance_properties/getter]: VNDetectBarcodesRequestRevision3 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNDetectBarcodesRequest */



