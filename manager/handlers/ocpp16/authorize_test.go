// SPDX-License-Identifier: Apache-2.0

package ocpp16_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	handlers "github.com/thoughtworks/maeve-csms/manager/handlers/ocpp16"
	types "github.com/thoughtworks/maeve-csms/manager/ocpp/ocpp16"
	"github.com/thoughtworks/maeve-csms/manager/services"
	"testing"
)

func TestAuthorizeKnownRfidCard(t *testing.T) {
	ah := handlers.AuthorizeHandler{
		TokenStore: services.InMemoryTokenStore{
			Tokens: map[string]*services.Token{
				"ISO14443:MYRFIDCARD": {
					Type: "ISO14443",
					Uid:  "MYRFIDCARD",
				},
			},
		},
	}

	req := &types.AuthorizeJson{
		IdTag: "MYRFIDCARD",
	}

	got, err := ah.HandleCall(context.Background(), "cs001", req)
	assert.NoError(t, err)

	want := &types.AuthorizeResponseJson{
		IdTagInfo: types.AuthorizeResponseJsonIdTagInfo{
			Status: types.AuthorizeResponseJsonIdTagInfoStatusAccepted,
		},
	}

	assert.Equal(t, want, got)
}

func TestAuthorizeWithUnknownRfidCard(t *testing.T) {
	ah := handlers.AuthorizeHandler{
		TokenStore: services.InMemoryTokenStore{
			Tokens: map[string]*services.Token{
				"ISO14443:MYRFIDCARD": {
					Type: "ISO14443",
					Uid:  "MYRFIDCARD",
				},
			},
		},
	}

	req := &types.AuthorizeJson{
		IdTag: "MYBADRFID",
	}

	got, err := ah.HandleCall(context.Background(), "cs001", req)
	assert.NoError(t, err)

	want := &types.AuthorizeResponseJson{
		IdTagInfo: types.AuthorizeResponseJsonIdTagInfo{
			Status: types.AuthorizeResponseJsonIdTagInfoStatusInvalid,
		},
	}

	assert.Equal(t, want, got)
}
