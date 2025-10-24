//go:build darwin && ios

// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for SFSafariViewControllerConfiguration


// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/Configuration-swift.class/activityButton
func (s_ SFSafariViewControllerConfiguration) ActivityButton() ISFSafariViewControllerActivityButton {
	rv := objc.Send[SFSafariViewControllerActivityButton](s_.ID, objc.Sel("activityButton"))
	return rv
}
func (s_ SFSafariViewControllerConfiguration) SetActivityButton(value ISFSafariViewControllerActivityButton) {
	s_.ID.Send(objc.RegisterName("setActivityButton:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/Configuration-swift.class/barCollapsingEnabled
func (s_ SFSafariViewControllerConfiguration) BarCollapsingEnabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("barCollapsingEnabled"))
	return rv
}
func (s_ SFSafariViewControllerConfiguration) SetBarCollapsingEnabled(value bool) {
	s_.ID.Send(objc.RegisterName("setBarCollapsingEnabled:"), value)
}

// A value that specifies whether Safari should enter Reader mode, if it is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/Configuration-swift.class/entersReaderIfAvailable
func (s_ SFSafariViewControllerConfiguration) EntersReaderIfAvailable() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("entersReaderIfAvailable"))
	return rv
}
func (s_ SFSafariViewControllerConfiguration) SetEntersReaderIfAvailable(value bool) {
	s_.ID.Send(objc.RegisterName("setEntersReaderIfAvailable:"), value)
}

// An object you use to send tap event attribution data to the browser for Private Click Measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/Configuration-swift.class/eventAttribution
func (s_ SFSafariViewControllerConfiguration) EventAttribution() EventAttribution /* not a class type */ {
	rv := objc.Send[EventAttribution](s_.ID, objc.Sel("eventAttribution"))
	return rv
}
func (s_ SFSafariViewControllerConfiguration) SetEventAttribution(value EventAttribution /* not a class type */) {
	s_.ID.Send(objc.RegisterName("setEventAttribution:"), value)
}





