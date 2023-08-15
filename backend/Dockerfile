FROM golang:1.20.5-alpine 

RUN go install github.com/cosmtrek/air@latest
COPY . /backend
WORKDIR /backend
CMD [ "air", "-c", ".air.toml" ]
