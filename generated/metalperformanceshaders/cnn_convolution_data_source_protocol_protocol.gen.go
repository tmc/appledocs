// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PCNNConvolutionDataSource is the MPSCNNConvolutionDataSource protocol interface.
//
// The protocol that provides convolution filter weights and bias terms.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//
// See: doc://com.apple.metalperformanceshaders/documentation/MetalPerformanceShaders/MPSCNNConvolutionDataSource
type PCNNConvolutionDataSource interface {
	// Required methods
	BiasTerms()/* debug [protocol_interface/required_method]: BiasTerms */
	Load()/* debug [protocol_interface/required_method]: Load */
	Descriptor()/* debug [protocol_interface/required_method]: Descriptor */
	Purge()/* debug [protocol_interface/required_method]: Purge */
	DataType()/* debug [protocol_interface/required_method]: DataType */
	Weights()/* debug [protocol_interface/required_method]: Weights */
	Label()/* debug [protocol_interface/required_method]: Label */
	UpdateWithCommandBufferGradientStateSourceState(commandBuffer unsafe.Pointer, gradientState ICNNConvolutionGradientState, sourceState ICNNConvolutionWeightsAndBiasesState) CNNConvolutionWeightsAndBiasesState/* debug [protocol_interface/required_method]: UpdateWithCommandBufferGradientStateSourceState */
	UpdateWithGradientStateSourceState(gradientState ICNNConvolutionGradientState, sourceState ICNNConvolutionWeightsAndBiasesState) bool/* debug [protocol_interface/required_method]: UpdateWithGradientStateSourceState */
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject/* debug [protocol_interface/required_method]: CopyWithZoneDevice */
	// Optional methods
	RangesForUInt8Kernel()
	HasRangesForUInt8Kernel() bool
	LookupTableForUInt8Kernel()
	HasLookupTableForUInt8Kernel() bool
	Update()
	HasUpdate() bool
	WeightsQuantizationType()
	HasWeightsQuantizationType() bool
	Copy()
	HasCopy() bool
	WeightsLayout()
	HasWeightsLayout() bool
	KernelWeightsDataType()
	HasKernelWeightsDataType() bool
}

// CNNConvolutionDataSource is a delegate implementation builder for the PCNNConvolutionDataSource protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CNNConvolutionDataSource struct {
	_RangesForUInt8Kernel func()
	_LookupTableForUInt8Kernel func()
	_Update func()
	_WeightsQuantizationType func()
	_Copy func()
	_WeightsLayout func()
	_KernelWeightsDataType func()
	_BiasTerms func()
	_Load func()
	_Descriptor func()
	_Purge func()
	_DataType func()
	_Weights func()
	_Label func()
	_UpdateWithCommandBufferGradientStateSourceState func(commandBuffer unsafe.Pointer, gradientState ICNNConvolutionGradientState, sourceState ICNNConvolutionWeightsAndBiasesState) CNNConvolutionWeightsAndBiasesState
	_UpdateWithGradientStateSourceState func(gradientState ICNNConvolutionGradientState, sourceState ICNNConvolutionWeightsAndBiasesState) bool
	_CopyWithZoneDevice func(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
}

// SetRangesForUInt8Kernel sets the handler for the RangesForUInt8Kernel delegate method.
func (d *CNNConvolutionDataSource) SetRangesForUInt8Kernel(f func()) {
	d._RangesForUInt8Kernel = f
}

// SetLookupTableForUInt8Kernel sets the handler for the LookupTableForUInt8Kernel delegate method.
func (d *CNNConvolutionDataSource) SetLookupTableForUInt8Kernel(f func()) {
	d._LookupTableForUInt8Kernel = f
}

// SetUpdate sets the handler for the Update delegate method.
func (d *CNNConvolutionDataSource) SetUpdate(f func()) {
	d._Update = f
}

// SetWeightsQuantizationType sets the handler for the WeightsQuantizationType delegate method.
func (d *CNNConvolutionDataSource) SetWeightsQuantizationType(f func()) {
	d._WeightsQuantizationType = f
}

// SetCopy sets the handler for the Copy delegate method.
func (d *CNNConvolutionDataSource) SetCopy(f func()) {
	d._Copy = f
}

// SetWeightsLayout sets the handler for the WeightsLayout delegate method.
func (d *CNNConvolutionDataSource) SetWeightsLayout(f func()) {
	d._WeightsLayout = f
}

// SetKernelWeightsDataType sets the handler for the KernelWeightsDataType delegate method.
func (d *CNNConvolutionDataSource) SetKernelWeightsDataType(f func()) {
	d._KernelWeightsDataType = f
}

// SetBiasTerms sets the handler for the BiasTerms delegate method.
func (d *CNNConvolutionDataSource) SetBiasTerms(f func()) {
	d._BiasTerms = f
}

// SetLoad sets the handler for the Load delegate method.
func (d *CNNConvolutionDataSource) SetLoad(f func()) {
	d._Load = f
}

