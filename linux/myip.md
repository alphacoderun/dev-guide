How Do I find my IP address in Windows Linux SubSystem?

Run the command

`ip addr | grep eth0`

Get the exmaple output 

```
4: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP group default qlen 1000
    inet 174.23.212.4/20 brd 174.23.243.255 scope global eth0
```

The IP address for wls in this example would be 174.23.212.4