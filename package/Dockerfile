FROM ubuntu:18.04
RUN sed -i 's:^path-exclude=/usr/share/man:#path-exclude=/usr/share/man:' /etc/dpkg/dpkg.cfg.d/excludes
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
        man \
        manpages-posix \
        mtr \
        net-tools \
        netcat \
        netcat-openbsd \
        openssl \
        openssh-client \
        psmisc \
        socat \
        tcpdump \
        telnet \
        tmux \
        traceroute \
        tcptraceroute \
        tree \
        ngrep \
        vim \
        wget && \
    rm -rf /var/lib/apt/lists/* && \
    mv /usr/sbin/tcpdump /usr/bin/tcpdump && \
    mv /usr/sbin/traceroute /usr/bin/traceroute && \
    curl -sLf https://get.docker.com/builds/Linux/x86_64/docker-1.10.3 > /usr/bin/docker && \
    chmod +x /usr/bin/docker && \
    curl -sLf https://storage.googleapis.com/kubernetes-release/release/v1.16.8/bin/linux/amd64/kubectl > /usr/local/bin/kubectl-1.16 && \
    curl -sLf https://storage.googleapis.com/kubernetes-release/release/v1.17.4/bin/linux/amd64/kubectl > /usr/local/bin/kubectl-1.17 && \
    curl -sLf https://storage.googleapis.com/kubernetes-release/release/v1.18.0/bin/linux/amd64/kubectl > /usr/local/bin/kubectl-1.18 && \
    chmod +x /usr/local/bin/kubectl* && \
    ln -s /usr/local/bin/kubectl-1.18 /usr/local/bin/kubectl && \
    mkdir /root/.kube

ENV LOGLEVEL_TAG v0.1
RUN curl -sfSL https://github.com/rancher/loglevel/releases/download/${LOGLEVEL_TAG}/loglevel-amd64-${LOGLEVEL_TAG}.tar.gz | tar xvzf - -C /usr/bin
COPY swiss-army-knife /usr/bin/
WORKDIR /root
CMD ["swiss-army-knife"]
