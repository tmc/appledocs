// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCaptureBracketedStillImageSettings] class.
var (
	aVCaptureBracketedStillImageSettingsClass     _AVCaptureBracketedStillImageSettingsClass
	aVCaptureBracketedStillImageSettingsClassOnce sync.Once
)

func getAVCaptureBracketedStillImageSettingsClass() _AVCaptureBracketedStillImageSettingsClass {
	aVCaptureBracketedStillImageSettingsClassOnce.Do(func() {
		aVCaptureBracketedStillImageSettingsClass = _AVCaptureBracketedStillImageSettingsClass{objc.GetClass("AVCaptureBracketedStillImageSettings")}
	})
	return aVCaptureBracketedStillImageSettingsClass
}

type _AVCaptureBracketedStillImageSettingsClass struct {
	class objc.Class
}

// An interface definition for the [AVCaptureBracketedStillImageSettings] class.
type IAVCaptureBracketedStillImageSettings interface {
	objectivec.IObject
}

// A parent class referenced by other AVFoundation classes. [Full Topic]
type AVCaptureBracketedStillImageSettings struct {
	objectivec.Object
}

// AVCaptureBracketedStillImageSettingsFrom constructs a [AVCaptureBracketedStillImageSettings] from an unsafe.Pointer.
//
// A parent class referenced by other AVFoundation classes.
func AVCaptureBracketedStillImageSettingsFrom(ptr unsafe.Pointer) AVCaptureBracketedStillImageSettings {
	return AVCaptureBracketedStillImageSettings{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVCaptureBracketedStillImageSettingsClass) Alloc() AVCaptureBracketedStillImageSettings {
	rv := objc.Send[AVCaptureBracketedStillImageSettings](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVCaptureBracketedStillImageSettingsClass) New() AVCaptureBracketedStillImageSettings {
	rv := objc.Send[AVCaptureBracketedStillImageSettings](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaptureBracketedStillImageSettings) Init() AVCaptureBracketedStillImageSettings {
	rv := objc.Send[AVCaptureBracketedStillImageSettings](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaptureBracketedStillImageSettings) Autorelease() AVCaptureBracketedStillImageSettings {
	rv := objc.Send[AVCaptureBracketedStillImageSettings](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureBracketedStillImageSettings creates a new AVCaptureBracketedStillImageSettings instance.
func NewAVCaptureBracketedStillImageSettings() AVCaptureBracketedStillImageSettings {
	return getAVCaptureBracketedStillImageSettingsClass().New()
}




