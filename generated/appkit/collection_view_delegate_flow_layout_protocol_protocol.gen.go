// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

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
	CollectionViewLayoutInsetForSectionAtIndex(collectionView objc.IObject /* cross-framework: CollectionView */, collectionViewLayout ICollectionViewLayout, section int) foundation.EdgeInsets
	HasCollectionViewLayoutInsetForSectionAtIndex() bool
	CollectionViewLayoutMinimumInteritemSpacingForSectionAtIndex(collectionView objc.IObject /* cross-framework: CollectionView */, collectionViewLayout ICollectionViewLayout, section int) float64
	HasCollectionViewLayoutMinimumInteritemSpacingForSectionAtIndex() bool
	CollectionViewLayoutMinimumLineSpacingForSectionAtIndex(collectionView objc.IObject /* cross-framework: CollectionView */, collectionViewLayout ICollectionViewLayout, section int) float64
	HasCollectionViewLayoutMinimumLineSpacingForSectionAtIndex() bool
	CollectionViewLayoutReferenceSizeForFooterInSection(collectionView objc.IObject /* cross-framework: CollectionView */, collectionViewLayout ICollectionViewLayout, section int) corefoundation.Size
	HasCollectionViewLayoutReferenceSizeForFooterInSection() bool
	CollectionViewLayoutReferenceSizeForHeaderInSection(collectionView objc.IObject /* cross-framework: CollectionView */, collectionViewLayout ICollectionViewLayout, section int) corefoundation.Size
	HasCollectionViewLayoutReferenceSizeForHeaderInSection() bool
	CollectionViewLayoutSizeForItemAtIndexPath(collectionView objc.IObject /* cross-framework: CollectionView */, collectionViewLayout ICollectionViewLayout, indexPath foundation.IndexPath) corefoundation.Size
	HasCollectionViewLayoutSizeForItemAtIndexPath() bool
}
