#impleting cors using static websitehosting in s3
https://docs.aws.amazon.com/AmazonS3/latest/userguide/HostingWebsiteOnS3Setup.html
#create a bucket

```sh 
aws s3 mb s3://cors-3825/
```

#edit block public access

```sh
aws s3api put-public-access-block \
    --bucket cors2-3825 \
    --public-access-block-configuration "BlockPublicAcls=true,IgnorePublicAcls=true,BlockPublicPolicy=false,RestrictPublicBuckets=false"
```

#put bucket policy to access by others 

```sh 
aws s3api put-bucket-policy --bucket cors2-3825 --policy file://policy.json
```

#create an index.html and enable static hosting and upload it to s3

```sh 
aws s3 cp index.html s3://cors-3825
aws s3 cp error.html s3://cors-3825
```

#configure bucket to lok for this index and error fiels
```sh
aws s3api put-bucket-website --bucket cors-3825 --website-configuration file://website.json
```

#set cors policy on bucket
```sh
aws s3api put-bucket-cors --bucket cors2-3825 --cors-configuration file://cors.json
```

#remove bucket
```sh 
aws s3 rm s3://cors-3825/ --recursive 
aws s3 rm s3://cors2-3825/ --recursive

aws s3 rb s3://cors-3825/
aws s3 rb s3://cors2-3825/
```


