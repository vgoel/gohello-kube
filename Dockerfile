FROM alpine:latest
RUN apk add curl
RUN mkdir /app
ADD helloWorld /app
WORKDIR /app
CMD ["/app/helloWorld"]
