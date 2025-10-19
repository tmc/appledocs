// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVAssetWriter] class.
var (
	aVAssetWriterClass     _AVAssetWriterClass
	aVAssetWriterClassOnce sync.Once
)

func getAVAssetWriterClass() _AVAssetWriterClass {
	aVAssetWriterClassOnce.Do(func() {
		aVAssetWriterClass = _AVAssetWriterClass{objc.GetClass("AVAssetWriter")}
	})
	return aVAssetWriterClass
}

type _AVAssetWriterClass struct {
	class objc.Class
}

// An interface definition for the [AVAssetWriter] class.
type IAVAssetWriter interface {
	objectivec.IObject
}

// An object that writes media data to a container file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter
type AVAssetWriter struct {
	objectivec.Object
}

// AVAssetWriterFrom constructs a [AVAssetWriter] from an unsafe.Pointer.
//
// An object that writes media data to a container file.
func AVAssetWriterFrom(ptr unsafe.Pointer) AVAssetWriter {
	return AVAssetWriter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVAssetWriterClass) Alloc() AVAssetWriter {
	rv := objc.Send[AVAssetWriter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVAssetWriterClass) New() AVAssetWriter {
	rv := objc.Send[AVAssetWriter](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVAssetWriter) Init() AVAssetWriter {
	rv := objc.Send[AVAssetWriter](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVAssetWriter) Autorelease() AVAssetWriter {
	rv := objc.Send[AVAssetWriter](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetWriter creates a new AVAssetWriter instance.
func NewAVAssetWriter() AVAssetWriter {
	return getAVAssetWriterClass().New()
}




