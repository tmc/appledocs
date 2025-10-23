// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/coreimage"
)

// The class instance for the [MetadataMachineReadableCodeObject] class.
var (
	MetadataMachineReadableCodeObjectClass     _MetadataMachineReadableCodeObjectClass
	MetadataMachineReadableCodeObjectClassOnce sync.Once
)

func getMetadataMachineReadableCodeObjectClass() _MetadataMachineReadableCodeObjectClass {
	MetadataMachineReadableCodeObjectClassOnce.Do(func() {
		MetadataMachineReadableCodeObjectClass = _MetadataMachineReadableCodeObjectClass{objc.GetClass("AVMetadataMachineReadableCodeObject")}
	})
	return MetadataMachineReadableCodeObjectClass
}

type _MetadataMachineReadableCodeObjectClass struct {
	class objc.Class
}

// An interface definition for the [MetadataMachineReadableCodeObject] class.
type IMetadataMachineReadableCodeObject interface {
	IMetadataObject
	Corners() coregraphics.CGPoint
	SetCorners(value coregraphics.CGPoint)
	Descriptor() coreimage.BarcodeDescriptor
	SetDescriptor(value coreimage.IBarcodeDescriptor)
	StringValue() string
	SetStringValue(value string)
}

// Barcode information detected by a metadata capture output.
//
// The class is a concrete subclass of defining the features of a detected one-dimensional or two-dimensional barcode. An instance represents a single detected machine readable code in an image.  It’s an immutable object describing the features and payload of a barcode. On supported platforms, the class outputs arrays of detected machine readable code objects.


// Barcode information detected by a metadata capture output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataMachineReadableCodeObject
type MetadataMachineReadableCodeObject struct {
	MetadataObject
}

// MetadataMachineReadableCodeObjectFrom constructs a [MetadataMachineReadableCodeObject] from an unsafe.Pointer.
//
// Barcode information detected by a metadata capture output.
func MetadataMachineReadableCodeObjectFrom(ptr unsafe.Pointer) MetadataMachineReadableCodeObject {
	return MetadataMachineReadableCodeObject{
		MetadataObject: MetadataObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MetadataMachineReadableCodeObjectClass) Alloc() MetadataMachineReadableCodeObject {
	rv := objc.Send[MetadataMachineReadableCodeObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MetadataMachineReadableCodeObjectClass) New() MetadataMachineReadableCodeObject {
	rv := objc.Send[MetadataMachineReadableCodeObject](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataMachineReadableCodeObject) Init() MetadataMachineReadableCodeObject {
	rv := objc.Send[MetadataMachineReadableCodeObject](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataMachineReadableCodeObject) Autorelease() MetadataMachineReadableCodeObject {
	rv := objc.Send[MetadataMachineReadableCodeObject](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataMachineReadableCodeObject creates a new MetadataMachineReadableCodeObject instance.
func NewMetadataMachineReadableCodeObject() MetadataMachineReadableCodeObject {
	return getMetadataMachineReadableCodeObjectClass().New()
}



// A Swift array of corner points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatamachinereadablecodeobject/corners-58qbe
func (m_ MetadataMachineReadableCodeObject) Corners() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](m_.ID, objc.Sel("corners"))
	return rv
}


// A Swift array of corner points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatamachinereadablecodeobject/corners-58qbe
func (m_ MetadataMachineReadableCodeObject) SetCorners(value coregraphics.CGPoint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCorners:"), value)
}


// A barcode description for use in Core Image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatamachinereadablecodeobject/descriptor
func (m_ MetadataMachineReadableCodeObject) Descriptor() coreimage.BarcodeDescriptor {
	rv := objc.Send[coreimage.BarcodeDescriptor](m_.ID, objc.Sel("descriptor"))
	return rv
}


// A barcode description for use in Core Image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatamachinereadablecodeobject/descriptor
func (m_ MetadataMachineReadableCodeObject) SetDescriptor(value coreimage.IBarcodeDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDescriptor:"), value)
}


// Returns the error-corrected data decoded into a human-readable string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatamachinereadablecodeobject/stringvalue
func (m_ MetadataMachineReadableCodeObject) StringValue() string {
	rv := objc.Send[string](m_.ID, objc.Sel("stringValue"))
	return rv
}


// Returns the error-corrected data decoded into a human-readable string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatamachinereadablecodeobject/stringvalue
func (m_ MetadataMachineReadableCodeObject) SetStringValue(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStringValue:"), objc.String(value))
}