// SetDescriptor sets the handler for the Descriptor delegate method.
func (d *CNNConvolutionDataSource) SetDescriptor(f func()) {
	d._Descriptor = f
}

// SetPurge sets the handler for the Purge delegate method.
func (d *CNNConvolutionDataSource) SetPurge(f func()) {
	d._Purge = f
}

// SetDataType sets the handler for the DataType delegate method.
func (d *CNNConvolutionDataSource) SetDataType(f func()) {
	d._DataType = f
}

// SetWeights sets the handler for the Weights delegate method.
func (d *CNNConvolutionDataSource) SetWeights(f func()) {
	d._Weights = f
}

// SetLabel sets the handler for the Label delegate method.
func (d *CNNConvolutionDataSource) SetLabel(f func()) {
	d._Label = f
}

// SetUpdateWithCommandBufferGradientStateSourceState sets the handler for the UpdateWithCommandBufferGradientStateSourceState delegate method.
func (d *CNNConvolutionDataSource) SetUpdateWithCommandBufferGradientStateSourceState(f func(commandBuffer unsafe.Pointer, gradientState ICNNConvolutionGradientState, sourceState ICNNConvolutionWeightsAndBiasesState) CNNConvolutionWeightsAndBiasesState) {
	d._UpdateWithCommandBufferGradientStateSourceState = f
}

// SetUpdateWithGradientStateSourceState sets the handler for the UpdateWithGradientStateSourceState delegate method.
func (d *CNNConvolutionDataSource) SetUpdateWithGradientStateSourceState(f func(gradientState ICNNConvolutionGradientState, sourceState ICNNConvolutionWeightsAndBiasesState) bool) {
	d._UpdateWithGradientStateSourceState = f
}

