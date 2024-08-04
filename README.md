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
