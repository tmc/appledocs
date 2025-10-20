// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StatusItem] class.
var (
	StatusItemClass     _StatusItemClass
	StatusItemClassOnce sync.Once
)

func getStatusItemClass() _StatusItemClass {
	StatusItemClassOnce.Do(func() {
		StatusItemClass = _StatusItemClass{objc.GetClass("NSStatusItem")}
	})
	return StatusItemClass
}

type _StatusItemClass struct {
	class objc.Class
}

// An interface definition for the [StatusItem] class.
type IStatusItem interface {
	objectivec.IObject
}

// An individual element displayed in the system menu bar.
//
// The method creates instances of this class and automatically adds them to the menu bar. Use the property to customize the appearance and behavior of the status item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem
type StatusItem struct {
	objectivec.Object
}

// StatusItemFrom constructs a [StatusItem] from an unsafe.Pointer.
//
// An individual element displayed in the system menu bar.
func StatusItemFrom(ptr unsafe.Pointer) StatusItem {
	return StatusItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StatusItemClass) Alloc() StatusItem {
	rv := objc.Send[StatusItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StatusItemClass) New() StatusItem {
	rv := objc.Send[StatusItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StatusItem) Init() StatusItem {
	rv := objc.Send[StatusItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StatusItem) Autorelease() StatusItem {
	rv := objc.Send[StatusItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStatusItem creates a new StatusItem instance.
func NewStatusItem() StatusItem {
	return getStatusItemClass().New()
}


// The alternate image to be displayed when a status bar item is highlighted.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/alternateImage
func (s_ StatusItem) AlternateImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("alternateImage"))
	return rv
}


// SetAlternateImage sets the value of the alternateImage property.
// The alternate image to be displayed when a status bar item is highlighted.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/alternateImage
func (s_ StatusItem) SetAlternateImage(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAlternateImage:"), value)
}


