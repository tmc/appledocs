// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [SharingServicePickerTouchBarItem] class.
var (
	SharingServicePickerTouchBarItemClass     _SharingServicePickerTouchBarItemClass
	SharingServicePickerTouchBarItemClassOnce sync.Once
)

func getSharingServicePickerTouchBarItemClass() _SharingServicePickerTouchBarItemClass {
	SharingServicePickerTouchBarItemClassOnce.Do(func() {
		SharingServicePickerTouchBarItemClass = _SharingServicePickerTouchBarItemClass{objc.GetClass("NSSharingServicePickerTouchBarItem")}
	})
	return SharingServicePickerTouchBarItemClass
}

type _SharingServicePickerTouchBarItemClass struct {
	class objc.Class
}





// An interface definition for the [SharingServicePickerTouchBarItem] class.
type ISharingServicePickerTouchBarItem interface {
	ITouchBarItem
	

	// properties:
	ButtonImage() IImage
	SetButtonImage(value IImage)
	ButtonTitle() foundation.foundation.INSString
	SetButtonTitle(value foundation.foundation.INSString)
	Enabled() bool
	SetEnabled(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _SharingServicePickerTouchBarItemClass) Alloc() SharingServicePickerTouchBarItem {
	rv := objc.Send[SharingServicePickerTouchBarItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SharingServicePickerTouchBarItemClass) New() SharingServicePickerTouchBarItem {
	rv := objc.Send[SharingServicePickerTouchBarItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SharingServicePickerTouchBarItem) Init() SharingServicePickerTouchBarItem {
	rv := objc.Send[SharingServicePickerTouchBarItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SharingServicePickerTouchBarItem) Autorelease() SharingServicePickerTouchBarItem {
	rv := objc.Send[SharingServicePickerTouchBarItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSharingServicePickerTouchBarItem creates a new SharingServicePickerTouchBarItem instance.
func NewSharingServicePickerTouchBarItem() SharingServicePickerTouchBarItem {
	return getSharingServicePickerTouchBarItemClass().New()
}





// A bar item that, along with its delegate, provides a list of objects eligible for sharing.


// A bar item that, along with its delegate, provides a list of objects eligible for sharing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem
type SharingServicePickerTouchBarItem struct {
	TouchBarItem
}

// SharingServicePickerTouchBarItemFrom constructs a [SharingServicePickerTouchBarItem] from an unsafe.Pointer.
//
// A bar item that, along with its delegate, provides a list of objects eligible for sharing.
func SharingServicePickerTouchBarItemFrom(ptr unsafe.Pointer) SharingServicePickerTouchBarItem {
	return SharingServicePickerTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

























// The image displayed in the sharing service picker item button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/buttonImage
func (s_ SharingServicePickerTouchBarItem) ButtonImage() IImage {
	rv := objc.Send[Image](s_.ID, objc.Sel("buttonImage"))
	return rv
}


// The image displayed in the sharing service picker item button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/buttonImage
func (s_ SharingServicePickerTouchBarItem) SetButtonImage(value IImage) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setButtonImage:"), value)
}


// The text displayed in the sharing service picker item button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/buttonTitle
func (s_ SharingServicePickerTouchBarItem) ButtonTitle() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("buttonTitle"))
	return rv
}


// The text displayed in the sharing service picker item button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/buttonTitle
func (s_ SharingServicePickerTouchBarItem) SetButtonTitle(value foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setButtonTitle:"), value)
}


// A Boolean value that specifies whether the sharing service picker item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/isEnabled
func (s_ SharingServicePickerTouchBarItem) Enabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("enabled"))
	return rv
}


// A Boolean value that specifies whether the sharing service picker item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/isEnabled
func (s_ SharingServicePickerTouchBarItem) SetEnabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEnabled:"), value)
}


// A Boolean value that specifies whether the sharing service picker item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/isenabled
func (s_ SharingServicePickerTouchBarItem) IsEnabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that specifies whether the sharing service picker item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/isenabled
func (s_ SharingServicePickerTouchBarItem) SetIsEnabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsEnabled:"), value)
}







