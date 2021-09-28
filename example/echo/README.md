An Instrumented Server Echo Example
==========================

An example of usage of Instana tracer to instrument an Echo application with [`github.com/instana/go-sensor/instrumentation/instaecho`](../../instrumentation/instaecho).

Usage
-----

To start an example locally on `localhost:8881` run:

```bash
go run main.go -l localhost:8881
```

In case when the port is already in use, please choose another one.

```
  -l string
        Server listen address
```

There will be two endpoints exposed:

```
GET    /myendpoint               
```       