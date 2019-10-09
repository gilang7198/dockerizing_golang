FROM golang

RUN mkdir /bunch
RUN mkdir /app

ENV GOPATH /bunch
WORKDIR /bunch
RUN go get github.com/dkulchenko/bunch
RUN go install github.com/dkulchenko/bunch
RUN pwd;ls bin
RUN cp /bunch/bin/bunch /usr/bin/bunch 
RUN which bunch

WORKDIR /app
ADD . /app
RUN bunch 


ENV GOPATH /app
RUN pwd

RUN bunch update 

RUN bunch rebuild

ENTRYPOINT bunch go run /app/main.go /app/model.go /app/mysql.go

EXPOSE 9091