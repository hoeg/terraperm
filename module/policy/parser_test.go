package policy

import (
	"bytes"
	"testing"
)

func TestParserOneItem(t *testing.T) {
	log := `
	2020-09-20T21:36:31.528+0200 [DEBUG] plugin.terraform-provider-aws_v3.7.0_x5: 2020/09/20 21:36:31 [DEBUG] [aws-sdk-go] DEBUG: Request s3/CreateBucket Details:
2020-09-20T21:36:31.528+0200 [DEBUG] plugin.terraform-provider-aws_v3.7.0_x5: ---[ REQUEST POST-SIGN ]-----------------------------
2020-09-20T21:36:31.528+0200 [DEBUG] plugin.terraform-provider-aws_v3.7.0_x5: PUT / HTTP/1.1
2020-09-20T21:36:31.528+0200 [DEBUG] plugin.terraform-provider-aws_v3.7.0_x5: Host: hoeg-tf-test-recordingstudio.s3.eu-west-1.amazonaws.com
2020-09-20T21:36:31.528+0200 [DEBUG] plugin.terraform-provider-aws_v3.7.0_x5: User-Agent: aws-sdk-go/1.34.21 (go1.14.5; linux; amd64) APN/1.0 HashiCorp/1.0 Terraform/0.13.2 (+https://www.terraform.io)
2020-09-20T21:36:31.528+0200 [DEBUG] plugin.terraform-provider-aws_v3.7.0_x5: Content-Length: 153
2020-09-20T21:36:31.528+0200 [DEBUG] plugin.terraform-provider-aws_v3.7.0_x5: Authorization: AWS4-HMAC-SHA256 Credential=AKIA4TV6MKMV6ZNQRJNN/20200920/eu-west-1/s3/aws4_request, SignedHeaders=content-length;host;x-amz-acl;x-amz-content-sha256;x-amz-date, Signature=134fd1bcb305ac797f652c2cd7a0987797ff45aa2e6e06c44ba15c4b6d998f4b
2020-09-20T21:36:31.528+0200 [DEBUG] plugin.terraform-provider-aws_v3.7.0_x5: X-Amz-Acl: private
2020-09-20T21:36:31.528+0200 [DEBUG] plugin.terraform-provider-aws_v3.7.0_x5: X-Amz-Content-Sha256: a2531158b25edb57200e54140cc1b12d0af943f596ea467c2fbbedcc16223cc5
2020-09-20T21:36:31.528+0200 [DEBUG] plugin.terraform-provider-aws_v3.7.0_x5: X-Amz-Date: 20200920T193631Z
2020-09-20T21:36:31.528+0200 [DEBUG] plugin.terraform-provider-aws_v3.7.0_x5: Accept-Encoding: gzip
2020-09-20T21:36:31.528+0200 [DEBUG] plugin.terraform-provider-aws_v3.7.0_x5: 
2020-09-20T21:36:31.528+0200 [DEBUG] plugin.terraform-provider-aws_v3.7.0_x5: <CreateBucketConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/"><LocationConstraint>eu-west-1</LocationConstraint></CreateBucketConfiguration>
2020-09-20T21:36:31.528+0200 [DEBUG] plugin.terraform-provider-aws_v3.7.0_x5: -----------------------------------------------------
	`

	trace := bytes.NewBuffer([]byte(log))
	p := NewParser(trace)
	reqs, err := p.Requests()
	if err != nil {
		t.Fatalf("failed parse: %v", err)
	}
	if len(reqs) != 1 {
		t.Fatalf("expected 1 request, got %d", len(reqs))
	}
	if reqs[0].apiKey != "s3/CreateBucket" {
		t.Fatalf("expected s3/CreateBucket request, got %q", reqs[0].apiKey)
	}

}
