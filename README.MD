# WIP: Compagny helper

![Go](https://github.com/jsenon/compagnyhelper/workflows/Go/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/jsenon/compagnyhelper)](https://goreportcard.com/report/github.com/jsenon/compagnyhelper)

Tired to find into your bookmarks a corporate link?

Why do not use a helper command line in order to retrieve for example the compagny kibana urls?

## WIP

Remain:

- Automatic packages

## How its work

A client cli will contact a server containing all the corporate usefull links

## Client configuration

You need to register the URL of the server, you can configure it in `$HOME/.compagnyhelper.yaml`

```yaml
url_server: "http://localhost:8080"
```

## Example

Get All application registered:

```sh
compagnyhelper get links
```

Get All application registered in dedicated namespace:

```sh
compagnyhelper get links -n dev
```

Get link of an application:

```sh
compagnyhelper get links Kibana -n dev
```

Describe an application link

```sh
compagnyhelper describe links Kibana -n dev
```

Open application in a borwser

```sh
compagnyhelper open links Kibana -n prod
```

## Server launch

Launch the server with following command

```sh
compagnyhelper server
```

Check additional option with `-help` command

### json application sample

```json
{
    "applications": [
        {
            "shortname": "Kibana",
            "appdesc": {
                "longname": "Kibana Dashboard for yourr logs",
                "link": "http://kibana.com",
                "env": "dev"
            }
        },
        {
            "shortname": "Grafana",
            "appdesc": {
                "longname": "Grafana Dashboard for yourr metrics",
                "link": "http://grafana.com",
                "env": "dev"
            }
        },
        {
            "shortname": "Kibana",
            "appdesc": {
                "longname": "Kibana Dashboard for yourr logs",
                "link": "http://kibana.com",
                "env": "prod"
            }
        }
    ]
}
```
