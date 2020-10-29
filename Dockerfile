FROM golang:1.15-alpine as build
WORKDIR /go/src/vue_back
COPY . .

RUN go get -d -v
RUN CGO_ENABLED=0 go build -v -o /go/bin/vue_back .

FROM scratch as bin
COPY --from=build /go/bin/vue_back .
ENV BLOG_SECRET=sykablyat
ENV DB_DSN=postgresql://blog:admin123@db/blog
ENV BLOG_ENV=dev
EXPOSE 9090

CMD ["./vue_back"]

