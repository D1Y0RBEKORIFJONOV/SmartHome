FROM golang:1.22.5-alpine AS build

WORKDIR /src

COPY . .

WORKDIR /src/cmd/device

RUN go mod download

RUN go build -o /bin/hello

FROM scratch

COPY --from=build /bin/hello /bin/hello

CMD [ "/bin/hello" ]