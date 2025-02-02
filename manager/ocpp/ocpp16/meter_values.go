// SPDX-License-Identifier: Apache-2.0

package ocpp16

type MeterValuesJson struct {
	// ConnectorId corresponds to the JSON schema field "connectorId".
	ConnectorId int `json:"connectorId" yaml:"connectorId" mapstructure:"connectorId"`

	// MeterValue corresponds to the JSON schema field "meterValue".
	MeterValue []MeterValuesJsonMeterValueElem `json:"meterValue" yaml:"meterValue" mapstructure:"meterValue"`

	// TransactionId corresponds to the JSON schema field "transactionId".
	TransactionId *int `json:"transactionId,omitempty" yaml:"transactionId,omitempty" mapstructure:"transactionId,omitempty"`
}

func (*MeterValuesJson) IsRequest() {}

type MeterValuesJsonMeterValueElem struct {
	// SampledValue corresponds to the JSON schema field "sampledValue".
	SampledValue []MeterValuesJsonMeterValueElemSampledValueElem `json:"sampledValue" yaml:"sampledValue" mapstructure:"sampledValue"`

	// Timestamp corresponds to the JSON schema field "timestamp".
	Timestamp string `json:"timestamp" yaml:"timestamp" mapstructure:"timestamp"`
}

type MeterValuesJsonMeterValueElemSampledValueElem struct {
	// Context corresponds to the JSON schema field "context".
	Context *MeterValuesJsonMeterValueElemSampledValueElemContext `json:"context,omitempty" yaml:"context,omitempty" mapstructure:"context,omitempty"`

	// Format corresponds to the JSON schema field "format".
	Format *MeterValuesJsonMeterValueElemSampledValueElemFormat `json:"format,omitempty" yaml:"format,omitempty" mapstructure:"format,omitempty"`

	// Location corresponds to the JSON schema field "location".
	Location *MeterValuesJsonMeterValueElemSampledValueElemLocation `json:"location,omitempty" yaml:"location,omitempty" mapstructure:"location,omitempty"`

	// Measurand corresponds to the JSON schema field "measurand".
	Measurand *MeterValuesJsonMeterValueElemSampledValueElemMeasurand `json:"measurand,omitempty" yaml:"measurand,omitempty" mapstructure:"measurand,omitempty"`

	// Phase corresponds to the JSON schema field "phase".
	Phase *MeterValuesJsonMeterValueElemSampledValueElemPhase `json:"phase,omitempty" yaml:"phase,omitempty" mapstructure:"phase,omitempty"`

	// Unit corresponds to the JSON schema field "unit".
	Unit *MeterValuesJsonMeterValueElemSampledValueElemUnit `json:"unit,omitempty" yaml:"unit,omitempty" mapstructure:"unit,omitempty"`

	// Value corresponds to the JSON schema field "value".
	Value string `json:"value" yaml:"value" mapstructure:"value"`
}

type MeterValuesJsonMeterValueElemSampledValueElemContext string

const MeterValuesJsonMeterValueElemSampledValueElemContextInterruptionBegin MeterValuesJsonMeterValueElemSampledValueElemContext = "Interruption.Begin"
const MeterValuesJsonMeterValueElemSampledValueElemContextInterruptionEnd MeterValuesJsonMeterValueElemSampledValueElemContext = "Interruption.End"
const MeterValuesJsonMeterValueElemSampledValueElemContextOther MeterValuesJsonMeterValueElemSampledValueElemContext = "Other"
const MeterValuesJsonMeterValueElemSampledValueElemContextSampleClock MeterValuesJsonMeterValueElemSampledValueElemContext = "Sample.Clock"
const MeterValuesJsonMeterValueElemSampledValueElemContextSamplePeriodic MeterValuesJsonMeterValueElemSampledValueElemContext = "Sample.Periodic"
const MeterValuesJsonMeterValueElemSampledValueElemContextTransactionBegin MeterValuesJsonMeterValueElemSampledValueElemContext = "Transaction.Begin"
const MeterValuesJsonMeterValueElemSampledValueElemContextTransactionEnd MeterValuesJsonMeterValueElemSampledValueElemContext = "Transaction.End"
const MeterValuesJsonMeterValueElemSampledValueElemContextTrigger MeterValuesJsonMeterValueElemSampledValueElemContext = "Trigger"

type MeterValuesJsonMeterValueElemSampledValueElemFormat string

const MeterValuesJsonMeterValueElemSampledValueElemFormatRaw MeterValuesJsonMeterValueElemSampledValueElemFormat = "Raw"
const MeterValuesJsonMeterValueElemSampledValueElemFormatSignedData MeterValuesJsonMeterValueElemSampledValueElemFormat = "SignedData"

type MeterValuesJsonMeterValueElemSampledValueElemLocation string

const MeterValuesJsonMeterValueElemSampledValueElemLocationBody MeterValuesJsonMeterValueElemSampledValueElemLocation = "Body"
const MeterValuesJsonMeterValueElemSampledValueElemLocationCable MeterValuesJsonMeterValueElemSampledValueElemLocation = "Cable"
const MeterValuesJsonMeterValueElemSampledValueElemLocationEV MeterValuesJsonMeterValueElemSampledValueElemLocation = "EV"
const MeterValuesJsonMeterValueElemSampledValueElemLocationInlet MeterValuesJsonMeterValueElemSampledValueElemLocation = "Inlet"
const MeterValuesJsonMeterValueElemSampledValueElemLocationOutlet MeterValuesJsonMeterValueElemSampledValueElemLocation = "Outlet"

