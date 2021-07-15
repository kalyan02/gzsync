package cmd

import (
	"bufio"
	"compress/gzip"
	"context"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"

	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync <source> <s3 bucket>",
	Short: "Sync directories and s3 prefixes to AWS S3",
	Long: `Command tries to emulate the use of aws s3 cli, however
aims to only support sync of local folder or files to remote s3
and not vice versa for the moment.
`,
	RunE: SyncCMDRun,
}

func SyncCMDRun(_ *cobra.Command, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("exactly 2 parameters are required [source] [destination]")
	}

	srcArg := args[0]
	_, err := os.Stat(srcArg)
	if err != nil {
		return fmt.Errorf("file path is invalid or does not exist: %v", srcArg)
	}

	destArg := args[1]
	u, err := url.Parse(destArg)
	if err != nil || (u.Scheme != "s3" && u.Scheme != "S3") {
		return fmt.Errorf("destination must be of the format s3://<bucket-name/<path>")
	}
	destBucket := u.Host

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return fmt.Errorf("could not find aws config at ~/.aws/config or environment variables: %v", err)
	}

	s3client := s3.NewFromConfig(cfg)

	s3uploader := manager.NewUploader(s3client)

	var sources []string

	srcPathAbs, _ := filepath.Abs(srcArg)

	err = filepath.WalkDir(srcPathAbs, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			sources = append(sources, path)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("could not traverse directory: %v", err)
	}

	fmt.Println("Starting Upload")
	for _, source := range sources {

		sourceTrimmed := strings.TrimPrefix(source, srcPathAbs)
		//destPath := path.Join(destArg, sourceTrimmed)
		fmt.Println(sourceTrimmed)

		srcContent, _ := ioutil.ReadFile(source)

		dstFile, _ := ioutil.TempFile("", "gzsync-*")
		dstFileWriter := gzip.NewWriter(dstFile)
		_, _ = dstFileWriter.Write(srcContent)
		dstFileWriter.Close()
		defer dstFile.Close()

		s3Input := &s3.PutObjectInput{
			Bucket:          aws.String(destBucket),
			Key:             aws.String(sourceTrimmed),
			Body:            bufio.NewReader(dstFile),
			ContentEncoding: aws.String("gzip"),
		}
		if acl != "" {
			s3Input.ACL = types.ObjectCannedACL(acl)
		}
		upResult, err := s3uploader.Upload(context.TODO(), s3Input)
		if err != nil {
			fmt.Printf("%s:\n\tFailed\n\t%v\n", sourceTrimmed, err)
		} else {
			fmt.Printf("%s:\n\tSuccess\n\t%s\n", sourceTrimmed, upResult.Location)
		}

	}

	return nil
}

var (
	acl             string
	gzipCompression bool
)

func init() {
	s3Cmd.AddCommand(syncCmd)
	syncCmd.Flags().StringVar(&acl, "acl", "", "Default ACL to be applied")
	syncCmd.Flags().BoolVar(&gzipCompression, "gzip", true, "Enable gzip compression")
}
