FROM golang:1.25-bookworm as go-build
COPY . /go/src/github.com/keel-hq/keel
WORKDIR /go/src/github.com/keel-hq/keel
RUN make install

FROM node:20-alpine as ui-build
WORKDIR /app
COPY ui /app
RUN npm install
RUN npm run lint --no-fix
RUN npm run build

FROM alpine:latest
ARG USERNAME=keel
ARG USER_ID=666
ARG GROUP_ID=$USER_ID

RUN apk --no-cache add ca-certificates
RUN addgroup --gid $GROUP_ID $USERNAME \
  && adduser --home /data --ingroup $USERNAME --disabled-password --uid $USER_ID $USERNAME \
  && mkdir -p /data && chown $USERNAME:0 /data && chmod g=u /data

COPY --from=go-build /go/bin/keel /bin/keel
COPY --from=ui-build /app/dist /www

USER $USER_ID

VOLUME /data
ENV XDG_DATA_HOME /data

ENTRYPOINT ["/bin/keel"]
EXPOSE 9300
