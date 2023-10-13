FROM golang:1.21
WORKDIR /dba

COPY . . 
RUN go build -o dba
EXPOSE 8080
CMD ["./dba"]