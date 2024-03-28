package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAwsVpcPrivateSubnetsOnly(t *testing.T) {
	t.Parallel()

	// Make a copy of the terraform module to a temporary directory. This allows running multiple tests in parallel
	// against the same terraform module.
	exampleFolder := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/private-only")

	// Give this VPC a unique ID for a name tag so we can distinguish it from any other VPC provisioned
	// in your AWS account
	vpcName := fmt.Sprintf("private-only-example-%s", strings.ToLower(random.UniqueId()))

	// Pick a random AWS region to test in. This helps ensure your code works in all regions.
	awsRegion := aws.GetRandomStableRegion(t, nil, nil)

	// Configure Terraform setting path to Terraform code, VPC name, and AWS Region. We also configure
	// the options with default retryable errors to handle the most common retryable errors encountered in
	// terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: exampleFolder,
		Vars: map[string]interface{}{
			"vpc_name": vpcName,
			"region":   awsRegion,
		},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Define outputs
	vpcId := terraform.Output(t, terraformOptions, "vpc_id")
	privateSubnetIds := terraform.OutputList(t, terraformOptions, "private_subnet_ids")

	// Check that the VPC was created and the name tag is correct
	vpc := aws.GetVpcById(t, vpcId, awsRegion)

	// Check that each subnet was created and is public/private
	publicSubnetCount := 0
	privateSubnetCount := 0
	for i := range vpc.Subnets {
		if aws.IsPublicSubnet(t, vpc.Subnets[i].Id, awsRegion) {
			publicSubnetCount++
		} else {
			assert.Contains(t, privateSubnetIds, vpc.Subnets[i].Id, "Private subnet not found in private_subnet_ids")
			privateSubnetCount++
		}
	}

	// Check that the number of public and private subnets match the expected number
	assert.Equal(t, 0, publicSubnetCount, "Number of public subnets does not match")
	assert.Equal(t, len(privateSubnetIds), privateSubnetCount, "Number of private subnets does not match")

}
