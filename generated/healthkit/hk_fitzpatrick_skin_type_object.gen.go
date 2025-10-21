// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKFitzpatrickSkinTypeObject] class.
var (
	HKFitzpatrickSkinTypeObjectClass     _HKFitzpatrickSkinTypeObjectClass
	HKFitzpatrickSkinTypeObjectClassOnce sync.Once
)

func getHKFitzpatrickSkinTypeObjectClass() _HKFitzpatrickSkinTypeObjectClass {
	HKFitzpatrickSkinTypeObjectClassOnce.Do(func() {
		HKFitzpatrickSkinTypeObjectClass = _HKFitzpatrickSkinTypeObjectClass{objc.GetClass("HKFitzpatrickSkinTypeObject")}
	})
	return HKFitzpatrickSkinTypeObjectClass
}

type _HKFitzpatrickSkinTypeObjectClass struct {
	class objc.Class
}

// An interface definition for the [HKFitzpatrickSkinTypeObject] class.
type IHKFitzpatrickSkinTypeObject interface {
	objectivec.IObject
}

// This class acts as a wrapper for the enumeration.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFitzpatrickSkinTypeObject
type HKFitzpatrickSkinTypeObject struct {
	objectivec.Object
}

// HKFitzpatrickSkinTypeObjectFrom constructs a [HKFitzpatrickSkinTypeObject] from an unsafe.Pointer.
//
// This class acts as a wrapper for the enumeration.
func HKFitzpatrickSkinTypeObjectFrom(ptr unsafe.Pointer) HKFitzpatrickSkinTypeObject {
	return HKFitzpatrickSkinTypeObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKFitzpatrickSkinTypeObjectClass) Alloc() HKFitzpatrickSkinTypeObject {
	rv := objc.Send[HKFitzpatrickSkinTypeObject](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKFitzpatrickSkinTypeObjectClass) New() HKFitzpatrickSkinTypeObject {
	rv := objc.Send[HKFitzpatrickSkinTypeObject](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKFitzpatrickSkinTypeObject) Init() HKFitzpatrickSkinTypeObject {
	rv := objc.Send[HKFitzpatrickSkinTypeObject](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKFitzpatrickSkinTypeObject) Autorelease() HKFitzpatrickSkinTypeObject {
	rv := objc.Send[HKFitzpatrickSkinTypeObject](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKFitzpatrickSkinTypeObject creates a new HKFitzpatrickSkinTypeObject instance.
func NewHKFitzpatrickSkinTypeObject() HKFitzpatrickSkinTypeObject {
	return getHKFitzpatrickSkinTypeObjectClass().New()
}


// The user’s skin type.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfitzpatrickskintypeobject/skintype
func (h_ HKFitzpatrickSkinTypeObject) SkinType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("skinType"))
	return rv
}


// SetSkinType sets the value of the skinType property.
// The user’s skin type.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfitzpatrickskintypeobject/skintype
func (h_ HKFitzpatrickSkinTypeObject) SetSkinType(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSkinType:"), value)
}



