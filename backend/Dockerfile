FROM golang:1.20.5-alpine 

RUN go install github.com/cosmtrek/air@latest
COPY . /backend
WORKDIR /backend
RUN mkdir dist
CMD [ "air", "-c", ".air.toml" ]
