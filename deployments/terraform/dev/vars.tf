variable "region" {
  description = "specifies aws region"
  default     = "us-east-1"
}

variable "artifact_bucket" {
  description = "the bucket for fetching the artifact"
  default     = "artifact-bucket-gin-test"
}

variable "artifact_binary_name" {
  description = "name of the zip file"
  default     = "main"
}

