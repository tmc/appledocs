//go:build darwin && ios

// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/audiotoolbox"
)

// iOS-only methods for InterAppAudioSwitcherView


// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioSwitcherView/isShowingAppNames
func (i_ InterAppAudioSwitcherView) ShowingAppNames() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("showingAppNames"))
	return rv
}
func (i_ InterAppAudioSwitcherView) SetShowingAppNames(value bool) {
	i_.ID.Send(objc.RegisterName("setShowingAppNames:"), value)
}





