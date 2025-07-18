#!/bin/bash
aws ec2 detach-internet-gateway --internet-gateway-id "$1" --vpc-id "$4"
aws ec2 delete-internet-gateway --internet-gateway-id "$1"
echo "deleted IGW"

assoc_id=$(aws ec2 describe-route-tables --route-table-ids "$3" --query "RouteTables[0].Associations[?SubnetId=='$2'].RouteTableAssociationId" --output text)
echo "$assoc_id"
if [ -n "$assoc_id" ]; then
  aws ec2 disassociate-route-table --association-id "$assoc_id"
else
  echo "No association found for subnet $2 in route table $3"
fi

aws ec2 delete-subnet --subnet-id "$2"
echo "deleted subnet"

aws ec2 delete-route-table --route-table-id "$3"
echo "deleted route table"

aws ec2 delete-vpc --vpc-id "$4"
echo "sucessfully deleted vpc"