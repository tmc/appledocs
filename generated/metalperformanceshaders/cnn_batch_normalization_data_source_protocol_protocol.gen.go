// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PCNNBatchNormalizationDataSource is the MPSCNNBatchNormalizationDataSource protocol interface.
//
// A protocol that defines methods that a batch normalization state uses to initialize scale factors, bias terms, and batch statistics.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 11.3+
//   - iPadOS 11.3+
//   - macOS 10.13.4+
//   - tvOS 11.3+
//   - visionOS 1.0+
//
// See: doc://com.apple.metalperformanceshaders/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationDataSource
type PCNNBatchNormalizationDataSource interface {
	// Required methods
	Load()
	Beta()
	Mean()
	Variance()
	NumberOfFeatureChannels()
	Gamma()
	Purge()
	EncodeWithCoder(aCoder foundation.Coder)
	InitWithCoder(aDecoder foundation.Coder) objectivec.IObject
	UpdateGammaAndBetaWithCommandBufferBatchNormalizationState(commandBuffer unsafe.Pointer, batchNormalizationState ICNNBatchNormalizationState) CNNNormalizationGammaAndBetaState
	Label()
	UpdateGammaAndBetaWithBatchNormalizationState(batchNormalizationState ICNNBatchNormalizationState) bool
	UpdateMeanAndVarianceWithBatchNormalizationState(batchNormalizationState ICNNBatchNormalizationState) bool
	UpdateMeanAndVarianceWithCommandBufferBatchNormalizationState(commandBuffer unsafe.Pointer, batchNormalizationState ICNNBatchNormalizationState) CNNNormalizationMeanAndVarianceState
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	// Optional methods
	Epsilon()
	HasEpsilon() bool
	Encode()
	HasEncode() bool
	UpdateGammaAndBeta()
	HasUpdateGammaAndBeta() bool
	UpdateMeanAndVariance()
	HasUpdateMeanAndVariance() bool
	Copy()
	HasCopy() bool
}

// CNNBatchNormalizationDataSource is a delegate implementation builder for the PCNNBatchNormalizationDataSource protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CNNBatchNormalizationDataSource struct {
	_Epsilon func()
	_Encode func()
	_UpdateGammaAndBeta func()
	_UpdateMeanAndVariance func()
	_Copy func()
	_Load func()
	_Beta func()
	_Mean func()
	_Variance func()
	_NumberOfFeatureChannels func()
	_Gamma func()
	_Purge func()
	_EncodeWithCoder func(aCoder foundation.Coder)
	_InitWithCoder func(aDecoder foundation.Coder) objectivec.IObject
	_UpdateGammaAndBetaWithCommandBufferBatchNormalizationState func(commandBuffer unsafe.Pointer, batchNormalizationState ICNNBatchNormalizationState) CNNNormalizationGammaAndBetaState
	_Label func()
	_UpdateGammaAndBetaWithBatchNormalizationState func(batchNormalizationState ICNNBatchNormalizationState) bool
	_UpdateMeanAndVarianceWithBatchNormalizationState func(batchNormalizationState ICNNBatchNormalizationState) bool
	_UpdateMeanAndVarianceWithCommandBufferBatchNormalizationState func(commandBuffer unsafe.Pointer, batchNormalizationState ICNNBatchNormalizationState) CNNNormalizationMeanAndVarianceState
	_CopyWithZoneDevice func(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
}

// SetEpsilon sets the handler for the Epsilon delegate method.
func (d *CNNBatchNormalizationDataSource) SetEpsilon(f func()) {
	d._Epsilon = f
}

// SetEncode sets the handler for the Encode delegate method.
func (d *CNNBatchNormalizationDataSource) SetEncode(f func()) {
	d._Encode = f
}

// SetUpdateGammaAndBeta sets the handler for the UpdateGammaAndBeta delegate method.
func (d *CNNBatchNormalizationDataSource) SetUpdateGammaAndBeta(f func()) {
	d._UpdateGammaAndBeta = f
}

// SetUpdateMeanAndVariance sets the handler for the UpdateMeanAndVariance delegate method.
func (d *CNNBatchNormalizationDataSource) SetUpdateMeanAndVariance(f func()) {
	d._UpdateMeanAndVariance = f
}

// SetCopy sets the handler for the Copy delegate method.
func (d *CNNBatchNormalizationDataSource) SetCopy(f func()) {
	d._Copy = f
}