type MeterValuesJsonMeterValueElemSampledValueElemMeasurand string

const MeterValuesJsonMeterValueElemSampledValueElemMeasurandCurrentExport MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Current.Export"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandCurrentImport MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Current.Import"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandCurrentOffered MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Current.Offered"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandEnergyActiveExportInterval MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Energy.Active.Export.Interval"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandEnergyActiveExportRegister MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Energy.Active.Export.Register"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandEnergyActiveImportInterval MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Energy.Active.Import.Interval"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandEnergyActiveImportRegister MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Energy.Active.Import.Register"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandEnergyReactiveExportInterval MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Energy.Reactive.Export.Interval"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandEnergyReactiveExportRegister MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Energy.Reactive.Export.Register"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandEnergyReactiveImportInterval MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Energy.Reactive.Import.Interval"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandEnergyReactiveImportRegister MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Energy.Reactive.Import.Register"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandFrequency MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Frequency"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandPowerActiveExport MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Power.Active.Export"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandPowerActiveImport MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Power.Active.Import"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandPowerFactor MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Power.Factor"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandPowerOffered MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Power.Offered"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandPowerReactiveExport MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Power.Reactive.Export"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandPowerReactiveImport MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Power.Reactive.Import"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandRPM MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "RPM"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandSoC MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "SoC"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandTemperature MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Temperature"
const MeterValuesJsonMeterValueElemSampledValueElemMeasurandVoltage MeterValuesJsonMeterValueElemSampledValueElemMeasurand = "Voltage"

type MeterValuesJsonMeterValueElemSampledValueElemPhase string

const MeterValuesJsonMeterValueElemSampledValueElemPhaseL1 MeterValuesJsonMeterValueElemSampledValueElemPhase = "L1"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL1L2 MeterValuesJsonMeterValueElemSampledValueElemPhase = "L1-L2"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL1N MeterValuesJsonMeterValueElemSampledValueElemPhase = "L1-N"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL2 MeterValuesJsonMeterValueElemSampledValueElemPhase = "L2"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL2L3 MeterValuesJsonMeterValueElemSampledValueElemPhase = "L2-L3"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL2N MeterValuesJsonMeterValueElemSampledValueElemPhase = "L2-N"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL3 MeterValuesJsonMeterValueElemSampledValueElemPhase = "L3"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL3L1 MeterValuesJsonMeterValueElemSampledValueElemPhase = "L3-L1"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL3N MeterValuesJsonMeterValueElemSampledValueElemPhase = "L3-N"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseN MeterValuesJsonMeterValueElemSampledValueElemPhase = "N"

type MeterValuesJsonMeterValueElemSampledValueElemUnit string

const MeterValuesJsonMeterValueElemSampledValueElemUnitA MeterValuesJsonMeterValueElemSampledValueElemUnit = "A"
const MeterValuesJsonMeterValueElemSampledValueElemUnitCelcius MeterValuesJsonMeterValueElemSampledValueElemUnit = "Celcius"
const MeterValuesJsonMeterValueElemSampledValueElemUnitCelsius MeterValuesJsonMeterValueElemSampledValueElemUnit = "Celsius"
const MeterValuesJsonMeterValueElemSampledValueElemUnitFahrenheit MeterValuesJsonMeterValueElemSampledValueElemUnit = "Fahrenheit"
const MeterValuesJsonMeterValueElemSampledValueElemUnitK MeterValuesJsonMeterValueElemSampledValueElemUnit = "K"
const MeterValuesJsonMeterValueElemSampledValueElemUnitKVA MeterValuesJsonMeterValueElemSampledValueElemUnit = "kVA"
const MeterValuesJsonMeterValueElemSampledValueElemUnitKW MeterValuesJsonMeterValueElemSampledValueElemUnit = "kW"
const MeterValuesJsonMeterValueElemSampledValueElemUnitKWh MeterValuesJsonMeterValueElemSampledValueElemUnit = "kWh"
const MeterValuesJsonMeterValueElemSampledValueElemUnitKvar MeterValuesJsonMeterValueElemSampledValueElemUnit = "kvar"
const MeterValuesJsonMeterValueElemSampledValueElemUnitKvarh MeterValuesJsonMeterValueElemSampledValueElemUnit = "kvarh"
const MeterValuesJsonMeterValueElemSampledValueElemUnitPercent MeterValuesJsonMeterValueElemSampledValueElemUnit = "Percent"
const MeterValuesJsonMeterValueElemSampledValueElemUnitV MeterValuesJsonMeterValueElemSampledValueElemUnit = "V"
const MeterValuesJsonMeterValueElemSampledValueElemUnitVA MeterValuesJsonMeterValueElemSampledValueElemUnit = "VA"
const MeterValuesJsonMeterValueElemSampledValueElemUnitVar MeterValuesJsonMeterValueElemSampledValueElemUnit = "var"
const MeterValuesJsonMeterValueElemSampledValueElemUnitW MeterValuesJsonMeterValueElemSampledValueElemUnit = "W"
const MeterValuesJsonMeterValueElemSampledValueElemUnitWh MeterValuesJsonMeterValueElemSampledValueElemUnit = "Wh"
