// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RectangleFeature] class.
var rectangleFeatureClass = _RectangleFeatureClass{objc.GetClass("CIRectangleFeature")}

type _RectangleFeatureClass struct {
	class objc.Class
}

// An interface definition for the [RectangleFeature] class.
type IRectangleFeature interface {
	IFeature
}

// Information about a rectangular region detected in a still or video image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature

type RectangleFeature struct {
	Feature
}

// RectangleFeatureFrom constructs a [RectangleFeature] from an unsafe.Pointer.
//
// Information about a rectangular region detected in a still or video image.
func RectangleFeatureFrom(ptr unsafe.Pointer) RectangleFeature {
	return RectangleFeature{
		Feature: FeatureFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (rc _RectangleFeatureClass) Alloc() RectangleFeature {
	rv := objc.Send[RectangleFeature](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (rc _RectangleFeatureClass) New() RectangleFeature {
	rv := objc.Send[RectangleFeature](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RectangleFeature) Init() RectangleFeature {
	rv := objc.Send[RectangleFeature](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RectangleFeature) Autorelease() RectangleFeature {
	rv := objc.Send[RectangleFeature](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRectangleFeature creates a new RectangleFeature instance.
func NewRectangleFeature() RectangleFeature {
	return rectangleFeatureClass.New()
}




