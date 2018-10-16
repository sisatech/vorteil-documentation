# Vorteil Service Configuration
```toml
# An absolute path to a license file.
license = ""

# The encryption key is used to encrypt and decrypt authentication tokens
# used by users interacting with the Vorteil service via a protected endpoint.
encryption-key = ""


[environment]
  # The 'home' directory for the Vorteil service .
  # Essential files and folders will be placed here.
  home = "/home/USER/.vorteild"

[kernel-sources]
  # Directory in which gathered kernels will be stored.
  directory = "/home/USER/.vorteild/kernels"

  # This directory is watched for changes in real-time by 
  # the Vorteil service. Kernel files can be manually dropped 
  # in to this directory and will be immediately available for use.
  drop-path = "/home/USER/.vorteild/kernels/watch"

  # An array of URLs from which kernels can be sourced. Can include:
  # - https://downloads.vorteil.io/system
  # - https://downloads.vorteil.io/system-release-candidates
  remote-repositories = ["https://downloads.vorteil.io/system"]

  # The strategy used to upgrade to more recent kernel versions when 
  # running apps with Vorteil.
  upgrade-strategy = "none"

[cookie]
  # Hash key used for generating secure cookies. If left blank, 
  # a value will be automatically generated for this field.
  hash-key = ""

  # Block key used for generating secure cookies. If left blank,
  # a value will be automatically generated for this field.
  block-key = ""

[storage]
  # Defines the storage driver used by the Vorteil service. Files such 
  # as repository packages, icons, etc. will be stored via the 
  # storage driver. Acceptable values: 
  # - local
  # - google
  type = "local"

  # This field is specific to the 'google' driver. It specifies an 
  # absolute path to a .json-formatted Google Service Account key file.
  service-account-key = ""

  # This field is specific to the 'google' driver. It specifies the 
  # name of the storage bucket used for file storage.
  bucket-name = ""

  # This field applies to both the 'local' and 'google' drivers.
  # Files will be stored within this directory.
  storage-directory = "/home/USER/.vorteild/storage"

[log]
  # Absolute path specifying where to write service logs.
  file = "/home/USER/.vorteild/log"

  # Output format. Acceptable values:
  # - "json"
  # - "" (empty string; default formatting)
  format = "json"

  # Minimum log severity level. Acceptable values: 
  # - "debug"
  # - "normal"
  # - "important"
  level = "debug"

  # Max file size (in MiB) before log files will be rotated out.
  # size = 20

  # Number of rotated log files to retain before erasing.
  # backups = 5

  # Max file age (in days) before log files will be rotated out.
  # age = 7

  # Enables or disables log file compression.
  # compressed = true

[database]
  # Defines the database driver used by the Vorteil service. 
  # Acceptable values:
  # - "local"
  # - "postgres" (based on CockroachDB v2)
  type = "local"

  # This field is specific to the 'local' driver. Absolute path to 
  # the database file created by the local database driver.
  path = "/home/USER/.vorteild/localdb.db"

  # This field is specific to the 'postgres' driver. Name of the 
  # database to create (if not exists) and use.
  # name = "vorteil"

  # This field is specific to the 'postgres' driver.
  # Used to establish a connection to the database.
  # url = "postgresql://root@127.0.0.1:7473/vorteil?sslmode=disable"

[access]
  # Defines the access driver used by the Vorteil service.
  # Acceptable values:
  # - "local"
  # - "postgres" (based on CockroachDB v2)
  driver = "local"

  # This field is specific to the 'postgres' driver. Name of the 
  # database to create (if not exists) and use.
  # name = "vauth"

  # This field is specific to the 'postgres' driver.
  # Used to establish a connection to the database.
  # url = "postgresql://root@127.0.0.1:7473/vauth?sslmode=disable"

  [access.default-key]
    # Enables or disables the use of a 'default' authorization key,
    # applied to all otherwise unauthenticated requests.
    enabled = false

    # Primary group applied to the default key.
    primarygroup = "public"

    # Array of groups applied to the default key.
    groups = ["public"]

[admin]
  # Enables or disables the admin endpoint.
  disabled = false

  # Address to bind the admin endpoint to.
  bind = "127.0.0.1:7471"

  # Enables or disables the usage of TLS protocol on the 
  # admin endpoint.
  disable-tls = true

  # Absolute path to a .pem encoded TLS certificate.
  tls-certificate = ""

  # Absolute path to a .pem encoded TLS private key.
  tls-key = ""

  # Absolute path to the web files served on the admin endpoint.
  website-files = "/opt/vorteil/web/admin"

[loopback]
  # Enables or disables the loopback endpoint.
  disabled = false

  # Address to bind the loopback endpoint to.
  bind = "127.0.0.1:7472"

  # Enables or disables the usage of TLS protocol on the 
  # loopback endpoint.
  disable-tls = true

  # Absolute path to a .pem encoded TLS certificate.
  tls-certificate = ""

  # Absolute path to a .pem encoded TLS private key.
  tls-key = ""

  # Absolute path to the web files served on the loopback endpoint.
  website-files = "/opt/vorteil/web/vrepo"

# Multiple protected 'network' endpoints can be configured.
# Incoming connections will be subject to user authentication.
[[network]]
  # Enables or disables this network endpoint.
  disabled = true

  # Address to bind this network endpoint to.
  bind = "0.0.0.0:7472"

  # Enables or disables the usage of TLS protocol on this
  # network endpoint.
  disable-tls = true

  # Absolute path to a .pem encoded TLS certificate.
  tls-certificate = ""

  # Absolute path to a .pem encoded TLS private key.
  tls-key = ""

  # Absolute path to the web files served on this network endpoint.
  website-files = "/opt/vorteil/web/vrepo"

```