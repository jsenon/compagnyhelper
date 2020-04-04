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
