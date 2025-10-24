// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [QuickLookPreviewItem] class.
var (
	QuickLookPreviewItemClass     _QuickLookPreviewItemClass
	QuickLookPreviewItemClassOnce sync.Once
)

func getQuickLookPreviewItemClass() _QuickLookPreviewItemClass {
	QuickLookPreviewItemClassOnce.Do(func() {
		QuickLookPreviewItemClass = _QuickLookPreviewItemClass{objc.GetClass("ARQuickLookPreviewItem")}
	})
	return QuickLookPreviewItemClass
}

type _QuickLookPreviewItemClass struct {
	class objc.Class
}

// An interface definition for the [QuickLookPreviewItem] class.
type IQuickLookPreviewItem interface {
	objectivec.IObject
	// properties:
	AllowsContentScaling() bool
	SetAllowsContentScaling(value bool)
	CanonicalWebPageURL() objc.IObject /* cross-framework: URL */
	SetCanonicalWebPageURL(value objc.IObject /* cross-framework: URL */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/ARQuickLookPreviewItem
type QuickLookPreviewItem struct {
	objectivec.Object
}

// QuickLookPreviewItemFrom constructs a [QuickLookPreviewItem] from an unsafe.Pointer.
func QuickLookPreviewItemFrom(ptr unsafe.Pointer) QuickLookPreviewItem {
	return QuickLookPreviewItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (qc _QuickLookPreviewItemClass) Alloc() QuickLookPreviewItem {
	rv := objc.Send[QuickLookPreviewItem](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (qc _QuickLookPreviewItemClass) New() QuickLookPreviewItem {
	rv := objc.Send[QuickLookPreviewItem](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QuickLookPreviewItem) Init() QuickLookPreviewItem {
	rv := objc.Send[QuickLookPreviewItem](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QuickLookPreviewItem) Autorelease() QuickLookPreviewItem {
	rv := objc.Send[QuickLookPreviewItem](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQuickLookPreviewItem creates a new QuickLookPreviewItem instance.
func NewQuickLookPreviewItem() QuickLookPreviewItem {
	return getQuickLookPreviewItemClass().New()
}



// Creates an object representing the 3D content that will be previewed in AR Quick Look.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/ARQuickLookPreviewItem/init(fileAt:)
func NewQuickLookPreviewItemWithFileAtURL(url objc.IObject /* cross-framework: NSURL */) QuickLookPreviewItem {
	instance := getQuickLookPreviewItemClass().Alloc()
	rv := objc.Send[QuickLookPreviewItem](instance.ID, objc.Sel("initWithFileAtURL:"), url)
	rv.Autorelease()
	return rv
}



// Whether or not AR Quick Look allows content scaling in AR mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklook/arquicklookpreviewitem/allowscontentscaling
func (q_ QuickLookPreviewItem) AllowsContentScaling() bool {
	rv := objc.Send[bool](q_.ID, objc.Sel("allowsContentScaling"))
	return rv
}


// Whether or not AR Quick Look allows content scaling in AR mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklook/arquicklookpreviewitem/allowscontentscaling
func (q_ QuickLookPreviewItem) SetAllowsContentScaling(value bool) {
	objc.Send[objc.ID](q_.ID, objc.Sel("setAllowsContentScaling:"), value)
}


// An optional canonical web page URL for the 3D content that will be shared.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklook/arquicklookpreviewitem/canonicalwebpageurl
func (q_ QuickLookPreviewItem) CanonicalWebPageURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](q_.ID, objc.Sel("canonicalWebPageURL"))
	return rv
}


// An optional canonical web page URL for the 3D content that will be shared.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklook/arquicklookpreviewitem/canonicalwebpageurl
func (q_ QuickLookPreviewItem) SetCanonicalWebPageURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](q_.ID, objc.Sel("setCanonicalWebPageURL:"), value)
}


