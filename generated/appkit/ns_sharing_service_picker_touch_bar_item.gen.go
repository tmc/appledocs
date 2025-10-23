// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	ButtonTitle() objc.IObject /* cross-framework: NSString */
	SetButtonTitle(value objc.IObject /* cross-framework: NSString */)
	ActivityItemsConfiguration() ActivityItemsConfigurationReading /* not a class type */
	SetActivityItemsConfiguration(value ActivityItemsConfigurationReading /* not a class type */)
	ButtonImage() IImage
	SetButtonImage(value IImage)
	Delegate() SharingServicePickerTouchBarItemDelegate /* not a class type */
	SetDelegate(value SharingServicePickerTouchBarItemDelegate /* not a class type */)
	IsEnabled() bool /* primitive/slice/pointer. */
	SetIsEnabled(value bool /* primitive/slice/pointer. */)
	// methods:
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



// The text displayed in the sharing service picker item button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/buttonTitle
func (s_ SharingServicePickerTouchBarItem) ButtonTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("buttonTitle"))
	return rv
}


// The text displayed in the sharing service picker item button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/buttonTitle
func (s_ SharingServicePickerTouchBarItem) SetButtonTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setButtonTitle:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/activityitemsconfiguration
func (s_ SharingServicePickerTouchBarItem) ActivityItemsConfiguration() ActivityItemsConfigurationReading /* not a class type */ {
	rv := objc.Send[ActivityItemsConfigurationReading](s_.ID, objc.Sel("activityItemsConfiguration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/activityitemsconfiguration
func (s_ SharingServicePickerTouchBarItem) SetActivityItemsConfiguration(value ActivityItemsConfigurationReading /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setActivityItemsConfiguration:"), value)
}


// The image displayed in the sharing service picker item button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/buttonimage
func (s_ SharingServicePickerTouchBarItem) ButtonImage() IImage {
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


// The object that acts as the delegate of the sharing service picker bar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/delegate
func (s_ SharingServicePickerTouchBarItem) Delegate() SharingServicePickerTouchBarItemDelegate /* not a class type */ {
	rv := objc.Send[SharingServicePickerTouchBarItemDelegate](s_.ID, objc.Sel("delegate"))
	return rv
}


// The object that acts as the delegate of the sharing service picker bar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/delegate
func (s_ SharingServicePickerTouchBarItem) SetDelegate(value SharingServicePickerTouchBarItemDelegate /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that specifies whether the sharing service picker item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/isenabled
func (s_ SharingServicePickerTouchBarItem) IsEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that specifies whether the sharing service picker item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/isenabled
func (s_ SharingServicePickerTouchBarItem) SetIsEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsEnabled:"), value)
}



