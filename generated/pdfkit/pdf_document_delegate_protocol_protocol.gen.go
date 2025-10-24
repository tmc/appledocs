// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PPDFDocumentDelegate is the PDFDocumentDelegate protocol interface.
//
// The delegate for the   object.
//
// Availability:
//   - Mac Catalyst 11.0+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.4+
//   - tvOS +
//   - visionOS 1.0+
//
// See: doc://com.apple.pdfkit/documentation/PDFKit/PDFDocumentDelegate
type PPDFDocumentDelegate interface {
	// Optional methods
	ClassForAnnotationClass(annotationClass objc.Class) objc.Class
	HasClassForAnnotationClass() bool
	ClassForAnnotationType(annotationType objc.IObject /* cross-framework: NSString */) objc.Class
	HasClassForAnnotationType() bool
	ClassForPage() objc.Class
	HasClassForPage() bool
	DidMatchString(instance IPDFSelection)
	HasDidMatchString() bool
	DocumentDidBeginDocumentFind(notification foundation.Notification)
	HasDocumentDidBeginDocumentFind() bool
	DocumentDidBeginPageFind(notification foundation.Notification)
	HasDocumentDidBeginPageFind() bool
	DocumentDidEndDocumentFind(notification foundation.Notification)
	HasDocumentDidEndDocumentFind() bool
	DocumentDidEndPageFind(notification foundation.Notification)
	HasDocumentDidEndPageFind() bool
	DocumentDidFindMatch(notification foundation.Notification)
	HasDocumentDidFindMatch() bool
	DocumentDidUnlock(notification foundation.Notification)
	HasDocumentDidUnlock() bool
}

// PDFDocumentDelegate is a delegate implementation builder for the PPDFDocumentDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PDFDocumentDelegate struct {
	_ClassForAnnotationClass func(annotationClass objc.Class) objc.Class
	_ClassForAnnotationType func(annotationType objc.IObject /* cross-framework: NSString */) objc.Class
	_ClassForPage func() objc.Class
	_DidMatchString func(instance IPDFSelection)
	_DocumentDidBeginDocumentFind func(notification foundation.Notification)
	_DocumentDidBeginPageFind func(notification foundation.Notification)
	_DocumentDidEndDocumentFind func(notification foundation.Notification)
	_DocumentDidEndPageFind func(notification foundation.Notification)
	_DocumentDidFindMatch func(notification foundation.Notification)
	_DocumentDidUnlock func(notification foundation.Notification)
}

// SetClassForAnnotationClass sets the handler for the ClassForAnnotationClass delegate method.
//
// Returns a   subclass for a class.
func (d *PDFDocumentDelegate) SetClassForAnnotationClass(f func(annotationClass objc.Class) objc.Class) {
	d._ClassForAnnotationClass = f
}

// SetClassForAnnotationType sets the handler for the ClassForAnnotationType delegate method.
//
// Returns a   subclass for an annotation type.
func (d *PDFDocumentDelegate) SetClassForAnnotationType(f func(annotationType objc.IObject /* cross-framework: NSString */) objc.Class) {
	d._ClassForAnnotationType = f
}

// SetClassForPage sets the handler for the ClassForPage delegate method.
//
// Returns a   subclass for a page object.
func (d *PDFDocumentDelegate) SetClassForPage(f func() objc.Class) {
	d._ClassForPage = f
}

// SetDidMatchString sets the handler for the DidMatchString delegate method.
//
// Called for every match found during a find operation.
func (d *PDFDocumentDelegate) SetDidMatchString(f func(instance IPDFSelection)) {
	d._DidMatchString = f
}

// SetDocumentDidBeginDocumentFind sets the handler for the DocumentDidBeginDocumentFind delegate method.
//
// Called when the   notification is posted.
func (d *PDFDocumentDelegate) SetDocumentDidBeginDocumentFind(f func(notification foundation.Notification)) {
	d._DocumentDidBeginDocumentFind = f
}

// SetDocumentDidBeginPageFind sets the handler for the DocumentDidBeginPageFind delegate method.
//
// Called when the   notification is posted.
func (d *PDFDocumentDelegate) SetDocumentDidBeginPageFind(f func(notification foundation.Notification)) {
	d._DocumentDidBeginPageFind = f
}

// SetDocumentDidEndDocumentFind sets the handler for the DocumentDidEndDocumentFind delegate method.
//
// Called when the   notification is posted.
func (d *PDFDocumentDelegate) SetDocumentDidEndDocumentFind(f func(notification foundation.Notification)) {
	d._DocumentDidEndDocumentFind = f
}

// SetDocumentDidEndPageFind sets the handler for the DocumentDidEndPageFind delegate method.
//
// Called when the   notification is posted.
func (d *PDFDocumentDelegate) SetDocumentDidEndPageFind(f func(notification foundation.Notification)) {
	d._DocumentDidEndPageFind = f
}

// SetDocumentDidFindMatch sets the handler for the DocumentDidFindMatch delegate method.
//
// Called when the   notification is posted.
func (d *PDFDocumentDelegate) SetDocumentDidFindMatch(f func(notification foundation.Notification)) {
	d._DocumentDidFindMatch = f
}

// SetDocumentDidUnlock sets the handler for the DocumentDidUnlock delegate method.
//
// Called when the   notification is posted.
func (d *PDFDocumentDelegate) SetDocumentDidUnlock(f func(notification foundation.Notification)) {
	d._DocumentDidUnlock = f
}

