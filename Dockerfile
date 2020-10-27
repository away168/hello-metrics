FROM scratch
ADD hello-metrics /
EXPOSE 8080
CMD ["/hello-metrics"]