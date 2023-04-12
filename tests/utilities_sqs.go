package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func testSQS(t *testing.T, variant string) {
	t.Parallel()

	terraformDir := fmt.Sprintf("../examples/%s", variant)

	terraformOptions := &terraform.Options{
		TerraformDir: terraformDir,
		LockTimeout:  "5m",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	awsAccountID := getAWSAccountID(t)
	awsRegion := getAWSRegion(t)

	expectedPartialName := fmt.Sprintf("example-tf-sqs-%s-", variant)

	name := terraform.Output(t, terraformOptions, "name")
	assert.Contains(t, name, expectedPartialName)

	expectedARN := fmt.Sprintf("arn:aws:sqs:%s:%s:%s", awsRegion, awsAccountID, name)

	arn := terraform.Output(t, terraformOptions, "arn")
	assert.Equal(t, expectedARN, arn)

	expectedURL := fmt.Sprintf("https://sqs.%s.amazonaws.com/%s/%s", awsRegion, awsAccountID, name)

	url := terraform.Output(t, terraformOptions, "url")
	assert.Contains(t, expectedURL, url)
}
