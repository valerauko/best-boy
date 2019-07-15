# Best Boy

Lambda code to start a transcoder job on receiving an S3 create event.

Based on @mmyoji's [guide](https://dev.to/mmyoji/video-processing-with-aws-lambda--elastic-transcoder-in-golang--hf2).

## Compiling
```
docker-compose run src
```

That's it.

## Set-up

The following environment variables can be used to configure:
* PIPELINE_ID
* PRESET_ID
* INPUT_PREFIX
* VIDEO_PREFIX
* THUMB_PREFIX

The pipeline and preset IDs are obviously required, and Go will probably throw some null pointer exception if the others are missing too.
