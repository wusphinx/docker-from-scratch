# UTS
UTS namespace主要用来隔离 nodename 和 domainname

# run
## before
```
ubuntu@engaging-lab:~/docker-from-scratch/uts$ hostname
engaging-lab
```

## after 
```
ubuntu@engaging-lab:~/docker-from-scratch/uts$ sudo go run main.go 
# hostname -b myself
# hostname
myself
# exit
ubuntu@engaging-lab:~/docker-from-scratch/uts$ hostname
engaging-lab
```

# 结论
fork出来的shell进程内对hostname的修改并不会影响父进程也就是宿主机的hostname