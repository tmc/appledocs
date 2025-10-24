// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFSafariSettings] class.
var (
	SFSafariSettingsClass     _SFSafariSettingsClass
	SFSafariSettingsClassOnce sync.Once
)

func getSFSafariSettingsClass() _SFSafariSettingsClass {
	SFSafariSettingsClassOnce.Do(func() {
		SFSafariSettingsClass = _SFSafariSettingsClass{objc.GetClass("SFSafariSettings")}
	})
	return SFSafariSettingsClass
}

type _SFSafariSettingsClass struct {
	class objc.Class
}

// An interface definition for the [SFSafariSettings] class.
type ISFSafariSettings interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariSettings
type SFSafariSettings struct {
	objectivec.Object
}

// SFSafariSettingsFrom constructs a [SFSafariSettings] from an unsafe.Pointer.
func SFSafariSettingsFrom(ptr unsafe.Pointer) SFSafariSettings {
	return SFSafariSettings{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFSafariSettingsClass) Alloc() SFSafariSettings {
	rv := objc.Send[SFSafariSettings](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFSafariSettingsClass) New() SFSafariSettings {
	rv := objc.Send[SFSafariSettings](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariSettings) Init() SFSafariSettings {
	rv := objc.Send[SFSafariSettings](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariSettings) Autorelease() SFSafariSettings {
	rv := objc.Send[SFSafariSettings](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariSettings creates a new SFSafariSettings instance.
func NewSFSafariSettings() SFSafariSettings {
	return getSFSafariSettingsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariSettings/openExportBrowsingDataSettings(completionHandler:)
func (sc _SFSafariSettingsClass) OpenExportBrowsingDataSettingsWithCompletionHandler(completionHandler func(unsafe.Pointer)) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("openExportBrowsingDataSettingsWithCompletionHandler:"), completionHandler)
}


