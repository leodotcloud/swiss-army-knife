FROM ubuntu:16.04
RUN apt-get update && \
    apt-get install --no-install-recommends -y \
        tree \
        arptables \
        ca-certificates \
        tmux \
        wget \
        curl \
        tree \
        iperf \
        iptables \
        conntrack \
        ipset \
        tcpdump \
        telnet \
        arping \
        net-tools \
        bridge-utils \
        ethtool \
        iproute2 \
        netcat-openbsd \
        ipsec-tools \
        psmisc \
        socat \
        dnsutils \
        vim \
        iperf3 \
        less \
        libpcap-dev \
        netcat \
        iputils-ping \
        jq \
        kmod \
        openssl && \
    rm -rf /var/lib/apt/lists/* && \
    mv /usr/sbin/tcpdump /usr/bin/tcpdump && \
    curl -sLf https://get.docker.com/builds/Linux/x86_64/docker-1.10.3 > /usr/bin/docker && \
    chmod +x /usr/bin/docker

ADD loglevel.tar.gz /usr/bin
COPY swiss-army-knife /usr/bin/
CMD ["swiss-army-knife"]
