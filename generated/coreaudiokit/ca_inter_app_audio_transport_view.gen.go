// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [InterAppAudioTransportView] class.
var (
	InterAppAudioTransportViewClass     _InterAppAudioTransportViewClass
	InterAppAudioTransportViewClassOnce sync.Once
)

func getInterAppAudioTransportViewClass() _InterAppAudioTransportViewClass {
	InterAppAudioTransportViewClassOnce.Do(func() {
		InterAppAudioTransportViewClass = _InterAppAudioTransportViewClass{objc.GetClass("CAInterAppAudioTransportView")}
	})
	return InterAppAudioTransportViewClass
}

type _InterAppAudioTransportViewClass struct {
	class objc.Class
}

// An interface definition for the [InterAppAudioTransportView] class.
type IInterAppAudioTransportView interface {
	appkit.IView
}

// A view that provides an audio transport user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView
type InterAppAudioTransportView struct {
	appkit.View
}

// InterAppAudioTransportViewFrom constructs a [InterAppAudioTransportView] from an unsafe.Pointer.
//
// A view that provides an audio transport user interface.
func InterAppAudioTransportViewFrom(ptr unsafe.Pointer) InterAppAudioTransportView {
	return InterAppAudioTransportView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _InterAppAudioTransportViewClass) Alloc() InterAppAudioTransportView {
	rv := objc.Send[InterAppAudioTransportView](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InterAppAudioTransportViewClass) New() InterAppAudioTransportView {
	rv := objc.Send[InterAppAudioTransportView](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InterAppAudioTransportView) Init() InterAppAudioTransportView {
	rv := objc.Send[InterAppAudioTransportView](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InterAppAudioTransportView) Autorelease() InterAppAudioTransportView {
	rv := objc.Send[InterAppAudioTransportView](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInterAppAudioTransportView creates a new InterAppAudioTransportView instance.
func NewInterAppAudioTransportView() InterAppAudioTransportView {
	return getInterAppAudioTransportViewClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/currentTimeLabelFont
func (i_ InterAppAudioTransportView) CurrentTimeLabelFont() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("currentTimeLabelFont"))
	return rv
}


// SetCurrentTimeLabelFont sets the value of the currentTimeLabelFont property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/currentTimeLabelFont
func (i_ InterAppAudioTransportView) SetCurrentTimeLabelFont(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCurrentTimeLabelFont:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/isConnected
func (i_ InterAppAudioTransportView) Connected() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("connected"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/isEnabled
func (i_ InterAppAudioTransportView) Enabled() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("enabled"))
	return rv
}


// SetEnabled sets the value of the enabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/isEnabled
func (i_ InterAppAudioTransportView) SetEnabled(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEnabled:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/isPlaying
func (i_ InterAppAudioTransportView) Playing() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("playing"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/labelColor
func (i_ InterAppAudioTransportView) LabelColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("labelColor"))
	return rv
}


// SetLabelColor sets the value of the labelColor property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/labelColor
func (i_ InterAppAudioTransportView) SetLabelColor(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLabelColor:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/recordButtonColor
func (i_ InterAppAudioTransportView) RecordButtonColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("recordButtonColor"))
	return rv
}


// SetRecordButtonColor sets the value of the recordButtonColor property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/recordButtonColor
func (i_ InterAppAudioTransportView) SetRecordButtonColor(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRecordButtonColor:"), value)
}


