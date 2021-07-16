# GZip S3 Sync

S3 while serving files does not perform gzip compression. Instead you are expected to upload an already gzipped file with the correct content-encoding set.

This utility performs gzip compression and upload to s3. This makes it convenient to push static site files directly and they can be loaded fast by the user.

## Usage

```shell
$ gzsync s3 sync ~/folder/to/sync s3://my-bucket/path --acl public-read

Starting Upload
/style.css:
  Success
  https://s3.us-east-1.amazonaws.com/my-bucket/style.css
/index.html:
  Success
  https://s3.us-east-1.amazonaws.com/my-bucket/index.html 

```

## Installation

```
brew tap kalyan02/tools
brew install gzsync
```