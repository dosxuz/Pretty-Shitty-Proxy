# Pretty-Shitty-Proxy
This is a proxy server written in golang. I'll use this for malware analysis and for analysing the requests.

## Setup

1) You need Golang for this to work<br>
2) You need gonfig for the proxy to be able to read from the configuration file<br>

### Installing gonfig

Gonfig can be installed as follows :

```
go get github.com/tkanos/gonfig
```

### Changing the config file

The config file looks as follows : 

```
{
	"Port" : 1331,
	"Interface" : "localhost"
}
```

The Port number and interface can be changed according to your own use. By default it is set to listen on localhost at port 1331.
## Usage 

```
go run psp.go
```

 

**Note** : It currently does not have support for HTTPS. 