// SetCopyWithZoneDevice sets the handler for the CopyWithZoneDevice delegate method.
func (d *CNNConvolutionDataSource) SetCopyWithZoneDevice(f func(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject) {
	d._CopyWithZoneDevice = f
}

// RangesForUInt8Kernel implements the PCNNConvolutionDataSource interface.
func (d *CNNConvolutionDataSource) RangesForUInt8Kernel() {
	if d._RangesForUInt8Kernel != nil {
		d._RangesForUInt8Kernel()
	}
}

// HasRangesForUInt8Kernel returns true if a handler for RangesForUInt8Kernel has been set.
func (d *CNNConvolutionDataSource) HasRangesForUInt8Kernel() bool {
	return d._RangesForUInt8Kernel != nil
}

// LookupTableForUInt8Kernel implements the PCNNConvolutionDataSource interface.
func (d *CNNConvolutionDataSource) LookupTableForUInt8Kernel() {
	if d._LookupTableForUInt8Kernel != nil {
		d._LookupTableForUInt8Kernel()
	}
}

// HasLookupTableForUInt8Kernel returns true if a handler for LookupTableForUInt8Kernel has been set.
func (d *CNNConvolutionDataSource) HasLookupTableForUInt8Kernel() bool {
	return d._LookupTableForUInt8Kernel != nil
}

// Update implements the PCNNConvolutionDataSource interface.
func (d *CNNConvolutionDataSource) Update() {
	if d._Update != nil {
		d._Update()
	}
}

// HasUpdate returns true if a handler for Update has been set.
func (d *CNNConvolutionDataSource) HasUpdate() bool {
	return d._Update != nil
}

// WeightsQuantizationType implements the PCNNConvolutionDataSource interface.
func (d *CNNConvolutionDataSource) WeightsQuantizationType() {
	if d._WeightsQuantizationType != nil {
		d._WeightsQuantizationType()
	}
}

// HasWeightsQuantizationType returns true if a handler for WeightsQuantizationType has been set.
func (d *CNNConvolutionDataSource) HasWeightsQuantizationType() bool {
	return d._WeightsQuantizationType != nil
}

// Copy implements the PCNNConvolutionDataSource interface.
func (d *CNNConvolutionDataSource) Copy() {
	if d._Copy != nil {
		d._Copy()
	}
}

// HasCopy returns true if a handler for Copy has been set.
func (d *CNNConvolutionDataSource) HasCopy() bool {
	return d._Copy != nil
}

// WeightsLayout implements the PCNNConvolutionDataSource interface.
func (d *CNNConvolutionDataSource) WeightsLayout() {
	if d._WeightsLayout != nil {
		d._WeightsLayout()
	}
}

// HasWeightsLayout returns true if a handler for WeightsLayout has been set.
func (d *CNNConvolutionDataSource) HasWeightsLayout() bool {
	return d._WeightsLayout != nil
}

// KernelWeightsDataType implements the PCNNConvolutionDataSource interface.
func (d *CNNConvolutionDataSource) KernelWeightsDataType() {
	if d._KernelWeightsDataType != nil {
		d._KernelWeightsDataType()
	}
}

// HasKernelWeightsDataType returns true if a handler for KernelWeightsDataType has been set.
func (d *CNNConvolutionDataSource) HasKernelWeightsDataType() bool {
	return d._KernelWeightsDataType != nil
}

// BiasTerms implements the PCNNConvolutionDataSource interface.
func (d *CNNConvolutionDataSource) BiasTerms() {
	if d._BiasTerms != nil {
		d._BiasTerms()
	}
}

// HasBiasTerms returns true if a handler for BiasTerms has been set.
func (d *CNNConvolutionDataSource) HasBiasTerms() bool {
	return d._BiasTerms != nil
}

// Load implements the PCNNConvolutionDataSource interface.
func (d *CNNConvolutionDataSource) Load() {
	if d._Load != nil {
		d._Load()
	}
}

// HasLoad returns true if a handler for Load has been set.
func (d *CNNConvolutionDataSource) HasLoad() bool {
	return d._Load != nil
}

// Descriptor implements the PCNNConvolutionDataSource interface.
func (d *CNNConvolutionDataSource) Descriptor() {
	if d._Descriptor != nil {
		d._Descriptor()
	}
}

// HasDescriptor returns true if a handler for Descriptor has been set.
func (d *CNNConvolutionDataSource) HasDescriptor() bool {
	return d._Descriptor != nil
}

// Purge implements the PCNNConvolutionDataSource interface.
func (d *CNNConvolutionDataSource) Purge() {
	if d._Purge != nil {
		d._Purge()
	}
}

// HasPurge returns true if a handler for Purge has been set.
func (d *CNNConvolutionDataSource) HasPurge() bool {
	return d._Purge != nil
}

// DataType implements the PCNNConvolutionDataSource interface.
func (d *CNNConvolutionDataSource) DataType() {
	if d._DataType != nil {
		d._DataType()
	}
}

// HasDataType returns true if a handler for DataType has been set.
func (d *CNNConvolutionDataSource) HasDataType() bool {
	return d._DataType != nil
}

// Weights implements the PCNNConvolutionDataSource interface.
func (d *CNNConvolutionDataSource) Weights() {
	if d._Weights != nil {
		d._Weights()
	}
}

// HasWeights returns true if a handler for Weights has been set.
func (d *CNNConvolutionDataSource) HasWeights() bool {
	return d._Weights != nil
}

// Label implements the PCNNConvolutionDataSource interface.
func (d *CNNConvolutionDataSource) Label() {
	if d._Label != nil {
		d._Label()
	}
}

// HasLabel returns true if a handler for Label has been set.
func (d *CNNConvolutionDataSource) HasLabel() bool {
	return d._Label != nil
}

// UpdateWithCommandBufferGradientStateSourceState implements the PCNNConvolutionDataSource interface.
func (d *CNNConvolutionDataSource) UpdateWithCommandBufferGradientStateSourceState(commandBuffer unsafe.Pointer, gradientState ICNNConvolutionGradientState, sourceState ICNNConvolutionWeightsAndBiasesState) CNNConvolutionWeightsAndBiasesState {
	if d._UpdateWithCommandBufferGradientStateSourceState != nil {
		return d._UpdateWithCommandBufferGradientStateSourceState(commandBuffer, gradientState, sourceState)
	}
	var zero CNNConvolutionWeightsAndBiasesState
	return zero
}

// HasUpdateWithCommandBufferGradientStateSourceState returns true if a handler for UpdateWithCommandBufferGradientStateSourceState has been set.
func (d *CNNConvolutionDataSource) HasUpdateWithCommandBufferGradientStateSourceState() bool {
	return d._UpdateWithCommandBufferGradientStateSourceState != nil
}

// UpdateWithGradientStateSourceState implements the PCNNConvolutionDataSource interface.
func (d *CNNConvolutionDataSource) UpdateWithGradientStateSourceState(gradientState ICNNConvolutionGradientState, sourceState ICNNConvolutionWeightsAndBiasesState) bool {
	if d._UpdateWithGradientStateSourceState != nil {
		return d._UpdateWithGradientStateSourceState(gradientState, sourceState)
	}
	var zero bool
	return zero
}

// HasUpdateWithGradientStateSourceState returns true if a handler for UpdateWithGradientStateSourceState has been set.
func (d *CNNConvolutionDataSource) HasUpdateWithGradientStateSourceState() bool {
	return d._UpdateWithGradientStateSourceState != nil
}

// CopyWithZoneDevice implements the PCNNConvolutionDataSource interface.
func (d *CNNConvolutionDataSource) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	if d._CopyWithZoneDevice != nil {
		return d._CopyWithZoneDevice(zone, device)
	}
	var zero objectivec.IObject
	return zero
}

// HasCopyWithZoneDevice returns true if a handler for CopyWithZoneDevice has been set.
func (d *CNNConvolutionDataSource) HasCopyWithZoneDevice() bool {
	return d._CopyWithZoneDevice != nil
}
