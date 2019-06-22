package main

import (
  "fmt"
  "log"
  "os"
  "strings"
  "path/filepath"

  "github.com/aws/aws-lambda-go/lambda"
  "github.com/aws/aws-lambda-go/events"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  transcoder "github.com/aws/aws-sdk-go/service/elastictranscoder"
)

func startJob(record events.S3Entity) (*transcoder.CreateJobResponse, error) {
  bucket := record.Bucket.Name

  input_prefix := os.Getenv("INPUT_PREFIX")
  video_prefix := os.Getenv("VIDEO_PREFIX")
  thumb_prefix := os.Getenv("THUMB_PREFIX")

  // the key gets URL encoded when passed as an AWS event
  // so if it contains any characters that would be changed
  // by URL encoding (space, non-ascii etc), the job fails
  input_key := record.Object.Key
  without_ext := strings.TrimSuffix(input_key, filepath.Ext(input_key))

  output_key := strings.Replace(input_key, input_prefix, video_prefix, 1)
  // the thumbnail then gets processed
  thumb_pattern := strings.Replace(
    fmt.Sprintf("%s-{count}", without_ext),
    input_prefix, thumb_prefix, 1,
  )

  log.Printf("Transcoding %s / %s\n", bucket, input_key)

  sess := session.Must(
    session.NewSession(
      &aws.Config{
        Region: aws.String("ap-northeast-1"),
      },
    ),
  )
  service := transcoder.New(sess)

  return service.CreateJob(
    &transcoder.CreateJobInput{
      Input: &transcoder.JobInput {
        Key: aws.String(input_key),
      },
      Outputs: []*transcoder.CreateJobOutput {
        &transcoder.CreateJobOutput {
          PresetId: aws.String(os.Getenv("PRESET_ID")),
          Key: aws.String(output_key),
          ThumbnailPattern: aws.String(thumb_pattern),
        },
      },
      PipelineId: aws.String(os.Getenv("PIPELINE_ID")),
    },
  )
}

func HandleLambdaEvent(event events.S3Event) {
  for _, record := range event.Records {
    resp, err := startJob(record.S3)
    if err != nil {
      log.Printf("Failed to create job: %v\n", err)
    } else {
      log.Printf("Job created: %v\n", resp)
    }
  }
}

func main() {
  lambda.Start(HandleLambdaEvent)
}
