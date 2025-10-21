// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [TextContainer] class.
var (
	TextContainerClass     _TextContainerClass
	TextContainerClassOnce sync.Once
)

func getTextContainerClass() _TextContainerClass {
	TextContainerClassOnce.Do(func() {
		TextContainerClass = _TextContainerClass{objc.GetClass("NSTextContainer")}
	})
	return TextContainerClass
}

type _TextContainerClass struct {
	class objc.Class
}

// An interface definition for the [TextContainer] class.
type ITextContainer interface {
	objectivec.IObject
}

// A region where text layout occurs.
//
// An uses to determine where to break lines, lay out portions of text, and so on. An object typically defines rectangular regions, but you can define exclusion paths inside the text container to create regions where text doesn’t flow. You can also subclass to create text containers with nonrectangular regions, such as circular regions, regions with holes in them, or regions that flow alongside graphics. You can access instances of the , , and classes from threads other than the main thread as long as the app guarantees access from only one thread at a time.
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


// The text container’s layout manager.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/layoutManager
func (t_ TextContainer) LayoutManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("layoutManager"))
	return rv
}


// SetLayoutManager sets the value of the layoutManager property.
// The text container’s layout manager.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/layoutManager
func (t_ TextContainer) SetLayoutManager(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutManager:"), value)
}



