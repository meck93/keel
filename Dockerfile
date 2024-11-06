FROM golang:1.23-bookworm
COPY . /go/src/github.com/keel-hq/keel
WORKDIR /go/src/github.com/keel-hq/keel
RUN make install

FROM node:20-alpine
WORKDIR /app
COPY ui /app
RUN npm install
RUN npm run lint --no-fix
RUN npm run build

FROM alpine:latest
RUN apk --no-cache add ca-certificates

VOLUME /data
ENV XDG_DATA_HOME /data

COPY --from=0 /go/bin/keel /bin/keel
COPY --from=1 /app/dist /www
ENTRYPOINT ["/bin/keel"]
EXPOSE 9300
