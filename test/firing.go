package main

import (
    "fmt"
    "time"
    "github.com/bona-ppetit/alertmanager-desktop-notifier/internal/alertparse"
    "github.com/bona-ppetit/alertmanager-desktop-notifier/internal/notifier"
)

var (
    verbose bool
    alert_name string
    summary string 
)

func main() {
    var raw_response []byte

    verbose = true

    raw_response = []byte(`{
  "status": "success",
  "data": {
    "alerts": [
      {
        "labels": {
          "alertname": "Test alert",
          "severity": "critical"
        },
        "annotations": {
          "description": "Server bonjour\n  VALUE = 1\n  LABELS = map[]",
          "summary": "Bonjour not working"
        },
        "state": "firing",
        "activeAt": "2024-01-16T10:56:54.448072255Z",
        "value": "1.0e+01"
      }
    ]
  }
}`)
    
    alerts := alertparse.ParseAlerts(raw_response)

    for i := 0; i < 3; i++ {
        for i := 0; i < len(alerts.Alerts); i++ {
            alert_name = alerts.Alerts[i].Labels["alertname"]
            summary = alerts.Alerts[i].Annotations["summary"]
            if verbose {
                fmt.Printf("Name: %#v\n", alert_name)
                fmt.Printf("Summary: %#v\n", summary)
                fmt.Printf("State: %#v\n", alerts.Alerts[i].State)
            }
  	    notifier.SendAlert(alert_name, summary, "printer-error")
	    time.Sleep(1 * time.Second)
        }
    }

}
