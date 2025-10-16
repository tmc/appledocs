
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [autosavesConfiguration] class.
var autosavesConfigurationClass _autosavesConfigurationClass

func init() {
	autosavesConfigurationClass = _autosavesConfigurationClass{objc.GetClass("autosavesConfiguration")}
}

type _autosavesConfigurationClass struct {
	objc.Class
}

// An interface definition for the [autosavesConfiguration] class.
type IautosavesConfiguration interface {
	ID() objc.ID
}

type autosavesConfiguration struct {
	id objc.ID
}

func autosavesConfigurationFrom(ptr unsafe.Pointer) autosavesConfiguration {
	return autosavesConfiguration{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ autosavesConfiguration) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _autosavesConfigurationClass) Alloc() autosavesConfiguration {
	rv := objc.Send[autosavesConfiguration](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _autosavesConfigurationClass) New() autosavesConfiguration {
	rv := objc.Send[autosavesConfiguration](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewautosavesConfiguration creates and returns a new initialized instance.
func NewautosavesConfiguration() autosavesConfiguration {
	return autosavesConfigurationClass.New()
}

// Init initializes the instance.
func (a_ autosavesConfiguration) Init() autosavesConfiguration {
	rv := objc.Send[autosavesConfiguration](a_.ID(), selInit)
	return rv
}
