// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [QCRenderer] class.
var (
	QCRendererClass     _QCRendererClass
	QCRendererClassOnce sync.Once
)

func getQCRendererClass() _QCRendererClass {
	QCRendererClassOnce.Do(func() {
		QCRendererClass = _QCRendererClass{objc.GetClass("QCRenderer")}
	})
	return QCRendererClass
}

type _QCRendererClass struct {
	class objc.Class
}

// An interface definition for the [QCRenderer] class.
type IQCRenderer interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A base class for low-level rendering.
//
// A class is designed for low-level rendering of Quartz Composer compositions. This is the class to use if you want to be in charge of rendering a composition to a specific OpenGL context—either using the class or a object. also allows you to load, play, and control a composition. To render a composition to a specific OpenGL context: Create an instance of using one of the initialization methods, such as . Render frames by calling the method If you use double buffering in OpenGL, you must swap the OpenGL buffers. Release the renderer when you no longer need it. This code snippet shows how to implement these tasks:


// A base class for low-level rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCRenderer
type QCRenderer struct {
	objectivec.Object
}

// QCRendererFrom constructs a [QCRenderer] from an unsafe.Pointer.
//
// A base class for low-level rendering.
func QCRendererFrom(ptr unsafe.Pointer) QCRenderer {
	return QCRenderer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (qc _QCRendererClass) Alloc() QCRenderer {
	rv := objc.Send[QCRenderer](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (qc _QCRendererClass) New() QCRenderer {
	rv := objc.Send[QCRenderer](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QCRenderer) Init() QCRenderer {
	rv := objc.Send[QCRenderer](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QCRenderer) Autorelease() QCRenderer {
	rv := objc.Send[QCRenderer](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQCRenderer creates a new QCRenderer instance.
func NewQCRenderer() QCRenderer {
	return getQCRendererClass().New()
}




