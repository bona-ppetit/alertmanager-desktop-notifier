// Copyright 2024 Guillaume Bonaparte
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
    "os"
    "io"
    "fmt"
    "net/http"
    "time"
    "log"

    "github.com/alecthomas/kingpin/v2"
    "github.com/bona-ppetit/alertmanager-desktop-notifier/internal/alertparse"
    "github.com/bona-ppetit/alertmanager-desktop-notifier/internal/notifier"
)


var myClient = &http.Client{Timeout: 10 * time.Second}

var debug string

var (
    app        = kingpin.New("alertmanager-notifier", "Alertmanager desktop notificaiton application.").UsageWriter(os.Stdout)
    server   = app.Arg("server", "Server address to query.").Required().String()
    interval = app.Flag("interval", "Query interval (in seconds). Default 30s").Default("30").Int()
    path     = app.Flag("path", "Prefix path. Default /").Default("").String()
    port     = app.Flag("port", "Port. Default 9090").Default("").String()
    verbose  = app.Flag("verbose", "Verbose mode.").Short('V').Default("false").Bool()
)

const (
    defaultPromPort      = "9090"
    defaultPromPath      = ""
    defaultPromApiv1path = "/api/v1"
)

func get_alerts(server string) {

    var prom_url string 
    if *port == "" {
        prom_url = fmt.Sprintf("%s:%s", server, defaultPromPort)
    } else {
        prom_url = fmt.Sprintf("%s:%s", server, *port)
    }

    if *path == "" {
        prom_url = fmt.Sprintf("%s%s%s", prom_url, defaultPromPath, defaultPromApiv1path)
    } else {
        prom_url = fmt.Sprintf("%s%s%s", prom_url, *path, defaultPromApiv1path)
    }

    prom_url = fmt.Sprintf("%s%s", prom_url, "/alerts")

    if *verbose {
        fmt.Println("Query server: ", prom_url)
    }

    client := &http.Client{}
    req, err := http.NewRequest(http.MethodGet, prom_url, nil)
    if err != nil {
        log.Fatal(err)
    }

    resp, err := client.Do(req)
    if err != nil {
        notifier.SendNot("alertmanager notifier", "Unable to connect\nExiting")
        log.Fatal(err)
    }

    bytes, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    if debug == "true" {
        fmt.Println("Response (string):", string(bytes))
    }

    alerts, err := alertparse.ParseAlerts(bytes)
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }

    if *verbose {
        fmt.Println("Total alerts:" )
    }


    for i := 0; i < len(alerts.Alerts); i++ {
        alert_name := alerts.Alerts[i].Labels["alertname"]
        summary := alerts.Alerts[i].Annotations["summary"]
        if debug == "true" {
	    alert_name = alert_name + " [DEBUG]"
	}
        if *verbose {
            fmt.Printf("Name: %#v\n", alert_name)
            fmt.Printf("Summary: %#v\n", summary)
            fmt.Printf("State: %#v\n", alerts.Alerts[i].State)
        }
        if alerts.Alerts[i].State == "firing" {
            notifier.SendAlert(alert_name, summary, "")
        }
    }

}

func main() {
    os.Exit(run())
}

func run() int {

    _, err := app.Parse(os.Args[1:])
    if err != nil {
    	fmt.Fprintln(os.Stderr, fmt.Errorf("Error parsing command line arguments: %w", err))
    	app.Usage(os.Args[1:])
    	os.Exit(2)
    }

    if (*interval < 10) {
	fmt.Println("Minimum query interval is 10 seconds")
        os.Exit(1)
    }

    var intervalSeconds = *interval * int(time.Second)

    if *verbose {
        fmt.Println("Alertmanager notifier starting...")
    }


    notifier.SendNot("Alertmanager notifier", "\nStarting...")

    for true {
        get_alerts(*server)
        time.Sleep(time.Duration(intervalSeconds))
    }

    return 0
}
