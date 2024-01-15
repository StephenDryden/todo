variable "account_id" {
  description = "aws account id"
  type        = string
  default     = "589113656956"
}

variable "region" {
  description = "aws region"
  type        = string
  default     = "eu-west-2"
}

variable "project_name" {
  description = "the name of the project used for naming resources"
  type        = string
  default     = "todo"
}

variable "table_name" {
  description = "the name of the table to create"
  type        = string
  default     = "todo"
}