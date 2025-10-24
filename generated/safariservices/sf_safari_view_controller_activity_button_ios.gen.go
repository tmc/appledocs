//go:build darwin && ios

// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for SFSafariViewControllerActivityButton


// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/ActivityButton/extensionIdentifier
func (s_ SFSafariViewControllerActivityButton) ExtensionIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("extensionIdentifier"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/ActivityButton/templateImage
func (s_ SFSafariViewControllerActivityButton) TemplateImage() appkit.Image {
	rv := objc.Send[appkit.Image](s_.ID, objc.Sel("templateImage"))
	return rv
}




