package main

import (
    "os"
    "fmt"

    "github.com/itera-io/taikungoclient"
    "github.com/itera-io/taikungoclient/client/users"
)

var resourceName = "User"
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

    id_raw := os.Getenv("RESOURCE_REF")

    params := users.NewUsersListParams().WithV(apiVersion).WithID(&id_raw)

    response, err := apiClient.Client.Users.UsersList(params, apiClient)

    if err != nil || len(response.Payload.Data) != 1 {
        fmt.Println("User not found !!!")
        return 1
    }

    return 0
}

func main() {
	init_provider()
        os.Exit(run_test_created())
}
