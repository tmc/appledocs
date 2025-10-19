// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FaceFeature] class.
var faceFeatureClass = _FaceFeatureClass{objc.GetClass("CIFaceFeature")}

type _FaceFeatureClass struct {
	class objc.Class
}

// An interface definition for the [FaceFeature] class.
type IFaceFeature interface {
	IFeature
}

// Information about a face detected in a still or video image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature

type FaceFeature struct {
	Feature
}

// FaceFeatureFrom constructs a [FaceFeature] from an unsafe.Pointer.
//
// Information about a face detected in a still or video image.
func FaceFeatureFrom(ptr unsafe.Pointer) FaceFeature {
	return FaceFeature{
		Feature: FeatureFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (fc _FaceFeatureClass) Alloc() FaceFeature {
	rv := objc.Send[FaceFeature](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (fc _FaceFeatureClass) New() FaceFeature {
	rv := objc.Send[FaceFeature](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FaceFeature) Init() FaceFeature {
	rv := objc.Send[FaceFeature](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FaceFeature) Autorelease() FaceFeature {
	rv := objc.Send[FaceFeature](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFaceFeature creates a new FaceFeature instance.
func NewFaceFeature() FaceFeature {
	return faceFeatureClass.New()
}




