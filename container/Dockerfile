FROM golang:1.12
LABEL maintainer="Piers <mofadeyunduo@qq.com>"

WORKDIR $HOME/go/enfi-tracker
COPY . .
RUN go install -v ./tracker/enfi-tracker.go
EXPOSE 8080
CMD ["enfi-tracker"]