#  K8s Deployment using K8s Helm 3

### Validate the Helm Chart (Pre-Validate prior to K8s Deployment)

Go to the root-level directory where the chart directory is (`inside k8s-helm directort') which is
the `go-gorilla-restsvc` directory and issue the following validation helm command:

```
helm template <release-name> <helm-chart-directory>
```
Actual command is:
```
helm template go-gorilla-restsvc go-gorilla-restsvc
``

If the validation is successful the fully rendered (locally rendered) helm templates merged with the values.yaml will parameterize the helm chart to show all the values are plugged in.

### Execute the service (service active in K8s)
```
curl localhost/engines
```
