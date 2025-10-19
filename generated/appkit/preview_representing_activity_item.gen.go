// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PreviewRepresentingActivityItem] class.
var (
	previewRepresentingActivityItemClass     _PreviewRepresentingActivityItemClass
	previewRepresentingActivityItemClassOnce sync.Once
)

func getPreviewRepresentingActivityItemClass() _PreviewRepresentingActivityItemClass {
	previewRepresentingActivityItemClassOnce.Do(func() {
		previewRepresentingActivityItemClass = _PreviewRepresentingActivityItemClass{objc.GetClass("NSPreviewRepresentingActivityItem")}
	})
	return previewRepresentingActivityItemClass
}

type _PreviewRepresentingActivityItemClass struct {
	class objc.Class
}

// An interface definition for the [PreviewRepresentingActivityItem] class.
type IPreviewRepresentingActivityItem interface {
	objectivec.IObject
}

// A type that adds metadata to an item you share using the macOS share sheet.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentingActivityItem
type PreviewRepresentingActivityItem struct {
	objectivec.Object
}

// PreviewRepresentingActivityItemFrom constructs a [PreviewRepresentingActivityItem] from an unsafe.Pointer.
//
// A type that adds metadata to an item you share using the macOS share sheet.
func PreviewRepresentingActivityItemFrom(ptr unsafe.Pointer) PreviewRepresentingActivityItem {
	return PreviewRepresentingActivityItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PreviewRepresentingActivityItemClass) Alloc() PreviewRepresentingActivityItem {
	rv := objc.Send[PreviewRepresentingActivityItem](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PreviewRepresentingActivityItemClass) New() PreviewRepresentingActivityItem {
	rv := objc.Send[PreviewRepresentingActivityItem](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewRepresentingActivityItem) Init() PreviewRepresentingActivityItem {
	rv := objc.Send[PreviewRepresentingActivityItem](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewRepresentingActivityItem) Autorelease() PreviewRepresentingActivityItem {
	rv := objc.Send[PreviewRepresentingActivityItem](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewRepresentingActivityItem creates a new PreviewRepresentingActivityItem instance.
func NewPreviewRepresentingActivityItem() PreviewRepresentingActivityItem {
	return getPreviewRepresentingActivityItemClass().New()
}




