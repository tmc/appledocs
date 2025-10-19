// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVContentKeySpecifier] class.
var (
	aVContentKeySpecifierClass     _AVContentKeySpecifierClass
	aVContentKeySpecifierClassOnce sync.Once
)

func getAVContentKeySpecifierClass() _AVContentKeySpecifierClass {
	aVContentKeySpecifierClassOnce.Do(func() {
		aVContentKeySpecifierClass = _AVContentKeySpecifierClass{objc.GetClass("AVContentKeySpecifier")}
	})
	return aVContentKeySpecifierClass
}

type _AVContentKeySpecifierClass struct {
	class objc.Class
}

// An interface definition for the [AVContentKeySpecifier] class.
type IAVContentKeySpecifier interface {
	objectivec.IObject
}

// An object that uniquely identifies a content key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySpecifier
type AVContentKeySpecifier struct {
	objectivec.Object
}

// AVContentKeySpecifierFrom constructs a [AVContentKeySpecifier] from an unsafe.Pointer.
//
// An object that uniquely identifies a content key.
func AVContentKeySpecifierFrom(ptr unsafe.Pointer) AVContentKeySpecifier {
	return AVContentKeySpecifier{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVContentKeySpecifierClass) Alloc() AVContentKeySpecifier {
	rv := objc.Send[AVContentKeySpecifier](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVContentKeySpecifierClass) New() AVContentKeySpecifier {
	rv := objc.Send[AVContentKeySpecifier](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVContentKeySpecifier) Init() AVContentKeySpecifier {
	rv := objc.Send[AVContentKeySpecifier](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVContentKeySpecifier) Autorelease() AVContentKeySpecifier {
	rv := objc.Send[AVContentKeySpecifier](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVContentKeySpecifier creates a new AVContentKeySpecifier instance.
func NewAVContentKeySpecifier() AVContentKeySpecifier {
	return getAVContentKeySpecifierClass().New()
}




