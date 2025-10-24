// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHProjectTextElement] class.
var (
	PHProjectTextElementClass     _PHProjectTextElementClass
	PHProjectTextElementClassOnce sync.Once
)

func getPHProjectTextElementClass() _PHProjectTextElementClass {
	PHProjectTextElementClassOnce.Do(func() {
		PHProjectTextElementClass = _PHProjectTextElementClass{objc.GetClass("PHProjectTextElement")}
	})
	return PHProjectTextElementClass
}

type _PHProjectTextElementClass struct {
	class objc.Class
}

// An interface definition for the [PHProjectTextElement] class.
type IPHProjectTextElement interface {
	IPHProjectElement
	AttributedText() foundation.AttributedString
	Text() string
	TextElementType() PHProjectTextElementType
}

// An element that represents text within project section content.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectTextElement
type PHProjectTextElement struct {
	PHProjectElement
}

// PHProjectTextElementFrom constructs a [PHProjectTextElement] from an unsafe.Pointer.
//
// An element that represents text within project section content.
func PHProjectTextElementFrom(ptr unsafe.Pointer) PHProjectTextElement {
	return PHProjectTextElement{
		PHProjectElement: PHProjectElementFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHProjectTextElementClass) Alloc() PHProjectTextElement {
	rv := objc.Send[PHProjectTextElement](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHProjectTextElementClass) New() PHProjectTextElement {
	rv := objc.Send[PHProjectTextElement](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHProjectTextElement) Init() PHProjectTextElement {
	rv := objc.Send[PHProjectTextElement](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHProjectTextElement) Autorelease() PHProjectTextElement {
	rv := objc.Send[PHProjectTextElement](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHProjectTextElement creates a new PHProjectTextElement instance.
func NewPHProjectTextElement() PHProjectTextElement {
	return getPHProjectTextElementClass().New()
}

// The stylized attributed string for the text element as presented to the user in Photos.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectTextElement/attributedText
func (p_ PHProjectTextElement) AttributedText() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](p_.ID, objc.Sel("attributedText"))
	return rv
}

// The raw unformatted string for the text element.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectTextElement/text
func (p_ PHProjectTextElement) Text() string {
	rv := objc.Send[string](p_.ID, objc.Sel("text"))
	return rv
}

// The enumerated type of the text element.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectTextElement/textElementType
func (p_ PHProjectTextElement) TextElementType() PHProjectTextElementType {
	rv := objc.Send[PHProjectTextElementType](p_.ID, objc.Sel("textElementType"))
	return rv
}
