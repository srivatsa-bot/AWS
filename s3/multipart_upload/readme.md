#this is used for files above 100mb
#we will dive the filei into multiple parts and uplaod each part, when we do that we will get an etag as output
#in a json file we will mention th eetag and the prt no so aws can assenmble it

#First create a session for multipart uplaod
```sh 
aws s3api create-multipart-upload --bucket fada7738 --key lorem.txt
```
#you iwll gt uplaod id as output

#use split cmmand to spli the files byte level(i am going with2 files)
```sh
split -n 2 lorem.txt lorem_part_
```

#now upload parts(note you will not see this ins3 console yet)

```sh
aws s3api upload-part --bucket fada7738 --key lorem.txt --part-number 01 --body lorem_part_aa --upload-id ""
```

#this will return etag of part(checksum)

#now create a json specifying part no with its respective id 

#now use complete-multipart-upload to assemble theparts (this will refelct the changes in s3 gui)

```sh 
aws s3api complete-multipart-upload --multipart-upload file://multipart.json --bucket fada7738 --key lorem.txt --upload-id ""
```

note each part should be more than 5mb except the last part. otherwise error

#use this to create null file with 10 mb
```sh
dd if=/dev/zero of=sample_file bs=1M count=10

```


