// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Value] class.
var (
	ValueClass     _ValueClass
	ValueClassOnce sync.Once
)

func getValueClass() _ValueClass {
	ValueClassOnce.Do(func() {
		ValueClass = _ValueClass{objc.GetClass("NSValue")}
	})
	return ValueClass
}

type _ValueClass struct {
	class objc.Class
}

// An interface definition for the [Value] class.
type IValue interface {
	objectivec.IObject
	// properties:
	Hash() int /* primitive/slice/pointer. */
	SetHash(value int /* primitive/slice/pointer. */)
	CaTransform3DValue() Transform3D /* not a class type */
	SetCaTransform3DValue(value Transform3D /* not a class type */)
	CgAffineTransformValue() IAffineTransform
	SetCgAffineTransformValue(value IAffineTransform)
	CgPointValue() objc.IObject /* cross-framework: Point */
	SetCgPointValue(value objc.IObject /* cross-framework: Point */)
	CgRectValue() objc.IObject /* cross-framework: Rect */
	SetCgRectValue(value objc.IObject /* cross-framework: Rect */)
	CgSizeValue() objc.IObject /* cross-framework: Size */
	SetCgSizeValue(value objc.IObject /* cross-framework: Size */)
	CgVectorValue() Vector /* not a class type */
	SetCgVectorValue(value Vector /* not a class type */)
	DirectionalEdgeInsetsValue() DirectionalEdgeInsets /* not a class type */
	SetDirectionalEdgeInsetsValue(value DirectionalEdgeInsets /* not a class type */)
	EdgeInsetsValue() NSEdgeInsets /* not a class type */
	SetEdgeInsetsValue(value NSEdgeInsets /* not a class type */)
	GcPoint2Value() unsafe.Pointer
	SetGcPoint2Value(value unsafe.Pointer)
	MkCoordinateSpanValue() unsafe.Pointer
	SetMkCoordinateSpanValue(value unsafe.Pointer)
	MkCoordinateValue() LocationCoordinate2D /* not a class type */
	SetMkCoordinateValue(value LocationCoordinate2D /* not a class type */)
	NonretainedObjectValue() unsafe.Pointer
	SetNonretainedObjectValue(value unsafe.Pointer)
	ObjCType() unsafe.Pointer
	SetObjCType(value unsafe.Pointer)
	PointValue() objc.IObject /* cross-framework: Point */
	SetPointValue(value objc.IObject /* cross-framework: Point */)
	PointerValue() unsafe.Pointer
	SetPointerValue(value unsafe.Pointer)
	RangeValue() NSRange /* not a class type */
	SetRangeValue(value NSRange /* not a class type */)
	RectValue() objc.IObject /* cross-framework: Rect */
	SetRectValue(value objc.IObject /* cross-framework: Rect */)
	ScnMatrix4Value() NMatrix4 /* not a class type */
	SetScnMatrix4Value(value NMatrix4 /* not a class type */)
	ScnVector3Value() NVector3 /* not a class type */
	SetScnVector3Value(value NVector3 /* not a class type */)
	ScnVector4Value() NVector4 /* not a class type */
	SetScnVector4Value(value NVector4 /* not a class type */)
	SizeValue() objc.IObject /* cross-framework: Size */
	SetSizeValue(value objc.IObject /* cross-framework: Size */)
	TimeMappingValue() TimeMapping /* not a class type */
	SetTimeMappingValue(value TimeMapping /* not a class type */)
	TimeRangeValue() TimeRange /* not a class type */
	SetTimeRangeValue(value TimeRange /* not a class type */)
	TimeValue() Time /* not a class type */
	SetTimeValue(value Time /* not a class type */)
	UiEdgeInsetsValue() UIEdgeInsets /* not a class type */
	SetUiEdgeInsetsValue(value UIEdgeInsets /* not a class type */)
	UiOffsetValue() Offset /* not a class type */
	SetUiOffsetValue(value Offset /* not a class type */)
	VideoDimensionsValue() VideoDimensions /* not a class type */
	SetVideoDimensionsValue(value VideoDimensions /* not a class type */)
	// methods:
}

// A simple container for a single C or Objective-C data item.
//
// An object can hold any of the scalar types such as , , and , as well as pointers, structures, and object references. Use this class to work with such data types in collections (such as and ), , and other APIs that require Objective-C objects. objects are always immutable.


// A simple container for a single C or Objective-C data item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue
type Value struct {
	objectivec.Object
}

