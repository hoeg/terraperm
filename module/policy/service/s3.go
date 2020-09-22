package service

type S3Action map[string]func(body string) (string, error)

var S3Actions S3Action = S3Action{
	"CreateBucket":     createBucket,
	"GetBucketTagging": getBucketTagging,
}

func createBucket(body string) (string, error) {
	//"Location"
	return "", nil
}

func getBucketTagging(body string) (string, error) {
	return "", nil
}
