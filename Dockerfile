FROM golang:1.16-alpine

ADD . /app

WORKDIR /app


RUN go get -u github.com/gofiber/fiber/v2

RUN go get -u gorm.io/gorm

RUN go get -d gorm.io/driver/mysql
RUN go install gorm.io/driver/mysql

EXPOSE 3000

CMD ["go","run","main.go"]
