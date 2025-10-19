// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextContainer] class.
var (
	textContainerClass     _TextContainerClass
	textContainerClassOnce sync.Once
)

func getTextContainerClass() _TextContainerClass {
	textContainerClassOnce.Do(func() {
		textContainerClass = _TextContainerClass{objc.GetClass("NSTextContainer")}
	})
	return textContainerClass
}

type _TextContainerClass struct {
	class objc.Class
}

// An interface definition for the [TextContainer] class.
type ITextContainer interface {
	objectivec.IObject
}

// A region where text layout occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer

type TextContainer struct {
	objectivec.Object
}

// TextContainerFrom constructs a [TextContainer] from an unsafe.Pointer.
//
// A region where text layout occurs.
func TextContainerFrom(ptr unsafe.Pointer) TextContainer {
	return TextContainer{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (tc _TextContainerClass) Alloc() TextContainer {
	rv := objc.Send[TextContainer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextContainerClass) New() TextContainer {
	rv := objc.Send[TextContainer](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextContainer) Init() TextContainer {
	rv := objc.Send[TextContainer](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextContainer) Autorelease() TextContainer {
	rv := objc.Send[TextContainer](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextContainer creates a new TextContainer instance.
func NewTextContainer() TextContainer {
	return getTextContainerClass().New()
}




