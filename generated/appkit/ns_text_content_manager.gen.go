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
	TextContentManagerClass     _TextContentManagerClass
	TextContentManagerClassOnce sync.Once
)

func getTextContentManagerClass() _TextContentManagerClass {
	TextContentManagerClassOnce.Do(func() {
		TextContentManagerClass = _TextContentManagerClass{objc.GetClass("NSTextContentManager")}
	})
	return TextContentManagerClass
}

type _TextContentManagerClass struct {
	class objc.Class
}

// An interface definition for the [TextContentManager] class.
type ITextContentManager interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other AppKit classes.


// A parent class referenced by other AppKit classes. [Full Topic]
type TextContentManager struct {
	objectivec.Object
}

// TextContentManagerFrom constructs a [TextContentManager] from an unsafe.Pointer.
//
// A parent class referenced by other AppKit classes.
func TextContentManagerFrom(ptr unsafe.Pointer) TextContentManager {
	return TextContentManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextContentManagerClass) Alloc() TextContentManager {
	rv := objc.Send[TextContentManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




