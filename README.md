## http-proxy

[![Documentation](https://godoc.org/github.com/jkernech/go-http-proxy?status.svg)](https://godoc.org/github.com/jkernech/go-http-proxy)
[![Sonar](https://sonarcloud.io/api/project_badges/measure?project=go-http-proxy&metric=alert_status)](https://sonarcloud.io/dashboard?id=go-http-proxy)
[![Sonar](https://sonarcloud.io/api/project_badges/measure?project=go-http-proxy&metric=coverage)](https://sonarcloud.io/dashboard?id=go-http-proxy)
[![Sonar](https://sonarcloud.io/api/project_badges/measure?project=go-http-proxy&metric=security_rating)](https://sonarcloud.io/dashboard?id=go-http-proxy)
[![Sonar](https://sonarcloud.io/api/project_badges/measure?project=go-http-proxy&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=go-http-proxy)
[![Sonar](https://sonarcloud.io/api/project_badges/measure?project=go-http-proxy&metric=code_smells)](https://sonarcloud.io/dashboard?id=go-http-proxy)
[![Sonar](https://sonarcloud.io/api/project_badges/measure?project=go-http-proxy&metric=sqale_index)](https://sonarcloud.io/dashboard?id=go-http-proxy)


HTTP proxy that forward requests, useful to expose secure endpoint.

### Demo

#### Standard queries
Specifying the URL in the path forward a given request, e.g

https://go-http-proxy.herokuapp.com/https://gocover.io/_badge/github.com/jkernech/go-http-proxy

#### Mapped queries
Setting the `PATH_MAPPING` env var allows to use shorthand path according to the host mapped, useful for websites that requires basic authentication), e.g

https://go-http-proxy.herokuapp.com/https://godoc.org/github.com/jkernech/go-http-proxy?status.svg (simple request forwarding)

https://go-http-proxy.herokuapp.com/sonar/api/badges/measure?key=service-analytics&metric=lines (request forwarding with basic authentication)

### Configuration
The following environment variables allow to customise the application
```
PORT=8080 // default port on which the http-proxy listen
PATH_MAPPING={} // key/value of paths matching hosts, default is an empty object
```

#### PORT
Can be replace by the value of your choice

#### PATH_MAPPING
Optional, JSON object of paths matching hosts.

For example, if you want to expose publicly Sonarqube badges since they require basic authentication:
```
PATH_MAPPING={"sonar": "http://<token>@sonar.your_domain.com"}
```

Now you can request the server without credentials.

 [http://localhost:8080/sonar/api/badges/measure?key=%your_project%&metric=lines]([http://localhost:8080/sonar/api/badges/measure?key=<your_project>&metric=lines)

Here is another request forwarded [http://localhost:8080/google/?q=test](http://localhost:8080/google/?q=test)
```
PATH_MAPPING={"google": "https://www.google.com"}
# forward to http://www.google.com/?q=test
```

### Setup

#### From the sources

*Golang development environment needs to be setup on your machine.*

Once you cloned the repository, you can run the application with the following command:

```
make run
```

Note that you can configure the app as needed by setting your environment variables.

We strongly recommend to use a `dotenv` file locally, the template provided is reusable so you can copy it and customize it as needed:

```
cp .env.tpl .env // customize the .env file as needed
make run
```

#### Docker

*Docker environment needs to be setup on your machine.*

Run it locally:
```
docker run -p 8080:8080 -e PATH_MAPPING='{"google": "https://www.google.com", "sonar": "http://<token>@sonar.your_domain.com"}' -it jkernech/http-proxy
```

#### Docker Compose

*Docker environment needs to be setup on your machine.*

Once you cloned the repository, make sure to update the environment variables in the `docker-compose.yml` file before executing the following command
```
docker-compose up
```
