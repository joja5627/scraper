provider "aws" {
    access_key = ""
    secret_key = ""
    region = "us-east-1"
}

resource "aws_instance" "example" {
    ami = "ami-00dc79254d0461090"
    instance_type = "t2.micro"
}