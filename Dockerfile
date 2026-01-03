# Start from a lightweight Debian base
FROM debian:stable-slim

# Copy the newly built Linux binary
COPY goserver /bin/goserver

# Optional: set environment variable for port
ENV PORT=8010

# Run the server
CMD ["/bin/goserver"]
