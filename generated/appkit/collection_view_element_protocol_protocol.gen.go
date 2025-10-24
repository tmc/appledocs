// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PCollectionViewElement is the NSCollectionViewElement protocol interface.
//
// A set of methods that you use to manage the content in a collection view.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSCollectionViewElement
type PCollectionViewElement interface {
	// Optional methods
	ApplyLayoutAttributes(layoutAttributes ICollectionViewLayoutAttributes)
	HasApplyLayoutAttributes() bool
	DidTransitionFromLayoutToLayout(oldLayout ICollectionViewLayout, newLayout ICollectionViewLayout)
	HasDidTransitionFromLayoutToLayout() bool
	PreferredLayoutAttributesFittingAttributes(layoutAttributes ICollectionViewLayoutAttributes) CollectionViewLayoutAttributes
	HasPreferredLayoutAttributesFittingAttributes() bool
	PrepareForReuse()
	HasPrepareForReuse() bool
	WillTransitionFromLayoutToLayout(oldLayout ICollectionViewLayout, newLayout ICollectionViewLayout)
	HasWillTransitionFromLayoutToLayout() bool
}
