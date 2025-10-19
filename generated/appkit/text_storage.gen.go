// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [TextStorage] class.
var textStorageClass = _TextStorageClass{objc.GetClass("NSTextStorage")}

type _TextStorageClass struct {
	class objc.Class
}

// An interface definition for the [TextStorage] class.
type ITextStorage interface {
	foundation.IMutableAttributedString
}

// The fundamental storage mechanism of TextKit that contains the text managed by the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage

type TextStorage struct {
	foundation.MutableAttributedString
}

// TextStorageFrom constructs a [TextStorage] from an unsafe.Pointer.
//
// The fundamental storage mechanism of TextKit that contains the text managed by the system.
func TextStorageFrom(ptr unsafe.Pointer) TextStorage {
	return TextStorage{
		MutableAttributedString: foundation.MutableAttributedStringFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (tc _TextStorageClass) Alloc() TextStorage {
	rv := objc.Send[TextStorage](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _TextStorageClass) New() TextStorage {
	rv := objc.Send[TextStorage](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextStorage) Init() TextStorage {
	rv := objc.Send[TextStorage](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextStorage) Autorelease() TextStorage {
	rv := objc.Send[TextStorage](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextStorage creates a new TextStorage instance.
func NewTextStorage() TextStorage {
	return textStorageClass.New()
}




