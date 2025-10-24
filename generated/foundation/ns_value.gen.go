// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	CATransform3DValue() objectivec.IObject
	EdgeInsetsValue() objc.IObject /* cross-framework: EdgeInsets */
	GCPoint2Value() unsafe.Pointer
	MKCoordinateSpanValue() unsafe.Pointer
	MKCoordinateValue() LocationCoordinate2D /* not a class type */
	NonretainedObjectValue() objc.ID
	ObjCType() unsafe.Pointer
	PointValue() objc.IObject /* cross-framework: Point */
	PointerValue() unsafe.Pointer
	RangeValue() objc.IObject /* cross-framework: Range */
	RectValue() objc.IObject /* cross-framework: Rect */
	SCNMatrix4Value() NMatrix4 /* not a class type */
	SCNVector3Value() NVector3 /* not a class type */
	SCNVector4Value() NVector4 /* not a class type */
	SizeValue() objc.IObject /* cross-framework: Size */
	CMTimeMappingValue() TimeMapping /* not a class type */
	CMTimeRangeValue() TimeRange /* not a class type */
	CMTimeValue() Time /* not a class type */
	CMVideoDimensionsValue() VideoDimensions /* not a class type */
	Hash() int
	SetHash(value int)
	TimeMappingValue() TimeMapping /* not a class type */
	SetTimeMappingValue(value TimeMapping /* not a class type */)
	TimeRangeValue() TimeRange /* not a class type */
	SetTimeRangeValue(value TimeRange /* not a class type */)
	TimeValue() Time /* not a class type */
	SetTimeValue(value Time /* not a class type */)
	VideoDimensionsValue() VideoDimensions /* not a class type */
	SetVideoDimensionsValue(value VideoDimensions /* not a class type */)
	// methods:
	GetValueSize(value unsafe.Pointer, size uint)
	IsEqualToValue(value IValue) bool
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



// Initializes a value object to contain the specified value, interpreted with the specified Objective-C type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(bytes:objCType:)
func NewValueWithBytesObjCType(value unsafe.Pointer, type_ unsafe.Pointer) Value {
	instance := getValueClass().Alloc()
	rv := objc.Send[Value](instance.ID, objc.Sel("initWithBytes:objCType:"), value, type_)
	rv.Autorelease()
	return rv
}


// Creates a new value object containing the specified CoreAnimation transform structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CATransform3D:)
func NewValueWithCATransform3D(t objectivec.IObject) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithCATransform3D:"), t)
	return rv
}


// Creates a new value object containing the specified CoreGraphics affine transform structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CGAffineTransform:)
func NewValueWithCGAffineTransform(transform corefoundation.CGAffineTransform) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithCGAffineTransform:"), transform)
	return rv
}


// Creates a new value object containing the specified CoreGraphics point structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CGPoint:)
func NewValueWithCGPoint(point objc.IObject /* cross-framework: Point */) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithCGPoint:"), point)
	return rv
}


// Creates a new value object containing the specified CoreGraphics rectangle structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CGRect:)
func NewValueWithCGRect(rect objc.IObject /* cross-framework: Rect */) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithCGRect:"), rect)
	return rv
}


// Creates a new value object containing the specified CoreGraphics size structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CGSize:)
func NewValueWithCGSize(size objc.IObject /* cross-framework: Size */) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithCGSize:"), size)
	return rv
}


// Creates a new value object containing the specified CoreGraphics vector structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CGVector:)
func NewValueWithCGVector(vector objc.IObject /* cross-framework: Vector */) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithCGVector:"), vector)
	return rv
}


// Creates a new value object containing the specified CoreMedia time structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CMTime:)
func NewValueWithCMTime(time Time /* not a class type */) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithCMTime:"), time)
	return rv
}


// Creates a new value object containing the specified CoreMedia time mapping structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CMTimeMapping:)
func NewValueWithCMTimeMapping(timeMapping TimeMapping /* not a class type */) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithCMTimeMapping:"), timeMapping)
	return rv
}


// Creates a new value object containing the specified CoreMedia time range structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CMTimeRange:)
func NewValueWithCMTimeRange(timeRange TimeRange /* not a class type */) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithCMTimeRange:"), timeRange)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CMVideoDimensions:)
func NewValueWithCMVideoDimensions(dimensions VideoDimensions /* not a class type */) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithCMVideoDimensions:"), dimensions)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(coder:)
func NewValueWithCoder(coder ICoder) Value {
	instance := getValueClass().Alloc()
	rv := objc.Send[Value](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(directionalEdgeInsets:)
func NewValueWithDirectionalEdgeInsets(insets objectivec.IObject) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithDirectionalEdgeInsets:"), insets)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(edgeInsets:)
func NewValueWithEdgeInsets(insets objc.IObject /* cross-framework: EdgeInsets */) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithEdgeInsets:"), insets)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(GCPoint2:)
func NewValueWithGCPoint2(point unsafe.Pointer) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithGCPoint2:"), point)
	return rv
}


// Creates a new value object containing the specified CoreLocation geographic coordinate structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(MKCoordinate:)
func NewValueWithMKCoordinate(coordinate LocationCoordinate2D /* not a class type */) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithMKCoordinate:"), coordinate)
	return rv
}


