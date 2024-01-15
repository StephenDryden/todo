# ./http-crud-tutorial-items
terraform {
  required_providers {
    aws = {
      version = "~> 4.59.0"
    }
  }
  backend "s3" {
    bucket         = "stephendryden-state-prd"
    key            = "scrud-go.key"
    region         = "eu-west-2"
    dynamodb_table = "dynamodb-state-locking"
  }
}

provider "aws" {
  region = "eu-west-2"
}