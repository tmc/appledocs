// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [showFeedback] class.
var (
	ShowFeedbackClass     _showFeedbackClass
	ShowFeedbackClassOnce sync.Once
)

func getshowFeedbackClass() _showFeedbackClass {
	ShowFeedbackClassOnce.Do(func() {
		ShowFeedbackClass = _showFeedbackClass{objc.GetClass("showFeedback")}
	})
	return ShowFeedbackClass
}

type _showFeedbackClass struct {
	class objc.Class
}

// An interface definition for the [showFeedback] class.
type IshowFeedback interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/showFeedback
type showFeedback struct {
	objectivec.Object
}

// showFeedbackFrom constructs a [showFeedback] from an unsafe.Pointer.
func showFeedbackFrom(ptr unsafe.Pointer) showFeedback {
	return showFeedback{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _showFeedbackClass) Alloc() showFeedback {
	rv := objc.Send[showFeedback](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _showFeedbackClass) New() showFeedback {
	rv := objc.Send[showFeedback](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ showFeedback) Init() showFeedback {
	rv := objc.Send[showFeedback](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ showFeedback) Autorelease() showFeedback {
	rv := objc.Send[showFeedback](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewshowFeedback creates a new showFeedback instance.
func NewshowFeedback() showFeedback {
	return getshowFeedbackClass().New()
}




