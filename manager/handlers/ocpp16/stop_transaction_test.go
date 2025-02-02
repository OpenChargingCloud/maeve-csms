// SPDX-License-Identifier: Apache-2.0

package ocpp16_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	handlers "github.com/thoughtworks/maeve-csms/manager/handlers/ocpp16"
	types "github.com/thoughtworks/maeve-csms/manager/ocpp/ocpp16"
	"github.com/thoughtworks/maeve-csms/manager/services"
	clockTest "k8s.io/utils/clock/testing"
	"testing"
	"time"
)

func TestStopTransactionHandler(t *testing.T) {
	tokenStore := services.InMemoryTokenStore{
		Tokens: map[string]*services.Token{
			"ISO14443:MYRFIDTAG": {
				Type: "ISO14443",
				Uid:  "MYRFIDTAG",
			},
		},
	}

	now, err := time.Parse(time.RFC3339, "2023-06-15T15:06:00+01:00")
	require.NoError(t, err)

	transactionStore := services.NewInMemoryTransactionStore()

	startContext := "Transaction.Begin"
	startMeasurand := "MeterValue"
	startLocation := "Outlet"
	transactionStore.CreateTransaction("cs001", handlers.ConvertToUUID(42), "MYRFIDTAG", "ISO14443",
		[]services.MeterValue{
			{
				SampledValues: []services.SampledValue{
					{
						Context:   &startContext,
						Measurand: &startMeasurand,
						Location:  &startLocation,
						Value:     50,
					},
				},
				Timestamp: now.Format(time.RFC3339),
			},
		}, 0, false)

	handler := handlers.StopTransactionHandler{
		Clock:            clockTest.NewFakePassiveClock(now),
		TokenStore:       tokenStore,
		TransactionStore: transactionStore,
	}

	idTag := "MYRFIDTAG"
	reason := types.StopTransactionJsonReasonEVDisconnected
	periodicSampleContext := types.StopTransactionJsonTransactionDataElemSampledValueElemContextSamplePeriodic
	energyRegisterMeasurand := types.StopTransactionJsonTransactionDataElemSampledValueElemMeasurandEnergyActiveImportRegister
	outletLocation := types.StopTransactionJsonTransactionDataElemSampledValueElemLocationOutlet
	req := &types.StopTransactionJson{
		IdTag:     &idTag,
		MeterStop: 200,
		Reason:    &reason,
		Timestamp: now.Format(time.RFC3339),
		TransactionData: []types.StopTransactionJsonTransactionDataElem{
			{
				SampledValue: []types.StopTransactionJsonTransactionDataElemSampledValueElem{
					{
						Context:   &periodicSampleContext,
						Measurand: &energyRegisterMeasurand,
						Location:  &outletLocation,
						Value:     "100",
					},
				},
				Timestamp: now.Format(time.RFC3339),
			},
		},
		TransactionId: 42,
	}

	got, err := handler.HandleCall(context.Background(), "cs001", req)
	require.NoError(t, err)

	want := &types.StopTransactionResponseJson{
		IdTagInfo: &types.StopTransactionResponseJsonIdTagInfo{
			Status: types.StopTransactionResponseJsonIdTagInfoStatusAccepted,
		},
	}

	assert.Equal(t, want, got)

	found, err := transactionStore.FindTransaction("cs001", handlers.ConvertToUUID(42))
	require.NoError(t, err)

	expectedTransactionEndContext := "Transaction.End"
	expectedPeriodicContext := "Sample.Periodic"
	expectedOutletLocation := "Outlet"
	expectedMeasurand := "Energy.Active.Import.Register"
	expected := &services.Transaction{
		ChargeStationId: "cs001",
		TransactionId:   handlers.ConvertToUUID(42),
		IdToken:         "MYRFIDTAG",
		TokenType:       "ISO14443",
		MeterValues: []services.MeterValue{
			{
				Timestamp: now.Format(time.RFC3339),
				SampledValues: []services.SampledValue{
					{
						Context:   &startContext,
						Location:  &startLocation,
						Measurand: &startMeasurand,
						Value:     50,
					},
				},
			},
			{
				Timestamp: now.Format(time.RFC3339),
				SampledValues: []services.SampledValue{
					{
						Context:   &expectedPeriodicContext,
						Location:  &expectedOutletLocation,
						Measurand: &expectedMeasurand,
						Value:     100,
					},
				},
			},
			{
				Timestamp: now.Format(time.RFC3339),
				SampledValues: []services.SampledValue{
					{
						Context:   &expectedTransactionEndContext,
						Location:  &expectedOutletLocation,
						Measurand: &expectedMeasurand,
						Value:     150,
					},
				},
			},
		},
		StartSeqNo:        0,
		EndedSeqNo:        1,
		UpdatedSeqNoCount: 0,
		Offline:           false,
	}

	assert.Equal(t, expected, found)
}
