// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PCollectionViewPrefetching is the NSCollectionViewPrefetching protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSCollectionViewPrefetching
type PCollectionViewPrefetching interface {
	// Required methods
	CollectionViewPrefetchItemsAtIndexPaths(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths []foundation.IndexPath)/* debug [protocol_interface/required_method]: CollectionViewPrefetchItemsAtIndexPaths */
	// Optional methods
	CollectionViewCancelPrefetchingForItemsAtIndexPaths(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths []foundation.IndexPath)
	HasCollectionViewCancelPrefetchingForItemsAtIndexPaths() bool
}
