Dnotify is a simple wrapper for the [gen2brain/beeep](https://github.com/gen2brain/beeep) .

## Install

```bash
go install github.com/gzj/dnotify/cmd/dnotify-alert@latest
go install github.com/gzj/dnotify/cmd/dnotify-info@latest
```

## Usage

```bash
dnotify-info -title "Info" -body "This is an informational message."
dnotify-alert -title "Alert" -body "This is an alert message."
```

## Third-Party Libraries

This project includes code from the following third-party libraries:
- **beeep**: BSD 2-Clause License. Copyright (c) 2017, Milan Nikolic <gen2brain> 
