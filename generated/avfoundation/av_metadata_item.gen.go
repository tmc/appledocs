// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MetadataItem] class.
var (
	MetadataItemClass     _MetadataItemClass
	MetadataItemClassOnce sync.Once
)

func getMetadataItemClass() _MetadataItemClass {
	MetadataItemClassOnce.Do(func() {
		MetadataItemClass = _MetadataItemClass{objc.GetClass("AVMetadataItem")}
	})
	return MetadataItemClass
}

type _MetadataItemClass struct {
	class objc.Class
}

// An interface definition for the [MetadataItem] class.
type IMetadataItem interface {
	objectivec.IObject
}

// A metadata item for an audiovisual asset or one of its tracks.
//
// To effectively use , you need to understand how organizes metadata. To simplify finding and filtering metadata items, the framework groups related metadata into key spaces: The framework defines several format-specific key spaces. They roughly correlate to a particular container or file format, such as QuickTime (QuickTime metadata and user data) or MP3 (ID3). However, a single asset may contain metadata values across multiple key spaces. To retrieve an asset’s complete collection of format-specific metadata, you use its property. There are several common metadata values, such as a movie’s creation date or description, that can exist across multiple key spaces. To help normalize access to this common metadata, the framework provides a common key space that gives access to a limited set of metadata values common to several key spaces. This makes it easy to retrieve commonly used metadata without concern for the specific format. To retrieve an asset’s collection of common metadata, you use its property. Metadata items have keys that accord with the specification of the container format from which they’re drawn. Full details of the metadata formats, metadata keys, and metadata key spaces supported by AVFoundation are available in and . To load values of a metadata item when you access them for the first time, use the methods from the protocol. The class and other classes in turn provide their metadata as needed so that you can obtain objects from those arrays without incurring overhead for items you don’t inspect. To filter arrays of metadata items, you use the methods of this class. For example, you can filter by key and key space, by locale, and by preferred language.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem
type MetadataItem struct {
	objectivec.Object
}

// MetadataItemFrom constructs a [MetadataItem] from an unsafe.Pointer.
//
// A metadata item for an audiovisual asset or one of its tracks.
func MetadataItemFrom(ptr unsafe.Pointer) MetadataItem {
	return MetadataItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MetadataItemClass) Alloc() MetadataItem {
	rv := objc.Send[MetadataItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MetadataItemClass) New() MetadataItem {
	rv := objc.Send[MetadataItem](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataItem) Init() MetadataItem {
	rv := objc.Send[MetadataItem](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataItem) Autorelease() MetadataItem {
	rv := objc.Send[MetadataItem](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataItem creates a new MetadataItem instance.
func NewMetadataItem() MetadataItem {
	return getMetadataItemClass().New()
}


// Returns metadata items whose locales match one of the specified language identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/metadataItems(from:filteredAndSortedAccordingToPreferredLanguages:)
func (mc _MetadataItemClass) MetadataItemsFromArrayFilteredAndSortedAccordingToPreferredLanguages(metadataItems unsafe.Pointer, preferredLanguages unsafe.Pointer) []MetadataItem {
	rv := objc.Send[[]MetadataItem](objc.ID(mc.class), objc.Sel("metadataItemsFromArray:filteredAndSortedAccordingToPreferredLanguages:"), metadataItems, preferredLanguages)
	return rv
}

// Returns metadata items for the specified identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/metadataItems(from:filteredByIdentifier:)
func (mc _MetadataItemClass) MetadataItemsFromArrayFilteredByIdentifier(metadataItems unsafe.Pointer, identifier unsafe.Pointer) []MetadataItem {
	rv := objc.Send[[]MetadataItem](objc.ID(mc.class), objc.Sel("metadataItemsFromArray:filteredByIdentifier:"), metadataItems, identifier)
	return rv
}



