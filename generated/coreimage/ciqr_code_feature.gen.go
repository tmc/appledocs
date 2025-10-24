// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [QRCodeFeature] class.
var (
	QRCodeFeatureClass     _QRCodeFeatureClass
	QRCodeFeatureClassOnce sync.Once
)

func getQRCodeFeatureClass() _QRCodeFeatureClass {
	QRCodeFeatureClassOnce.Do(func() {
		QRCodeFeatureClass = _QRCodeFeatureClass{objc.GetClass("CIQRCodeFeature")}
	})
	return QRCodeFeatureClass
}

type _QRCodeFeatureClass struct {
	class objc.Class
}





// An interface definition for the [QRCodeFeature] class.
type IQRCodeFeature interface {
	IFeature
	

	// properties:
	BottomLeft() corefoundation.CGPoint
	BottomRight() corefoundation.CGPoint
	Bounds() corefoundation.CGRect
	MessageString() objc.IObject /* cross-framework: NSString */
	SymbolDescriptor() ICIQRCodeDescriptor
	TopLeft() corefoundation.CGPoint
	TopRight() corefoundation.CGPoint
	CIDetectorTypeQRCode() objc.IObject /* cross-framework: NSString */


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (qc _QRCodeFeatureClass) Alloc() QRCodeFeature {
	rv := objc.Send[QRCodeFeature](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (qc _QRCodeFeatureClass) New() QRCodeFeature {
	rv := objc.Send[QRCodeFeature](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QRCodeFeature) Init() QRCodeFeature {
	rv := objc.Send[QRCodeFeature](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QRCodeFeature) Autorelease() QRCodeFeature {
	rv := objc.Send[QRCodeFeature](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQRCodeFeature creates a new QRCodeFeature instance.
func NewQRCodeFeature() QRCodeFeature {
	return getQRCodeFeatureClass().New()
}





// Information about a Quick Response code detected in a still or video image.
//
// A QR code is a two-dimensional barcode using the ISO/IEC 18004:2006 standard. The properties of a CIQRCodeFeature object identify the corners of the barcode in the image perspective and provide the decoded message. To detect QR codes in an image or video, choose type when initializing a object.


// Information about a Quick Response code detected in a still or video image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature
type QRCodeFeature struct {
	Feature
}

// QRCodeFeatureFrom constructs a [QRCodeFeature] from an unsafe.Pointer.
//
// Information about a Quick Response code detected in a still or video image.
func QRCodeFeatureFrom(ptr unsafe.Pointer) QRCodeFeature {
	return QRCodeFeature{
		Feature: FeatureFrom(ptr),
	}
}

























// The image coordinate of the lower-left corner of the detected QR code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/bottomLeft-swift.property
func (q_ QRCodeFeature) BottomLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](q_.ID, objc.Sel("bottomLeft"))
	return rv
}


// The image coordinate of the lower-right corner of the detected QR code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/bottomRight-swift.property
func (q_ QRCodeFeature) BottomRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](q_.ID, objc.Sel("bottomRight"))
	return rv
}


// A rectangle that indicates the position and extent of the QR code feature in image coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/bounds-swift.property
func (q_ QRCodeFeature) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](q_.ID, objc.Sel("bounds"))
	return rv
}


// The string decoded from the detected barcode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/messageString
func (q_ QRCodeFeature) MessageString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](q_.ID, objc.Sel("messageString"))
	return rv
}


// An abstract representation of a QR Code symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/symbolDescriptor-swift.property
func (q_ QRCodeFeature) SymbolDescriptor() ICIQRCodeDescriptor {
	rv := objc.Send[QRCodeDescriptor](q_.ID, objc.Sel("symbolDescriptor"))
	return rv
}


// The image coordinate of the upper-left corner of the detected QR code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/topLeft-swift.property
func (q_ QRCodeFeature) TopLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](q_.ID, objc.Sel("topLeft"))
	return rv
}


// The image coordinate of the upper-right corner of the detected QR code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/topRight-swift.property
func (q_ QRCodeFeature) TopRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](q_.ID, objc.Sel("topRight"))
	return rv
}


// A detector that searches for Quick Response codes (a type of 2D barcode) in a still image or video, returning
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidetectortypeqrcode
func (q_ QRCodeFeature) CIDetectorTypeQRCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](q_.ID, objc.Sel("CIDetectorTypeQRCode"))
	return rv
}








