// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation


// _undefined is the base type for all undefined types referenced in documentation
// but not defined. These types are typically enums, options, or type aliases that
// weren't extracted from the documentation.
type _undefined = int

// Undefined types - referenced but not defined in documentation
type (
	BaselineAdjustment = _undefined // referenced in Foundation
	BinarySearchingOptions = _undefined // referenced in Foundation
	Bool = _undefined // referenced in Foundation
	CChar = _undefined // referenced in Foundation
	CalendarDate = _undefined // referenced in Foundation
	CharacterSet = _undefined // referenced in Foundation
	Color = _undefined // referenced in Foundation
	Comparator = _undefined // referenced in Foundation
	ComparisonResult = _undefined // referenced in Foundation
	Decimal = _undefined // referenced in Foundation
	DirectionalEdgeInsets = _undefined // referenced in Foundation
	EdgeInsets = _undefined // referenced in Foundation
	EnumerationOptions = _undefined // referenced in Foundation
	Enumerator = _undefined // referenced in Foundation
	Error = _undefined // referenced in Foundation
	FastEnumerationState = _undefined // referenced in Foundation
	Font = _undefined // referenced in Foundation
	IndexSet = _undefined // referenced in Foundation
	Int16 = _undefined // referenced in Foundation
	Int32 = _undefined // referenced in Foundation
	Int64 = _undefined // referenced in Foundation
	KeyType = _undefined // referenced in Foundation
	KeyValueObservingOptions = _undefined // referenced in Foundation
	LineBreakMode = _undefined // referenced in Foundation
	LinguisticTagScheme = _undefined // referenced in Foundation
	LinguisticTaggerOptions = _undefined // referenced in Foundation
	Locale = _undefined // referenced in Foundation
	OSType = _undefined // referenced in Foundation
	ObjectType = _undefined // referenced in Foundation
	Offset = _undefined // referenced in Foundation
	OrderedCollectionDifference = _undefined // referenced in Foundation
	OrderedCollectionDifferenceCalculationOptions = _undefined // referenced in Foundation
	Orthography = _undefined // referenced in Foundation
	PlaygroundQuickLook = _undefined // referenced in Foundation
	Point = _undefined // referenced in Foundation
	Predicate = _undefined // referenced in Foundation
	RandomSource = _undefined // referenced in Foundation
	Range = _undefined // referenced in Foundation
	RangePointer = _undefined // referenced in Foundation
	Rect = _undefined // referenced in Foundation
	SRAbsoluteTime = _undefined // referenced in Foundation
	SRSensor = _undefined // referenced in Foundation
	Size = _undefined // referenced in Foundation
	SortDescriptor = _undefined // referenced in Foundation
	SortOptions = _undefined // referenced in Foundation
	StringDrawingContext = _undefined // referenced in Foundation
	StringEncoding = _undefined // referenced in Foundation
	StringTransform = _undefined // referenced in Foundation
	TextAlignment = _undefined // referenced in Foundation
	Time = _undefined // referenced in Foundation
	TimeMapping = _undefined // referenced in Foundation
	TimeRange = _undefined // referenced in Foundation
	TimeZone = _undefined // referenced in Foundation
	UInt16 = _undefined // referenced in Foundation
	UInt32 = _undefined // referenced in Foundation
	UInt8 = _undefined // referenced in Foundation
	UInteger = _undefined // referenced in Foundation
	Vector = _undefined // referenced in Foundation
	Zone = _undefined // referenced in Foundation
)



// C struct types
// NSAffineTransformStruct - A structure that defines the three-by-three matrix that performs an affine transform between two coordinate systems.
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSAffineTransformStruct
type NSAffineTransformStruct struct {
	m11 float64 // An element of the transform matrix that contributes scaling, rotation, and shear.
	m12 float64 // An element of the transform matrix that contributes scaling, rotation, and shear.
	m21 float64 // An element of the transform matrix that contributes scaling, rotation, and shear.
	m22 float64 // An element of the transform matrix that contributes scaling, rotation, and shear.
	tX float64 // An element of the transform matrix that contributes translation.
	tY float64 // An element of the transform matrix that contributes translation.
}
// AffineTransformStruct is a type alias for NSAffineTransformStruct for use in objc.Send[T] calls.
type AffineTransformStruct = NSAffineTransformStruct





