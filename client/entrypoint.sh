#!/bin/sh

start_client() {

  CMDSTRING="ls"

  echo $CMDSTRING
}

start_client 

# SIGUSR1-handler
my_handler() {
  exit 143; # 128 + 15 -- SIGTERM
}

# SIGTERM-handler
term_handler() {
  exit 143; # 128 + 15 -- SIGTERM
}

# setup handlers
# on callback, kill the last background process, which is `tail -f /dev/null` and execute the specified handler
trap 'kill ${!}; my_handler' SIGUSR1
trap 'kill ${!}; term_handler' SIGTERM

while true
do
  tail -f /dev/null & wait ${!}
done