create a yaml fle containg resources and with create-stack command launch the reources

```sh 
aws cloudformation create-stack --stack-name temp-prac --template-body file://VPC/clodformation/resource.yaml --region ap-south-1
```

note templete contains 3 main parts

```sh 
---
AWSTemplateFormatVersion: version date

Description:
  String

Metadata:
  template metadata

Parameters:
  set of parameters

Rules:
  set of rules

Mappings:
  set of mappings

Conditions:
  set of conditions

Transform:
  set of transforms

Resources:
  set of resources
  resorce logical name1:

  resoucelogic name2:


Outputs:
  set of outputs

```

abovve command will create a stack which you can review and to apply this stack 

```sh
aws cloudformation update-stack \
  --stack-name temp-prac \
  --template-body file://VPC/clodformation/resource.yaml \
  --region ap-south-1
```

or you can check how the stack looks

```sh
aws cloudformation describe-stacks --stack-name temp-prac --region ap-south-1
```

to delete the resorces we deployed we just need to delete the stack cloud formation will delet it the resources automatically

```sh
aws cloudformation delete-stack --stack-name temp-prac --region ap-south-1
```

You can also update your satck
```sh
   aws cloudformation update-stack \
     --stack-name <your-stack-name> \
     --template-body file://VPC/clodformation/resource.yaml \
     --region <your-region>
```

#note differnce between stack vs stack set
```sh
Use a stack for single-account, single-region deployments.
Use a stack set for deploying the same resources to multiple accounts and/or regions automatically.
```