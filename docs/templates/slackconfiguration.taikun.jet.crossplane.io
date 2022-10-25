apiVersion: slackconfiguration.taikun.jet.crossplane.io/v1alpha1
kind: Configuration
metadata:
  name: SLACK_CONFIGURATION
spec:
  forProvider:
    name: "SLACK_CONFIGURATION"
    organizationIdRef:
        name: "ORGANIZATION_REF"
    channel: "CHANNEL"
    type: "TYPE"
    url: "URL"
  providerConfigRef:
    name: default