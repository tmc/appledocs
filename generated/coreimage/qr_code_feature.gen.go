// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [QRCodeFeature] class.
var qRCodeFeatureClass = _QRCodeFeatureClass{objc.GetClass("CIQRCodeFeature")}

type _QRCodeFeatureClass struct {
	class objc.Class
}

// An interface definition for the [QRCodeFeature] class.
type IQRCodeFeature interface {
	IFeature
}

// Information about a Quick Response code detected in a still or video image. [Full Topic]
//
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
// Alloc allocates a new instance without initialization.
func (qc _QRCodeFeatureClass) Alloc() QRCodeFeature {
	rv := objc.Send[QRCodeFeature](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
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
	return qRCodeFeatureClass.New()
}




