output "invoke_url" {
  description = "The url to invoke for the todo application"
  value       = module.apigateway.invoke_url
}