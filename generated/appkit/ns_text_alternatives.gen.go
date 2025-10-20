// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextAlternatives] class.
var (
	TextAlternativesClass     _TextAlternativesClass
	TextAlternativesClassOnce sync.Once
)

func getTextAlternativesClass() _TextAlternativesClass {
	TextAlternativesClassOnce.Do(func() {
		TextAlternativesClass = _TextAlternativesClass{objc.GetClass("NSTextAlternatives")}
	})
	return TextAlternativesClass
}

type _TextAlternativesClass struct {
	class objc.Class
}

// An interface definition for the [TextAlternatives] class.
type ITextAlternatives interface {
	objectivec.IObject
}

// A list of alternative strings for a piece of text.
//
// is an immutable value class that stores a list of alternatives for a piece of text and communicates the user’s selection of an alternative via a notification to your app. To support dictation, for example, you might use to present a list of alternative interpretations for a word or phrase the user speaks. If the user chooses to replace the initial interpretation with an alternative, notifies you of the choice so that you can update the text appropriately. instances are attached to attributed strings as the value of a text attribute, .
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getTextAlternativesClass().New()
}


// The text that was initially chosen as the input string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAlternatives/primaryString
func (t_ TextAlternatives) PrimaryString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("primaryString"))
	return rv
}



