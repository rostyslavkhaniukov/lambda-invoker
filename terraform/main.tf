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
  invoker_lambda_name = format("%s-invoker", var.lambda_function_name)
}

resource "aws_lambda_function" "web_lambda" {
  function_name = local.web_lambda_name
  package_type  = "Image"
  image_uri     = "499781533590.dkr.ecr.us-east-1.amazonaws.com/test-lambda:master_c1ef7f"

  role = aws_iam_role.test_lambda_role.arn

  memory_size = 1024
  timeout     = 28
}

resource "aws_lambda_function" "invoker_lambda" {
  function_name = local.invoker_lambda_name
  package_type  = "Image"
  image_uri     = "499781533590.dkr.ecr.us-east-1.amazonaws.com/lambda-invoker:master_c1ef7f"

  role = aws_iam_role.invoker_lambda_role.arn

  memory_size = 1024
  timeout     = 28
}
