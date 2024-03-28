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

func TestTerraformAwsVpcPartialAvailabilityZones(t *testing.T) {
	t.Parallel()

	// Make a copy of the terraform module to a temporary directory. This allows running multiple tests in parallel
	// against the same terraform module.
	exampleFolder := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/partial-azs")

	// Give this VPC a unique ID for a name tag so we can distinguish it from any other VPC provisioned
	// in your AWS account
	vpcName := fmt.Sprintf("partial-azs-example-%s", strings.ToLower(random.UniqueId()))

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
	publicSubnetIds := terraform.OutputList(t, terraformOptions, "public_subnet_ids")
	privateSubnetIds := terraform.OutputList(t, terraformOptions, "private_subnet_ids")
	availabilityZones := terraform.OutputList(t, terraformOptions, "availability_zones")

	// Check that only some of the availability zones were used in the VPC
	availabilityZonesInRegion := aws.GetAvailabilityZones(t, awsRegion)
	assert.True(t, len(availabilityZones) < len(availabilityZonesInRegion), "Number of availability zones used in the VPC is not less than the number of availability zones in the region")

	// Check that the number of subnets created matches the number of
	// availability zones used in the VPC
	assert.Equal(t, len(publicSubnetIds), len(availabilityZones))
	assert.Equal(t, len(privateSubnetIds), len(availabilityZones))
}
