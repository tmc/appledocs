// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OutputSettingsAssistant] class.
var (
	OutputSettingsAssistantClass     _OutputSettingsAssistantClass
	OutputSettingsAssistantClassOnce sync.Once
)

func getOutputSettingsAssistantClass() _OutputSettingsAssistantClass {
	OutputSettingsAssistantClassOnce.Do(func() {
		OutputSettingsAssistantClass = _OutputSettingsAssistantClass{objc.GetClass("AVOutputSettingsAssistant")}
	})
	return OutputSettingsAssistantClass
}

type _OutputSettingsAssistantClass struct {
	class objc.Class
}

// An interface definition for the [OutputSettingsAssistant] class.
type IOutputSettingsAssistant interface {
	objectivec.IObject
}

// An object that builds audio and video output settings dictionaries.
//
// Use an output settings assistant to create the audio and video settings that you use to configure instances of and . You create an assistant with a specific preset configuration, such as or . You can accept the settings dictionaries as is to generate a file that conforms to the criteria that the preset implies. You may also use the dictionaries it generates as a base configuration that you can customize as you require. Providing the assistant additional details about your source media helps it generate more complete results. For example, setting a value for its property ensures that the assistant generates settings that don’t scale up video frames from a smaller size.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant
type OutputSettingsAssistant struct {
	objectivec.Object
}

// OutputSettingsAssistantFrom constructs a [OutputSettingsAssistant] from an unsafe.Pointer.
//
// An object that builds audio and video output settings dictionaries.
func OutputSettingsAssistantFrom(ptr unsafe.Pointer) OutputSettingsAssistant {
	return OutputSettingsAssistant{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OutputSettingsAssistantClass) Alloc() OutputSettingsAssistant {
	rv := objc.Send[OutputSettingsAssistant](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OutputSettingsAssistantClass) New() OutputSettingsAssistant {
	rv := objc.Send[OutputSettingsAssistant](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OutputSettingsAssistant) Init() OutputSettingsAssistant {
	rv := objc.Send[OutputSettingsAssistant](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OutputSettingsAssistant) Autorelease() OutputSettingsAssistant {
	rv := objc.Send[OutputSettingsAssistant](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOutputSettingsAssistant creates a new OutputSettingsAssistant instance.
func NewOutputSettingsAssistant() OutputSettingsAssistant {
	return getOutputSettingsAssistantClass().New()
}




