# Configuring a User-Specific Installation
This guide details more specifically the components and configuration of a Vorteil environment. By the end of it, users should be able to configure a user-specific Vorteil environment in a way that other users may set up independent environments on the same machine without conflicts.

The guide details the steps required to configure the environment on a Linux machine, but the components and general principles of the process remain the same across all supported operating systems.

## Overview
At a minimum, a functional Vorteil environment comprises of the DBMS and the Vorteil daemon. Currently, only CockroachDB is supported as a DBMS. There are 2 services that keep the Vorteil environment running (a default installation will automatically generate these and name them 'vorteil-cockroach' and 'vorteil-daemon', respectively).

## Daemon Configuration
The daemon configuration file must be set up to avoid sharing resources with other users on the system. For this reason, it is suggested that a copy of the default configuration file be places somewhere in the user's home directory.

```toml
license = ""
encryption-key = ""

[environment]
home = "/home/USER_NAME/.vorteild"

[cookie]
hash-key = ""
block-key = ""

[storage]
type = "local"
service-account-key = ""
bucket-name = ""
storage-directory = "/home/USER_NAME/.vorteild/storage"

[log]
file = PATH_TO_LOG_FILE
format = "json"
level = "debug"

[database]
type = "postgres"
name = "vorteil"
url = "postgresql://root@127.0.0.1:7473/vorteil?sslmode=disable"

[access]
driver = "database"
type = "postgres"
name = "vauth"
url = "postgresql://root@127.0.0.1:7473/vauth?sslmode=disable"

[access.default-key]
enabled = false
readonly = true
primarygroup = "public"
groups = ["public"]
comment = "DefaultKey"

[admin]
disabled = false
bind = "127.0.0.1:7471"
disable-tls = true
tls-certificate = ""
tls-key = ""
website-files = "/opt/vorteil/web/admin"

[loopback]
disabled = false
bind = "127.0.0.1:7472"
disable-tls = true
tls-certificate = ""
tls-key = ""
website-files = "/opt/vorteil/web/vrepo"

[[network]]
disabled = true
bind = "0.0.0.0:7472"
disable-tls = true
tls-certificate = ""
tls-key = ""
website-files = "/opt/vorteil/web/vrepo"

```
Replace all instances of 'USER_NAME' with the name of the user the service is being created for. Replace 'PATH_TO_LOG_FILE' with a path at which the daemon log file will be created.

## Services
*Note: the service files will target the binaries included in a standard Vorteil installation ("/opt/vorteil/bin/vorteil", and "/opt/vorteil/bin/cockroach", respectively).*

### Cockroach

```systemd
[Unit]
Description=Vorteil Daemon Cockroach DB

[Install]
WantedBy=default.target

[Service]
ExecStart=/opt/vorteil/bin/cockroach start -p 7473 --http-port 7474 --insecure --http-host 127.0.0.1 --host 127.0.0.1 --store /home/USER_NAME/.vorteild/cockroach
ExecStop=/opt/vorteil/bin/stop-cockroach.sh
Restart=always
RestartSec=5
RemainAfterExit=no
```

Replace 'USER_NAME' with the name of the user the service is being configured for.

### Daemon
```systemd
[Unit]
Description=Vorteil Daemon

[Install]
WantedBy=default.target

[Service]
Environment=DISPLAY=:0
UMask=0000
ExecStart=/opt/vorteil/bin/vorteil daemon PATH_TO_CONFIG_FILE
ExecStop=/opt/vorteil/bin/stop-daemon.sh
Restart=always
RestartSec=5
RemainAfterExit=no
```

Replace 'PATH_TO_CONFIG_FILE' with the path to your Vorteil daemon configuration file (default location is "/opt/vorteil/conf/conf.toml").

Save these files in the "~/.config/systemd/user/" directory.

**Note that both service files have had the User/Group fields removed, and the WantedBy field altered!**

## Finishing Up
Disable the existing services from the default Vorteil installation.
```bash
$ systemctl stop vorteil-daemon.service
$ systemctl stop vorteil-cockroach.service
$ systemctl disable vorteil-daemon.service
$ systemctl disable vorteil-cockroach.service
```

Enable the new user-specific services:
```bash
$ systemctl --user enable vorteil-cockroach.service
$ systemctl --user enable vorteil-daemon.service
$ systemctl --user start vorteil-cockroach.service
$ systemctl --user start vorteil-daemon.service
```