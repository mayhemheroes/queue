FROM golang:1.18 as builder

COPY . /queue
WORKDIR /queue/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzQueue.a && \
    clang -fsanitize=fuzzer fuzzQueue.a -o fuzzQueue.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /queue/mayhem/fuzzQueue.libfuzzer /