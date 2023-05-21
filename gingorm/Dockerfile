FROM golang:1.20.3

WORKDIR /home
COPY ./pkg /home

RUN cd /home && go build -o gingorm

CMD ["/home/gingorm"]