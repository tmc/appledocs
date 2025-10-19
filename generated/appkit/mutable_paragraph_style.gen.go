// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableParagraphStyle] class.
var (
	mutableParagraphStyleClass     _MutableParagraphStyleClass
	mutableParagraphStyleClassOnce sync.Once
)

func getMutableParagraphStyleClass() _MutableParagraphStyleClass {
	mutableParagraphStyleClassOnce.Do(func() {
		mutableParagraphStyleClass = _MutableParagraphStyleClass{objc.GetClass("NSMutableParagraphStyle")}
	})
	return mutableParagraphStyleClass
}

type _MutableParagraphStyleClass struct {
	class objc.Class
}

// An interface definition for the [MutableParagraphStyle] class.
type IMutableParagraphStyle interface {
	IParagraphStyle
}

// An object for changing the values of the subattributes in a paragraph style attribute.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle
type MutableParagraphStyle struct {
	ParagraphStyle
}

// MutableParagraphStyleFrom constructs a [MutableParagraphStyle] from an unsafe.Pointer.
//
// An object for changing the values of the subattributes in a paragraph style attribute.
func MutableParagraphStyleFrom(ptr unsafe.Pointer) MutableParagraphStyle {
	return MutableParagraphStyle{
		ParagraphStyle: ParagraphStyleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableParagraphStyleClass) Alloc() MutableParagraphStyle {
	rv := objc.Send[MutableParagraphStyle](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableParagraphStyleClass) New() MutableParagraphStyle {
	rv := objc.Send[MutableParagraphStyle](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableParagraphStyle) Init() MutableParagraphStyle {
	rv := objc.Send[MutableParagraphStyle](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableParagraphStyle) Autorelease() MutableParagraphStyle {
	rv := objc.Send[MutableParagraphStyle](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableParagraphStyle creates a new MutableParagraphStyle instance.
func NewMutableParagraphStyle() MutableParagraphStyle {
	return getMutableParagraphStyleClass().New()
}




