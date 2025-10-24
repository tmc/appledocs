//go:build darwin && ios

// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for WebExtensionCommand

// iOS-only properties

// A key command representation of the web extension command for use in the responder chain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command/keyCommand
func (w_ WebExtensionCommand) KeyCommand() KeyCommand /* not a class type */ {
	rv := objc.Send[KeyCommand](w_.ID, objc.Sel("keyCommand"))
	return rv
}
