// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PresentationIntent] class.
var (
	presentationIntentClass     _PresentationIntentClass
	presentationIntentClassOnce sync.Once
)

func getPresentationIntentClass() _PresentationIntentClass {
	presentationIntentClassOnce.Do(func() {
		presentationIntentClass = _PresentationIntentClass{objc.GetClass("NSPresentationIntent")}
	})
	return presentationIntentClass
}

type _PresentationIntentClass struct {
	class objc.Class
}

// An interface definition for the [PresentationIntent] class.
type IPresentationIntent interface {
	objectivec.IObject
}

// A type that contains the Markdown formatting for blocks of text, like paragraphs, lists, code blocks, and parts of tables.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent
type PresentationIntent struct {
	objectivec.Object
}

// PresentationIntentFrom constructs a [PresentationIntent] from an unsafe.Pointer.
//
// A type that contains the Markdown formatting for blocks of text, like paragraphs, lists, code blocks, and parts of tables.
func PresentationIntentFrom(ptr unsafe.Pointer) PresentationIntent {
	return PresentationIntent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PresentationIntentClass) Alloc() PresentationIntent {
	rv := objc.Send[PresentationIntent](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PresentationIntentClass) New() PresentationIntent {
	rv := objc.Send[PresentationIntent](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PresentationIntent) Init() PresentationIntent {
	rv := objc.Send[PresentationIntent](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PresentationIntent) Autorelease() PresentationIntent {
	rv := objc.Send[PresentationIntent](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPresentationIntent creates a new PresentationIntent instance.
func NewPresentationIntent() PresentationIntent {
	return getPresentationIntentClass().New()
}




