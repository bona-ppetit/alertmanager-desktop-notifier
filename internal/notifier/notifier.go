package notifier

import (
    "github.com/gen2brain/beeep"
)

func SendNot(title string, body string) {

    err := beeep.Notify(title, body, "assets/favicon.ico")
    if err != nil {
        panic(err)
    }
}

func SendAlert(title string, body string, icon string) {

    err := beeep.Alert(title, body, icon)
    if err != nil {
        panic(err)
    }
}
