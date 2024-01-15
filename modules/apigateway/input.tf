variable "account_id" {
  description = "aws account id"
  type        = string
}

variable "region" {
  description = "aws region"
  type        = string
  default     = "eu-west-2"
}

variable "project_name" {
  description = "the name of the project used for naming resources"
  type        = string
}

variable "invoke_arn" {
  description = "arn of lambda to invoke"
  type        = string
}

variable "function_name" {
  description = "name of the lambda function"
  type        = string
}

variable "table_name" {
  description = "the name of the table to create"
  type        = string
}