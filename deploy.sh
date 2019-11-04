#!/usr/bin/env bash
TAG=$(git log --pretty=format:'%h' -n 1)

GOOS=linux go build -o main cmd/ticketAPI/main.go
zip main.zip cmd/userAPI/userAPI
# drone-lambda --region us-east-1 \
#   --access-key AKIAVGPWU3NHGX6TA6VV \
#   --secret-key 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 \
#   --function-name upload-s3 \
#   --zip-file deployment.zip
aws s3 cp main.zip s3://hex-lambda-1/$TAG/main.zip
//aws s3api create-bucket --bucket my-bucket --region us-east-1
cd terraform/prod/

terraform apply -var "app_version=$TAG" -auto-approve

cd ../../
rm -rf main.zip