// ClassForAnnotationClass implements the PPDFDocumentDelegate interface.
func (d *PDFDocumentDelegate) ClassForAnnotationClass(annotationClass objc.Class) objc.Class {
	if d._ClassForAnnotationClass != nil {
		return d._ClassForAnnotationClass(annotationClass)
	}
	var zero objc.Class
	return zero
}

// HasClassForAnnotationClass returns true if a handler for ClassForAnnotationClass has been set.
func (d *PDFDocumentDelegate) HasClassForAnnotationClass() bool {
	return d._ClassForAnnotationClass != nil
}

// ClassForAnnotationType implements the PPDFDocumentDelegate interface.
func (d *PDFDocumentDelegate) ClassForAnnotationType(annotationType objc.IObject /* cross-framework: NSString */) objc.Class {
	if d._ClassForAnnotationType != nil {
		return d._ClassForAnnotationType(annotationType)
	}
	var zero objc.Class
	return zero
}

// HasClassForAnnotationType returns true if a handler for ClassForAnnotationType has been set.
func (d *PDFDocumentDelegate) HasClassForAnnotationType() bool {
	return d._ClassForAnnotationType != nil
}

// ClassForPage implements the PPDFDocumentDelegate interface.
func (d *PDFDocumentDelegate) ClassForPage() objc.Class {
	if d._ClassForPage != nil {
		return d._ClassForPage()
	}
	var zero objc.Class
	return zero
}

// HasClassForPage returns true if a handler for ClassForPage has been set.
func (d *PDFDocumentDelegate) HasClassForPage() bool {
	return d._ClassForPage != nil
}

// DidMatchString implements the PPDFDocumentDelegate interface.
func (d *PDFDocumentDelegate) DidMatchString(instance IPDFSelection) {
	if d._DidMatchString != nil {
		d._DidMatchString(instance)
	}
}

// HasDidMatchString returns true if a handler for DidMatchString has been set.
func (d *PDFDocumentDelegate) HasDidMatchString() bool {
	return d._DidMatchString != nil
}

// DocumentDidBeginDocumentFind implements the PPDFDocumentDelegate interface.
func (d *PDFDocumentDelegate) DocumentDidBeginDocumentFind(notification foundation.Notification) {
	if d._DocumentDidBeginDocumentFind != nil {
		d._DocumentDidBeginDocumentFind(notification)
	}
}

// HasDocumentDidBeginDocumentFind returns true if a handler for DocumentDidBeginDocumentFind has been set.
func (d *PDFDocumentDelegate) HasDocumentDidBeginDocumentFind() bool {
	return d._DocumentDidBeginDocumentFind != nil
}

// DocumentDidBeginPageFind implements the PPDFDocumentDelegate interface.
func (d *PDFDocumentDelegate) DocumentDidBeginPageFind(notification foundation.Notification) {
	if d._DocumentDidBeginPageFind != nil {
		d._DocumentDidBeginPageFind(notification)
	}
}

// HasDocumentDidBeginPageFind returns true if a handler for DocumentDidBeginPageFind has been set.
func (d *PDFDocumentDelegate) HasDocumentDidBeginPageFind() bool {
	return d._DocumentDidBeginPageFind != nil
}

// DocumentDidEndDocumentFind implements the PPDFDocumentDelegate interface.
func (d *PDFDocumentDelegate) DocumentDidEndDocumentFind(notification foundation.Notification) {
	if d._DocumentDidEndDocumentFind != nil {
		d._DocumentDidEndDocumentFind(notification)
	}
}

// HasDocumentDidEndDocumentFind returns true if a handler for DocumentDidEndDocumentFind has been set.
func (d *PDFDocumentDelegate) HasDocumentDidEndDocumentFind() bool {
	return d._DocumentDidEndDocumentFind != nil
}

// DocumentDidEndPageFind implements the PPDFDocumentDelegate interface.
func (d *PDFDocumentDelegate) DocumentDidEndPageFind(notification foundation.Notification) {
	if d._DocumentDidEndPageFind != nil {
		d._DocumentDidEndPageFind(notification)
	}
}

// HasDocumentDidEndPageFind returns true if a handler for DocumentDidEndPageFind has been set.
func (d *PDFDocumentDelegate) HasDocumentDidEndPageFind() bool {
	return d._DocumentDidEndPageFind != nil
}

// DocumentDidFindMatch implements the PPDFDocumentDelegate interface.
func (d *PDFDocumentDelegate) DocumentDidFindMatch(notification foundation.Notification) {
	if d._DocumentDidFindMatch != nil {
		d._DocumentDidFindMatch(notification)
	}
}

// HasDocumentDidFindMatch returns true if a handler for DocumentDidFindMatch has been set.
func (d *PDFDocumentDelegate) HasDocumentDidFindMatch() bool {
	return d._DocumentDidFindMatch != nil
}

// DocumentDidUnlock implements the PPDFDocumentDelegate interface.
func (d *PDFDocumentDelegate) DocumentDidUnlock(notification foundation.Notification) {
	if d._DocumentDidUnlock != nil {
		d._DocumentDidUnlock(notification)
	}
}

// HasDocumentDidUnlock returns true if a handler for DocumentDidUnlock has been set.
func (d *PDFDocumentDelegate) HasDocumentDidUnlock() bool {
	return d._DocumentDidUnlock != nil
}
