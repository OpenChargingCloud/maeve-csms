// SPDX-License-Identifier: Apache-2.0

package registry

type SecurityProfile int

const (
	UnsecuredTransportWithBasicAuth SecurityProfile = iota
	TLSWithBasicAuth
	TLSWithClientSideCertificates
)

type ChargeStation struct {
	ClientId             string
	SecurityProfile      SecurityProfile
	Base64SHA256Password string
}

type DeviceRegistry interface {
	LookupChargeStation(clientId string) *ChargeStation
}
