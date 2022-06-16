# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /usr/local/bin/app

COPY go.mod ./
COPY go.sum ./
COPY requirements.txt ./
COPY *.go ./
COPY *.py ./
COPY autos.db ./
COPY autoParking ./

#RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add -
#RUN sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google-chrome.list'
RUN apt-get -y update
RUN apt-get upgrade -y


#RUN apt-get install -y google-chrome-stable



RUN apt-get install -y python3
#RUN apt-get install -y python
RUN apt-get install -y python3-pip
#RUN apt-get install -y python-pip


#RUN apt-get install -yqq unzip
#RUN wget -O /tmp/chromedriver.zip http://chromedriver.storage.googleapis.com/`curl -sS chromedriver.storage.googleapis.com/LATEST_RELEASE`/chromedriver_linux64.zip
#RUN unzip /tmp/chromedriver.zip chromedriver -d /usr/local/bin/
ENV DISPLAY=:99

RUN pip install --upgrade pip
RUN python3 -m pip install -r requirements.txt
#RUN apt install python-is-python3

RUN echo "deb http://deb.debian.org/debian/ unstable main contrib non-free" >> /etc/apt/sources.list.d/debian.list
RUN apt-get update
RUN apt-get install -y --no-install-recommends firefox

#commands for installing go
#RUN curl -OL https://golang.org/dl/go1.16.7.linux-amd64.tar.gz
#RUN sha256sum go1.16.7.linux-amd64.tar.gz
#RUN tar -C /usr/local -xvf go1.16.7.linux-amd64.tar.gz


#CMD ["sleep","3600"]
RUN go build -o autoParking .
CMD ["./autoParking"]
