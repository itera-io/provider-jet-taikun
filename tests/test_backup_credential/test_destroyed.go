package main

import (
    "os"
    "fmt"
    "strconv"

    "github.com/itera-io/taikungoclient"
    "github.com/itera-io/taikungoclient/client/s3_credentials"
)

var resourceName = "BackupCredential"
var apiVersion = "1"

var taikunEmail string
var taikunPassword string
var useKeycloak = false
var taikunApiHost = "api.taikun.dev"

func init_provider() {

    taikunEmail = os.Getenv("TAIKUN_EMAIL")
    taikunPassword = os.Getenv("TAIKUN_PASSWORD")
}

func run_test_destroyed() int {

    apiClient, _ := taikungoclient.NewClientFromCredentials(taikunEmail, taikunPassword, useKeycloak, taikunApiHost)

    id_raw, _ := strconv.Atoi(os.Getenv("RESOURCE_REF"))
    id := int32(id_raw)

    params := s3_credentials.NewS3CredentialsListParams().WithV(apiVersion).WithID(&id)

    response, err := apiClient.Client.S3Credentials.S3CredentialsList(params, apiClient)
    if err == nil && response.Payload.TotalCount == 1 {
        fmt.Println("Backup credential still exists !!!")
        return 1
    }

    return 0
}

func main() {
	init_provider()
        os.Exit(run_test_destroyed())
}
