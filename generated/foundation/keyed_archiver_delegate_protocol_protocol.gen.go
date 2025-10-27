// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PKeyedArchiverDelegate is the NSKeyedArchiverDelegate protocol interface.
//
// The optional methods implemented by the delegate of a keyed archiver.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/NSKeyedArchiverDelegate
type PKeyedArchiverDelegate interface {
	// Optional methods
	ArchiverDidEncodeObject(archiver IKeyedArchiver, object objectivec.IObject)
	HasArchiverDidEncodeObject() bool
	ArchiverWillEncodeObject(archiver IKeyedArchiver, object objectivec.IObject) objc.ID
	HasArchiverWillEncodeObject() bool
	ArchiverWillReplaceObjectWithObject(archiver IKeyedArchiver, object objectivec.IObject, newObject objectivec.IObject)
	HasArchiverWillReplaceObjectWithObject() bool
	ArchiverDidFinish(archiver IKeyedArchiver)
	HasArchiverDidFinish() bool
	ArchiverWillFinish(archiver IKeyedArchiver)
	HasArchiverWillFinish() bool
}

// KeyedArchiverDelegate is a delegate implementation builder for the PKeyedArchiverDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type KeyedArchiverDelegate struct {
	_ArchiverDidEncodeObject func(archiver IKeyedArchiver, object objectivec.IObject)
	_ArchiverWillEncodeObject func(archiver IKeyedArchiver, object objectivec.IObject) objc.ID
	_ArchiverWillReplaceObjectWithObject func(archiver IKeyedArchiver, object objectivec.IObject, newObject objectivec.IObject)
	_ArchiverDidFinish func(archiver IKeyedArchiver)
	_ArchiverWillFinish func(archiver IKeyedArchiver)
}

// SetArchiverDidEncodeObject sets the handler for the ArchiverDidEncodeObject delegate method.
//
// Informs the delegate that a given object has been encoded.
func (d *KeyedArchiverDelegate) SetArchiverDidEncodeObject(f func(archiver IKeyedArchiver, object objectivec.IObject)) {
	d._ArchiverDidEncodeObject = f
}

// SetArchiverWillEncodeObject sets the handler for the ArchiverWillEncodeObject delegate method.
//
// Informs the delegate that   is about to be encoded.
func (d *KeyedArchiverDelegate) SetArchiverWillEncodeObject(f func(archiver IKeyedArchiver, object objectivec.IObject) objc.ID) {
	d._ArchiverWillEncodeObject = f
}

// SetArchiverWillReplaceObjectWithObject sets the handler for the ArchiverWillReplaceObjectWithObject delegate method.
//
// Informs the delegate that one given object is being substituted for another given object.
func (d *KeyedArchiverDelegate) SetArchiverWillReplaceObjectWithObject(f func(archiver IKeyedArchiver, object objectivec.IObject, newObject objectivec.IObject)) {
	d._ArchiverWillReplaceObjectWithObject = f
}

// SetArchiverDidFinish sets the handler for the ArchiverDidFinish delegate method.
//
// Notifies the delegate that encoding has finished.
func (d *KeyedArchiverDelegate) SetArchiverDidFinish(f func(archiver IKeyedArchiver)) {
	d._ArchiverDidFinish = f
}

// SetArchiverWillFinish sets the handler for the ArchiverWillFinish delegate method.
//
// Notifies the delegate that encoding is about to finish.
func (d *KeyedArchiverDelegate) SetArchiverWillFinish(f func(archiver IKeyedArchiver)) {
	d._ArchiverWillFinish = f
}

// ArchiverDidEncodeObject implements the PKeyedArchiverDelegate interface.
func (d *KeyedArchiverDelegate) ArchiverDidEncodeObject(archiver IKeyedArchiver, object objectivec.IObject) {
	if d._ArchiverDidEncodeObject != nil {
		d._ArchiverDidEncodeObject(archiver, object)
	}
}

// HasArchiverDidEncodeObject returns true if a handler for ArchiverDidEncodeObject has been set.
func (d *KeyedArchiverDelegate) HasArchiverDidEncodeObject() bool {
	return d._ArchiverDidEncodeObject != nil
}

// ArchiverWillEncodeObject implements the PKeyedArchiverDelegate interface.
func (d *KeyedArchiverDelegate) ArchiverWillEncodeObject(archiver IKeyedArchiver, object objectivec.IObject) objc.ID {
	if d._ArchiverWillEncodeObject != nil {
		return d._ArchiverWillEncodeObject(archiver, object)
	}
	var zero objc.ID
	return zero
}

// HasArchiverWillEncodeObject returns true if a handler for ArchiverWillEncodeObject has been set.
func (d *KeyedArchiverDelegate) HasArchiverWillEncodeObject() bool {
	return d._ArchiverWillEncodeObject != nil
}

