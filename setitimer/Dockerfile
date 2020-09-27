FROM ubuntu
RUN apt-get -y update
RUN apt-get -y install build-essential
WORKDIR /src
ADD *.c /src/
RUN gcc -o main ./main.c -lpthread
