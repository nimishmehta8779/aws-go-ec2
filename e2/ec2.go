package e2

import (
	"fmt"

	"github.com/nimishmehta8779/aws-go-obj/util"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ec2Input struct {
	Size     string
	SubnetID string
}

type Ec2Output struct {
	e2 *ec2.Instance
}

type IdOutput struct {
	ID string
}

func (in *Ec2Input) Validate() error {
	if len(in.Size) == 0 {
		return fmt.Errorf("no ec2 instance are provided")
	}
	return nil
}

func NewEc2(ctx *pulumi.Context, name string, in *Ec2Input) (*Ec2Output, error) {
	var err error
	if err := in.Validate(); err != nil {
		return nil, fmt.Errorf("while validating ec2 instance")

	}

	output := &Ec2Output{}
	ami, _ := AmiSearch(ctx)

	output.e2, err = ec2.NewInstance(ctx, name, &ec2.InstanceArgs{
		InstanceType: pulumi.String(in.Size),
		Ami:          pulumi.String(ami.ID),
		SubnetId:     pulumi.StringPtr(in.SubnetID),
		Tags:         pulumi.ToStringMap(util.NewNameTags(ctx, name)),
	})
	if err != nil {
		return nil, err
	}
	ctx.Export("Instance-id", output.e2.ID())
	return output, nil
}

func AmiSearch(ctx *pulumi.Context) (*IdOutput, error) {
	var recent bool = true
	id, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
		Filters: []ec2.GetAmiFilter{
			ec2.GetAmiFilter{
				Name: "name",
				Values: []string{
					"amzn-ami-hvm-*-x86_64-ebs",
				},
			},
		},
		Owners: []string{
			"137112412989",
		},
		MostRecent: &recent,
	}, nil)
	if err != nil {
		return nil, err
	}
	return &IdOutput{
		id.Id,
	}, err
}
