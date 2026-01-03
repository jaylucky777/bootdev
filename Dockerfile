FROM debian:stable-slim
WORKDIR /app
COPY goserver /bin/goserver
ENV PORT=8010
CMD ["/bin/goserver"]
