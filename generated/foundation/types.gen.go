// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation
import (
	"unsafe"
)

// CFPlugInRef is a CoreGraphics opaque type.
type CFPlugInRef unsafe.Pointer

// CFRunLoopRef is a CoreGraphics opaque type.
type CFRunLoopRef unsafe.Pointer


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
type AbsoluteTime = int

type AccessibilityAssistiveTechnologyIdentifier = int

type AccessibilityNotifications = int

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

type Comparator = int

type ComparatorFunction = int

type ComparisonResult = int

type CompositingOperation = int

type DateFormatterKey = int

type Decimal = int

type DictionaryKeyCallBacks = int

type DictionaryValueCallBacks = int

type EdgeInsets = int

type FileDescriptorNativeDescriptor = int

type FontTraitMask = int

type FormattingUnitStyle = int

type HashCode = int

type HashEnumerator = int

type HashTableCallBacks = int

type ImageSymbolWeight = int

type KeyValueObservingOptions = int

type LocaleIdentifier = int

type MapEnumerator = int

type MapTableKeyCallBacks = int

type MapTableValueCallBacks = int

type PlugInRef = int

type RunLoopMode = int

type RunLoopRef = int

type StreamClientContext = int

type StreamError = int

type StreamPropertyKey = int

type StreamStatus = int

type UIEdgeInsets = int

type UInteger = int

type UUIDBytes = int

type Zone = int



