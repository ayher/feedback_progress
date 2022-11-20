FROM centos:7

# 安装 wget
RUN yum -y install wget

# 安装 node
RUN cd /usr/local/ && mkdir node && \
    wget https://npm.taobao.org/mirrors/node/v16.0.0/node-v16.0.0-linux-x64.tar.gz && \
    mv node-v16.0.0-linux-x64.tar.gz  /usr/local/node/  && ls && \
    cd /usr/local/node/ && tar -xzvf node-v16.0.0-linux-x64.tar.gz && ls && \
    # 创建软链接
    ln -s /usr/local/node/node-v16.0.0-linux-x64/bin/node /usr/local/bin/node && \
    ln -s /usr/local/node/node-v16.0.0-linux-x64/bin/npm /usr/local/bin/npm

# 安装 go
ENV GOROOT=/usr/local/go GOPATH=/data/gopath PATH=$GOROOT/bin:$PATH
RUN wget https://storage.googleapis.com/golang/go1.16.linux-amd64.tar.gz && \
    tar --remove-files -C /usr/local/ -zxf go1.16.linux-amd64.tar.gz && \
    mkdir -p /data/go && ln -sv /usr/local/go/bin/go /bin && mkdir -p /data/gopath/src && mkdir -p /data/gopath/pkg

CMD ["tail","-f","/dev/null"]