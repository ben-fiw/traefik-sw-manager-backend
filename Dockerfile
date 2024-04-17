FROM golang:latest

# SSH keys for private git repo
ARG ssh_prv_key
ARG ssh_pub_key
ARG ssh_known_hosts

# Install required packages
RUN apt-get update && apt-get install -y \
    git \
    openssh-server \
    openssh-client

# Setup private git repo
RUN mkdir -p /root/.ssh/ && \
    echo "${ssh_prv_key}" > /root/.ssh/id_ed25519 && \
    chmod 600 /root/.ssh/id_ed25519 && \
    echo "${ssh_pub_key}" > /root/.ssh/id_ed25519.pub && \
    chmod 644 /root/.ssh/id_ed25519.pub && \
    echo "${ssh_known_hosts}" > /root/.ssh/known_hosts && \
    echo "[url \"git@github.com:\"]" >> /root/.gitconfig && \
    echo "    insteadOf = https://github.com/" >> /root/.gitconfig

# Start ssh-agent
RUN eval $(ssh-agent -s) && \
    ssh-add /root/.ssh/id_ed25519

# Set the Current Working Directory inside the container
WORKDIR /app

# Install gin for live reloading
RUN go install github.com/codegangsta/gin@latest

# Set GOFLAGS to disable VCS
RUN go env -w GOFLAGS="-buildvcs=false"
RUN go env -w GOPRIVATE="github.com"

EXPOSE 4115

CMD ["gin", "-p", "4116", "-a", "4115", "-i", "run", "main.go"]
