//go:build darwin && ios

// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for WebExtensionContext

// Performs the command associated with the given key command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/performCommand(for:)-25rd1
func (w_ WebExtensionContext) PerformCommandForKeyCommand(keyCommand KeyCommand /* not a class type */) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("performCommandForKeyCommand:"), keyCommand)
	return rv
}

// iOS-only properties
