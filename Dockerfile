FROM tinygo/tinygo:0.24.0 AS builder

RUN apt-get install make -y

COPY . .

RUN make build

FROM alpine

COPY --from=builder ./direct-response.wasm ./plugin.wasm

