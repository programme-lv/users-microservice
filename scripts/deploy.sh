#!/bin/bash

set -e

# Variables
S3_BUCKET="your-s3-bucket-name"
STACK_NAME="your-stack-name"
REGION="your-region"

# Build the project
./scripts/build.sh

# Package the SAM template
sam package --template-file template.yaml --output-template-file packaged.yaml --s3-bucket $S3_BUCKET --region $REGION

# Deploy the SAM stack
sam deploy --template-file packaged.yaml --stack-name $STACK_NAME --capabilities CAPABILITY_IAM --region $REGION
