// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PreviewSceneOptions] class.
var (
	PreviewSceneOptionsClass     _PreviewSceneOptionsClass
	PreviewSceneOptionsClassOnce sync.Once
)

func getPreviewSceneOptionsClass() _PreviewSceneOptionsClass {
	PreviewSceneOptionsClassOnce.Do(func() {
		PreviewSceneOptionsClass = _PreviewSceneOptionsClass{objc.GetClass("QLPreviewSceneOptions")}
	})
	return PreviewSceneOptionsClass
}

type _PreviewSceneOptionsClass struct {
	class objc.Class
}

// An interface definition for the [PreviewSceneOptions] class.
type IPreviewSceneOptions interface {
	objectivec.IObject
}

// A class that represents the configuration for a preview scene activation.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewSceneActivationConfiguration/Options
type PreviewSceneOptions struct {
	objectivec.Object
}

// PreviewSceneOptionsFrom constructs a [PreviewSceneOptions] from an unsafe.Pointer.
//
// A class that represents the configuration for a preview scene activation.
func PreviewSceneOptionsFrom(ptr unsafe.Pointer) PreviewSceneOptions {
	return PreviewSceneOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PreviewSceneOptionsClass) Alloc() PreviewSceneOptions {
	rv := objc.Send[PreviewSceneOptions](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PreviewSceneOptionsClass) New() PreviewSceneOptions {
	rv := objc.Send[PreviewSceneOptions](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewSceneOptions) Init() PreviewSceneOptions {
	rv := objc.Send[PreviewSceneOptions](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewSceneOptions) Autorelease() PreviewSceneOptions {
	rv := objc.Send[PreviewSceneOptions](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewSceneOptions creates a new PreviewSceneOptions instance.
func NewPreviewSceneOptions() PreviewSceneOptions {
	return getPreviewSceneOptionsClass().New()
}


// The index of the item to preview.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewSceneActivationConfiguration/Options/initialPreviewIndex
func (p_ PreviewSceneOptions) InitialPreviewIndex() int {
	rv := objc.Send[int](p_.ID, objc.Sel("initialPreviewIndex"))
	return rv
}


// SetInitialPreviewIndex sets the value of the initialPreviewIndex property.
// The index of the item to preview.

//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewSceneActivationConfiguration/Options/initialPreviewIndex
func (p_ PreviewSceneOptions) SetInitialPreviewIndex(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInitialPreviewIndex:"), value)
}


