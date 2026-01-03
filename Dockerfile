# Start from a lightweight Debian base
FROM debian:stable-slim

# Copy the compiled Go server into /bin
COPY goserver /bin/goserver

# Set the PORT environment variable inside the container
ENV PORT=8991

# Start the Go server
CMD ["/bin/goserver"]
