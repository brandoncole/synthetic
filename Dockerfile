FROM alpine
ADD bin/synthetic /synthetic
RUN /synthetic --help