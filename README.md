# A simple script to upload file to Juniper SFTP [https://kb.juniper.net/InfoCenter/index?page=content&id=KB23337]
```
 ./upload-jtac --help
Usage of ./upload-jtac:
  -caseid string
        Input Your caseid.
  -file string
        Input Your file like /tmp/log.tar.gz

./upload-jtac -caseid 2021-1125-369000 -file /tmp/re-0-var-log.tar 
INFO[0000] case id:2021-1125-369000                     
INFO[0001] Trying to connect to sftp.juniper.net        
INFO[0003] trying to create file on sftp /pub/incoming/2021-1125-369000/re-0-var-log.tar 
INFO[0003] trying to upload file to sftp /pub/incoming/2021-1125-369000 
19228160 bytes copied
```

### How to collect Juniper log ###
1. capture /var/log
```
start shell user root
tar -zcvf /var/tmp/varlog-mem0.tar.gz /var/log/*
```
2. copy a coredump from a junos host
```
start shell user root
request app-engine host-shell
request app-engine file-copy from-jhost erequest app-engine file-copy from-jhost e01.aaa-node.dcpfe.9765.1638630532.core.tgz to-vjunos /var/tmp/ crash
and find your file in /var/tmp/
```
# Credit goes to [http://networkbit.ch/golang-sftp-client/]
