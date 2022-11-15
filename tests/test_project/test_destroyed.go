package main

import (
    "os"
    "fmt"
    "strconv"

    "github.com/itera-io/taikungoclient"
    "github.com/itera-io/taikungoclient/client/projects"
)

var resourceName = "Project"
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

    params := projects.NewProjectsListParams().WithV(apiVersion).WithID(&id)

    response, err := apiClient.Client.Projects.ProjectsList(params, apiClient)
    if err == nil && len(response.Payload.Data) == 1 {
        fmt.Println("Project still exists !!!")
        return 1
    }

    return 0
}

func main() {
	init_provider()
        os.Exit(run_test_destroyed())
}
