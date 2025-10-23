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
	Connected() bool
	CurrentTimeLabelFont() appkit.Font
	SetCurrentTimeLabelFont(value appkit.Font)
	IsConnected() bool
	SetIsConnected(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsPlaying() bool
	SetIsPlaying(value bool)
	IsRecording() bool
	SetIsRecording(value bool)
	LabelColor() appkit.Color
	SetLabelColor(value appkit.Color)
	PauseButtonColor() appkit.Color
	SetPauseButtonColor(value appkit.Color)
	PlayButtonColor() appkit.Color
	SetPlayButtonColor(value appkit.Color)
	RecordButtonColor() appkit.Color
	SetRecordButtonColor(value appkit.Color)
	RewindButtonColor() appkit.Color
	SetRewindButtonColor(value appkit.Color)
}

// A view that provides an audio transport user interface.


// A view that provides an audio transport user interface.
//
// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/isConnected
func (i_ InterAppAudioTransportView) Connected() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("connected"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/currenttimelabelfont
func (i_ InterAppAudioTransportView) CurrentTimeLabelFont() appkit.Font {
	rv := objc.Send[appkit.Font](i_.ID, objc.Sel("currentTimeLabelFont"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/currenttimelabelfont
func (i_ InterAppAudioTransportView) SetCurrentTimeLabelFont(value appkit.Font) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCurrentTimeLabelFont:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/isconnected
func (i_ InterAppAudioTransportView) IsConnected() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isConnected"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/isconnected
func (i_ InterAppAudioTransportView) SetIsConnected(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsConnected:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/isenabled
func (i_ InterAppAudioTransportView) IsEnabled() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/isenabled
func (i_ InterAppAudioTransportView) SetIsEnabled(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/isplaying
func (i_ InterAppAudioTransportView) IsPlaying() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isPlaying"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/isplaying
func (i_ InterAppAudioTransportView) SetIsPlaying(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsPlaying:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/isrecording
func (i_ InterAppAudioTransportView) IsRecording() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isRecording"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/isrecording
func (i_ InterAppAudioTransportView) SetIsRecording(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsRecording:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/labelcolor
func (i_ InterAppAudioTransportView) LabelColor() appkit.Color {
	rv := objc.Send[appkit.Color](i_.ID, objc.Sel("labelColor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/labelcolor
func (i_ InterAppAudioTransportView) SetLabelColor(value appkit.Color) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLabelColor:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/pausebuttoncolor
func (i_ InterAppAudioTransportView) PauseButtonColor() appkit.Color {
	rv := objc.Send[appkit.Color](i_.ID, objc.Sel("pauseButtonColor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/pausebuttoncolor
func (i_ InterAppAudioTransportView) SetPauseButtonColor(value appkit.Color) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPauseButtonColor:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/playbuttoncolor
func (i_ InterAppAudioTransportView) PlayButtonColor() appkit.Color {
	rv := objc.Send[appkit.Color](i_.ID, objc.Sel("playButtonColor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/playbuttoncolor
func (i_ InterAppAudioTransportView) SetPlayButtonColor(value appkit.Color) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPlayButtonColor:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/recordbuttoncolor
func (i_ InterAppAudioTransportView) RecordButtonColor() appkit.Color {
	rv := objc.Send[appkit.Color](i_.ID, objc.Sel("recordButtonColor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/recordbuttoncolor
func (i_ InterAppAudioTransportView) SetRecordButtonColor(value appkit.Color) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRecordButtonColor:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/rewindbuttoncolor
func (i_ InterAppAudioTransportView) RewindButtonColor() appkit.Color {
	rv := objc.Send[appkit.Color](i_.ID, objc.Sel("rewindButtonColor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/cainterappaudiotransportview/rewindbuttoncolor
func (i_ InterAppAudioTransportView) SetRewindButtonColor(value appkit.Color) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRewindButtonColor:"), value)
}



