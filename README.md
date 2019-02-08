# Http Task Runner

This is a daemon.
It listens 56565 port and run a command.
All commands must be defined in file /etc/httptaskrunner.yml

It is usefull for run command on the remote server via http(s).

I use it for auto deploy docker-composer images, for pull and restart docker-compose servicess.


example of call:
```
http://localhost:56565?cmd=ls
```

## Build this app
```
make
```