// Creates a new value object containing the specified MapKit coordinate span structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(MKCoordinateSpan:)
func NewValueWithMKCoordinateSpan(span unsafe.Pointer) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithMKCoordinateSpan:"), span)
	return rv
}


// Creates a value object containing the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(nonretainedObject:)
func NewValueWithNonretainedObject(anObject objc.IObject) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithNonretainedObject:"), anObject)
	return rv
}


// Creates a value object containing the specified value, interpreted with the specified Objective-C type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(_:withObjCType:)
func NewValueWithObjCType(value unsafe.Pointer, type_ unsafe.Pointer) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("value:withObjCType:"), value, type_)
	return rv
}


// Creates a new value object containing the specified Foundation point structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(point:)
func NewValueWithPoint(point objc.IObject /* cross-framework: Point */) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithPoint:"), point)
	return rv
}


// Creates a value object containing the specified pointer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(pointer:)
func NewValueWithPointer(pointer unsafe.Pointer) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithPointer:"), pointer)
	return rv
}


// Creates a new value object containing the specified Foundation range structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(range:)
func NewValueWithRange(range_ objc.IObject /* cross-framework: Range */) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithRange:"), range_)
	return rv
}


// Creates a new value object containing the specified Foundation rectangle structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(rect:)
func NewValueWithRect(rect objc.IObject /* cross-framework: Rect */) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithRect:"), rect)
	return rv
}


// Creates a value object that contains the specified SceneKit 4 x 4 matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(SCNMatrix4:)
func NewValueWithSCNMatrix4(v NMatrix4 /* not a class type */) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithSCNMatrix4:"), v)
	return rv
}


// Creates a value object that contains the specified three-element SceneKit vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(SCNVector3:)
func NewValueWithSCNVector3(v NVector3 /* not a class type */) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithSCNVector3:"), v)
	return rv
}


// Creates a value object that contains the specified four-element SceneKit vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(SCNVector4:)
func NewValueWithSCNVector4(v NVector4 /* not a class type */) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithSCNVector4:"), v)
	return rv
}


// Creates a new value object containing the specified Foundation size structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(size:)
func NewValueWithSize(size objc.IObject /* cross-framework: Size */) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithSize:"), size)
	return rv
}


// Creates a new value object containing the specified UIKit edge insets structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(UIEdgeInsets:)
func NewValueWithUIEdgeInsets(insets EdgeInsets /* not a class type */) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithUIEdgeInsets:"), insets)
	return rv
}


// Creates a new value object containing the specified UIKit offset structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(UIOffset:)
func NewValueWithUIOffset(insets objectivec.IObject) Value {
	rv := objc.Send[Value](objc.ID(getValueClass().class), objc.Sel("valueWithUIOffset:"), insets)
	return rv
}



// Creates a new value object containing the specified CoreAnimation transform structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CATransform3D:)
func (vc _ValueClass) ValueWithCATransform3D(t objectivec.IObject) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithCATransform3D:"), t)
	return rv
}


// Creates a new value object containing the specified CoreGraphics affine transform structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CGAffineTransform:)
func (vc _ValueClass) ValueWithCGAffineTransform(transform corefoundation.CGAffineTransform) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithCGAffineTransform:"), transform)
	return rv
}


// Creates a new value object containing the specified CoreGraphics point structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CGPoint:)
func (vc _ValueClass) ValueWithCGPoint(point objc.IObject /* cross-framework: Point */) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithCGPoint:"), point)
	return rv
}


// Creates a new value object containing the specified CoreGraphics rectangle structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CGRect:)
func (vc _ValueClass) ValueWithCGRect(rect objc.IObject /* cross-framework: Rect */) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithCGRect:"), rect)
	return rv
}


// Creates a new value object containing the specified CoreGraphics size structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CGSize:)
func (vc _ValueClass) ValueWithCGSize(size objc.IObject /* cross-framework: Size */) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithCGSize:"), size)
	return rv
}


// Creates a new value object containing the specified CoreGraphics vector structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CGVector:)
func (vc _ValueClass) ValueWithCGVector(vector objc.IObject /* cross-framework: Vector */) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithCGVector:"), vector)
	return rv
}


