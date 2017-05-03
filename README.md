Dummy Artifacts Builder
===============================

This is a builder that outputs artifact file paths and claims to be any builder. It is meant to provide a quicker way to test post-processors instead of needing to wait for a build.


Use the builder as follows:
```
  "builders": [
    {
      "type":            "dummy-artifacts",
      "builder_id":      "mitchellh.vmware-esx",
      "files": [
        "packer-ubuntu/packer-ubuntu/packer-ubuntu-disk1.vmdk",
        "packer-ubuntu/packer-ubuntu/packer-ubuntu.mf",
        "packer-ubuntu/packer-ubuntu/packer-ubuntu.ovf"
      ]
    }
  ]
```

The following configuration options are available:

| Attribute        | Description                                                             | Required/Optional |
| ---------------- | ----------------------------------------------------------------------- | ----------------- |
| builder_id       | What builder to imitate (e.g mitchellh.vmware-esx)                      | required          |
| files            | List of files                                                           | required          |
