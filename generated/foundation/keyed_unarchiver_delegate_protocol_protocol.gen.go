// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PKeyedUnarchiverDelegate is the NSKeyedUnarchiverDelegate protocol interface.
//
// The optional methods implemented by the delegate of a keyed unarchiver.
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
// See: doc://com.apple.foundation/documentation/Foundation/NSKeyedUnarchiverDelegate
type PKeyedUnarchiverDelegate interface {
	// Optional methods
	UnarchiverCannotDecodeObjectOfClassNameOriginalClasses(unarchiver IKeyedUnarchiver, name IString, classNames []string) objc.Class
	HasUnarchiverCannotDecodeObjectOfClassNameOriginalClasses() bool
	UnarchiverDidDecodeObject(unarchiver IKeyedUnarchiver, object objc.IObject) objc.ID
	HasUnarchiverDidDecodeObject() bool
	UnarchiverWillReplaceObjectWithObject(unarchiver IKeyedUnarchiver, object objc.IObject, newObject objc.IObject)
	HasUnarchiverWillReplaceObjectWithObject() bool
	UnarchiverDidFinish(unarchiver IKeyedUnarchiver)
	HasUnarchiverDidFinish() bool
	UnarchiverWillFinish(unarchiver IKeyedUnarchiver)
	HasUnarchiverWillFinish() bool
}

// KeyedUnarchiverDelegate is a delegate implementation builder for the PKeyedUnarchiverDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type KeyedUnarchiverDelegate struct {
	_UnarchiverCannotDecodeObjectOfClassNameOriginalClasses func(unarchiver IKeyedUnarchiver, name IString, classNames []string) objc.Class
	_UnarchiverDidDecodeObject func(unarchiver IKeyedUnarchiver, object objc.IObject) objc.ID
	_UnarchiverWillReplaceObjectWithObject func(unarchiver IKeyedUnarchiver, object objc.IObject, newObject objc.IObject)
	_UnarchiverDidFinish func(unarchiver IKeyedUnarchiver)
	_UnarchiverWillFinish func(unarchiver IKeyedUnarchiver)
}

// SetUnarchiverCannotDecodeObjectOfClassNameOriginalClasses sets the handler for the UnarchiverCannotDecodeObjectOfClassNameOriginalClasses delegate method.
//
// Informs the delegate that the class with a given name is not available during decoding.
func (d *KeyedUnarchiverDelegate) SetUnarchiverCannotDecodeObjectOfClassNameOriginalClasses(f func(unarchiver IKeyedUnarchiver, name IString, classNames []string) objc.Class) {
	d._UnarchiverCannotDecodeObjectOfClassNameOriginalClasses = f
}

// SetUnarchiverDidDecodeObject sets the handler for the UnarchiverDidDecodeObject delegate method.
//
// Informs the delegate that a given object has been decoded.
func (d *KeyedUnarchiverDelegate) SetUnarchiverDidDecodeObject(f func(unarchiver IKeyedUnarchiver, object objc.IObject) objc.ID) {
	d._UnarchiverDidDecodeObject = f
}

// SetUnarchiverWillReplaceObjectWithObject sets the handler for the UnarchiverWillReplaceObjectWithObject delegate method.
//
// Informs the delegate that one object is being substituted for another.
func (d *KeyedUnarchiverDelegate) SetUnarchiverWillReplaceObjectWithObject(f func(unarchiver IKeyedUnarchiver, object objc.IObject, newObject objc.IObject)) {
	d._UnarchiverWillReplaceObjectWithObject = f
}

// SetUnarchiverDidFinish sets the handler for the UnarchiverDidFinish delegate method.
//
// Notifies the delegate that decoding has finished.
func (d *KeyedUnarchiverDelegate) SetUnarchiverDidFinish(f func(unarchiver IKeyedUnarchiver)) {
	d._UnarchiverDidFinish = f
}

