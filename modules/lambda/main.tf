data "archive_file" "lambda" {
  type        = "zip"
  source_file = "src/bin/bootstrap"
  output_path = "src/bin/bootstrap.zip"
}

resource "aws_lambda_function" "todo" {
  filename      = "src/bin/bootstrap.zip"
  function_name = "${var.project_name}-function"
  role          = var.role_arn
  handler       = "Handler"

  source_code_hash = data.archive_file.lambda.output_base64sha256

  runtime = "provided.al2"
}

output "invoke_arn" {
  value = aws_lambda_function.todo.invoke_arn
}

output "function_name" {
  value = aws_lambda_function.todo.function_name
}