package alertparse

import (
    "fmt"
    "encoding/json"
)

var debug string

func ParseAlerts(bytes []byte) (alerts Alerts, err error) {

    var response Data
    err = json.Unmarshal(bytes, &response)
    if err != nil {
        fmt.Printf("Could not unmarshal:\n - %s\n", err)
    } else {

      if debug == "true" {
          fmt.Printf("json struct: %#v\n", response)
          fmt.Printf("Status: %#v\n", response.Status)
          fmt.Printf("Data: %#v\n", response.Data)
      }

      alerts = *response.Data

    }

    return alerts, err
}
