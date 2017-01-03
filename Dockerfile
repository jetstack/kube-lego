FROM scratch
ADD dist/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY _build/kube-lego /kube-lego
COPY README.md /README.md
CMD ["/kube-lego"]
# ARG VCS_REF
LABEL org.label-schema.vcs-url="https://github.com/unya/kube-lego" \
      org.label-schema.license="Apache-2.0"
