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
	CollectionViewPrefetchItemsAtIndexPaths(collectionView CollectionView /* not a class type */, indexPaths []foundation.IndexPath)
	// Optional methods
	CollectionViewCancelPrefetchingForItemsAtIndexPaths(collectionView CollectionView /* not a class type */, indexPaths []foundation.IndexPath)
	HasCollectionViewCancelPrefetchingForItemsAtIndexPaths() bool
}