// Creates a new value object containing the specified CoreMedia time structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CMTime:)
func (vc _ValueClass) ValueWithCMTime(time Time /* not a class type */) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithCMTime:"), time)
	return rv
}


// Creates a new value object containing the specified CoreMedia time mapping structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CMTimeMapping:)
func (vc _ValueClass) ValueWithCMTimeMapping(timeMapping TimeMapping /* not a class type */) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithCMTimeMapping:"), timeMapping)
	return rv
}


// Creates a new value object containing the specified CoreMedia time range structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CMTimeRange:)
func (vc _ValueClass) ValueWithCMTimeRange(timeRange TimeRange /* not a class type */) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithCMTimeRange:"), timeRange)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(CMVideoDimensions:)
func (vc _ValueClass) ValueWithCMVideoDimensions(dimensions VideoDimensions /* not a class type */) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithCMVideoDimensions:"), dimensions)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(GCPoint2:)
func (vc _ValueClass) ValueWithGCPoint2(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("valueWithGCPoint2:"), point)
	return rv
}


// Creates a new value object containing the specified CoreLocation geographic coordinate structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(MKCoordinate:)
func (vc _ValueClass) ValueWithMKCoordinate(coordinate LocationCoordinate2D /* not a class type */) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithMKCoordinate:"), coordinate)
	return rv
}


// Creates a new value object containing the specified MapKit coordinate span structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(MKCoordinateSpan:)
func (vc _ValueClass) ValueWithMKCoordinateSpan(span unsafe.Pointer) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithMKCoordinateSpan:"), span)
	return rv
}


// Creates a value object that contains the specified SceneKit 4 x 4 matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(SCNMatrix4:)
func (vc _ValueClass) ValueWithSCNMatrix4(v NMatrix4 /* not a class type */) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithSCNMatrix4:"), v)
	return rv
}


// Creates a value object that contains the specified three-element SceneKit vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(SCNVector3:)
func (vc _ValueClass) ValueWithSCNVector3(v NVector3 /* not a class type */) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithSCNVector3:"), v)
	return rv
}


// Creates a value object that contains the specified four-element SceneKit vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(SCNVector4:)
func (vc _ValueClass) ValueWithSCNVector4(v NVector4 /* not a class type */) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithSCNVector4:"), v)
	return rv
}


// Creates a new value object containing the specified UIKit edge insets structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(UIEdgeInsets:)
func (vc _ValueClass) ValueWithUIEdgeInsets(insets EdgeInsets /* not a class type */) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithUIEdgeInsets:"), insets)
	return rv
}


// Creates a new value object containing the specified UIKit offset structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(UIOffset:)
func (vc _ValueClass) ValueWithUIOffset(insets objectivec.IObject) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithUIOffset:"), insets)
	return rv
}


// Creates a value object containing the specified value, interpreted with the specified Objective-C type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(_:withObjCType:)
func (vc _ValueClass) ValueWithObjCType(value unsafe.Pointer, type_ unsafe.Pointer) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("value:withObjCType:"), value, type_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(directionalEdgeInsets:)
func (vc _ValueClass) ValueWithDirectionalEdgeInsets(insets objectivec.IObject) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithDirectionalEdgeInsets:"), insets)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(edgeInsets:)
func (vc _ValueClass) ValueWithEdgeInsets(insets objc.IObject /* cross-framework: EdgeInsets */) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithEdgeInsets:"), insets)
	return rv
}


// Creates a value object containing the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(nonretainedObject:)
func (vc _ValueClass) ValueWithNonretainedObject(anObject objc.IObject) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithNonretainedObject:"), anObject)
	return rv
}


// Creates a new value object containing the specified Foundation point structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(point:)
func (vc _ValueClass) ValueWithPoint(point objc.IObject /* cross-framework: Point */) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithPoint:"), point)
	return rv
}


// Creates a value object containing the specified pointer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(pointer:)
func (vc _ValueClass) ValueWithPointer(pointer unsafe.Pointer) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithPointer:"), pointer)
	return rv
}


// Creates a new value object containing the specified Foundation range structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(range:)
func (vc _ValueClass) ValueWithRange(range_ objc.IObject /* cross-framework: Range */) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithRange:"), range_)
	return rv
}


// Creates a new value object containing the specified Foundation rectangle structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(rect:)
func (vc _ValueClass) ValueWithRect(rect objc.IObject /* cross-framework: Rect */) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithRect:"), rect)
	return rv
}


