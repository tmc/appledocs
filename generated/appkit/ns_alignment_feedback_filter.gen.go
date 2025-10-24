// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AlignmentFeedbackFilter] class.
var (
	AlignmentFeedbackFilterClass     _AlignmentFeedbackFilterClass
	AlignmentFeedbackFilterClassOnce sync.Once
)

func getAlignmentFeedbackFilterClass() _AlignmentFeedbackFilterClass {
	AlignmentFeedbackFilterClassOnce.Do(func() {
		AlignmentFeedbackFilterClass = _AlignmentFeedbackFilterClass{objc.GetClass("NSAlignmentFeedbackFilter")}
	})
	return AlignmentFeedbackFilterClass
}

type _AlignmentFeedbackFilterClass struct {
	class objc.Class
}

// An interface definition for the [AlignmentFeedbackFilter] class.
type IAlignmentFeedbackFilter interface {
	objectivec.IObject
	// properties:
	// methods:
}

// An object that can filter the movement of an object and provides haptic feedback when alignment occurs.
//
// With a Force Touch trackpad, apps can produce tactile feedback to complement user actions. If your app implements alignment features, you can use the class to filter object movements and provide haptic feedback to the user at appropriate times. As the user drags objects into alignment with a guide or another object, the user actually feels a physical bump as the object snaps into place.


// An object that can filter the movement of an object and provides haptic feedback when alignment occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter
type AlignmentFeedbackFilter struct {
	objectivec.Object
}

// AlignmentFeedbackFilterFrom constructs a [AlignmentFeedbackFilter] from an unsafe.Pointer.
//
// An object that can filter the movement of an object and provides haptic feedback when alignment occurs.
func AlignmentFeedbackFilterFrom(ptr unsafe.Pointer) AlignmentFeedbackFilter {
	return AlignmentFeedbackFilter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AlignmentFeedbackFilterClass) Alloc() AlignmentFeedbackFilter {
	rv := objc.Send[AlignmentFeedbackFilter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AlignmentFeedbackFilterClass) New() AlignmentFeedbackFilter {
	rv := objc.Send[AlignmentFeedbackFilter](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AlignmentFeedbackFilter) Init() AlignmentFeedbackFilter {
	rv := objc.Send[AlignmentFeedbackFilter](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AlignmentFeedbackFilter) Autorelease() AlignmentFeedbackFilter {
	rv := objc.Send[AlignmentFeedbackFilter](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAlignmentFeedbackFilter creates a new AlignmentFeedbackFilter instance.
func NewAlignmentFeedbackFilter() AlignmentFeedbackFilter {
	return getAlignmentFeedbackFilterClass().New()
}




