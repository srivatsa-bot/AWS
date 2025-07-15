create buckert
aws s3api create-bucket --bucket mybucky --create-bucket-configuration LocationConstraint=ap-south-1

#to check for publi caccess block 
aws s3api get-public-access-block --bucket mybucky3825

# to turn off public acl block
aws s3api put-public-access-block --bucket mybucky3825 --public-access-block-configuration "BlockPublicAcls=false IgnorePublicAcls=false,BlockPublicPolicy=true,RestrictPublicBuckets=true"

# acl also depends upon the bucket ownership cpntrol

aws s3api put-bucket-ownership-controls \
    --bucket amzn-s3-demo-bucket \
    --ownership-controls="Rules=[{ObjectOwnership=BucketOwnerPreferred}]"

# change acl using thins
aws s3api put-bucket-acl --bucket mybucky3825 --acl privatea


#d elete all objects in s3 usling cli 



