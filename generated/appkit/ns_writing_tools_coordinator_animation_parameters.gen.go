// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WritingToolsCoordinatorAnimationParameters] class.
var (
	WritingToolsCoordinatorAnimationParametersClass     _WritingToolsCoordinatorAnimationParametersClass
	WritingToolsCoordinatorAnimationParametersClassOnce sync.Once
)

func getWritingToolsCoordinatorAnimationParametersClass() _WritingToolsCoordinatorAnimationParametersClass {
	WritingToolsCoordinatorAnimationParametersClassOnce.Do(func() {
		WritingToolsCoordinatorAnimationParametersClass = _WritingToolsCoordinatorAnimationParametersClass{objc.GetClass("NSWritingToolsCoordinatorAnimationParameters")}
	})
	return WritingToolsCoordinatorAnimationParametersClass
}

type _WritingToolsCoordinatorAnimationParametersClass struct {
	class objc.Class
}

// An interface definition for the [WritingToolsCoordinatorAnimationParameters] class.
type IWritingToolsCoordinatorAnimationParameters interface {
	objectivec.IObject
}

// An object you use to configure additional tasks or animations to run alongside the Writing Tools animations.
//
// When Writing Tools replaces text in one of your context objects, it provides an object for you to use to configure any additional animations. During a Writing Tools session, you hide the text under evaluation and provide a targeted preview of your content. Writing Tools animations changes to that preview, but you might need to provide additional animations for other parts of your view’s content. For example, you might need to animate any layout changes caused by the insertion or removal of text in other parts of your view. Use this object to configure those animations. You don’t create an object directly. Instead, the system creates one and passes it to the method of your object. Use that object to specify the blocks to run during and after the system animations.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/AnimationParameters
type WritingToolsCoordinatorAnimationParameters struct {
	objectivec.Object
}

// WritingToolsCoordinatorAnimationParametersFrom constructs a [WritingToolsCoordinatorAnimationParameters] from an unsafe.Pointer.
//
// An object you use to configure additional tasks or animations to run alongside the Writing Tools animations.
func WritingToolsCoordinatorAnimationParametersFrom(ptr unsafe.Pointer) WritingToolsCoordinatorAnimationParameters {
	return WritingToolsCoordinatorAnimationParameters{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WritingToolsCoordinatorAnimationParametersClass) Alloc() WritingToolsCoordinatorAnimationParameters {
	rv := objc.Send[WritingToolsCoordinatorAnimationParameters](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WritingToolsCoordinatorAnimationParametersClass) New() WritingToolsCoordinatorAnimationParameters {
	rv := objc.Send[WritingToolsCoordinatorAnimationParameters](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WritingToolsCoordinatorAnimationParameters) Init() WritingToolsCoordinatorAnimationParameters {
	rv := objc.Send[WritingToolsCoordinatorAnimationParameters](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WritingToolsCoordinatorAnimationParameters) Autorelease() WritingToolsCoordinatorAnimationParameters {
	rv := objc.Send[WritingToolsCoordinatorAnimationParameters](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWritingToolsCoordinatorAnimationParameters creates a new WritingToolsCoordinatorAnimationParameters instance.
func NewWritingToolsCoordinatorAnimationParameters() WritingToolsCoordinatorAnimationParameters {
	return getWritingToolsCoordinatorAnimationParametersClass().New()
}
