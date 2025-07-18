```sh
https://docs.aws.amazon.com/cli/latest/
```
#creating a vpc using aws cli(vpc commands are under ec3 command reference)
`
#create vpc
```sh
aws ec2 create-vpc\
    --cidr-block 10.0.0.0/16 \
    --tag-specifications 'ResourceType=vpc,Tags=[{Key=vpcc,Value=aws-exam}]'\
    --output text\
    --query Vpc.VpcId
``
#create a igw
```sh
aws ec2 create-internet-gateway --tag-specifications 'ResourceType=internet-gateway,Tags=[{Key="aw-exam",Value=IGW}]'
```
#attach igw to vpc
```sh 
aws ec2 attach-internet-gateway --vpc-id "vpc-06990" --internet-gateway-id "igw-029f"
```
#create subnet 
```sh  
aws ec2 create-subnet --vpc-id vpc-06990a161551220ea --cidr-block "10.0.0.0/18" --tag-specifications ResourceType=subnet,Tags=[{Key="aws-exam,Value=subnet"}]
```

#associate route table with subnet
```sh
 aws ec2 associate-route-table --route-table-id rtb-0458f61012ca6e596 --subnet-id subnet-0c8af89f18b6a1cba
```

#create a routre conneting igw and local subnet
```sh
aws ec2 create-route --route-table-id rtb-0458f61012ca6e596 --destination-cidr-block 0.0.0.0/0 --gateway-id igw-029f37b3ac4e601e5
```

#turn ondns hostname so instances inside vpc can have public dns
#enable dns hostname
```sh
aws ec2 modify-vpc-attribute --vpc-id vpc-06990a161551220ea --enable-dns-hostnames
```
#enable dns resolution
```sh 
aws ec2 modify-vpc-attribute --vpc-id <your-vpc-id> --enable-dns-support
```


#to show all vpc(describe can be used for other vpc elemnts)
```sh 
aws ec2 describe-vpcs
```

#to run this script turn off aws cli auto prompt
```sh
export AWS_CLI_AUTO_PROMPT=off
```

#note by default this will create a private subnet if you want to make it public use this 

```sh
aws ec2 modify-subnet-attribute --subnet-id "$subID" --map-public-ip-on-launch
```