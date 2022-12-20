package main

import (
    "os"
    "fmt"
    "strconv"

    "github.com/itera-io/taikungoclient"
    "github.com/itera-io/taikungoclient/showbackclient/showback_rules"
)

var resourceName = "ShowbackRule"
var apiVersion = "1"

var taikunEmail string
var taikunPassword string
var useKeycloak = false
var taikunApiHost = "api.taikun.dev"

func init_provider() {

    taikunEmail = os.Getenv("TAIKUN_EMAIL")
    taikunPassword = os.Getenv("TAIKUN_PASSWORD")
}

func run_test_created() int {

    apiClient, _ := taikungoclient.NewClientFromCredentials(taikunEmail, taikunPassword, useKeycloak, taikunApiHost)

    id_raw, _ := strconv.Atoi(os.Getenv("RESOURCE_REF"))
    id := int32(id_raw)

    params := showback_rules.NewShowbackRulesListParams().WithV(apiVersion).WithID(&id)

    response, err := apiClient.ShowbackClient.ShowbackRules.ShowbackRulesList(params, apiClient)

    if err != nil || len(response.Payload.Data) != 1 {
        fmt.Println("Showback rule not found !!!")
        return 1
    }

    return 0
}

func main() {
	init_provider()
        os.Exit(run_test_created())
}
