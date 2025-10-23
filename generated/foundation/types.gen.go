// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation
import (
	"unsafe"
)

// CFPropertyListRef is a CoreGraphics opaque type.
type CFPropertyListRef unsafe.Pointer

// CFRunLoopSourceRef is a CoreGraphics opaque type.
type CFRunLoopSourceRef unsafe.Pointer


// Foundation-specific types

// Foundation geometry types - compatible with NSPoint, NSSize, NSRect, NSRange
type Point struct {
	X float64
	Y float64
}

type Size struct {
	Width  float64
	Height float64
}

type Rect struct {
	Origin Point
	Size   Size
}

type Range struct {
	Location int
	Length   int
}


// Fallback type aliases for undefined types
// These types are referenced in method signatures but not fully documented.
// Using inferred base types as fallback to allow code generation.
type AccessibilityNotifications = int

type AccessibilityRole = int

type AccessibilitySubrole = int

type AccessibilityZoomType = int

type AffineTransformStruct = int

type AllocatorContext = int

type ArrayCallBacks = int

type AttributedStringCompletionHandler = int

type BagCallBacks = int

type BinaryHeapCallBacks = int

type BinaryHeapCompareContext = int

type Bool = int

type CFString = int

type CFURL = int

type ColorRenderingIntent = int

type Comparator = int

type CompositingOperation = int

type Decimal = int

type DictionaryKeyCallBacks = int

type DictionaryValueCallBacks = int

type EdgeInsets = int

type FileDescriptorNativeDescriptor = int

type FontManagerScope = int

type FontTraitMask = int

type FormattingUnitStyle = int

type GregorianDate = int

type GregorianUnits = int

type HashEnumerator = int

type HashTableCallBacks = int

type ImageSymbolWeight = int

type KeyValueObservingOptions = int

type MapEnumerator = int

type MapTableKeyCallBacks = int

type MapTableValueCallBacks = int

type MultibyteGlyphPacking = int

type NetworkReachabilityCallBack = int

type NetworkReachabilityContext = int

type Offset = int

type OpenGLGlobalOption = int

type PropertyListRef = int

type RunLoopSourceRef = int

type SocketCallBack = int

type SocketContext = int

type SocketError = int

type SocketSignature = int

type StreamClientContext = int

type StreamError = int

type UIEdgeInsets = int

type UInteger = int

type UUIDBytes = int

type Zone = int



