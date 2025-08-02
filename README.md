# What is this?

DockHome is a lighweight, simple dashboard for docker containers.

# How to use this?

1. Run the dashboard using:
  ```bash
  docker run -d --name DockHome -v /var/run/docker.sock:/var/run/docker.sock -p 80:80 -e DOMAIN=localhost ghcr.io/cyberl1/dockhome
  ```

# How to use this with Dockport?

1. Run DockHome connected to the Dockport network:
  ```bash
  docker run -d --name DockHome --network Dockport -v /var/run/docker.sock:/var/run/docker.sock -l Dockport.alias=home -e DOMAIN=localhost ghcr.io/cyberl1/dockhome
  ```