// ValueFrom constructs a [Value] from an unsafe.Pointer.
//
// A simple container for a single C or Objective-C data item.
func ValueFrom(ptr unsafe.Pointer) Value {
	return Value{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _ValueClass) Alloc() Value {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _ValueClass) New() Value {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ Value) Init() Value {
	rv := objc.Send[Value](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ Value) Autorelease() Value {
	rv := objc.Send[Value](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewValue creates a new Value instance.
func NewValue() Value {
	return getValueClass().New()
}



// Returns an integer that can be used as a table address in a hash table structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash
func (v_ Value) Hash() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](v_.ID, objc.Sel("hash"))
	return rv
}


// Returns an integer that can be used as a table address in a hash table structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash
func (v_ Value) SetHash(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setHash:"), value)
}


// The CoreAnimation transform structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/catransform3dvalue
func (v_ Value) CaTransform3DValue() Transform3D /* not a class type */ {
	rv := objc.Send[Transform3D](v_.ID, objc.Sel("caTransform3DValue"))
	return rv
}


// The CoreAnimation transform structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/catransform3dvalue
func (v_ Value) SetCaTransform3DValue(value Transform3D /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCaTransform3DValue:"), value)
}


// Returns the CoreGraphics affine transform representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/cgaffinetransformvalue
func (v_ Value) CgAffineTransformValue() IAffineTransform {
	rv := objc.Send[AffineTransform](v_.ID, objc.Sel("cgAffineTransformValue"))
	return rv
}


// Returns the CoreGraphics affine transform representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/cgaffinetransformvalue
func (v_ Value) SetCgAffineTransformValue(value IAffineTransform) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCgAffineTransformValue:"), value)
}


// Returns the CoreGraphics point structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/cgpointvalue
func (v_ Value) CgPointValue() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[Point](v_.ID, objc.Sel("cgPointValue"))
	return rv
}


// Returns the CoreGraphics point structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/cgpointvalue
func (v_ Value) SetCgPointValue(value objc.IObject /* cross-framework: Point */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCgPointValue:"), value)
}


// Returns the CoreGraphics rectangle structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/cgrectvalue
func (v_ Value) CgRectValue() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("cgRectValue"))
	return rv
}


// Returns the CoreGraphics rectangle structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/cgrectvalue
func (v_ Value) SetCgRectValue(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCgRectValue:"), value)
}


// Returns the CoreGraphics size structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/cgsizevalue
func (v_ Value) CgSizeValue() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[Size](v_.ID, objc.Sel("cgSizeValue"))
	return rv
}


// Returns the CoreGraphics size structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/cgsizevalue
func (v_ Value) SetCgSizeValue(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCgSizeValue:"), value)
}


// Returns the CoreGraphics vector structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/cgvectorvalue
func (v_ Value) CgVectorValue() Vector /* not a class type */ {
	rv := objc.Send[Vector](v_.ID, objc.Sel("cgVectorValue"))
	return rv
}


// Returns the CoreGraphics vector structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/cgvectorvalue
func (v_ Value) SetCgVectorValue(value Vector /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCgVectorValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/directionaledgeinsetsvalue
func (v_ Value) DirectionalEdgeInsetsValue() DirectionalEdgeInsets /* not a class type */ {
	rv := objc.Send[DirectionalEdgeInsets](v_.ID, objc.Sel("directionalEdgeInsetsValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/directionaledgeinsetsvalue
func (v_ Value) SetDirectionalEdgeInsetsValue(value DirectionalEdgeInsets /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDirectionalEdgeInsetsValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/edgeinsetsvalue
func (v_ Value) EdgeInsetsValue() NSEdgeInsets /* not a class type */ {
	rv := objc.Send[EdgeInsets](v_.ID, objc.Sel("edgeInsetsValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/edgeinsetsvalue
func (v_ Value) SetEdgeInsetsValue(value NSEdgeInsets /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setEdgeInsetsValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/gcpoint2value
func (v_ Value) GcPoint2Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("gcPoint2Value"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/gcpoint2value
func (v_ Value) SetGcPoint2Value(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setGcPoint2Value:"), value)
}


// The MapKit coordinate span structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/mkcoordinatespanvalue
func (v_ Value) MkCoordinateSpanValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("mkCoordinateSpanValue"))
	return rv
}


// The MapKit coordinate span structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/mkcoordinatespanvalue
func (v_ Value) SetMkCoordinateSpanValue(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMkCoordinateSpanValue:"), value)
}


// The CoreLocation geographic coordinate structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/mkcoordinatevalue
func (v_ Value) MkCoordinateValue() LocationCoordinate2D /* not a class type */ {
	rv := objc.Send[LocationCoordinate2D](v_.ID, objc.Sel("mkCoordinateValue"))
	return rv
}


// The CoreLocation geographic coordinate structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/mkcoordinatevalue
func (v_ Value) SetMkCoordinateValue(value LocationCoordinate2D /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMkCoordinateValue:"), value)
}


// The value as a non-retained pointer to an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/nonretainedobjectvalue
func (v_ Value) NonretainedObjectValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("nonretainedObjectValue"))
	return rv
}


// The value as a non-retained pointer to an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/nonretainedobjectvalue
func (v_ Value) SetNonretainedObjectValue(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNonretainedObjectValue:"), value)
}


// A C string containing the Objective-C type of the data contained in the value object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/objctype
func (v_ Value) ObjCType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("objCType"))
	return rv
}


// A C string containing the Objective-C type of the data contained in the value object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/objctype
func (v_ Value) SetObjCType(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setObjCType:"), value)
}


// The Foundation point structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/pointvalue
func (v_ Value) PointValue() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[Point](v_.ID, objc.Sel("pointValue"))
	return rv
}


// The Foundation point structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/pointvalue
func (v_ Value) SetPointValue(value objc.IObject /* cross-framework: Point */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPointValue:"), value)
}


// Returns the value as an untyped pointer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/pointervalue
func (v_ Value) PointerValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("pointerValue"))
	return rv
}


// Returns the value as an untyped pointer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/pointervalue
func (v_ Value) SetPointerValue(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPointerValue:"), value)
}


// The Foundation range structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/rangevalue
func (v_ Value) RangeValue() NSRange /* not a class type */ {
	rv := objc.Send[Range](v_.ID, objc.Sel("rangeValue"))
	return rv
}


// The Foundation range structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/rangevalue
func (v_ Value) SetRangeValue(value NSRange /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setRangeValue:"), value)
}


// The Foundation rectangle structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/rectvalue
func (v_ Value) RectValue() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("rectValue"))
	return rv
}


