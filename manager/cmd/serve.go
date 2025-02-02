// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"github.com/thoughtworks/maeve-csms/manager/mqtt"
	v201 "github.com/thoughtworks/maeve-csms/manager/ocpp/ocpp201"
	"github.com/thoughtworks/maeve-csms/manager/server"
	"github.com/thoughtworks/maeve-csms/manager/services"
)

var mqttAddr string
var mqttPrefix string
var mqttGroup string
var apiAddr string
var v2gCertPEMFiles []string
var hubjectToken string
var hubjectUrl string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Long: `Starts the server which will subscribe to messages from
the gateway and send appropriate responses.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		transactionStore := services.NewRedisTransactionStore(redisAddr)
		if transactionStore == nil {
			return errors.New("unable to connect to transaction store at address " + redisAddr)
		}

		brokerUrl, err := url.Parse(mqttAddr)
		if err != nil {
			return fmt.Errorf("parsing mqtt broker url: %v", err)
		}

		apiServer := server.New("api", apiAddr, nil, server.NewApiHandler(transactionStore))

		var v2gCertificates []*x509.Certificate
		for _, pemFile := range v2gCertPEMFiles {
			parsedCerts, err := readCertificatesFromPEMFile(pemFile)
			if err != nil {
				return fmt.Errorf("reading certificates from PEM file: %s: %v", pemFile, err)
			}
			v2gCertificates = append(v2gCertificates, parsedCerts...)
		}

		tokenStore := services.InMemoryTokenStore{
			Tokens: map[string]*services.Token{
				"ISO14443:DEADBEEF": {
					Type: string(v201.IdTokenEnumTypeISO14443),
					Uid:  "DEADBEEF",
				},
				"eMAID:EMP77TWTW00002": {
					Type: string(v201.IdTokenEnumTypeEMAID),
					Uid:  "EMP77TWTW00002",
				},
				"eMAID:EMP77TWTW00003": {
					Type: string(v201.IdTokenEnumTypeEMAID),
					Uid:  "EMP77TWTW00003",
				},
				"eMAID:EMP77TWTW00005": {
					Type: string(v201.IdTokenEnumTypeEMAID),
					Uid:  "EMP77TWTW00005",
				},
				"eMAID:EMP77TWTW99999": {
					Type: string(v201.IdTokenEnumTypeEMAID),
					Uid:  "EMP77TWTW99999",
				},
			},
		}

		tariffService := services.BasicKwhTariffService{}
		certValidationService := services.OnlineCertificateValidationService{
			RootCertificates: v2gCertificates,
			MaxOCSPAttempts:  3,
		}

		var certSignerService services.CertificateSignerService
		var certProviderService services.EvCertificateProvider
		if hubjectToken != "" && hubjectUrl != "" {
			certSignerService = services.HubjectCertificateSignerService{
				BaseURL:     hubjectUrl,
				BearerToken: hubjectToken,
				ISOVersion:  services.ISO15118V2,
			}
			certProviderService = services.HubjectEvCertificateProvider{
				BaseURL:     hubjectUrl,
				BearerToken: hubjectToken,
			}
		}

		mqttHandler := mqtt.NewHandler(
			mqtt.WithMqttBrokerUrl(brokerUrl),
			mqtt.WithMqttPrefix(mqttPrefix),
			mqtt.WithMqttGroup(mqttGroup),
			mqtt.WithTokenStore(tokenStore),
			mqtt.WithTransactionStore(transactionStore),
			mqtt.WithTariffService(tariffService),
			mqtt.WithCertValidationService(certValidationService),
			mqtt.WithCertSignerService(certSignerService),
			mqtt.WithCertificateProviderService(certProviderService),
		)

		errCh := make(chan error, 1)
		apiServer.Start(errCh)
		mqttHandler.Connect(errCh)

		err = <-errCh
		return err
	},
}

func readCertificatesFromPEMFile(pemFile string) ([]*x509.Certificate, error) {
	//#nosec G304 - only files specified by the person running the application will be loaded
	pemData, err := os.ReadFile(pemFile)
	if err != nil {
		return nil, err
	}
	return parseCertificates(pemData)
}

func parseCertificates(pemData []byte) ([]*x509.Certificate, error) {
	var certs []*x509.Certificate
	for {
		cert, rest, err := parseCertificate(pemData)
		if err != nil {
			return nil, err
		}
		if cert == nil {
			break
		}
		certs = append(certs, cert)
		pemData = rest
	}
	return certs, nil
}

func parseCertificate(pemData []byte) (cert *x509.Certificate, rest []byte, err error) {
	block, rest := pem.Decode(pemData)
	if block == nil {
		return
	}
	if block.Type != "CERTIFICATE" {
		return
	}
	cert, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		cert = nil
		return
	}
	return
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(&mqttAddr, "mqtt-addr", "m", "mqtt://127.0.0.1:1883",
		"The address of the MQTT broker, e.g. mqtt://127.0.0.1:1883")
	serveCmd.Flags().StringVar(&mqttPrefix, "mqtt-prefix", "cs",
		"The MQTT topic prefix that the manager will subscribe to, e.g. cs")
	serveCmd.Flags().StringVar(&mqttGroup, "mqtt-group", "manager",
		"The MQTT group to use for the shared subscription, e.g. manager")
	serveCmd.Flags().StringVarP(&apiAddr, "api-addr", "a", "127.0.0.1:9410",
		"The address that the API server will listen on for connections, e.g. 127.0.0.1:9410")
	serveCmd.Flags().StringSliceVar(&v2gCertPEMFiles, "v2g-pem-file", []string{},
		"The set of PEM files containing trusted V2G certificates")
	serveCmd.Flags().StringVarP(&redisAddr, "redis-addr", "r", "127.0.0.1:6379",
		"The address of the Redis store, e.g. 127.0.0.1:6379")
	serveCmd.Flags().StringVar(&hubjectToken, "hubject-token", "",
		"The Hubject Bearer token to use")
	serveCmd.Flags().StringVar(&hubjectUrl, "hubject-url", "https://open.plugncharge-test.hubject.com",
		"The Hubject Environment URL")
}
