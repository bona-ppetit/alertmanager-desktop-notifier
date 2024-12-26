package main

import (
    "github.com/bona-ppetit/alertmanager-desktop-notifier/internal/notifier"
)

func main() {
    notifier.SendNot("Alertmanager notifier", "test")
    notifier.SendAlert("Alertmanager notifier", "test", "icon")
}