// The Foundation rectangle structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/rectvalue
func (v_ Value) SetRectValue(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setRectValue:"), value)
}


// The Scene Kit 4 x 4 matrix representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/scnmatrix4value
func (v_ Value) ScnMatrix4Value() NMatrix4 /* not a class type */ {
	rv := objc.Send[NMatrix4](v_.ID, objc.Sel("scnMatrix4Value"))
	return rv
}


// The Scene Kit 4 x 4 matrix representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/scnmatrix4value
func (v_ Value) SetScnMatrix4Value(value NMatrix4 /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setScnMatrix4Value:"), value)
}


// The three-element Scene Kit vector representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/scnvector3value
func (v_ Value) ScnVector3Value() NVector3 /* not a class type */ {
	rv := objc.Send[NVector3](v_.ID, objc.Sel("scnVector3Value"))
	return rv
}


// The three-element Scene Kit vector representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/scnvector3value
func (v_ Value) SetScnVector3Value(value NVector3 /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setScnVector3Value:"), value)
}


// The four-element Scene Kit vector representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/scnvector4value
func (v_ Value) ScnVector4Value() NVector4 /* not a class type */ {
	rv := objc.Send[NVector4](v_.ID, objc.Sel("scnVector4Value"))
	return rv
}


// The four-element Scene Kit vector representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/scnvector4value
func (v_ Value) SetScnVector4Value(value NVector4 /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setScnVector4Value:"), value)
}


// The Foundation size structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/sizevalue
func (v_ Value) SizeValue() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[Size](v_.ID, objc.Sel("sizeValue"))
	return rv
}


// The Foundation size structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/sizevalue
func (v_ Value) SetSizeValue(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSizeValue:"), value)
}


// The CoreMedia time mapping structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/timemappingvalue
func (v_ Value) TimeMappingValue() TimeMapping /* not a class type */ {
	rv := objc.Send[TimeMapping](v_.ID, objc.Sel("timeMappingValue"))
	return rv
}


// The CoreMedia time mapping structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/timemappingvalue
func (v_ Value) SetTimeMappingValue(value TimeMapping /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setTimeMappingValue:"), value)
}


// The CoreMedia time range structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/timerangevalue
func (v_ Value) TimeRangeValue() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](v_.ID, objc.Sel("timeRangeValue"))
	return rv
}


// The CoreMedia time range structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/timerangevalue
func (v_ Value) SetTimeRangeValue(value TimeRange /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setTimeRangeValue:"), value)
}


// The CoreMedia time structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/timevalue
func (v_ Value) TimeValue() Time /* not a class type */ {
	rv := objc.Send[Time](v_.ID, objc.Sel("timeValue"))
	return rv
}


// The CoreMedia time structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/timevalue
func (v_ Value) SetTimeValue(value Time /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setTimeValue:"), value)
}


// Returns the UIKit edge insets structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/uiedgeinsetsvalue
func (v_ Value) UiEdgeInsetsValue() UIEdgeInsets /* not a class type */ {
	rv := objc.Send[EdgeInsets](v_.ID, objc.Sel("uiEdgeInsetsValue"))
	return rv
}


// Returns the UIKit edge insets structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/uiedgeinsetsvalue
func (v_ Value) SetUiEdgeInsetsValue(value UIEdgeInsets /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setUiEdgeInsetsValue:"), value)
}


// Returns the UIKit offset structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/uioffsetvalue
func (v_ Value) UiOffsetValue() Offset /* not a class type */ {
	rv := objc.Send[Offset](v_.ID, objc.Sel("uiOffsetValue"))
	return rv
}


// Returns the UIKit offset structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/uioffsetvalue
func (v_ Value) SetUiOffsetValue(value Offset /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setUiOffsetValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/videodimensionsvalue
func (v_ Value) VideoDimensionsValue() VideoDimensions /* not a class type */ {
	rv := objc.Send[VideoDimensions](v_.ID, objc.Sel("videoDimensionsValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsvalue/videodimensionsvalue
func (v_ Value) SetVideoDimensionsValue(value VideoDimensions /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVideoDimensionsValue:"), value)
}



