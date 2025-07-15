create a bucket 

```sh
aws s3 mb s3://buckt-policy-3825/
```

to get existing policy 
````sh
aws s3api get-bucket-policy --bucket buckt-policy-3825 
 ````


mkae a policy for other aws accounts to access your bucket

```sh
aws s3api put-bucket-policy --bucket buckt-policy-3825 --policy file://policy.json
```

while setting this if other user wants to list your bucket resource is bucket itself, my-bucky
if other user wants t get objects then the resouce isevethign in the bucket my-bucky/*


cleanup

```sh
    aws s3 rm s3://buckt-policy-3825/ --recursive 
```