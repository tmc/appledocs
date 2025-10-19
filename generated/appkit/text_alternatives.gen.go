// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextAlternatives] class.
var textAlternativesClass = _TextAlternativesClass{objc.GetClass("NSTextAlternatives")}

type _TextAlternativesClass struct {
	class objc.Class
}

// An interface definition for the [TextAlternatives] class.
type ITextAlternatives interface {
	objectivec.IObject
}

// A list of alternative strings for a piece of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAlternatives

type TextAlternatives struct {
	objectivec.Object
}

// TextAlternativesFrom constructs a [TextAlternatives] from an unsafe.Pointer.
//
// A list of alternative strings for a piece of text.
func TextAlternativesFrom(ptr unsafe.Pointer) TextAlternatives {
	return TextAlternatives{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (tc _TextAlternativesClass) Alloc() TextAlternatives {
	rv := objc.Send[TextAlternatives](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _TextAlternativesClass) New() TextAlternatives {
	rv := objc.Send[TextAlternatives](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextAlternatives) Init() TextAlternatives {
	rv := objc.Send[TextAlternatives](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextAlternatives) Autorelease() TextAlternatives {
	rv := objc.Send[TextAlternatives](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextAlternatives creates a new TextAlternatives instance.
func NewTextAlternatives() TextAlternatives {
	return textAlternativesClass.New()
}




