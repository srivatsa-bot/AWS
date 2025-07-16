#we will see 4 tyoes of encryption for s3 bucks

#types of encryption
serevr side encryption, client side encryption

types of sse
1)ss3 - s3 : default encyption used by aws where key is mangaed by aws and data is encrpted ysing aes 256 at rest and tls at tramsit

2)ss3-kms : here key is mangaed by key mangement service

3)ss3 -c : custommer will provide the key to encrpt and decrypt the data by using aes 256. key is never stored in aws its customer responsibulity

4)dss-kms : s3 will request key from kms and encryptsthe file with that key and stores the encypted key in file metadat, and both file(not metadat)  is encrypted using other key same way

#client side - user encrypts and decrypts key is user responsibilty 
--can be done by usng sdk 

I am installing go sdk ...
