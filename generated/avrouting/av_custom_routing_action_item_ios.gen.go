//go:build darwin && ios

// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

// iOS-only methods for CustomRoutingActionItem


// iOS-only properties

// A string to use to override the title of the item’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingActionItem/overrideTitle
func (c_ CustomRoutingActionItem) OverrideTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("overrideTitle"))
	return rv
}
func (c_ CustomRoutingActionItem) SetOverrideTitle(value objc.IObject /* cross-framework: NSString */) {
	c_.ID.Send(objc.RegisterName("setOverrideTitle:"), value)
}

// A type with an identifier that matches a value in the app’s configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingActionItem/type
func (c_ CustomRoutingActionItem) Type() objc.IObject /* cross-framework: UTType */ {
	rv := objc.Send[uniformtypeidentifiers.UTType](c_.ID, objc.Sel("type"))
	return rv
}
func (c_ CustomRoutingActionItem) SetType(value objc.IObject /* cross-framework: UTType */) {
	c_.ID.Send(objc.RegisterName("setType:"), value)
}





