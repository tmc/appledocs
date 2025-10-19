// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ParagraphStyle] class.
var paragraphStyleClass = _ParagraphStyleClass{objc.GetClass("NSParagraphStyle")}

type _ParagraphStyleClass struct {
	class objc.Class
}

// An interface definition for the [ParagraphStyle] class.
type IParagraphStyle interface {
	objectivec.IObject
}

// The paragraph or ruler attributes for an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle

type ParagraphStyle struct {
	objectivec.Object
}

// ParagraphStyleFrom constructs a [ParagraphStyle] from an unsafe.Pointer.
//
// The paragraph or ruler attributes for an attributed string.
func ParagraphStyleFrom(ptr unsafe.Pointer) ParagraphStyle {
	return ParagraphStyle{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (pc _ParagraphStyleClass) Alloc() ParagraphStyle {
	rv := objc.Send[ParagraphStyle](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _ParagraphStyleClass) New() ParagraphStyle {
	rv := objc.Send[ParagraphStyle](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ParagraphStyle) Init() ParagraphStyle {
	rv := objc.Send[ParagraphStyle](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ParagraphStyle) Autorelease() ParagraphStyle {
	rv := objc.Send[ParagraphStyle](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewParagraphStyle creates a new ParagraphStyle instance.
func NewParagraphStyle() ParagraphStyle {
	return paragraphStyleClass.New()
}




