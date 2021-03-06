package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Test the Terraform module in examples/complete using Terratest.
func TestExamplesComplete(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../../examples/complete",
		Upgrade:      true,
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"fixtures.us-east-2.tfvars"},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of an output variable
	vpcCidr := terraform.Output(t, terraformOptions, "vpc_cidr")
	// Verify we're getting back the outputs we expect
	assert.Equal(t, "172.16.0.0/16", vpcCidr)

	// Run `terraform output` to get the value of an output variable
	privateSubnetCidrs := terraform.OutputList(t, terraformOptions, "private_subnet_cidrs")
	// Verify we're getting back the outputs we expect
	assert.Equal(t, []string{"172.16.0.0/19", "172.16.32.0/19"}, privateSubnetCidrs)

	// Run `terraform output` to get the value of an output variable
	publicSubnetCidrs := terraform.OutputList(t, terraformOptions, "public_subnet_cidrs")
	// Verify we're getting back the outputs we expect
	assert.Equal(t, []string{"172.16.96.0/19", "172.16.128.0/19"}, publicSubnetCidrs)

	/* TODO: re-enable when bucket encryption issue is resolved for NLBs
	// Run `terraform output` to get the value of an output variable
	accessLogsBucketId := terraform.Output(t, terraformOptions, "access_logs_bucket_id")
	// Verify we're getting back the outputs we expect
	assert.Equal(t, "eg-test-nlb-nlb-access-logs", accessLogsBucketId)
	*/

	// Run `terraform output` to get the value of an output variable
	nlbName := terraform.Output(t, terraformOptions, "nlb_name")
	// Verify we're getting back the outputs we expect
	assert.Equal(t, "eg-test-nlb", nlbName)

	// Run `terraform output` to get the value of an output variable
	defaultTargetGroupArn := terraform.Output(t, terraformOptions, "default_target_group_arn")
	// Verify we're getting back the outputs we expect
	assert.Contains(t, defaultTargetGroupArn, ":targetgroup/eg-test-nlb-default")

	// Run `terraform output` to get the value of an output variable
	defaultListenerArn := terraform.Output(t, terraformOptions, "default_listener_arn")
	// Verify we're getting back the outputs we expect
	assert.Contains(t, defaultListenerArn, ":listener/net/eg-test-nlb")
}
