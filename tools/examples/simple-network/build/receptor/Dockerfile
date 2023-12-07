FROM quay.io/distronode/receptor

RUN set -x \
    # Set fastest repo
    echo 'fastestmirror=1' >> /etc/dnf/dnf.conf \
    # Add debug tools
    && yum install -y iputils nano htop vim
