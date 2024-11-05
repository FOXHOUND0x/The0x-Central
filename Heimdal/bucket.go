package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createBucket(ctx *pulumi.Context) (*s3.Bucket, error) {
	bucket, err := s3.NewBucket(ctx, "ragnarBucket", &s3.BucketArgs{
		Bucket: pulumi.String("ragnar-bucket"),
	})
	errorHandler(err)

	_, err = s3.NewBucketOwnershipControls(ctx, "ragnarBucketOwnershipControls", &s3.BucketOwnershipControlsArgs{
		Bucket: bucket.ID(),
		Rule: &s3.BucketOwnershipControlsRuleArgs{
			ObjectOwnership: pulumi.String("BucketOwnerPreferred"),
		},
	})

	_, err = s3.NewBucketPublicAccessBlock(ctx, "ragnarBucketPublicAccessBlock", &s3.BucketPublicAccessBlockArgs{
		Bucket:                bucket.ID(),
		BlockPublicAcls:       pulumi.Bool(true),
		BlockPublicPolicy:     pulumi.Bool(true),
		IgnorePublicAcls:      pulumi.Bool(true),
		RestrictPublicBuckets: pulumi.Bool(true),
	})
	errorHandler(err)

	_, err = s3.NewBucketVersioningV2(ctx, "versioning_enabled", &s3.BucketVersioningV2Args{
		Bucket: bucket.ID(),
		VersioningConfiguration: &s3.BucketVersioningV2VersioningConfigurationArgs{
			Status: pulumi.String("Enabled"),
		},
	})
	errorHandler(err)
	return bucket, nil
}