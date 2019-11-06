variable "app_version" {
}
variable "redis_url" {
}
variable "redis_pswd" {
}

provider "aws" {
  region = "us-east-2"
}

module "hex-api" {
  source = "../modules/api-lambda"

  name = "gira"
  display_name = "Hex Example"
  bucket = "hex-lambda-1"
  app_version = "${var.app_version}"
  stage = "qa"
  env_vars = {
      DATABASE_URL = "${var.redis_url}"
      REDIS_PASSWORD = "${var.redis_pswd}"
  }
}

output "url" {
  value = "${module.hex-api.base_url}"
}
