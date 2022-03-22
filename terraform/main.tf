terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
  access_key = var.access_key
  secret_key = var.secret_key
}

locals {
  web_lambda_name = format("%s-web", var.lambda_function_name)
}

resource "aws_iam_role" "lambda_role" {
  name = "${var.environment_name}-role-addons"

  assume_role_policy = file(format("%s/templates/assume-role-policy.json", path.module))
}

resource "aws_iam_policy" "lambda_policy" {
  name   = "lambda-policy-addons"
  policy = templatefile(format("%s/templates/role-policy.json.tmpl", path.module), {})
}

resource "aws_iam_role_policy_attachment" "lambda_policy_attach" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = aws_iam_policy.lambda_policy.arn
}

resource "aws_lambda_function" "web_lambda" {
  function_name = local.web_lambda_name
  package_type  = "Image"
  image_uri     = "499781533590.dkr.ecr.us-east-1.amazonaws.com/test-lambda:master_f3a4ca"
  
  role = aws_iam_role.lambda_role.arn

  memory_size = 1024
  timeout     = 28
}
