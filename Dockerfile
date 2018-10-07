FROM centos:7
ARG krsv
LABEL version=$krsv \
      description="Kubernetes Resource Stats (krs) tool" \
      maintainer="michael.hausenblas@gmail.com"

WORKDIR /app
RUN chown -R 1001:1 /app
USER 1001
COPY out/krs_linux /app/
RUN mv /app/krs_linux /app/krs && \
    curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.11.3/bin/linux/amd64/kubectl && \
    chmod +x ./kubectl
ENTRYPOINT ["/app/krs"]
