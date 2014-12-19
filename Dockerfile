FROM centos:centos6

RUN yum -y install ipa-gothic-fonts ipa-mincho-fonts epel-release
RUN yum -y install http://downloads.sourceforge.net/project/wkhtmltopdf/0.12.1/wkhtmltox-0.12.1_linux-centos6-amd64.rpm
RUN yum -y install git golang

ENV GOPATH /usr/local
RUN go get github.com/hayajo/md2pdf

ENTRYPOINT ["md2pdf"]

CMD ["--help"]
