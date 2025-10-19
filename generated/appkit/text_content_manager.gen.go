// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextContentManager] class.
var (
	textContentManagerClass     _TextContentManagerClass
	textContentManagerClassOnce sync.Once
)

func getTextContentManagerClass() _TextContentManagerClass {
	textContentManagerClassOnce.Do(func() {
		textContentManagerClass = _TextContentManagerClass{objc.GetClass("NSTextContentManager")}
	})
	return textContentManagerClass
}

type _TextContentManagerClass struct {
	class objc.Class
}

// An interface definition for the [TextContentManager] class.
type ITextContentManager interface {
	objectivec.IObject
}

// An abstract class that defines the interface and a default implementation for managing the text document contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager

type TextContentManager struct {
	objectivec.Object
}

// TextContentManagerFrom constructs a [TextContentManager] from an unsafe.Pointer.
//
// An abstract class that defines the interface and a default implementation for managing the text document contents.
func TextContentManagerFrom(ptr unsafe.Pointer) TextContentManager {
	return TextContentManager{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (tc _TextContentManagerClass) Alloc() TextContentManager {
	rv := objc.Send[TextContentManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _TextContentManagerClass) New() TextContentManager {
	rv := objc.Send[TextContentManager](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextContentManager) Init() TextContentManager {
	rv := objc.Send[TextContentManager](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextContentManager) Autorelease() TextContentManager {
	rv := objc.Send[TextContentManager](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextContentManager creates a new TextContentManager instance.
func NewTextContentManager() TextContentManager {
	return getTextContentManagerClass().New()
}




