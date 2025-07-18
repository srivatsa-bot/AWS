
#delete igw
```sh
aws ec2 delete-internet-gateway --internet-gateway-id "$igwID"
```
#delete subnet
```sh
aws ec2 delete-subnet --subnet-id "$subID"
```

#delete route table
```sh
aws ec2 delete-route-table --route-table-id "$routeID"
```
#delete vpc
```sh
aws ec2 delete-vpc --vpc-id "$vpcID"
```


#Correct Deletion Order
```sh
Detach and delete the internet gateway
Disassociate the route table (if association exists)
Delete the subnet
Delete the route table
Delete the VPC
```