// Creates a new value object containing the specified Foundation size structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/init(size:)
func (vc _ValueClass) ValueWithSize(size objc.IObject /* cross-framework: Size */) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithSize:"), size)
	return rv
}


// Creates a value object containing the specified value, interpreted with the specified Objective-C type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/valueWithBytes:objCType:
func (vc _ValueClass) ValueWithBytesObjCType(value unsafe.Pointer, type_ unsafe.Pointer) IValue {
	rv := objc.Send[Value](objc.ID(vc.class), objc.Sel("valueWithBytes:objCType:"), value, type_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/getValue(_:size:)
func (v_ Value) GetValueSize(value unsafe.Pointer, size uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("getValue:size:"), value, size)
}


// Returns a Boolean value that indicates whether the value object and another value object are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/isEqual(to:)
func (v_ Value) IsEqualToValue(value IValue) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isEqualToValue:"), value)
	return rv
}


// The CoreAnimation transform structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/caTransform3DValue
func (v_ Value) CATransform3DValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("CATransform3DValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/edgeInsetsValue
func (v_ Value) EdgeInsetsValue() objc.IObject /* cross-framework: EdgeInsets */ {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("edgeInsetsValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/gcPoint2Value
func (v_ Value) GCPoint2Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("GCPoint2Value"))
	return rv
}


// The MapKit coordinate span structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/mkCoordinateSpanValue
func (v_ Value) MKCoordinateSpanValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("MKCoordinateSpanValue"))
	return rv
}


// The CoreLocation geographic coordinate structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/mkCoordinateValue
func (v_ Value) MKCoordinateValue() LocationCoordinate2D /* not a class type */ {
	rv := objc.Send[LocationCoordinate2D](v_.ID, objc.Sel("MKCoordinateValue"))
	return rv
}


// The value as a non-retained pointer to an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/nonretainedObjectValue
func (v_ Value) NonretainedObjectValue() objc.ID {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("nonretainedObjectValue"))
	return rv
}


// A C string containing the Objective-C type of the data contained in the value object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/objCType
func (v_ Value) ObjCType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("objCType"))
	return rv
}


// The Foundation point structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/pointValue
func (v_ Value) PointValue() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("pointValue"))
	return rv
}


// Returns the value as an untyped pointer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/pointerValue
func (v_ Value) PointerValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("pointerValue"))
	return rv
}


// The Foundation range structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/rangeValue
func (v_ Value) RangeValue() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("rangeValue"))
	return rv
}


// The Foundation rectangle structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/rectValue
func (v_ Value) RectValue() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("rectValue"))
	return rv
}


// The Scene Kit 4 x 4 matrix representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/scnMatrix4Value
func (v_ Value) SCNMatrix4Value() NMatrix4 /* not a class type */ {
	rv := objc.Send[NMatrix4](v_.ID, objc.Sel("SCNMatrix4Value"))
	return rv
}


// The three-element Scene Kit vector representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/scnVector3Value
func (v_ Value) SCNVector3Value() NVector3 /* not a class type */ {
	rv := objc.Send[NVector3](v_.ID, objc.Sel("SCNVector3Value"))
	return rv
}


// The four-element Scene Kit vector representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/scnVector4Value
func (v_ Value) SCNVector4Value() NVector4 /* not a class type */ {
	rv := objc.Send[NVector4](v_.ID, objc.Sel("SCNVector4Value"))
	return rv
}


// The Foundation size structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/sizeValue
func (v_ Value) SizeValue() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("sizeValue"))
	return rv
}


// The CoreMedia time mapping structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/timeMappingValue
func (v_ Value) CMTimeMappingValue() TimeMapping /* not a class type */ {
	rv := objc.Send[TimeMapping](v_.ID, objc.Sel("CMTimeMappingValue"))
	return rv
}


// The CoreMedia time range structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/timeRangeValue
func (v_ Value) CMTimeRangeValue() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](v_.ID, objc.Sel("CMTimeRangeValue"))
	return rv
}


// The CoreMedia time structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/timeValue
func (v_ Value) CMTimeValue() Time /* not a class type */ {
	rv := objc.Send[Time](v_.ID, objc.Sel("CMTimeValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/videoDimensionsValue
func (v_ Value) CMVideoDimensionsValue() VideoDimensions /* not a class type */ {
	rv := objc.Send[VideoDimensions](v_.ID, objc.Sel("CMVideoDimensionsValue"))
	return rv
}


// Returns an integer that can be used as a table address in a hash table structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash
func (v_ Value) Hash() int {
	rv := objc.Send[int](v_.ID, objc.Sel("hash"))
	return rv
}


// Returns an integer that can be used as a table address in a hash table structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash
func (v_ Value) SetHash(value int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setHash:"), value)
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


