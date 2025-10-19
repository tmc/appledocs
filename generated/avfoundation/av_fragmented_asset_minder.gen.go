// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVFragmentedAssetMinder] class.
var aVFragmentedAssetMinderClass = _AVFragmentedAssetMinderClass{objc.GetClass("AVFragmentedAssetMinder")}

type _AVFragmentedAssetMinderClass struct {
	class objc.Class
}

// An interface definition for the [AVFragmentedAssetMinder] class.
type IAVFragmentedAssetMinder interface {
	objectivec.IObject
}

// A parent class referenced by other AVFoundation classes. [Full Topic]

type AVFragmentedAssetMinder struct {
	objectivec.Object
}

// AVFragmentedAssetMinderFrom constructs a [AVFragmentedAssetMinder] from an unsafe.Pointer.
//
// A parent class referenced by other AVFoundation classes.
func AVFragmentedAssetMinderFrom(ptr unsafe.Pointer) AVFragmentedAssetMinder {
	return AVFragmentedAssetMinder{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ac _AVFragmentedAssetMinderClass) Alloc() AVFragmentedAssetMinder {
	rv := objc.Send[AVFragmentedAssetMinder](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVFragmentedAssetMinderClass) New() AVFragmentedAssetMinder {
	rv := objc.Send[AVFragmentedAssetMinder](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVFragmentedAssetMinder) Init() AVFragmentedAssetMinder {
	rv := objc.Send[AVFragmentedAssetMinder](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVFragmentedAssetMinder) Autorelease() AVFragmentedAssetMinder {
	rv := objc.Send[AVFragmentedAssetMinder](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVFragmentedAssetMinder creates a new AVFragmentedAssetMinder instance.
func NewAVFragmentedAssetMinder() AVFragmentedAssetMinder {
	return aVFragmentedAssetMinderClass.New()
}




