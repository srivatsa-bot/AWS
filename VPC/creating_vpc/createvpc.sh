#!/bin/bash

vpcID=$(
    aws ec2 create-vpc\
    --cidr-block 10.0.0.0/16 \
    --tag-specifications 'ResourceType=vpc,Tags=[{Key=ENV,Value=shelll}]'\
    --output text\
    --query Vpc.VpcId
)
echo "VPC_ID: $vpcID"

igwID=$(
    aws ec2 create-internet-gateway --tag-specifications 'ResourceType=internet-gateway,Tags=[{Key=ENV,Value=shelll}]' --query InternetGateway.InternetGatewayId --output text
)
echo "IGW_ID: $igwID"

aws ec2 attach-internet-gateway --vpc-id "$vpcID" --internet-gateway-id "$igwID"
echo "Attached igw to vpc"

subID=$(
    aws ec2 create-subnet --vpc-id "$vpcID" --cidr-block 10.0.0.0/18 --tag-specifications 'ResourceType=subnet,Tags=[{Key=ENV,Value=shell}]' --query Subnet.SubnetId --output text
)
echo "subnet_id: $subID"

routeID=$(
    aws ec2 describe-route-tables \
  --filters Name=vpc-id,Values="$vpcID" \
  --query 'RouteTables[*].RouteTableId' \
  --output text
)
echo "rotetable_id: $routeID"

aws ec2 associate-route-table --route-table-id "$routeID" --subnet-id "$subID"
echo "created association between route table and subnet"

aws ec2 create-route --route-table-id "$routeID" --destination-cidr-block 0.0.0.0/0 --gateway-id "$igwID"
echo "created association between routetable and igw"


aws ec2 modify-vpc-attribute --vpc-id "$vpcID" --enable-dns-hostnames
echo "enabled vpc hostname for vpc"

echo "vpc created sucesfully"


echo "To dlete the vpc use this command with the scriptS"
echo "./deletevpc.sh $igwID $subID $routeID $vpcID"


aws ec2 describe-route-tables --route-table-ids rtb-0b9c00811501d8c03 --query "RouteTables[0].Associations[?SubnetId=='subnet-0c1b758b1d34e7224'].RouteTableAssociationId" --output text