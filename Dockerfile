FROM brew.registry.redhat.io/rh-osbs/openshift-golang-builder AS build-env
USER root

WORKDIR /cosign
COPY . .
RUN git config --global --add safe.directory /cosign
RUN make cosign

# Install Cosign
FROM registry.access.redhat.com/ubi9/go-toolset@sha256:9a0f860e143f2f771bee92ab3b0161e10e2390370152e07e6bcf8105242cee13
USER root

LABEL description="Cosign is a container signing tool that leverages simple, secure, and auditable signatures based on simple primitives and best practices."
LABEL io.k8s.description="Cosign is a container signing tool that leverages simple, secure, and auditable signatures based on simple primitives and best practices."
LABEL io.openshift.tags="cosign trusted-signer"
LABEL summary="Provides the cosign CLI binary for signing and verifying container images."
LABEL com.redhat.component="cosign"


COPY --from=build-env /cosign/cosign /usr/local/bin/cosign
ENV HOME=/home

RUN git config --global --add safe.directory /cosign && \
    cd /cosign && make cosign && mv /cosign/cosign /usr/local/bin/cosign && \
    chown root:0 /usr/local/bin/cosign && chmod g+wx /usr/local/bin/cosign  && \
    chgrp -R 0 /${HOME} && chmod -R g=u /${HOME}

# Makes sure the container stays running
CMD ["tail", "-f", "/dev/null"]