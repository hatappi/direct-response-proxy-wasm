# direct-response-proxy-wasm
The direct-response-proxy-wasm returns a static response without proxying the request to cluster if request headers don't match the following conditions:
- "x-foo: 1"
- "x-bar: 1"

## Setup
### Local

```sh
$ docker-compose up
```

### Istio

```yaml
apiVersion: extensions.istio.io/v1alpha1
kind: WasmPlugin
metadata:
  name: direct-response
spec:
  selector:
    matchLabels:
      istio: ingressgateway
  url: oci://ghcr.io/hatappi/direct-response-proxy-wasm/direct-response-oci:v0.0.1
  imagePullPolicy: IfNotPresent
```

## Usage

```sh
$ curl localhost:18000
Hello!
➜ curl localhost:18000 -H "x-foo: 1" -H "x-bar: 1"
OK
```