// SetLoad sets the handler for the Load delegate method.
func (d *CNNBatchNormalizationDataSource) SetLoad(f func()) {
	d._Load = f
}

// SetBeta sets the handler for the Beta delegate method.
func (d *CNNBatchNormalizationDataSource) SetBeta(f func()) {
	d._Beta = f
}

// SetMean sets the handler for the Mean delegate method.
func (d *CNNBatchNormalizationDataSource) SetMean(f func()) {
	d._Mean = f
}

// SetVariance sets the handler for the Variance delegate method.
func (d *CNNBatchNormalizationDataSource) SetVariance(f func()) {
	d._Variance = f
}

// SetNumberOfFeatureChannels sets the handler for the NumberOfFeatureChannels delegate method.
func (d *CNNBatchNormalizationDataSource) SetNumberOfFeatureChannels(f func()) {
	d._NumberOfFeatureChannels = f
}

// SetGamma sets the handler for the Gamma delegate method.
func (d *CNNBatchNormalizationDataSource) SetGamma(f func()) {
	d._Gamma = f
}

// SetPurge sets the handler for the Purge delegate method.
func (d *CNNBatchNormalizationDataSource) SetPurge(f func()) {
	d._Purge = f
}

// SetEncodeWithCoder sets the handler for the EncodeWithCoder delegate method.
func (d *CNNBatchNormalizationDataSource) SetEncodeWithCoder(f func(aCoder foundation.Coder)) {
	d._EncodeWithCoder = f
}

// SetInitWithCoder sets the handler for the InitWithCoder delegate method.
func (d *CNNBatchNormalizationDataSource) SetInitWithCoder(f func(aDecoder foundation.Coder) objectivec.IObject) {
	d._InitWithCoder = f
}

// SetUpdateGammaAndBetaWithCommandBufferBatchNormalizationState sets the handler for the UpdateGammaAndBetaWithCommandBufferBatchNormalizationState delegate method.
func (d *CNNBatchNormalizationDataSource) SetUpdateGammaAndBetaWithCommandBufferBatchNormalizationState(f func(commandBuffer unsafe.Pointer, batchNormalizationState ICNNBatchNormalizationState) CNNNormalizationGammaAndBetaState) {
	d._UpdateGammaAndBetaWithCommandBufferBatchNormalizationState = f
}

// SetLabel sets the handler for the Label delegate method.
func (d *CNNBatchNormalizationDataSource) SetLabel(f func()) {
	d._Label = f
}

// SetUpdateGammaAndBetaWithBatchNormalizationState sets the handler for the UpdateGammaAndBetaWithBatchNormalizationState delegate method.
func (d *CNNBatchNormalizationDataSource) SetUpdateGammaAndBetaWithBatchNormalizationState(f func(batchNormalizationState ICNNBatchNormalizationState) bool) {
	d._UpdateGammaAndBetaWithBatchNormalizationState = f
}

// SetUpdateMeanAndVarianceWithBatchNormalizationState sets the handler for the UpdateMeanAndVarianceWithBatchNormalizationState delegate method.
func (d *CNNBatchNormalizationDataSource) SetUpdateMeanAndVarianceWithBatchNormalizationState(f func(batchNormalizationState ICNNBatchNormalizationState) bool) {
	d._UpdateMeanAndVarianceWithBatchNormalizationState = f
}

// SetUpdateMeanAndVarianceWithCommandBufferBatchNormalizationState sets the handler for the UpdateMeanAndVarianceWithCommandBufferBatchNormalizationState delegate method.
func (d *CNNBatchNormalizationDataSource) SetUpdateMeanAndVarianceWithCommandBufferBatchNormalizationState(f func(commandBuffer unsafe.Pointer, batchNormalizationState ICNNBatchNormalizationState) CNNNormalizationMeanAndVarianceState) {
	d._UpdateMeanAndVarianceWithCommandBufferBatchNormalizationState = f
}

// SetCopyWithZoneDevice sets the handler for the CopyWithZoneDevice delegate method.
func (d *CNNBatchNormalizationDataSource) SetCopyWithZoneDevice(f func(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject) {
	d._CopyWithZoneDevice = f
}

// Epsilon implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) Epsilon() {
	if d._Epsilon != nil {
		d._Epsilon()
	}
}

// HasEpsilon returns true if a handler for Epsilon has been set.
func (d *CNNBatchNormalizationDataSource) HasEpsilon() bool {
	return d._Epsilon != nil
}

