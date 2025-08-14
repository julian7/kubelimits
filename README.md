# Set kubernetes memory and CPU limits

Go by itself is not handling CPU and memory limits on its own very well. If you set such limits, it still defaults to use all of the VM's resources. This can cause CPU throttling and memory issues. Usually, it's easier to let the go program know about these limits, which will play nicer in its cage.

There are nice, well-made tools for this: [automaxprocs](https://github.com/uber-go/automaxprocs/) and [automemlimit](https://github.com/KimMachineGun/automemlimit/).

Then, there's a built-in method for defining the container:

```yaml
env:
- name: GOMEMLIMIT
  valueFrom:
    resourceFieldRef:
      resource: limits.memory
- name: GOMAXPROCS
  valueFrom:
    resourceFieldRef:
      resource: limits.cpu
```

Call me a reinvent guy, but hear me out:

- Adding variables to kubernetes is nice, but this should be handled by the application instead.
- These libraries are nice, but they feel like being seriously overengineered.
- Probably supporting cgroups v2 only is not going to cause any issues anymore.

## GOMAXPROCS and Golang 1.25

Go 1.25 introduced a [container-aware](https://tip.golang.org/doc/go1.25#container-aware-gomaxprocs) `GOMAXPROCS` feature, which makes this part of this library unnecessary. In fact, this built-in feature is also following the allocated quota, which was not implemented in this library.

Therefore, I recommend upgrading to Go 1.25 (or newer), especially if you'd like to leverage [Dynamic Resource Allocation](https://kubernetes.io/docs/concepts/scheduling-eviction/dynamic-resource-allocation/) which became beta in Kubernetes v1.32.
