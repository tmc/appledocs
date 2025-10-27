// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"
)

// PCollectionViewDelegateFlowLayout is the NSCollectionViewDelegateFlowLayout protocol interface.
//
// A set of methods that a delegate implements to provide layout information to a flow layout object in a collection view.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSCollectionViewDelegateFlowLayout
type PCollectionViewDelegateFlowLayout interface {
	// Optional methods
	CollectionViewLayoutInsetForSectionAtIndex(collectionView CollectionView /* not a class type */, collectionViewLayout ICollectionViewLayout, section int) foundation.EdgeInsets
	HasCollectionViewLayoutInsetForSectionAtIndex() bool
	CollectionViewLayoutMinimumInteritemSpacingForSectionAtIndex(collectionView CollectionView /* not a class type */, collectionViewLayout ICollectionViewLayout, section int) float64
	HasCollectionViewLayoutMinimumInteritemSpacingForSectionAtIndex() bool
	CollectionViewLayoutMinimumLineSpacingForSectionAtIndex(collectionView CollectionView /* not a class type */, collectionViewLayout ICollectionViewLayout, section int) float64
	HasCollectionViewLayoutMinimumLineSpacingForSectionAtIndex() bool
	CollectionViewLayoutReferenceSizeForFooterInSection(collectionView CollectionView /* not a class type */, collectionViewLayout ICollectionViewLayout, section int) corefoundation.CGSize
	HasCollectionViewLayoutReferenceSizeForFooterInSection() bool
	CollectionViewLayoutReferenceSizeForHeaderInSection(collectionView CollectionView /* not a class type */, collectionViewLayout ICollectionViewLayout, section int) corefoundation.CGSize
	HasCollectionViewLayoutReferenceSizeForHeaderInSection() bool
	CollectionViewLayoutSizeForItemAtIndexPath(collectionView CollectionView /* not a class type */, collectionViewLayout ICollectionViewLayout, indexPath foundation.foundation.INSIndexPath) corefoundation.CGSize
	HasCollectionViewLayoutSizeForItemAtIndexPath() bool
}
