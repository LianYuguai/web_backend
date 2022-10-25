FROM loads/alpine:3.8

# ENV CGO_ENABLED=0
# ENV GOOS=linux
# ENV GOARCH=amd64
# RUN go build -o main main.go

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR                 /app
ADD resource                $WORKDIR/
ADD ./main $WORKDIR/main
RUN chmod +x $WORKDIR/main
EXPOSE 8888
###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main