// Encode implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) Encode() {
	if d._Encode != nil {
		d._Encode()
	}
}

// HasEncode returns true if a handler for Encode has been set.
func (d *CNNBatchNormalizationDataSource) HasEncode() bool {
	return d._Encode != nil
}

// UpdateGammaAndBeta implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) UpdateGammaAndBeta() {
	if d._UpdateGammaAndBeta != nil {
		d._UpdateGammaAndBeta()
	}
}

// HasUpdateGammaAndBeta returns true if a handler for UpdateGammaAndBeta has been set.
func (d *CNNBatchNormalizationDataSource) HasUpdateGammaAndBeta() bool {
	return d._UpdateGammaAndBeta != nil
}

// UpdateMeanAndVariance implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) UpdateMeanAndVariance() {
	if d._UpdateMeanAndVariance != nil {
		d._UpdateMeanAndVariance()
	}
}

// HasUpdateMeanAndVariance returns true if a handler for UpdateMeanAndVariance has been set.
func (d *CNNBatchNormalizationDataSource) HasUpdateMeanAndVariance() bool {
	return d._UpdateMeanAndVariance != nil
}

// Copy implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) Copy() {
	if d._Copy != nil {
		d._Copy()
	}
}

// HasCopy returns true if a handler for Copy has been set.
func (d *CNNBatchNormalizationDataSource) HasCopy() bool {
	return d._Copy != nil
}

// Load implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) Load() {
	if d._Load != nil {
		d._Load()
	}
}

// HasLoad returns true if a handler for Load has been set.
func (d *CNNBatchNormalizationDataSource) HasLoad() bool {
	return d._Load != nil
}

// Beta implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) Beta() {
	if d._Beta != nil {
		d._Beta()
	}
}

// HasBeta returns true if a handler for Beta has been set.
func (d *CNNBatchNormalizationDataSource) HasBeta() bool {
	return d._Beta != nil
}

// Mean implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) Mean() {
	if d._Mean != nil {
		d._Mean()
	}
}

// HasMean returns true if a handler for Mean has been set.
func (d *CNNBatchNormalizationDataSource) HasMean() bool {
	return d._Mean != nil
}

// Variance implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) Variance() {
	if d._Variance != nil {
		d._Variance()
	}
}

// HasVariance returns true if a handler for Variance has been set.
func (d *CNNBatchNormalizationDataSource) HasVariance() bool {
	return d._Variance != nil
}

// NumberOfFeatureChannels implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) NumberOfFeatureChannels() {
	if d._NumberOfFeatureChannels != nil {
		d._NumberOfFeatureChannels()
	}
}

// HasNumberOfFeatureChannels returns true if a handler for NumberOfFeatureChannels has been set.
func (d *CNNBatchNormalizationDataSource) HasNumberOfFeatureChannels() bool {
	return d._NumberOfFeatureChannels != nil
}

// Gamma implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) Gamma() {
	if d._Gamma != nil {
		d._Gamma()
	}
}

// HasGamma returns true if a handler for Gamma has been set.
func (d *CNNBatchNormalizationDataSource) HasGamma() bool {
	return d._Gamma != nil
}

// Purge implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) Purge() {
	if d._Purge != nil {
		d._Purge()
	}
}

// HasPurge returns true if a handler for Purge has been set.
func (d *CNNBatchNormalizationDataSource) HasPurge() bool {
	return d._Purge != nil
}

// EncodeWithCoder implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) EncodeWithCoder(aCoder foundation.Coder) {
	if d._EncodeWithCoder != nil {
		d._EncodeWithCoder(aCoder)
	}
}

// HasEncodeWithCoder returns true if a handler for EncodeWithCoder has been set.
func (d *CNNBatchNormalizationDataSource) HasEncodeWithCoder() bool {
	return d._EncodeWithCoder != nil
}

// InitWithCoder implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) InitWithCoder(aDecoder foundation.Coder) objectivec.IObject {
	if d._InitWithCoder != nil {
		return d._InitWithCoder(aDecoder)
	}
	var zero objectivec.IObject
	return zero
}

// HasInitWithCoder returns true if a handler for InitWithCoder has been set.
func (d *CNNBatchNormalizationDataSource) HasInitWithCoder() bool {
	return d._InitWithCoder != nil
}

