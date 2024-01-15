# todo WIP

todo is a bare bones simple serverless todo list app built using dynamodb, lambda and api gateway, deployed via Terraform.

Simply export your AWS credentails for a user that has access to dynamodb, lambda, cloudwatch and api gateway and then run your normal terraform commands. Do not under any circumstance commit your AWS credentials.

```
export AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
export AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
export AWS_DEFAULT_REGION=us-west-2
terraform plan
```
