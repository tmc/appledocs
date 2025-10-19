// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Feature] class.
var (
	featureClass     _FeatureClass
	featureClassOnce sync.Once
)

func getFeatureClass() _FeatureClass {
	featureClassOnce.Do(func() {
		featureClass = _FeatureClass{objc.GetClass("CIFeature")}
	})
	return featureClass
}

type _FeatureClass struct {
	class objc.Class
}

// An interface definition for the [Feature] class.
type IFeature interface {
	objectivec.IObject
}

// The abstract superclass for objects representing notable features detected in an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFeature
type Feature struct {
	objectivec.Object
}

// FeatureFrom constructs a [Feature] from an unsafe.Pointer.
//
// The abstract superclass for objects representing notable features detected in an image.
func FeatureFrom(ptr unsafe.Pointer) Feature {
	return Feature{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FeatureClass) Alloc() Feature {
	rv := objc.Send[Feature](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FeatureClass) New() Feature {
	rv := objc.Send[Feature](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ Feature) Init() Feature {
	rv := objc.Send[Feature](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ Feature) Autorelease() Feature {
	rv := objc.Send[Feature](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFeature creates a new Feature instance.
func NewFeature() Feature {
	return getFeatureClass().New()
}




