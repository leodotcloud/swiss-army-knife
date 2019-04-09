FROM ubuntu:18.04
RUN apt-get update && \
    apt-get install --no-install-recommends -y \
        arping \
        arptables \
        bridge-utils \
        ca-certificates \
        conntrack \
        curl \
        dnsutils \
        ethtool \
        iperf \
        iperf3 \
        iproute2 \
        ipsec-tools \
        ipset \
        iptables \
        iputils-ping \
        jq \
        kmod \
        ldap-utils \
        less \
        libpcap-dev \
        net-tools \
        netcat \
        netcat-openbsd \
        openssl \
        psmisc \
        socat \
        tcpdump \
        telnet \
        tmux \
        traceroute \
        tcptraceroute \
        tree \
        tree \
        vim \
        wget && \
    rm -rf /var/lib/apt/lists/* && \
    mv /usr/sbin/tcpdump /usr/bin/tcpdump && \
    mv /usr/sbin/traceroute /usr/bin/traceroute && \
    curl -sLf https://get.docker.com/builds/Linux/x86_64/docker-1.10.3 > /usr/bin/docker && \
    chmod +x /usr/bin/docker && \
    curl -sLf https://storage.googleapis.com/kubernetes-release/release/v1.9.11/bin/linux/amd64/kubectl > /usr/local/bin/kubectl-1.9 && \
    curl -sLf https://storage.googleapis.com/kubernetes-release/release/v1.10.13/bin/linux/amd64/kubectl > /usr/local/bin/kubectl-1.10 && \
    curl -sLf https://storage.googleapis.com/kubernetes-release/release/v1.11.9/bin/linux/amd64/kubectl > /usr/local/bin/kubectl-1.11 && \
    curl -sLf https://storage.googleapis.com/kubernetes-release/release/v1.12.7/bin/linux/amd64/kubectl > /usr/local/bin/kubectl-1.12 && \
    curl -sLf https://storage.googleapis.com/kubernetes-release/release/v1.13.5/bin/linux/amd64/kubectl > /usr/local/bin/kubectl-1.13 && \
    curl -sLf https://storage.googleapis.com/kubernetes-release/release/v1.14.1/bin/linux/amd64/kubectl > /usr/local/bin/kubectl-1.14 && \
    chmod +x /usr/local/bin/kubectl* && \
    ln -s /usr/local/bin/kubectl-1.14 /usr/local/bin/kubectl && \
    mkdir /root/.kube

ADD loglevel.tar.gz /usr/bin
COPY swiss-army-knife /usr/bin/
WORKDIR /root
CMD ["swiss-army-knife"]
