// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PreviewProvider] class.
var (
	PreviewProviderClass     _PreviewProviderClass
	PreviewProviderClassOnce sync.Once
)

func getPreviewProviderClass() _PreviewProviderClass {
	PreviewProviderClassOnce.Do(func() {
		PreviewProviderClass = _PreviewProviderClass{objc.GetClass("QLPreviewProvider")}
	})
	return PreviewProviderClass
}

type _PreviewProviderClass struct {
	class objc.Class
}

// An interface definition for the [PreviewProvider] class.
type IPreviewProvider interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewProvider
type PreviewProvider struct {
	objectivec.Object
}

// PreviewProviderFrom constructs a [PreviewProvider] from an unsafe.Pointer.
func PreviewProviderFrom(ptr unsafe.Pointer) PreviewProvider {
	return PreviewProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PreviewProviderClass) Alloc() PreviewProvider {
	rv := objc.Send[PreviewProvider](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PreviewProviderClass) New() PreviewProvider {
	rv := objc.Send[PreviewProvider](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewProvider) Init() PreviewProvider {
	rv := objc.Send[PreviewProvider](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewProvider) Autorelease() PreviewProvider {
	rv := objc.Send[PreviewProvider](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewProvider creates a new PreviewProvider instance.
func NewPreviewProvider() PreviewProvider {
	return getPreviewProviderClass().New()
}




