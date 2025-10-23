// Code generated from Apple documentation for LinkPresentation. DO NOT EDIT.

package linkpresentation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [LPLinkView] class.
var (
	LPLinkViewClass     _LPLinkViewClass
	LPLinkViewClassOnce sync.Once
)

func getLPLinkViewClass() _LPLinkViewClass {
	LPLinkViewClassOnce.Do(func() {
		LPLinkViewClass = _LPLinkViewClass{objc.GetClass("LPLinkView")}
	})
	return LPLinkViewClass
}

type _LPLinkViewClass struct {
	class objc.Class
}

// An interface definition for the [LPLinkView] class.
type ILPLinkView interface {
	appkit.IView
	// properties:
	Metadata() ILPLinkMetadata
	SetMetadata(value ILPLinkMetadata)
	// methods:
}

// A rich visual representation of a link.
//
// presents a link based on its available metadata. Use it to show a link’s title and icon, associated images, inline audio, video playback, and maps in a familiar and consistent style.


// A rich visual representation of a link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkView
type LPLinkView struct {
	appkit.View
}

// LPLinkViewFrom constructs a [LPLinkView] from an unsafe.Pointer.
//
// A rich visual representation of a link.
func LPLinkViewFrom(ptr unsafe.Pointer) LPLinkView {
	return LPLinkView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (lc _LPLinkViewClass) Alloc() LPLinkView {
	rv := objc.Send[LPLinkView](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LPLinkViewClass) New() LPLinkView {
	rv := objc.Send[LPLinkView](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LPLinkView) Init() LPLinkView {
	rv := objc.Send[LPLinkView](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LPLinkView) Autorelease() LPLinkView {
	rv := objc.Send[LPLinkView](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLPLinkView creates a new LPLinkView instance.
func NewLPLinkView() LPLinkView {
	return getLPLinkViewClass().New()
}



// Initializes a link view with specified metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkView/init(metadata:)
func NewLPLinkViewWithMetadata(metadata ILPLinkMetadata) LPLinkView {
	instance := getLPLinkViewClass().Alloc()
	rv := objc.Send[LPLinkView](instance.ID, objc.Sel("initWithMetadata:"), metadata)
	rv.Autorelease()
	return rv
}


// Initializes a placeholder link view without metadata for a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkView/init(url:)
func NewLPLinkViewWithURL(URL foundation.URL) LPLinkView {
	instance := getLPLinkViewClass().Alloc()
	rv := objc.Send[LPLinkView](instance.ID, objc.Sel("initWithURL:"), URL)
	rv.Autorelease()
	return rv
}



// The metadata from which to generate a rich presentation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkView/metadata
func (l_ LPLinkView) Metadata() ILPLinkMetadata {
	rv := objc.Send[LPLinkMetadata](l_.ID, objc.Sel("metadata"))
	return rv
}


// The metadata from which to generate a rich presentation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPLinkView/metadata
func (l_ LPLinkView) SetMetadata(value ILPLinkMetadata) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMetadata:"), value)
}


