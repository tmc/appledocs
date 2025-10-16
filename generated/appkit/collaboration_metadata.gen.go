
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [collaborationMetadata] class.
var collaborationMetadataClass _collaborationMetadataClass

func init() {
	collaborationMetadataClass = _collaborationMetadataClass{objc.GetClass("collaborationMetadata")}
}

type _collaborationMetadataClass struct {
	objc.Class
}

// An interface definition for the [collaborationMetadata] class.
type IcollaborationMetadata interface {
	ID() objc.ID
}

type collaborationMetadata struct {
	id objc.ID
}

func collaborationMetadataFrom(ptr unsafe.Pointer) collaborationMetadata {
	return collaborationMetadata{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ collaborationMetadata) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _collaborationMetadataClass) Alloc() collaborationMetadata {
	rv := objc.Send[collaborationMetadata](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _collaborationMetadataClass) New() collaborationMetadata {
	rv := objc.Send[collaborationMetadata](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcollaborationMetadata creates and returns a new initialized instance.
func NewcollaborationMetadata() collaborationMetadata {
	return collaborationMetadataClass.New()
}

// Init initializes the instance.
func (c_ collaborationMetadata) Init() collaborationMetadata {
	rv := objc.Send[collaborationMetadata](c_.ID(), selInit)
	return rv
}
