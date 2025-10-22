package utils

import (
	"encoding/json"
	"log"
	"os"
)

type FirebaseCredentials struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
}

func BuildJSONFromEnvironment() []byte {
	conf := &FirebaseCredentials{
		Type:                    os.Getenv("FB_TYPE"),
		ProjectId:               os.Getenv("FB_PROJECT_ID"),
		PrivateKeyId:            os.Getenv("FB_PRIVATE_KEY_ID"),
		PrivateKey:              os.Getenv("FB_PRIVATE_KEY"),
		ClientEmail:             os.Getenv("FB_CLIENT_EMAIL"),
		ClientId:                os.Getenv("FB_CLIENT_ID"),
		AuthUri:                 os.Getenv("FB_AUTH_URI"),
		TokenUri:                os.Getenv("FB_TOKEN_URI"),
		AuthProviderX509CertUrl: os.Getenv("FB_AUTH_PROVIDER_x509_CERT_URL"),
		ClientX509CertUrl:       os.Getenv("FB_CLIENT_x509_CERT_URL"),
	}
	jsonData, err := json.Marshal(conf)
	if err != nil {
		log.Fatal("Unable to load firebase config, exiting...")
	}
	return jsonData
}
