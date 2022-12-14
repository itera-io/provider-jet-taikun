apiVersion: cloudcredentialopenstack.taikun.jet.crossplane.io/v1alpha1
kind: CredentialOpenstack
metadata:
  name: OPENSTACK
spec:
  forProvider:
    name: "OPENSTACK"
    user: "USER"
    passwordSecretRef:
      key: "password"
      name: "openstack-credentials"
      namespace: "crossplane-system"
    organizationIdRef:
        name: "ORGANIZATION_REF"
    url: "URL"
    domain: "DOMAIN"
    projectName: "PROJECT"
    region: "REGION"
    publicNetworkName: "PUBLIC_NETWORK"
    lock: LOCK
    availabilityZone: "AVAILABILITY"