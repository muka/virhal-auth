FROM scratch
COPY ./essence-auth /essence-auth
ENTRYPOINT ["/essence-auth"]
