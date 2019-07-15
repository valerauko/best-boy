# Best Boy

Lambda code to start a transcoder job on receiving an S3 create event.

Based on @mmyoji's [guide](https://dev.to/mmyoji/video-processing-with-aws-lambda--elastic-transcoder-in-golang--hf2).

## Compiling
```
docker-compose up
```

That's it.

## Deploying

Best Boy uses [Serverless](https://serverless.com/) to manage its deployment.

Make an environment config YAML based on the [sample config](https://github.com/valerauko/best-boy/blob/master/conf/env.yml.sample). Name it `dev.yml`, since `dev` is the default stage. Staging would be `stg` and production `prd`.

Then you can use `serverless` commands. A `dev` deployment would be `serverless deploy`. For the other stages, the `--stage` option is needed: `serverless deploy --stage stg`

### Limitations

Serverless currently has no way to attach events to already existing S3 buckets (see [relevant issue](https://github.com/serverless/serverless/issues/3257)), so you'll have to create that manually.

On the AWS console, find the input S3 bucket, go to Properties, then "Add notification" under Events.

* `Name`: really whatever
* `Events`: as needed. I use "All object create events"
* `Prefix` should be the same as you configured Best Boy to use.
* `Send to`: the `Lambda function` Serverless created (would be `video-transcoder-dev-videoTranscoder` for example).

That's it!
