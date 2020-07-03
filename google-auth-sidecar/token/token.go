// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package token

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/lestrrat/go-jwx/jwa"
	"github.com/lestrrat/go-jwx/jwt"
	"github.com/srinandan/sample-apps/common"
	types "github.com/srinandan/samples-apps/google-auth-sidecar/types"
)

var serviceAccount = types.ServiceAccount{}

func getPrivateKey() (interface{}, error) {
	pemPrivateKey := fmt.Sprintf("%v", serviceAccount.PrivateKey)
	block, _ := pem.Decode([]byte(pemPrivateKey))
	privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		common.Error.Println("error parsing Private Key: ", err)
		return nil, err
	}

	return privKey, nil
}

func ReadServiceAccount() (err error) {
	content, err := ioutil.ReadFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	if err != nil {
		common.Error.Println("service account was not found")
		return err
	}

	if err = json.Unmarshal(content, &serviceAccount); err != nil {
		common.Error.Println("parsing error")
		return err
	}
	return nil
}

func generateJWT() (string, error) {
	const aud = "https://www.googleapis.com/oauth2/v4/token"
	const scope = "https://www.googleapis.com/auth/cloud-platform"

	privKey, err := getPrivateKey()

	if err != nil {
		return "", err
	}

	now := time.Now()
	token := jwt.New()

	_ = token.Set(jwt.AudienceKey, aud)
	_ = token.Set(jwt.IssuerKey, serviceAccount.ClientEmail)
	_ = token.Set("scope", scope)
	_ = token.Set(jwt.IssuedAtKey, now.Unix())
	_ = token.Set(jwt.ExpirationKey, now.Unix())

	payload, err := token.Sign(jwa.RS256, privKey)
	if err != nil {
		common.Error.Println("error parsing Private Key: ", err)
		return "", err
	}
	common.Info.Println("jwt token: ", string(payload))
	return string(payload), nil
}

//GenerateAccessToken generates a Google OAuth access token
func GenerateAccessToken() (types.OAuthAccessToken, error) {
	const tokenEndpoint = "https://www.googleapis.com/oauth2/v4/token"
	const grantType = "urn:ietf:params:oauth:grant-type:jwt-bearer"
	accessToken := types.OAuthAccessToken{}

	token, err := generateJWT()
	if err != nil {
		return accessToken, err
	}

	form := url.Values{}
	form.Add("grant_type", grantType)
	form.Add("assertion", token)

	client := &http.Client{}
	req, err := http.NewRequest("POST", tokenEndpoint, strings.NewReader(form.Encode()))
	if err != nil {
		common.Error.Println("error in client: ", err)
		return accessToken, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	resp, err := client.Do(req)
	if err != nil {
		common.Error.Println("failed to generate oauth token: ", err)
		return accessToken, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		common.Error.Println("error in response: ", string(bodyBytes))
		return accessToken, errors.New("error in response")
	}
	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(&accessToken); err != nil {
		common.Error.Println("error in response: ", err)
		return accessToken, errors.New("error in response")
	}

	common.Info.Println("access token: ", accessToken)
	return accessToken, nil
}
