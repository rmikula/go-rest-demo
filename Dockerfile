FROM golang

WORKDIR /app

COPY ./hello .

EXPOSE 8080

CMD [ "./hello" ]
