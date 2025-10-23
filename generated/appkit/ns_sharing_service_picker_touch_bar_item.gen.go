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
	ActivityItemsConfiguration() objc.ID
	SetActivityItemsConfiguration(value objc.ID)
	Enabled() bool
	SetEnabled(value bool)
	ButtonImage() Image
	SetButtonImage(value IImage)
	ButtonTitle() string
	SetButtonTitle(value string)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	IsEnabled() bool
	SetIsEnabled(value bool)
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

// Alloc allocates a new instance without initialization.
func (sc _SharingServicePickerTouchBarItemClass) Alloc() SharingServicePickerTouchBarItem {
	rv := objc.Send[SharingServicePickerTouchBarItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/activityItemsConfiguration
func (s_ SharingServicePickerTouchBarItem) ActivityItemsConfiguration() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("activityItemsConfiguration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/activityItemsConfiguration
func (s_ SharingServicePickerTouchBarItem) SetActivityItemsConfiguration(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setActivityItemsConfiguration:"), value)
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


// The image displayed in the sharing service picker item button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/buttonimage
func (s_ SharingServicePickerTouchBarItem) ButtonImage() Image {
	rv := objc.Send[Image](s_.ID, objc.Sel("buttonImage"))
	return rv
}


// The image displayed in the sharing service picker item button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/buttonimage
func (s_ SharingServicePickerTouchBarItem) SetButtonImage(value IImage) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setButtonImage:"), value)
}


// The text displayed in the sharing service picker item button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/buttontitle
func (s_ SharingServicePickerTouchBarItem) ButtonTitle() string {
	rv := objc.Send[string](s_.ID, objc.Sel("buttonTitle"))
	return rv
}


// The text displayed in the sharing service picker item button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/buttontitle
func (s_ SharingServicePickerTouchBarItem) SetButtonTitle(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setButtonTitle:"), objc.String(value))
}


// The object that acts as the delegate of the sharing service picker bar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/delegate
func (s_ SharingServicePickerTouchBarItem) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}


// The object that acts as the delegate of the sharing service picker bar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/delegate
func (s_ SharingServicePickerTouchBarItem) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
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



