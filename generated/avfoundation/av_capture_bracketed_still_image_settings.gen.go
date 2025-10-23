// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CaptureBracketedStillImageSettings] class.
var (
	CaptureBracketedStillImageSettingsClass     _CaptureBracketedStillImageSettingsClass
	CaptureBracketedStillImageSettingsClassOnce sync.Once
)

func getCaptureBracketedStillImageSettingsClass() _CaptureBracketedStillImageSettingsClass {
	CaptureBracketedStillImageSettingsClassOnce.Do(func() {
		CaptureBracketedStillImageSettingsClass = _CaptureBracketedStillImageSettingsClass{objc.GetClass("AVCaptureBracketedStillImageSettings")}
	})
	return CaptureBracketedStillImageSettingsClass
}

type _CaptureBracketedStillImageSettingsClass struct {
	class objc.Class
}

// An interface definition for the [CaptureBracketedStillImageSettings] class.
type ICaptureBracketedStillImageSettings interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other AVFoundation classes.


// A parent class referenced by other AVFoundation classes. [Full Topic]
type CaptureBracketedStillImageSettings struct {
	objectivec.Object
}

// CaptureBracketedStillImageSettingsFrom constructs a [CaptureBracketedStillImageSettings] from an unsafe.Pointer.
//
// A parent class referenced by other AVFoundation classes.
func CaptureBracketedStillImageSettingsFrom(ptr unsafe.Pointer) CaptureBracketedStillImageSettings {
	return CaptureBracketedStillImageSettings{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureBracketedStillImageSettingsClass) Alloc() CaptureBracketedStillImageSettings {
	rv := objc.Send[CaptureBracketedStillImageSettings](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureBracketedStillImageSettingsClass) New() CaptureBracketedStillImageSettings {
	rv := objc.Send[CaptureBracketedStillImageSettings](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureBracketedStillImageSettings) Init() CaptureBracketedStillImageSettings {
	rv := objc.Send[CaptureBracketedStillImageSettings](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureBracketedStillImageSettings) Autorelease() CaptureBracketedStillImageSettings {
	rv := objc.Send[CaptureBracketedStillImageSettings](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureBracketedStillImageSettings creates a new CaptureBracketedStillImageSettings instance.
func NewCaptureBracketedStillImageSettings() CaptureBracketedStillImageSettings {
	return getCaptureBracketedStillImageSettingsClass().New()
}




