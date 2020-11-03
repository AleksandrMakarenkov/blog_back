FROM golang:1.15 as build
RUN useradd -u 10001 scratchuser
WORKDIR /go/src/vue_back
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 go install -v ./...

FROM scratch as bin
COPY --from=build /go/bin/vue_back .
COPY --from=build /etc/passwd /etc/passwd
USER scratchuser
ENV BLOG_SECRET=sykablyat
ENV BLOG_ENV=dev
EXPOSE 9090

CMD ["./vue_back"]

