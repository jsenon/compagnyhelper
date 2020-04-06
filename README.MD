# Compagny helper

![Go](https://github.com/jsenon/compagnyhelper/workflows/Go/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/jsenon/compagnyhelper)](https://goreportcard.com/report/github.com/jsenon/compagnyhelper)

Tired to find into your bookmarks a corporate link?

Why do not use a helper command line in order to retrieve for example the compagny kibana urls?

## How its work

A client cli will contact a server containing all the corporate usefull links

## Client configuration

### How to install

You may also download the appropriate release for your platform from the [releases page](https://github.com/jsenon/compagnyhelper/releases).

You need to register the URL of the server, you can configure it in `$HOME/.compagnyhelper.yaml`

```yaml
url_server: "http://localhost:8080"
```

**Note:** For mac user, as the binary is not signed and notorized you have to authorize it in "Security & Privacy" control panel.

### Client Example

Get All application registered:

```sh
compagnyhelper get link
```

Get All application registered in dedicated namespace:

```sh
compagnyhelper get link --n dev
```

Get link of an application:

```sh
compagnyhelper get link Kibana --n dev
```

Describe an application link

```sh
compagnyhelper describe link Kibana --n dev
```

Open application in a borwser

```sh
compagnyhelper open link Kibana --n prod
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

### Deployment on Kubernetes

1. Create the configmap

    ```sh
    kubectl create configmap compagnyhelper-cm --from-file=apps.json --dry-run -o yaml > deployment/configmap.yml
    ```

1. Apply the deployment

    ```sh
    kubectl apply -f deployment
    ```

1. Create en external service with LB, or ingress, or gateway in order to publish your endpoint
