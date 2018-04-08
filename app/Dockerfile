FROM centos:centos7
RUN rpm --import https://mirror.go-repo.io/centos/RPM-GPG-KEY-GO-REPO
RUN curl -s https://mirror.go-repo.io/centos/go-repo.repo | tee /etc/yum.repos.d/go-repo.repo
RUN yum update -y
RUN yum install -y mysql-devel vim git golang nscd
RUN mkdir -p $HOME/go
RUN sed -i -e 's|^PATH=\(.*\)|GOPATH=$HOME/go\nPATH=\1:$GOPATH/bin|' $HOME/.bash_profile
RUN source $HOME/.bash_profile
RUN go get -u github.com/golang/dep/cmd/dep

ENTRYPOINT tail -f /dev/null
