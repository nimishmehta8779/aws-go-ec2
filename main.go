package main

import (
	"github.com/nimishmehta8779/aws-go-ec2/e2"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
		// Export the name of the bucket
		//		ctx.Export("bucketName", ec2.ID())

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		subnetid, err := ec2.LookupSubnet(ctx, &ec2.LookupSubnetArgs{
			Filters: []ec2.GetSubnetFilter{
				ec2.GetSubnetFilter{
					Name: "tag:Name",
					Values: []string{
						"aws-go-obj-aws-go-private-subnet-0",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}

		// Create an AWS resource (ec2 Bucket)
		_, err = e2.NewEc2(ctx, "myinstance", &e2.Ec2Input{
			Size:     "t2.medium",
			SubnetID: subnetid.Id,
		})
		if err != nil {
			return err
		}

		return nil
	})
}
