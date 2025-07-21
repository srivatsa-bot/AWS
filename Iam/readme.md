#its best to make iam polcies,roles and suers in cli with gui json policy maker

#here we will discuss how we can mangae 2 aws iam accounst

#name a profile so it can be stored in .aws/creds
```sh
aws configure --profile user-s3
```

#change from one user to other user 2ways 
```sh
aws s3 ls --profile user-s3 #specify in wvey command the username
```
or swith to other use using env varibales

```sh
export AWS_PROFILE=user-s3
```

#use this command to list all users
```sh
aws configure list-profiles
```


#lets discuss about iam policy structure
```sh
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "OptionalStatementID",
      "Effect": "Allow" | "Deny",
      "Action": "service:operation",  #specifies aws service api action can be s3 listObejcts or ec2 describe vpcs
      "Resource": "arn:aws:service:region:account-id:resource", #point stowards the specific resorce that user can acces. if user can access all the put "*"
      "Condition": {
        "StringEquals": {
          "aws:username": "Srivatsa"
        }
      }
    }
  ]
}

```

**[policysym](policysym.aws.amazon.com)** nice website to check if policies are enforced