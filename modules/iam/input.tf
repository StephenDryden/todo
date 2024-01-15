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