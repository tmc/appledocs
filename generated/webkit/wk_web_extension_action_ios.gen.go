//go:build darwin && ios

// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for WebExtensionAction


// iOS-only properties

// A view controller that presents a web view loaded with the pop-up page for this action, or if no popup is specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/popupViewController
func (w_ WebExtensionAction) PopupViewController() appkit.ViewController {
	rv := objc.Send[appkit.ViewController](w_.ID, objc.Sel("popupViewController"))
	return rv
}





