
# Alertmanager desktop notifier

An unofficial desktop notification application for Prometheus.

### Usage

```
alertmanager-notifier https://localhost
```

### Help

```
usage: alertmanager-notifier [<flags>] <server>

Alertmanager desktop notificaiton application.


Flags:
      --[no-]help     Show context-sensitive help (also try --help-long and --help-man).
      --interval=30   Query interval (in seconds). Default 30s
      --path=""       Prefix path. Default /
      --port=""       Port. Default 9090
  -V, --[no-]verbose  Verbose mode.

Args:
  <server>  Server address to query.
```


### Building from source

To build alertmanager-desktop-notifier from source code, You need:

* Go [version 1.17 or greater](https://golang.org/doc/install).

```bash
make
```


### Install

```bash
sudo make install
```
