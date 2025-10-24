//go:build darwin && ios

// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for SFSafariViewController


// iOS-only properties

// A copy of the Safari view controller’s initialized configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/configuration-swift.property
func (s_ SFSafariViewController) Configuration() ISFSafariViewControllerConfiguration {
	rv := objc.Send[SFSafariViewControllerConfiguration](s_.ID, objc.Sel("configuration"))
	return rv
}

// An object that provides behavior for the Safari view controller’s Done and Action buttons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/delegate
func (s_ SFSafariViewController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}
func (s_ SFSafariViewController) SetDelegate(value unsafe.Pointer) {
	s_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// The style of dismiss button to use in the navigation bar to close the Safari view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/dismissButtonStyle-swift.property
func (s_ SFSafariViewController) DismissButtonStyle() SFSafariViewControllerDismissButtonStyle {
	rv := objc.Send[SFSafariViewControllerDismissButtonStyle](s_.ID, objc.Sel("dismissButtonStyle"))
	return rv
}
func (s_ SFSafariViewController) SetDismissButtonStyle(value SFSafariViewControllerDismissButtonStyle) {
	s_.ID.Send(objc.RegisterName("setDismissButtonStyle:"), value)
}

// The color to tint the background of the navigation bar and the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/preferredBarTintColor
func (s_ SFSafariViewController) PreferredBarTintColor() appkit.Color {
	rv := objc.Send[appkit.Color](s_.ID, objc.Sel("preferredBarTintColor"))
	return rv
}
func (s_ SFSafariViewController) SetPreferredBarTintColor(value appkit.Color) {
	s_.ID.Send(objc.RegisterName("setPreferredBarTintColor:"), value)
}

// The color to tint the control buttons on the navigation bar and the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/preferredControlTintColor
func (s_ SFSafariViewController) PreferredControlTintColor() appkit.Color {
	rv := objc.Send[appkit.Color](s_.ID, objc.Sel("preferredControlTintColor"))
	return rv
}
func (s_ SFSafariViewController) SetPreferredControlTintColor(value appkit.Color) {
	s_.ID.Send(objc.RegisterName("setPreferredControlTintColor:"), value)
}




