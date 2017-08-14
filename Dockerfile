FROM golang:1.8-alpine AS build

ARG TAG
ARG BUILD
RUN apk add -U git make curl build-base bash git autoconf automake libtool unzip file
RUN git clone https://github.com/google/protobuf /tmp/protobuf && \
    cd /tmp/protobuf && \
    ./autogen.sh && \
    ./configure && make install
RUN go get github.com/LK4D4/vndr
RUN go get github.com/golang/protobuf/protoc-gen-go
RUN go get github.com/gogo/protobuf/protoc-gen-gofast
RUN go get github.com/gogo/protobuf/proto
RUN go get github.com/gogo/protobuf/gogoproto
RUN go get github.com/gogo/protobuf/protoc-gen-gogo
RUN go get github.com/gogo/protobuf/protoc-gen-gogofast
RUN go get github.com/stevvooe/protobuild
ENV APP element
ENV REPO ehazlett/$APP
WORKDIR /go/src/github.com/$REPO
COPY . /go/src/github.com/$REPO
RUN make TAG=$TAG BUILD=$BUILD generate build

FROM alpine:latest
WORKDIR /bin
ENV APP element
ENV REPO ehazlett/$APP
COPY --from=build /go/src/github.com/${REPO}/cmd/${APP}/${APP} /bin/${APP}
EXPOSE 8080
ENTRYPOINT ["/bin/element"]
