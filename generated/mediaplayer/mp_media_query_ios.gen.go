//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MediaQuery


// Adds a media property predicate to a query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/addFilterPredicate(_:)
func (m_ MediaQuery) AddFilterPredicate(predicate IMPMediaPredicate) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addFilterPredicate:"), predicate)
}

// Removes a filter predicate from a query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/removeFilterPredicate(_:)
func (m_ MediaQuery) RemoveFilterPredicate(predicate IMPMediaPredicate) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeFilterPredicate:"), predicate)
}

// iOS-only properties

// An array representing the section grouping of the query’s specified media item collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/collectionSections
func (m_ MediaQuery) CollectionSections() []IMediaQuerySection {
	rv := objc.Send[[]MediaQuerySection](m_.ID, objc.Sel("collectionSections"))
	return rv
}

// An array of media item collections whose contained items match the query’s media property predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/collections
func (m_ MediaQuery) Collections() []IMediaItemCollection {
	rv := objc.Send[[]MediaItemCollection](m_.ID, objc.Sel("collections"))
	return rv
}

// The media property predicates of the media query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/filterPredicates
func (m_ MediaQuery) FilterPredicates() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("filterPredicates"))
	return rv
}
func (m_ MediaQuery) SetFilterPredicates(value unsafe.Pointer) {
	m_.ID.Send(objc.RegisterName("setFilterPredicates:"), value)
}

// The grouping for collections retrieved with the media query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/groupingType
func (m_ MediaQuery) GroupingType() MediaGrouping {
	rv := objc.Send[MediaGrouping](m_.ID, objc.Sel("groupingType"))
	return rv
}
func (m_ MediaQuery) SetGroupingType(value MediaGrouping) {
	m_.ID.Send(objc.RegisterName("setGroupingType:"), value)
}

// An array representing the section grouping of the query’s specified media items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/itemSections
func (m_ MediaQuery) ItemSections() []IMediaQuerySection {
	rv := objc.Send[[]MediaQuerySection](m_.ID, objc.Sel("itemSections"))
	return rv
}

// An array of media items that match the media query’s predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/items
func (m_ MediaQuery) Items() []IMediaItem {
	rv := objc.Send[[]MediaItem](m_.ID, objc.Sel("items"))
	return rv
}




