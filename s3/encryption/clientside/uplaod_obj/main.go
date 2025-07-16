package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {

	if len(os.Args) != 3 {
		exitErrorf("bucket and file name required\nUsage: %s bucket_name filename",
			os.Args[0])
	}

	bucket := os.Args[1]
	filename := os.Args[2]

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1")},
	)
	if err != nil {
		exitErrorf("no seesion is created %v", err)
	}

	objUploadencrypt(bucket, filename, sess)
	downObjEncrypt(bucket, filename, "output.txt", sess)

}

func objUploadencrypt(bucket string, filename string, sess *session.Session) {
	// generate customer key
	//compute base64 encoded key and md5
	//Base64 encoding is used to safely transmit binary data over systems that expect text, especially when those systems are limited to ASCII characters. Http only supports string
	//mds is used as checksum for key
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		exitErrorf("random key not genrated")
	}
	fmt.Printf("customer key is %v\n", key)

	keyMD5 := md5.Sum(key)
	fmt.Printf("customer key checksum %v\n", keyMD5)

	//base64 of key for http ffriendly protocol
	keyB64 := base64.StdEncoding.EncodeToString(key)
	fmt.Printf("customer key base64 %v\n", keyB64)

	//base 64 of checksum for http freindly protocol
	keyMD5B64 := base64.StdEncoding.EncodeToString(keyMD5[:])
	fmt.Printf("customer keychecksum base64 %v\n", keyMD5B64)

	// - are not valis in shell use _
	err = os.Setenv("key_value_3825", *aws.String(string(key)))
	if err != nil {
		exitErrorf(err.Error())
	}
	err = os.Setenv("md5_value_3825", *aws.String(string(keyMD5B64)))
	if err != nil {
		exitErrorf(err.Error())
	}

	file, err := os.Open(filename)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", err)
	}

	defer file.Close()

	// Setup the S3 Upload Manager. Also see the SDK doc for the Upload Manager
	// for more information on configuring part size, and concurrency.
	//
	// http://docs.aws.amazon.com/sdk-for-go/api/service/s3/s3manager/#NewUploader
	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(filename),
		Body:                 file,
		SSECustomerAlgorithm: aws.String("AES256"),
		SSECustomerKey:       aws.String(string(key)),       //aws takes care of handling encoding
		SSECustomerKeyMD5:    aws.String(string(keyMD5B64)), //md5 hash we need to encode
	})
	if err != nil {
		// Print the error and exit.
		exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)

	sec1 := os.Getenv("key_value_3825")

	sec2 := os.Getenv("md5_value_3825")

	fmt.Println(sec1)
	fmt.Println(sec2)
}

func downObjEncrypt(bucket string, item string, fileclient string, sess *session.Session) {

	downloader := s3manager.NewDownloader(sess)

	key := os.Getenv("key_value_3825")
	keyMD5B64 := os.Getenv("md5_value_3825")

	if key == "" || keyMD5B64 == "" {
		exitErrorf("Encryption key not found in environment variables.")
	}

	//create filename
	file, err := os.Create(fileclient)
	if err != nil {
		exitErrorf("Unable to create file %q, %v", fileclient, err)
	}
	defer file.Close()

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket:               aws.String(bucket),
			Key:                  aws.String(item),
			SSECustomerAlgorithm: aws.String("AES256"),
			SSECustomerKey:       aws.String(string(key)),
			SSECustomerKeyMD5:    aws.String(string(keyMD5B64)),
		})
	if err != nil {
		exitErrorf("Unable to download item %q, %v", item, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}
