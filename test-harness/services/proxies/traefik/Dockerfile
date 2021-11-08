FROM traefik:2.4.6
COPY traefik.toml /etc/traefik/traefik.toml
RUN mkdir /etc/traefik/dynconf
COPY dynconf.toml /etc/traefik/dynconf/dynconf.toml
