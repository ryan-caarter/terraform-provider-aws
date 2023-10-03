// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package transfer

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	transfer_sdkv1 "github.com/aws/aws-sdk-go/service/transfer"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceServer,
			TypeName: "aws_transfer_server",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAccess,
			TypeName: "aws_transfer_access",
		},
		{
			Factory:  ResourceAgreement,
			TypeName: "aws_transfer_agreement",
			Name:     "Agreement",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceCertificate,
			TypeName: "aws_transfer_certificate",
			Name:     "Certificate",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceConnector,
			TypeName: "aws_transfer_connector",
			Name:     "Connector",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceProfile,
			TypeName: "aws_transfer_profile",
			Name:     "Profile",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceServer,
			TypeName: "aws_transfer_server",
			Name:     "Server",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceSSHKey,
			TypeName: "aws_transfer_ssh_key",
		},
		{
			Factory:  ResourceTag,
			TypeName: "aws_transfer_tag",
		},
		{
			Factory:  ResourceUser,
			TypeName: "aws_transfer_user",
			Name:     "User",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceWorkflow,
			TypeName: "aws_transfer_workflow",
			Name:     "Workflow",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Transfer
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*transfer_sdkv1.Transfer, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return transfer_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
