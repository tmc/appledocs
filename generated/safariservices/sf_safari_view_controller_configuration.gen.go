// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [SFSafariViewControllerConfiguration] class.
var (
	SFSafariViewControllerConfigurationClass     _SFSafariViewControllerConfigurationClass
	SFSafariViewControllerConfigurationClassOnce sync.Once
)

func getSFSafariViewControllerConfigurationClass() _SFSafariViewControllerConfigurationClass {
	SFSafariViewControllerConfigurationClassOnce.Do(func() {
		SFSafariViewControllerConfigurationClass = _SFSafariViewControllerConfigurationClass{objc.GetClass("SFSafariViewControllerConfiguration")}
	})
	return SFSafariViewControllerConfigurationClass
}

type _SFSafariViewControllerConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [SFSafariViewControllerConfiguration] class.
type ISFSafariViewControllerConfiguration interface {
	objectivec.IObject
}

// A configuration object that defines how a Safari view controller should be initialized.
//
// Use a configuration object with the method to initialize your view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/Configuration-swift.class
type SFSafariViewControllerConfiguration struct {
	objectivec.Object
}

// SFSafariViewControllerConfigurationFrom constructs a [SFSafariViewControllerConfiguration] from an unsafe.Pointer.
//
// A configuration object that defines how a Safari view controller should be initialized.
func SFSafariViewControllerConfigurationFrom(ptr unsafe.Pointer) SFSafariViewControllerConfiguration {
	return SFSafariViewControllerConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFSafariViewControllerConfigurationClass) Alloc() SFSafariViewControllerConfiguration {
	rv := objc.Send[SFSafariViewControllerConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFSafariViewControllerConfigurationClass) New() SFSafariViewControllerConfiguration {
	rv := objc.Send[SFSafariViewControllerConfiguration](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariViewControllerConfiguration) Init() SFSafariViewControllerConfiguration {
	rv := objc.Send[SFSafariViewControllerConfiguration](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariViewControllerConfiguration) Autorelease() SFSafariViewControllerConfiguration {
	rv := objc.Send[SFSafariViewControllerConfiguration](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariViewControllerConfiguration creates a new SFSafariViewControllerConfiguration instance.
func NewSFSafariViewControllerConfiguration() SFSafariViewControllerConfiguration {
	return getSFSafariViewControllerConfigurationClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/Configuration-swift.class/activityButton
func (s_ SFSafariViewControllerConfiguration) ActivityButton() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("activityButton"))
	return rv
}


// SetActivityButton sets the value of the activityButton property.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/Configuration-swift.class/activityButton
func (s_ SFSafariViewControllerConfiguration) SetActivityButton(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setActivityButton:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/Configuration-swift.class/barCollapsingEnabled
func (s_ SFSafariViewControllerConfiguration) BarCollapsingEnabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("barCollapsingEnabled"))
	return rv
}


// SetBarCollapsingEnabled sets the value of the barCollapsingEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/Configuration-swift.class/barCollapsingEnabled
func (s_ SFSafariViewControllerConfiguration) SetBarCollapsingEnabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBarCollapsingEnabled:"), value)
}

// A value that specifies whether Safari should enter Reader mode, if it is available.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/Configuration-swift.class/entersReaderIfAvailable
func (s_ SFSafariViewControllerConfiguration) EntersReaderIfAvailable() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("entersReaderIfAvailable"))
	return rv
}


// SetEntersReaderIfAvailable sets the value of the entersReaderIfAvailable property.
// A value that specifies whether Safari should enter Reader mode, if it is available.

//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/Configuration-swift.class/entersReaderIfAvailable
func (s_ SFSafariViewControllerConfiguration) SetEntersReaderIfAvailable(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEntersReaderIfAvailable:"), value)
}

// An object you use to send tap event attribution data to the browser for Private Click Measurement.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/Configuration-swift.class/eventAttribution
func (s_ SFSafariViewControllerConfiguration) EventAttribution() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("eventAttribution"))
	return rv
}


// SetEventAttribution sets the value of the eventAttribution property.
// An object you use to send tap event attribution data to the browser for Private Click Measurement.

//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/Configuration-swift.class/eventAttribution
func (s_ SFSafariViewControllerConfiguration) SetEventAttribution(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEventAttribution:"), value)
}



