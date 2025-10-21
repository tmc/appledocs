// Code generated from Apple documentation for QuickLookUI. DO NOT EDIT.

package quicklookui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// A class that you subclass to provide a data-based Quick Look preview extension.
//
// When you subclass , conform your subclass . To provide a data-based Quick Look extension, make the following modifications to your Info.plist file: Set the Boolean key to . Add the type identifiers for your extension’s supported content types to the array. Change the value of to the name of your subclass. For example, if you named your subclass , set the value to . After updating the extension’s file, implement the method to return a for the provided .
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewProvider
type PreviewProvider struct {
	objectivec.Object
}

// PreviewProviderFrom constructs a [PreviewProvider] from an unsafe.Pointer.
//
// A class that you subclass to provide a data-based Quick Look preview extension.
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




