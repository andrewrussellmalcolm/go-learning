FROM scratch


COPY server .
COPY server.key .
COPY server.pem .

#CMD ./server ${port}  ${auth} ${frames}
ENTRYPOINT ["./server"]

EXPOSE 80
