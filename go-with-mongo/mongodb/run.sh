[ "$1" = "mongod" ] || exec "$@" || exit $?

# checks if the user that owns the /data/db directory is "mongodb" and, if not, it changes the ownership of the /data/db directory and all of its contents to "mongodb".
[ "$(stat -c %U /data/db)" = mongodb ] || chown -R mongodb /data/db

# exec provided command as user mongodb
cmd=exec; for i; do cmd="$cmd '$i'"; done
exec su -s /bin/sh -c "$cmd" mongodb