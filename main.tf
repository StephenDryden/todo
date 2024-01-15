# ./modules/dynamodb/main.tf
module "dynamodb" {
  source = "./modules/dynamodb"
  table_name = var.table_name
}

# ./modules/iam/main.tf
module "iam" {
  source = "./modules/iam"
  account_id = var.account_id
  region = var.region
  project_name = var.project_name
}

# ./modules/lambda/main.tf
module "lambda" {
  source = "./modules/lambda"
  role_arn = module.iam.role_arn
  project_name = var.project_name
}

# ./modules/apigateway/main.tf
module "apigateway" {
  source = "./modules/apigateway"
  account_id = var.account_id
  region = var.region
  project_name = var.project_name
  table_name = var.table_name
  function_name = module.lambda.function_name
  invoke_arn = module.lambda.invoke_arn
}