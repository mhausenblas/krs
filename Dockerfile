FROM centos:7
ARG krsv
LABEL version=$krsv \
      description="Kubernetes Resource Stats (krs) tool" \
      maintainer="michael.hausenblas@gmail.com"

WORKDIR /app
RUN chown -R 1001:1 /app
USER 1001
COPY out/krs_linux /app/
RUN mv /app/krs_linux /app/krs
ENTRYPOINT ["/app/krs"]
