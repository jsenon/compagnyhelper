# WIP: Compagny helper

Tired to find into your bookmarks a corporate link?

Why do not use a helper command line in order to retrieve for example the compagny kibana urls?

## WIP

Remain:

- Server: Load json from file

## How its work

A client cli will contact a server containing all the corporate usefull links

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
