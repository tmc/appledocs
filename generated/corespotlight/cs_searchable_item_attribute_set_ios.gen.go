//go:build darwin && ios

// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

// iOS-only methods for CSSearchableItemAttributeSet


// iOS-only properties

// The identifiers that specify custom actions the app supports for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/actionIdentifiers
func (c_ CSSearchableItemAttributeSet) ActionIdentifiers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("actionIdentifiers"))
	return rv
}
func (c_ CSSearchableItemAttributeSet) SetActionIdentifiers(value []string) {
	c_.ID.Send(objc.RegisterName("setActionIdentifiers:"), value)
}

// The file type of the item to enable the user to share items from Spotlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/sharedItemContentType
func (c_ CSSearchableItemAttributeSet) SharedItemContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[uniformtypeidentifiers.UTType](c_.ID, objc.Sel("sharedItemContentType"))
	return rv
}
func (c_ CSSearchableItemAttributeSet) SetSharedItemContentType(value uniformtypeidentifiers.UTType) {
	c_.ID.Send(objc.RegisterName("setSharedItemContentType:"), value)
}




