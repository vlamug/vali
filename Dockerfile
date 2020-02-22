FROM debian:stretch-slim

COPY vali /bin/vali

ENTRYPOINT [ "/bin/vali" ]