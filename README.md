Dnotify is a simple wrapper for the [gen2brain/beeep](https://github.com/gen2brain/beeep) .

## Example
```shell
go run ./dnotify_info.go -title "Info" -body "This is an informational message." -icon "info.png"
go run ./dnotify_alert.go -title "Alert" -body "This is an alert message." -icon "alert.png"
```

## Third-Party Libraries

This project includes code from the following third-party libraries:
- **beeep**: BSD 2-Clause License. Copyright (c) 2017, Milan Nikolic <gen2brain> 
