output "base_url" {
  value = "${aws_api_gateway_deployment.gw_deploy.invoke_url}"
}
output "s3_artifact_arn2" {
  value = "${aws_s3_bucket.artifact_bucket.arn}"
}
