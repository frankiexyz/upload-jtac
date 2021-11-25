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

# Credit goes to [http://networkbit.ch/golang-sftp-client/]
