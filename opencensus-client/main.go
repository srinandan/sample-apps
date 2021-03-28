// Copyright 2021 Google LLC
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

package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

func invokeMessage(projectid string, endpoint string, apikey string) {
	var payload map[string]interface{}

	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: projectid,
	})
	if err != nil {
		fmt.Printf("failed to initialize Stackdriver exporter: %+v\n", err)
	} else {
		trace.RegisterExporter(exporter)
		trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
		fmt.Println("registered Stackdriver tracing")
		defer exporter.Flush()
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: &ochttp.Transport{
			Base:        tr,
			Propagation: &propagation.HTTPFormat{},
		},
	}

	req, err := http.NewRequest("GET", endpoint, nil)

	if apikey != "" {
		req.Header.Add("x-api-key", apikey)
	}

	if err != nil {
		fmt.Println("Error: %v", err)
		return
	}

	ctx, span := trace.StartSpan(context.Background(), "opencensus-client start")
	// It is imperative that req.WithContext is used to
	// propagate context and use it in the request.
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: %v", err)
		return
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	span.End()

	if resp.StatusCode != 200 {
		fmt.Printf("endpoint return http error %d", resp.StatusCode)
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: %v", err)
		return
	}

	err = json.Unmarshal(respBody, &payload)
	if err != nil {
		fmt.Println("Error: %v", err)
		return
	}

	fmt.Println(payload)
}

func main() {

	projectid := os.Getenv("projectid")
	endpoint := os.Getenv("endpoint")
	apikey := os.Getenv("apikey")

	fmt.Printf("Endpoint is %s\n", endpoint)
	fmt.Printf("apikey is %s\n", apikey)
	fmt.Printf("Project id is %s\n", projectid)

	invokeMessage(projectid, endpoint, apikey)

	<-time.After(5 * time.Second)
}
