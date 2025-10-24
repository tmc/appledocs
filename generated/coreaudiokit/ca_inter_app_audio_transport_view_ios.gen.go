//go:build darwin && ios

// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// iOS-only methods for InterAppAudioTransportView


// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/isConnected
func (i_ InterAppAudioTransportView) Connected() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("connected"))
	return rv
}





