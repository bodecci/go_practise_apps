FROM golang:1.20 AS build
WORKDIR /src
ADD . .
RUN CGO_ENABLED=0 go build -o app


FROM alpine:latest
COPY --from=build /src/app app
EXPOSE 3000
CMD [ "./app" ]