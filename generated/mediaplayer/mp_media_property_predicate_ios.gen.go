//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MediaPropertyPredicate


// iOS-only properties

// The type of matching comparison that the media property predicate performs when you invoke a query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPropertyPredicate/comparisonType
func (m_ MediaPropertyPredicate) ComparisonType() MediaPredicateComparison {
	rv := objc.Send[MediaPredicateComparison](m_.ID, objc.Sel("comparisonType"))
	return rv
}

// The property that the media property predicate uses when you invoke a query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPropertyPredicate/property
func (m_ MediaPropertyPredicate) Property() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("property"))
	return rv
}

// The value that the media property predicate matches against when you invoke a query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPropertyPredicate/value
func (m_ MediaPropertyPredicate) Value() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("value"))
	return rv
}




