//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for CXSetTranslatingCallAction


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetTranslatingCallAction/fulfill(using:)
func (c_ CXSetTranslatingCallAction) FulfillUsingTranslationEngine(translationEngine CXTranslationEngine) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fulfillUsingTranslationEngine:"), translationEngine)
}

// iOS-only properties

// A value that indicates whether translation is active for a call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetTranslatingCallAction/isTranslating
func (c_ CXSetTranslatingCallAction) IsTranslating() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isTranslating"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetTranslatingCallAction/localLanguage
func (c_ CXSetTranslatingCallAction) LocalLanguage() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("localLanguage"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetTranslatingCallAction/remoteLanguage
func (c_ CXSetTranslatingCallAction) RemoteLanguage() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("remoteLanguage"))
	return rv
}




