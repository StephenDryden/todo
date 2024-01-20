data "aws_iam_policy_document" "lambda-trust" {
  statement {
    actions    = ["sts:AssumeRole"]
    effect     = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "lambda" {
  statement {
    actions = [
        "logs:CreateLogGroup",
    ]

    resources = [
      "arn:aws:logs:${var.region}:${var.account_id}:*",
    ]
  }
  statement {
    actions = [
        "logs:CreateLogStream",
        "logs:PutLogEvents",
    ]

    resources = [
      "arn:aws:logs:${var.region}:${var.account_id}:log-group:/aws/lambda/${var.project_name}function:*",
    ]
  }

  statement {
    actions = [
        "dynamodb:DeleteItem",
        "dynamodb:GetItem",
        "dynamodb:Scan",
        "dynamodb:PutItem",
        "dynamodb:UpdateItem",
    ]

    resources = [
      "arn:aws:dynamodb:${var.region}:${var.account_id}:table/*",
    ]
  }  
}

resource "aws_iam_policy" "lambda" {
  name   = "${var.project_name}-lambda-role-policy"
  path   = "/"
  policy = data.aws_iam_policy_document.lambda.json
}

resource "aws_iam_role" "lambda" {
  name               = "${var.project_name}-lambda-role"
  assume_role_policy = "${data.aws_iam_policy_document.lambda-trust.json}"
}

resource "aws_iam_role_policy_attachment" "lambda-attachment" {
  role       = "${aws_iam_role.lambda.name}"
  policy_arn = aws_iam_policy.lambda.arn
}

output "role_arn" {
  value = aws_iam_role.lambda.arn
}