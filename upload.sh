export AWS_PROFILE=expensetracker
export AWS_REGION=eu-central-1
aws sts get-caller-identity | jq -r ".Arn"
aws iam list-account-aliases | jq -r ".AccountAliases[0]"
GOOS=linux GOARCH=amd64 go build -o expensetracker
zip expensetracker expensetracker
aws lambda update-function-code --function-name expense-tracker --zip-file fileb://expensetracker.zip
rm expensetracker
rm expensetracker.zip
