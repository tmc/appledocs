//go:build darwin && ios

// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for Coder


// Decodes and returns the Core Graphics affine transform structure associated with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCGAffineTransform(forKey:)
func (c_ Coder) DecodeCGAffineTransformForKey(key IString) IAffineTransform {
	rv := objc.Send[AffineTransform](c_.ID, objc.Sel("decodeCGAffineTransformForKey:"), key)
	return rv
}

// Decodes and returns the Core Graphics point structure associated with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCGPoint(forKey:)
func (c_ Coder) DecodeCGPointForKey(key IString) objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[Point](c_.ID, objc.Sel("decodeCGPointForKey:"), key)
	return rv
}

// Decodes and returns the Core Graphics rectangle structure associated with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCGRect(forKey:)
func (c_ Coder) DecodeCGRectForKey(key IString) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[Rect](c_.ID, objc.Sel("decodeCGRectForKey:"), key)
	return rv
}

// Decodes and returns the Core Graphics size structure associated with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCGSize(forKey:)
func (c_ Coder) DecodeCGSizeForKey(key IString) objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[Size](c_.ID, objc.Sel("decodeCGSizeForKey:"), key)
	return rv
}

// Decodes and returns the Core Graphics vector data associated with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCGVector(forKey:)
func (c_ Coder) DecodeCGVectorForKey(key IString) Vector /* not a class type */ {
	rv := objc.Send[Vector](c_.ID, objc.Sel("decodeCGVectorForKey:"), key)
	return rv
}

// Decodes and returns the UIKit directional edge insets structure associated with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeDirectionalEdgeInsets(forKey:)
func (c_ Coder) DecodeDirectionalEdgeInsetsForKey(key IString) DirectionalEdgeInsets /* not a class type */ {
	rv := objc.Send[DirectionalEdgeInsets](c_.ID, objc.Sel("decodeDirectionalEdgeInsetsForKey:"), key)
	return rv
}

// Decodes and returns the UIKit edge insets structure associated with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeUIEdgeInsets(forKey:)
func (c_ Coder) DecodeUIEdgeInsetsForKey(key IString) UIEdgeInsets /* not a class type */ {
	rv := objc.Send[EdgeInsets](c_.ID, objc.Sel("decodeUIEdgeInsetsForKey:"), key)
	return rv
}

// Decodes and returns the UIKit offset structure associated with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/decodeUIOffset(forKey:)
func (c_ Coder) DecodeUIOffsetForKey(key IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("decodeUIOffsetForKey:"), key)
	return rv
}

// Encodes a rectangle and associates it with the specified key in the receiver’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-10qhm
func (c_ Coder) EncodeCGRectForKey(rect objc.IObject /* cross-framework: Rect */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeCGRect:forKey:"), rect, key)
}

// Encodes vector data and associates it with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-26fxa
func (c_ Coder) EncodeCGVectorForKey(vector Vector /* not a class type */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeCGVector:forKey:"), vector, key)
}

// Encodes an affine transform and associates it with the specified key in the receiver’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-29jyx
func (c_ Coder) EncodeCGAffineTransformForKey(transform IAffineTransform, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeCGAffineTransform:forKey:"), transform, key)
}

// Encodes edge inset data and associates it with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-44zsc
func (c_ Coder) EncodeUIEdgeInsetsForKey(insets EdgeInsets /* not a class type */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeUIEdgeInsets:forKey:"), insets, key)
}

// Encodes size information and associates it with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-6wq3n
func (c_ Coder) EncodeCGSizeForKey(size objc.IObject /* cross-framework: Size */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeCGSize:forKey:"), size, key)
}

// Encodes directional edge inset data and associates it with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-7oo2n
func (c_ Coder) EncodeDirectionalEdgeInsetsForKey(insets DirectionalEdgeInsets /* not a class type */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeDirectionalEdgeInsets:forKey:"), insets, key)
}

// Encodes a point and associates it with the specified key in the receiver’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-7z9kc
func (c_ Coder) EncodeCGPointForKey(point objc.IObject /* cross-framework: Point */, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeCGPoint:forKey:"), point, key)
}

// Encodes offset data and associates it with the specified key in the coder’s archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-9d1qy
func (c_ Coder) EncodeUIOffsetForKey(offset objectivec.IObject, key IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeUIOffset:forKey:"), offset, key)
}

// iOS-only properties