// UpdateGammaAndBetaWithCommandBufferBatchNormalizationState implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) UpdateGammaAndBetaWithCommandBufferBatchNormalizationState(commandBuffer unsafe.Pointer, batchNormalizationState ICNNBatchNormalizationState) CNNNormalizationGammaAndBetaState {
	if d._UpdateGammaAndBetaWithCommandBufferBatchNormalizationState != nil {
		return d._UpdateGammaAndBetaWithCommandBufferBatchNormalizationState(commandBuffer, batchNormalizationState)
	}
	var zero CNNNormalizationGammaAndBetaState
	return zero
}

// HasUpdateGammaAndBetaWithCommandBufferBatchNormalizationState returns true if a handler for UpdateGammaAndBetaWithCommandBufferBatchNormalizationState has been set.
func (d *CNNBatchNormalizationDataSource) HasUpdateGammaAndBetaWithCommandBufferBatchNormalizationState() bool {
	return d._UpdateGammaAndBetaWithCommandBufferBatchNormalizationState != nil
}

// Label implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) Label() {
	if d._Label != nil {
		d._Label()
	}
}

// HasLabel returns true if a handler for Label has been set.
func (d *CNNBatchNormalizationDataSource) HasLabel() bool {
	return d._Label != nil
}

// UpdateGammaAndBetaWithBatchNormalizationState implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) UpdateGammaAndBetaWithBatchNormalizationState(batchNormalizationState ICNNBatchNormalizationState) bool {
	if d._UpdateGammaAndBetaWithBatchNormalizationState != nil {
		return d._UpdateGammaAndBetaWithBatchNormalizationState(batchNormalizationState)
	}
	var zero bool
	return zero
}

// HasUpdateGammaAndBetaWithBatchNormalizationState returns true if a handler for UpdateGammaAndBetaWithBatchNormalizationState has been set.
func (d *CNNBatchNormalizationDataSource) HasUpdateGammaAndBetaWithBatchNormalizationState() bool {
	return d._UpdateGammaAndBetaWithBatchNormalizationState != nil
}

// UpdateMeanAndVarianceWithBatchNormalizationState implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) UpdateMeanAndVarianceWithBatchNormalizationState(batchNormalizationState ICNNBatchNormalizationState) bool {
	if d._UpdateMeanAndVarianceWithBatchNormalizationState != nil {
		return d._UpdateMeanAndVarianceWithBatchNormalizationState(batchNormalizationState)
	}
	var zero bool
	return zero
}

// HasUpdateMeanAndVarianceWithBatchNormalizationState returns true if a handler for UpdateMeanAndVarianceWithBatchNormalizationState has been set.
func (d *CNNBatchNormalizationDataSource) HasUpdateMeanAndVarianceWithBatchNormalizationState() bool {
	return d._UpdateMeanAndVarianceWithBatchNormalizationState != nil
}

// UpdateMeanAndVarianceWithCommandBufferBatchNormalizationState implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) UpdateMeanAndVarianceWithCommandBufferBatchNormalizationState(commandBuffer unsafe.Pointer, batchNormalizationState ICNNBatchNormalizationState) CNNNormalizationMeanAndVarianceState {
	if d._UpdateMeanAndVarianceWithCommandBufferBatchNormalizationState != nil {
		return d._UpdateMeanAndVarianceWithCommandBufferBatchNormalizationState(commandBuffer, batchNormalizationState)
	}
	var zero CNNNormalizationMeanAndVarianceState
	return zero
}

// HasUpdateMeanAndVarianceWithCommandBufferBatchNormalizationState returns true if a handler for UpdateMeanAndVarianceWithCommandBufferBatchNormalizationState has been set.
func (d *CNNBatchNormalizationDataSource) HasUpdateMeanAndVarianceWithCommandBufferBatchNormalizationState() bool {
	return d._UpdateMeanAndVarianceWithCommandBufferBatchNormalizationState != nil
}

// CopyWithZoneDevice implements the PCNNBatchNormalizationDataSource interface.
func (d *CNNBatchNormalizationDataSource) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	if d._CopyWithZoneDevice != nil {
		return d._CopyWithZoneDevice(zone, device)
	}
	var zero objectivec.IObject
	return zero
}

// HasCopyWithZoneDevice returns true if a handler for CopyWithZoneDevice has been set.
func (d *CNNBatchNormalizationDataSource) HasCopyWithZoneDevice() bool {
	return d._CopyWithZoneDevice != nil
}