// SetUnarchiverWillFinish sets the handler for the UnarchiverWillFinish delegate method.
//
// Notifies the delegate that decoding is about to finish.
func (d *KeyedUnarchiverDelegate) SetUnarchiverWillFinish(f func(unarchiver IKeyedUnarchiver)) {
	d._UnarchiverWillFinish = f
}

// UnarchiverCannotDecodeObjectOfClassNameOriginalClasses implements the PKeyedUnarchiverDelegate interface.
func (d *KeyedUnarchiverDelegate) UnarchiverCannotDecodeObjectOfClassNameOriginalClasses(unarchiver IKeyedUnarchiver, name IString, classNames []string) objc.Class {
	if d._UnarchiverCannotDecodeObjectOfClassNameOriginalClasses != nil {
		return d._UnarchiverCannotDecodeObjectOfClassNameOriginalClasses(unarchiver, name, classNames)
	}
	var zero objc.Class
	return zero
}

// HasUnarchiverCannotDecodeObjectOfClassNameOriginalClasses returns true if a handler for UnarchiverCannotDecodeObjectOfClassNameOriginalClasses has been set.
func (d *KeyedUnarchiverDelegate) HasUnarchiverCannotDecodeObjectOfClassNameOriginalClasses() bool {
	return d._UnarchiverCannotDecodeObjectOfClassNameOriginalClasses != nil
}

// UnarchiverDidDecodeObject implements the PKeyedUnarchiverDelegate interface.
func (d *KeyedUnarchiverDelegate) UnarchiverDidDecodeObject(unarchiver IKeyedUnarchiver, object objc.IObject) objc.ID {
	if d._UnarchiverDidDecodeObject != nil {
		return d._UnarchiverDidDecodeObject(unarchiver, object)
	}
	var zero objc.ID
	return zero
}

// HasUnarchiverDidDecodeObject returns true if a handler for UnarchiverDidDecodeObject has been set.
func (d *KeyedUnarchiverDelegate) HasUnarchiverDidDecodeObject() bool {
	return d._UnarchiverDidDecodeObject != nil
}

// UnarchiverWillReplaceObjectWithObject implements the PKeyedUnarchiverDelegate interface.
func (d *KeyedUnarchiverDelegate) UnarchiverWillReplaceObjectWithObject(unarchiver IKeyedUnarchiver, object objc.IObject, newObject objc.IObject) {
	if d._UnarchiverWillReplaceObjectWithObject != nil {
		d._UnarchiverWillReplaceObjectWithObject(unarchiver, object, newObject)
	}
}

// HasUnarchiverWillReplaceObjectWithObject returns true if a handler for UnarchiverWillReplaceObjectWithObject has been set.
func (d *KeyedUnarchiverDelegate) HasUnarchiverWillReplaceObjectWithObject() bool {
	return d._UnarchiverWillReplaceObjectWithObject != nil
}

// UnarchiverDidFinish implements the PKeyedUnarchiverDelegate interface.
func (d *KeyedUnarchiverDelegate) UnarchiverDidFinish(unarchiver IKeyedUnarchiver) {
	if d._UnarchiverDidFinish != nil {
		d._UnarchiverDidFinish(unarchiver)
	}
}

// HasUnarchiverDidFinish returns true if a handler for UnarchiverDidFinish has been set.
func (d *KeyedUnarchiverDelegate) HasUnarchiverDidFinish() bool {
	return d._UnarchiverDidFinish != nil
}

// UnarchiverWillFinish implements the PKeyedUnarchiverDelegate interface.
func (d *KeyedUnarchiverDelegate) UnarchiverWillFinish(unarchiver IKeyedUnarchiver) {
	if d._UnarchiverWillFinish != nil {
		d._UnarchiverWillFinish(unarchiver)
	}
}

// HasUnarchiverWillFinish returns true if a handler for UnarchiverWillFinish has been set.
func (d *KeyedUnarchiverDelegate) HasUnarchiverWillFinish() bool {
	return d._UnarchiverWillFinish != nil
}
