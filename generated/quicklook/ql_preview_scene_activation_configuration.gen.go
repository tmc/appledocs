// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PreviewSceneActivationConfiguration] class.
var (
	PreviewSceneActivationConfigurationClass     _PreviewSceneActivationConfigurationClass
	PreviewSceneActivationConfigurationClassOnce sync.Once
)

func getPreviewSceneActivationConfigurationClass() _PreviewSceneActivationConfigurationClass {
	PreviewSceneActivationConfigurationClassOnce.Do(func() {
		PreviewSceneActivationConfigurationClass = _PreviewSceneActivationConfigurationClass{objc.GetClass("QLPreviewSceneActivationConfiguration")}
	})
	return PreviewSceneActivationConfigurationClass
}

type _PreviewSceneActivationConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [PreviewSceneActivationConfiguration] class.
type IPreviewSceneActivationConfiguration interface {
	objectivec.IObject
}

// A scene configuration to preview items at the specified URLs.
//
// This class provides the configuration for a prominent scene presentation of a preview, either from a swipe gesture or a menu action. The user can detach the prominent Quick Look window and display it independently. To provide a preview from a swipe gesture, use an instance of this class with . To provide a preview from a menu action, use an instance of this class with .
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewSceneActivationConfiguration
type PreviewSceneActivationConfiguration struct {
	objectivec.Object
}

// PreviewSceneActivationConfigurationFrom constructs a [PreviewSceneActivationConfiguration] from an unsafe.Pointer.
//
// A scene configuration to preview items at the specified URLs.
func PreviewSceneActivationConfigurationFrom(ptr unsafe.Pointer) PreviewSceneActivationConfiguration {
	return PreviewSceneActivationConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PreviewSceneActivationConfigurationClass) Alloc() PreviewSceneActivationConfiguration {
	rv := objc.Send[PreviewSceneActivationConfiguration](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PreviewSceneActivationConfigurationClass) New() PreviewSceneActivationConfiguration {
	rv := objc.Send[PreviewSceneActivationConfiguration](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewSceneActivationConfiguration) Init() PreviewSceneActivationConfiguration {
	rv := objc.Send[PreviewSceneActivationConfiguration](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewSceneActivationConfiguration) Autorelease() PreviewSceneActivationConfiguration {
	rv := objc.Send[PreviewSceneActivationConfiguration](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewSceneActivationConfiguration creates a new PreviewSceneActivationConfiguration instance.
func NewPreviewSceneActivationConfiguration() PreviewSceneActivationConfiguration {
	return getPreviewSceneActivationConfigurationClass().New()
}




