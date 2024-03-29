## backyards install

Install Backyards

### Synopsis

Installs Backyards.

The command automatically applies the resources.
It can only dump the applicable resources with the '--dump-resources' option.

The command can install every component at once with the '--install-everything' option.

```
backyards install [flags]
```

### Examples

```
  # Default install.
  backyards install

  # Install Backyards into a non-default namespace.
  backyards install -n backyards-system
```

### Options

```
  -d, --dump-resources           Dump resources to stdout instead of applying them
  -h, --help                     help for install
      --install-canary           Install Canary feature as well
      --install-demoapp          Install Demo application as well
  -a, --install-everything       Install every component at once
      --install-istio            Install Istio mesh as well
      --istio-namespace string   Namespace of Istio sidecar injector (default "istio-system")
      --release-name string      Name of the release (default "backyards")
      --run-demo                 Send load to demo application and opens up dashboard
```

### Options inherited from parent commands

```
      --context string      Name of the kubeconfig context to use
  -c, --kubeconfig string   Path to the kubeconfig file to use for CLI requests
  -n, --namespace string    Namespace in which Backyards is installed [$BACKYARDS_NAMESPACE] (default "backyards-system")
  -v, --verbose             Turn on debug logging
```

### SEE ALSO

* [backyards](backyards.md)	 - Install and manage Backyards

###### Auto generated by spf13/cobra on 10-Sep-2019
