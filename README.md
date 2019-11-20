# Recipes proxy

## Original requirements...
[...could be found here.](TASK.md)

## External dependencies
In this project I tried to use only the standard library. A couple of exceptions for configuration management:
* github.com/joho/godotenv
* github.com/kelseyhightower/envconfig

## Project structure
#### **cmd/server**
Main package for the proxy's binary.
The server, handler, async code and other utilities can be found here.

#### **internal/client**
HTTP client, initialization and configuration.

#### **internal/config**
Configs for the entire project.

#### **internal/types**
Collection of structs, which reflect the entities for the upstream API.

#### **test**
Files with JSON representation of the recipes, needed for the tests.

## Configuration
The configuration parameters are in the [.env](.env) file.
This file must be in the same directory as the binary. There are a bunch of variables, namely:

* `ADDR=":8080"`
    
    The address our server will be listening on.
    By default, this is **"localhost:8080"**.

* `PROXY_ENDPOINT="https://s3-eu-west-1.amazonaws.com/test-golang-recipes"`

    Upstream API endpoint.

* `CERT_PATH=""`

    It's possible to verify upstream's certificate chain.  
    For this purpose you should define path to the cert file.
    By default, proxy doesn't verify certificate.

* `CONCURRENCY_LIMIT=10`

	Setup a limitation for the number of concurrent requests.
	This helps to avoid resources exhaustion (e.g. file descriptors) and improve performance due to decreasing context switching.

* `MAX_IDLE_CONNECTIONS=0`

	Controls the maximum idle (keep-alive) connections to keep.

* `IDLE_CONN_TIMEOUT=500`

	The maximum amount of time an idle (keep-alive) connection will remain idle before closing itself.

* `CLIENT_TIMEOUT=1000`

    Timeout per request to the proxied endpoint (in milliseconds).

* `SERVER_TIMEOUT=800`

    Total timeout for the request to the proxy (in milliseconds).

* `DEFAULT_TOP_VALUE=100`

    Defines default **top** value.

* `DEFAULT_SKIP_VALUE=0`

    Defines default **skip** value.

## Prerequisites
* [Docker](https://www.docker.com/) >= 18.09.9
* [Docker Compose](https://docs.docker.com/compose/) >= 1.23.2
* Valid configuration file (`.env`)

## Running the server
```
$ sudo docker-compose build
$ sudo docker-compose up web
```

## Benchmarks
100 recipes per request (without losses)
```
Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /recipes?top=100
Document Length:        226723 bytes

Concurrency Level:      1
Time taken for tests:   8.215 seconds
Complete requests:      10
Failed requests:        0
   (Connect: 0, Receive: 0, Length: 0, Exceptions: 0)
Total transferred:      2254365 bytes
HTML transferred:       2253335 bytes
Requests per second:    1.22 [#/sec] (mean)
Time per request:       821.481 [ms] (mean)
Time per request:       821.481 [ms] (mean, across all concurrent requests)
Transfer rate:          268.00 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:   708  821  77.8    832     939
Waiting:      706  820  77.9    830     938
Total:        708  821  77.8    832     939
```

1000 recipes per request (with losses, limited by 1 second)
```
Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /recipes?top=1000
Document Length:        234882 bytes

Concurrency Level:      1
Time taken for tests:   9.208 seconds
Complete requests:      10
Failed requests:        9
   (Connect: 0, Receive: 0, Length: 9, Exceptions: 0)
Total transferred:      2246693 bytes
HTML transferred:       2245663 bytes
Requests per second:    1.09 [#/sec] (mean)
Time per request:       920.846 [ms] (mean)
Time per request:       920.846 [ms] (mean, across all concurrent requests)
Transfer rate:          238.26 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       0
Processing:   849  920  88.5    890    1156
Waiting:      848  919  88.2    888    1153
Total:        849  921  88.5    890    1156
```

10000 recipes per request (with losses, limited by 1 second)
```
Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /recipes?top=10000
Document Length:        208668 bytes

Concurrency Level:      1
Time taken for tests:   9.917 seconds
Complete requests:      10
Failed requests:        9
   (Connect: 0, Receive: 0, Length: 9, Exceptions: 0)
Total transferred:      2197325 bytes
HTML transferred:       2196295 bytes
Requests per second:    1.01 [#/sec] (mean)
Time per request:       991.715 [ms] (mean)
Time per request:       991.715 [ms] (mean, across all concurrent requests)
Transfer rate:          216.38 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       0
Processing:   873  991 249.6    918    1698
Waiting:      872  990 249.8    914    1697
Total:        873  992 249.6    918    1698
```

## Things I was lazy to implement
* Adequate logger
* Adequate responses on errors - at the moment the proxy just returns what it managed to pull out.
* Custom errors (wrapped)
* Handler-level integration tests.