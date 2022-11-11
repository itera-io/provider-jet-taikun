package main

import (
    "os"
    "fmt"
    "strconv"

    "github.com/itera-io/taikungoclient"
    "github.com/itera-io/taikungoclient/client/ops_credentials"
    "github.com/itera-io/taikungoclient/models"
)

var resourceName = "BillingCredential"
var apiVersion = "1"

var taikunEmail string
var taikunPassword string
var useKeycloak = false
var taikunApiHost = "api.taikun.dev"

func init_provider() {

    taikunEmail = os.Getenv("TAIKUN_EMAIL")
    taikunPassword = os.Getenv("TAIKUN_PASSWORD")
}

func resourceTaikunBillingCredentialFind(id int32, apiClient *taikungoclient.Client) (*models.OperationCredentialsListDto, error) {
	params := ops_credentials.NewOpsCredentialsListParams().WithV(apiVersion)
	var offset int32 = 0

	for {
		response, err := apiClient.Client.OpsCredentials.OpsCredentialsList(params, apiClient)
		if err != nil {
			return nil, err
		}

		for _, billingCredential := range response.Payload.Data {
			if billingCredential.ID == id {
				return billingCredential, nil
			}
		}

		offset += int32(len(response.Payload.Data))
		if offset == response.Payload.TotalCount {
			break
		}

		params = params.WithOffset(&offset)
	}

	return nil, nil
}

func run_test_created() int {

    apiClient, _ := taikungoclient.NewClientFromCredentials(taikunEmail, taikunPassword, useKeycloak, taikunApiHost)

    id_raw, _ := strconv.Atoi(os.Getenv("RESOURCE_REF"))
    id := int32(id_raw)

    resource, err := resourceTaikunBillingCredentialFind(id, apiClient)

    if err != nil || resource == nil {
        fmt.Println("Billing credential not found !!!")
        return 1
    }

    return 0
}

func main() {
	init_provider()
        os.Exit(run_test_created())
}
