# Systemd Service File

Use `xuanzang.service` as a systemd service file. It defaults to using a "xuanzang" user with
a homedir of `/var/lib/xuanzang` and the binary lives in `/usr/bin` and the config in
`/etc/xuanzang/xuanzang.yaml`.

In order to work, the systemd unit needs a user named `xuanzang`, an handy way to provide
it is to use the `systemd-sysusers` service by installing the `xuanzang-sysusers.conf` file in the
`sysusers.d` folder (e.g: `/usr/lib/sysusers.d`).