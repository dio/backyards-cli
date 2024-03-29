## backyards uninstall

Uninstall Backyards

### Synopsis

Uninstall Backyards

The command automatically removes the resources.
It can only dump the removable resources with the '--dump-resources' option.

```
backyards uninstall [flags]
```

### Examples

```
  # Default uninstall
  backyards uninstall

  # Uninstall Backyards from a non-default namespace
  backyards uninstall -n backyards-system
```

### Options

```
  -d, --dump-resources           Dump resources to stdout instead of applying them
  -h, --help                     help for uninstall
      --istio-namespace string   Namespace of Istio sidecar injector (default "istio-system")
      --release-name string      Name of the release (default "backyards")
      --uninstall-canary         Uninstall Canary feature as well
      --uninstall-demoapp        Uninstall Demo application as well
  -a, --uninstall-everything     Uninstall every component at once
      --uninstall-istio          Uninstall Istio mesh as well
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