// ArchiverWillReplaceObjectWithObject implements the PKeyedArchiverDelegate interface.
func (d *KeyedArchiverDelegate) ArchiverWillReplaceObjectWithObject(archiver IKeyedArchiver, object objectivec.IObject, newObject objectivec.IObject) {
	if d._ArchiverWillReplaceObjectWithObject != nil {
		d._ArchiverWillReplaceObjectWithObject(archiver, object, newObject)
	}
}

// HasArchiverWillReplaceObjectWithObject returns true if a handler for ArchiverWillReplaceObjectWithObject has been set.
func (d *KeyedArchiverDelegate) HasArchiverWillReplaceObjectWithObject() bool {
	return d._ArchiverWillReplaceObjectWithObject != nil
}

// ArchiverDidFinish implements the PKeyedArchiverDelegate interface.
func (d *KeyedArchiverDelegate) ArchiverDidFinish(archiver IKeyedArchiver) {
	if d._ArchiverDidFinish != nil {
		d._ArchiverDidFinish(archiver)
	}
}

// HasArchiverDidFinish returns true if a handler for ArchiverDidFinish has been set.
func (d *KeyedArchiverDelegate) HasArchiverDidFinish() bool {
	return d._ArchiverDidFinish != nil
}

// ArchiverWillFinish implements the PKeyedArchiverDelegate interface.
func (d *KeyedArchiverDelegate) ArchiverWillFinish(archiver IKeyedArchiver) {
	if d._ArchiverWillFinish != nil {
		d._ArchiverWillFinish(archiver)
	}
}

// HasArchiverWillFinish returns true if a handler for ArchiverWillFinish has been set.
func (d *KeyedArchiverDelegate) HasArchiverWillFinish() bool {
	return d._ArchiverWillFinish != nil
}

// KeyedArchiverDelegateObject wraps an existing Objective-C object that conforms to the PKeyedArchiverDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type KeyedArchiverDelegateObject struct {
	objectivec.Object
}

// NewKeyedArchiverDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSKeyedArchiverDelegate protocol.
func NewKeyedArchiverDelegateObject(obj objectivec.Object) *KeyedArchiverDelegateObject {
	return &KeyedArchiverDelegateObject{obj}
}

// Make sure KeyedArchiverDelegateObject implements PKeyedArchiverDelegate.
var _ PKeyedArchiverDelegate = (*KeyedArchiverDelegateObject)(nil)

// ArchiverDidEncodeObject implements the PKeyedArchiverDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *KeyedArchiverDelegateObject) ArchiverDidEncodeObject(archiver IKeyedArchiver, object objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("archiver:didEncodeObject:"), archiver, object)
}

// HasArchiverDidEncodeObject returns true; this is a placeholder for optional method checks.
func (o *KeyedArchiverDelegateObject) HasArchiverDidEncodeObject() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ArchiverWillEncodeObject implements the PKeyedArchiverDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *KeyedArchiverDelegateObject) ArchiverWillEncodeObject(archiver IKeyedArchiver, object objectivec.IObject) objc.ID {
	return objc.Send[objc.ID](o.ID, objc.Sel("archiver:willEncodeObject:"), archiver, object)
}

// HasArchiverWillEncodeObject returns true; this is a placeholder for optional method checks.
func (o *KeyedArchiverDelegateObject) HasArchiverWillEncodeObject() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ArchiverWillReplaceObjectWithObject implements the PKeyedArchiverDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *KeyedArchiverDelegateObject) ArchiverWillReplaceObjectWithObject(archiver IKeyedArchiver, object objectivec.IObject, newObject objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("archiver:willReplaceObject:withObject:"), archiver, object, newObject)
}

// HasArchiverWillReplaceObjectWithObject returns true; this is a placeholder for optional method checks.
func (o *KeyedArchiverDelegateObject) HasArchiverWillReplaceObjectWithObject() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ArchiverDidFinish implements the PKeyedArchiverDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *KeyedArchiverDelegateObject) ArchiverDidFinish(archiver IKeyedArchiver) {
	objc.Send[objc.ID](o.ID, objc.Sel("archiverDidFinish:"), archiver)
}

// HasArchiverDidFinish returns true; this is a placeholder for optional method checks.
func (o *KeyedArchiverDelegateObject) HasArchiverDidFinish() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ArchiverWillFinish implements the PKeyedArchiverDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *KeyedArchiverDelegateObject) ArchiverWillFinish(archiver IKeyedArchiver) {
	objc.Send[objc.ID](o.ID, objc.Sel("archiverWillFinish:"), archiver)
}

// HasArchiverWillFinish returns true; this is a placeholder for optional method checks.
func (o *KeyedArchiverDelegateObject) HasArchiverWillFinish() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
