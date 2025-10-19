// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEDNSSettings] class.
var nEDNSSettingsClass = _NEDNSSettingsClass{objc.GetClass("NEDNSSettings")}

type _NEDNSSettingsClass struct {
	class objc.Class
}

// An interface definition for the [NEDNSSettings] class.
type INEDNSSettings interface {
	objectivec.IObject
}

// A parent class referenced by other NetworkExtension classes. [Full Topic]

type NEDNSSettings struct {
	objectivec.Object
}

// NEDNSSettingsFrom constructs a [NEDNSSettings] from an unsafe.Pointer.
//
// A parent class referenced by other NetworkExtension classes.
func NEDNSSettingsFrom(ptr unsafe.Pointer) NEDNSSettings {
	return NEDNSSettings{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (nc _NEDNSSettingsClass) Alloc() NEDNSSettings {
	rv := objc.Send[NEDNSSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _NEDNSSettingsClass) New() NEDNSSettings {
	rv := objc.Send[NEDNSSettings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEDNSSettings) Init() NEDNSSettings {
	rv := objc.Send[NEDNSSettings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEDNSSettings) Autorelease() NEDNSSettings {
	rv := objc.Send[NEDNSSettings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSSettings creates a new NEDNSSettings instance.
func NewNEDNSSettings() NEDNSSettings {
	return nEDNSSettingsClass.New()
}




