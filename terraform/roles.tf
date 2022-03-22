resource "aws_iam_role" "test_lambda_role" {
  name = "${var.environment_name}-role-test"

  assume_role_policy = file(format("%s/templates/assume-role-policy.json", path.module))
}

resource "aws_iam_role" "invoker_lambda_role" {
  name = "${var.environment_name}-role-invoker"

  assume_role_policy = file(format("%s/templates/assume-role-policy.json", path.module))
}

resource "aws_iam_policy" "test_lambda_policy" {
  name   = "test-lambda-policy"
  policy = templatefile(format("%s/templates/test-role-policy.json.tmpl", path.module), {})
}

resource "aws_iam_policy" "invoker_lambda_policy" {
  name   = "invoker-lambda-policy"
  policy = templatefile(format("%s/templates/invoker-role-policy.json.tmpl", path.module), {})
}

resource "aws_iam_role_policy_attachment" "test_lambda_policy_attach" {
  role       = aws_iam_role.test_lambda_role.name
  policy_arn = aws_iam_policy.test_lambda_policy.arn
}

resource "aws_iam_role_policy_attachment" "invoker_lambda_policy_attach" {
  role       = aws_iam_role.invoker_lambda_role.name
  policy_arn = aws_iam_policy.invoker_lambda_policy.arn
}
