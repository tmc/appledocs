// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFSafariViewControllerActivityButton] class.
var (
	SFSafariViewControllerActivityButtonClass     _SFSafariViewControllerActivityButtonClass
	SFSafariViewControllerActivityButtonClassOnce sync.Once
)

func getSFSafariViewControllerActivityButtonClass() _SFSafariViewControllerActivityButtonClass {
	SFSafariViewControllerActivityButtonClassOnce.Do(func() {
		SFSafariViewControllerActivityButtonClass = _SFSafariViewControllerActivityButtonClass{objc.GetClass("SFSafariViewControllerActivityButton")}
	})
	return SFSafariViewControllerActivityButtonClass
}

type _SFSafariViewControllerActivityButtonClass struct {
	class objc.Class
}

// An interface definition for the [SFSafariViewControllerActivityButton] class.
type ISFSafariViewControllerActivityButton interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/ActivityButton
type SFSafariViewControllerActivityButton struct {
	objectivec.Object
}

// SFSafariViewControllerActivityButtonFrom constructs a [SFSafariViewControllerActivityButton] from an unsafe.Pointer.
func SFSafariViewControllerActivityButtonFrom(ptr unsafe.Pointer) SFSafariViewControllerActivityButton {
	return SFSafariViewControllerActivityButton{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFSafariViewControllerActivityButtonClass) Alloc() SFSafariViewControllerActivityButton {
	rv := objc.Send[SFSafariViewControllerActivityButton](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFSafariViewControllerActivityButtonClass) New() SFSafariViewControllerActivityButton {
	rv := objc.Send[SFSafariViewControllerActivityButton](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariViewControllerActivityButton) Init() SFSafariViewControllerActivityButton {
	rv := objc.Send[SFSafariViewControllerActivityButton](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariViewControllerActivityButton) Autorelease() SFSafariViewControllerActivityButton {
	rv := objc.Send[SFSafariViewControllerActivityButton](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariViewControllerActivityButton creates a new SFSafariViewControllerActivityButton instance.
func NewSFSafariViewControllerActivityButton() SFSafariViewControllerActivityButton {
	return getSFSafariViewControllerActivityButtonClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/ActivityButton/init(templateImage:extensionIdentifier:)
func NewSFSafariViewControllerActivityButtonWithTemplateImageExtensionIdentifier(templateImage appkit.IImage, extensionIdentifier appkit.string) SFSafariViewControllerActivityButton {
	instance := getSFSafariViewControllerActivityButtonClass().Alloc()
	rv := objc.Send[SFSafariViewControllerActivityButton](instance.ID, objc.Sel("initWithTemplateImage:extensionIdentifier:"), templateImage, extensionIdentifier)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/ActivityButton/extensionIdentifier
func (s_ SFSafariViewControllerActivityButton) ExtensionIdentifier() appkit.string {
	rv := objc.Send[appkit.string](s_.ID, objc.Sel("extensionIdentifier"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfsafariviewcontroller/activitybutton/templateimage
func (s_ SFSafariViewControllerActivityButton) TemplateImage() appkit.Image {
	rv := objc.Send[appkit.Image](s_.ID, objc.Sel("templateImage"))
	return rv
}


// SetTemplateImage sets the value of the templateImage property.
//
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfsafariviewcontroller/activitybutton/templateimage
func (s_ SFSafariViewControllerActivityButton) SetTemplateImage(value appkit.IImage) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTemplateImage:"), value)
}


