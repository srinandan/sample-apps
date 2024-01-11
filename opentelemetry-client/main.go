package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	common "github.com/srinandan/sample-apps/common"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

// Version
var Version string

func invokeMessage(endpoint string, apikey string, enableTLS bool, skipVerify bool) error {
	var transport *http.Transport

	if enableTLS {
		transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: skipVerify},
		}
	} else {
		transport = &http.Transport{}
	}

	client := http.Client{Transport: otelhttp.NewTransport(transport)}

	bag, _ := baggage.Parse("sample_name=sample_value")
	ctx := baggage.ContextWithBaggage(context.Background(), bag)

	var body []byte

	tr := otel.Tracer("OpenTelemetry-Client")
	err := func(ctx context.Context) error {
		ctx, span := tr.Start(ctx, "invokeMessage", trace.WithAttributes(semconv.PeerServiceKey.String("OpenTelemetry-Client")))
		defer span.End()
		req, _ := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
		if apikey != "" {
			req.Header.Add("x-api-key", apikey)
		}

		common.Info.Printf("Sending request...\n")
		res, err := client.Do(req)
		if err != nil {
			return err
		}
		body, err = ioutil.ReadAll(res.Body)
		_ = res.Body.Close()

		return err
	}(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("Response Received: %s\n\n\n", body)
	common.Info.Printf("Waiting for few seconds to export spans ...\n\n")
	return nil
}

func usage() {
	fmt.Println("")
	fmt.Println("opentelemetry-client version ", Version)
	fmt.Println("")
	fmt.Println("Usage: opentelemetry-client --endpoint=<endpoint> --project=<path>")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("endpoint    = hostname or ip of websocket server; default=http://localhost:8080/test1")
	// fmt.Println("project     = project id to publish stackdriver trace")
	// fmt.Println("token       = OAuth Bearer token")
	fmt.Println("traceType   = 0 = stdout, 1 = zipkin; default = 0")
	fmt.Println("apikey      = API Key")
	fmt.Println("skipVerify  = Skip verification of server cert; default=false")
}

func main() {
	var err error

	common.InitLog()

	tp, err := common.InitTracer(common.STDOUT_TRACE)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			çommon.Info.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	var endpoint, apiKey string // token
	var enableTLS, skipVerify, help bool
	var traceType int
	var u *url.URL

	flag.StringVar(&endpoint, "endpoint", "http://localhost:8080/test1", "API endpoint")
	// flag.StringVar(&projectid, "project", "gcp-project", "GCP Project")
	// flag.StringVar(&token, "token", "", "OAuth Bearer Token")
	flag.StringVar(&apiKey, "apikey", "", "API Key")
	flag.IntVar(&traceType, "traceType", 0, "Trace exporter type: 0 = stdout, 1 = zipkin; default = 0")
	flag.BoolVar(&skipVerify, "skipVerify", false, "SKip TLS verification")
	flag.BoolVar(&help, "help", false, "Display usage")

	// Parse commandline parameters
	flag.Parse()

	if help {
		usage()
		os.Exit(0)
	}

	/*if projectid == "" {
		projectid = os.Getenv("projectid")
	}*/

	if endpoint == "" {
		endpoint = os.Getenv("endpoint")
	}

	if apiKey == "" {
		apiKey = os.Getenv("apikey")
	}

	if u, err = url.Parse(endpoint); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if u.IsAbs() {
		if u.Scheme == "http" {
			enableTLS = false
		} else if u.Scheme == "https" {
			enableTLS = true
		} else {
			common.Error.Println("Unsupported scheme in url")
			os.Exit(1)
		}
	} else {
		enableTLS = false
	}

	fmt.Printf("Sending HTTP GET %s -H \"-x-api-key: %s\"\n", endpoint, apiKey)

	if err = invokeMessage(endpoint, apiKey, enableTLS, skipVerify); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	<-time.After(4 * time.Second)
	fmt.Printf("Inspect traces on stdout\n")
}
