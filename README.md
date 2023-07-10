# subscription-manager like fact collection

This repository contains fact collectors written in Go, with focus to 1:1 map to facts collected by `subscription-manager`.

<!-- 
The facts are tracked in our internal document:
https://docs.google.com/spreadsheets/d/1Y2TTLz_1sWKm6uLO4h9-Dl0z3qhZ7QFxnoKcNePnuks/edit

The table below only lists the current requirements.
What the future will require of us no one knows; if each service built their own rhc worker, our tooling could focus on other stuff.
-->


## Running

```console
$ make run
```


## The table

| fact                                            | required | implemented       |
|-------------------------------------------------|----------|-------------------|
| aws\_account\_id                                | yes      |
| aws\_billing\_products                          | yes      |
| aws\_instance\_id                               | yes      |
| aws\_marketplace\_product\_codes                | yes      |
| azure\_instance\_id                             | yes      |
| azure\_offer                                    | yes      |
| azure\_sku                                      | yes      |
| azure\_offer                                    | yes      |
| band.storage.usage                              | yes      | 3rd party (?)     |
| cpu.core(s)\_per\_socket                        | yes      |
| cpu.cpu(s)                                      | yes      |
| cpu.cpu\_socket(s)                              | yes      |
| cpu.thread(s)\_per\_core                        |          |
| cpu.topology\_source                            |          |
| dev_sku                                         | yes      | 3rd party (?)     |
| distribution.id                                 | yes      | `distribution.go` |
| distribution.name                               | yes      | `distribution.go` |
| distribution.version                            | yes      | `distribution.go` |
| distribution.version.modifier                   |          |
| distributor_version                             | yes      | 3rd party (?)     |
| dmi.baseboard.chassis\_handle                   |          |
| dmi.baseboard.contained\_object\_handles        |          |
| dmi.baseboard.manufacturer                      |          |
| dmi.baseboard.product\_name                     |          |
| dmi.baseboard.serial\_number                    |          |
| dmi.baseboard.type                              |          |
| dmi.baseboard.version                           |          |
| dmi.bios.address                                |          |
| dmi.bios.bios\_revision                         |          |
| dmi.bios.currently\_installed\_language         |          |
| dmi.bios.firmware\_revision                     |          |
| dmi.bios.installable\_languages                 |          |
| dmi.bios.language\_description\_format          |          |
| dmi.bios.release\_date                          |          |
| dmi.bios.rom\_size                              |          |
| dmi.bios.runtime\_size                          |          |
| dmi.bios.vendor                                 | yes      |
| dmi.bios.version                                | yes      |
| dmi.chassis.asset\_tag                          | yes      |
| dmi.chassis.contained\_elements                 |          |
| dmi.chassis.lock                                |          |
| dmi.chassis.manufacturer                        | yes      |
| dmi.chassis.oem\_information                    |          |
| dmi.chassis.serial\_number                      |          |
| dmi.chassis.type                                |          |
| dmi.chassis.version                             |          |
| dmi.connector.external\_connector\_type         |          |
| dmi.connector.external\_reference\_designator   |          |
| dmi.connector.internal\_connector\_type         |          |
| dmi.connector.port\_type                        |          |
| dmi.memory.array\_handle                        |          |
| dmi.memory.asset\_tag                           |          |
| dmi.memory.bank\_locator                        |          |
| dmi.memory.cache\_size                          |          |
| dmi.memory.configured\_memory\_speed            |          |
| dmi.memory.configured\_voltage                  |          |
| dmi.memory.data\_width                          |          |
| dmi.memory.error\_correction\_type              |          |
| dmi.memory.error\_information\_handle           |          |
| dmi.memory.form\_factor                         |          |
| dmi.memory.location                             |          |
| dmi.memory.locator                              |          |
| dmi.memory.logical\_size                        |          |
| dmi.memory.manufacturer                         |          |
| dmi.memory.maximum\_capacity                    |          |
| dmi.memory.memory\_operating\_mode\_capability  |          |
| dmi.memory.memory\_technology                   |          |
| dmi.memory.module\_manufacturer\_id             |          |
| dmi.memory.non-volatile\_size                   |          |
| dmi.memory.number\_of\_devices                  |          |
| dmi.memory.part\_number                         |          |
| dmi.memory.rank                                 |          |
| dmi.memory.serial\_number                       |          |
| dmi.memory.set                                  |          |
| dmi.memory.size                                 |          |
| dmi.memory.speed                                |          |
| dmi.memory.total\_width                         |          |
| dmi.memory.type                                 |          |
| dmi.memory.type\_detail                         |          |
| dmi.memory.use                                  |          |
| dmi.memory.volatile\_size                       |          |
| dmi.meta.cpu\_socket\_count                     |          |
| dmi.processor.asset\_tag                        |          |
| dmi.processor.core\_count                       |          |
| dmi.processor.core\_enabled                     |          |
| dmi.processor.current\_speed                    |          |
| dmi.processor.external\_clock                   |          |
| dmi.processor.family                            |          |
| dmi.processor.id                                |          |
| dmi.processor.l1\_cache\_handle                 |          |
| dmi.processor.l2\_cache\_handle                 |          |
| dmi.processor.l3\_cache\_handle                 |          |
| dmi.processor.manufacturer                      |          |
| dmi.processor.max\_speed                        |          |
| dmi.processor.part\_number                      |          |
| dmi.processor.serial\_number                    |          |
| dmi.processor.signature                         |          |
| dmi.processor.socket\_designation               |          |
| dmi.processor.status                            |          |
| dmi.processor.thread\_count                     |          |
| dmi.processor.type                              |          |
| dmi.processor.upgrade                           |          |
| dmi.processor.version                           |          |
| dmi.processor.voltage                           |          |
| dmi.slot.bus\_address                           |          |
| dmi.slot.characteristics                        |          |
| dmi.slot.current\_usage                         |          |
| dmi.slot.designation                            |          |
| dmi.slot.length                                 |          |
| dmi.slot.type                                   |          |
| dmi.system.family                               |          |
| dmi.system.manufacturer                         | yes      |
| dmi.system.product\_name                        |          |
| dmi.system.serial\_number                       |          |
| dmi.system.sku\_number                          |          |
| dmi.system.uuid                                 | yes      |
| dmi.system.version                              |          |
| dmi.system.wake-up\_type                        |          |
| gcp\_instance\_id                               | yes      |
| gcp\_license\_codes                             | yes      |
| gcp\_project\_id                                | yes      |
| gcp\_project\_number                            | yes      |
| insights_id                                     | yes      |
| last\_boot                                      |          |
| lscpu.address\_sizes                            |          |
| lscpu.architecture                              |          |
| lscpu.bios\_cpu\_family                         |          |
| lscpu.bios\_model\_name                         |          |
| lscpu.bios\_vendor\_id                          |          |
| lscpu.bogomips                                  |          |
| lscpu.byte\_order                               |          |
| lscpu.core(s)\_per\_socket                      |          |
| lscpu.cpu(s)                                    |          |
| lscpu.cpu(s)\_scaling\_mhz                      |          |
| lscpu.cpu\_family                               |          |
| lscpu.cpu\_max\_mhz                             |          |
| lscpu.cpu\_min\_mhz                             |          |
| lscpu.cpu\_op-mode(s)                           |          |
| lscpu.flags                                     |          |
| lscpu.l1d\_cache                                |          |
| lscpu.l1i\_cache                                |          |
| lscpu.l2\_cache                                 |          |
| lscpu.l3\_cache                                 |          |
| lscpu.model                                     |          |
| lscpu.model\_name                               |          |
| lscpu.numa\_node(s)                             |          |
| lscpu.numa\_node0\_cpu(s)                       |          |
| lscpu.on-line\_cpu(s)\_list                     |          |
| lscpu.socket(s)                                 |          |
| lscpu.stepping                                  |          |
| lscpu.thread(s)\_per\_core                      |          |
| lscpu.vendor\_id                                |          |
| lscpu.virtualization                            |          |
| lscpu.vulnerability\_itlb\_multihit             |          |
| lscpu.vulnerability\_l1tf                       |          |
| lscpu.vulnerability\_mds                        |          |
| lscpu.vulnerability\_meltdown                   |          |
| lscpu.vulnerability\_mmio\_stale\_data          |          |
| lscpu.vulnerability\_retbleed                   |          |
| lscpu.vulnerability\_spec\_store\_bypass        |          |
| lscpu.vulnerability\_spectre\_v1                |          |
| lscpu.vulnerability\_spectre\_v2                |          |
| lscpu.vulnerability\_srbds                      |          |
| lscpu.vulnerability\_tsx\_async\_abort          |          |
| memory.memtotal                                 | yes      | `memory.go`       |
| memory.swaptotal                                |          | `memory.go`       |
| net.interface.$IFACE.ipv4\_address              | yes      | `network.go`      |
| net.interface.$IFACE.ipv4\_address\_list        | yes      | `network.go`      |
| net.interface.$IFACE.ipv4\_broadcast            |          |                   |
| net.interface.$IFACE.ipv4\_broadcast\_list      |          |                   |
| net.interface.$IFACE.ipv4\_netmask              |          |                   |
| net.interface.$IFACE.ipv4\_netmask\_list        |          |                   |
| net.interface.$IFACE.ipv6\_address.global       | yes      | `network.go`      |
| net.interface.$IFACE.ipv6\_address.global\_list | yes      | `network.go`      |
| net.interface.$IFACE.ipv6\_address.host         |          |                   |
| net.interface.$IFACE.ipv6\_address.host\_list   |          |                   |
| net.interface.$IFACE.ipv6\_address.link         | yes      | `network.go`      |
| net.interface.$IFACE.ipv6\_address.link\_list   | yes      | `network.go`      |
| net.interface.$IFACE.ipv6\_netmask.global       |          |                   |
| net.interface.$IFACE.ipv6\_netmask.global\_list |          |                   |
| net.interface.$IFACE.ipv6\_netmask.host         |          |                   |
| net.interface.$IFACE.ipv6\_netmask.host\_list   |          |                   |
| net.interface.$IFACE.ipv6\_netmask.link         |          |                   |
| net.interface.$IFACE.ipv6\_netmask.link\_list   |          |                   |
| net.interface.$IFACE.mac\_address               | yes      | `network.go`      |
| network.fqdn                                    | yes      | `network.go`      |
| network.hostname                                | yes      | `network.go`      |
| network.ipv4\_address                           |          |                   |
| network.ipv6\_address                           |          |                   |
| ocm.units                                       | yes      | 3rd party (?)     |
| proc\_cpuinfo.common.address\_sizes             |          |
| proc\_cpuinfo.common.bogomips                   |          |
| proc\_cpuinfo.common.bugs                       |          |
| proc\_cpuinfo.common.cache\_alignment           |          |
| proc\_cpuinfo.common.cache\_size                |          |
| proc\_cpuinfo.common.clflush\_size              |          |
| proc\_cpuinfo.common.cpu\_cores                 |          |
| proc\_cpuinfo.common.cpu\_family                |          |
| proc\_cpuinfo.common.cpuid\_level               |          |
| proc\_cpuinfo.common.flags                      |          |
| proc\_cpuinfo.common.fpu                        |          |
| proc\_cpuinfo.common.fpu\_exception             |          |
| proc\_cpuinfo.common.microcode                  |          |
| proc\_cpuinfo.common.model                      |          |
| proc\_cpuinfo.common.model\_name                |          |
| proc\_cpuinfo.common.physical\_id               |          |
| proc\_cpuinfo.common.power\_management          |          |
| proc\_cpuinfo.common.siblings                   |          |
| proc\_cpuinfo.common.stepping                   |          |
| proc\_cpuinfo.common.vendor\_id                 |          |
| proc\_cpuinfo.common.vmx\_flags                 |          |
| proc\_cpuinfo.common.wp                         |          |
| proc\_stat.btime                                |          |
| supported_architectures                         | yes      | 3rd party (?)     |
| system.certificate\_version                     | yes      | `system.go`       |
| system.default\_locale                          |          |
| uname.machine                                   | yes      | `uname.go`        |
| uname.nodename                                  | yes      | `uname.go`        |
| uname.release                                   | yes      | `uname.go`        |
| uname.sysname                                   | yes      | `uname.go`        |
| uname.version                                   |          | `uname.go`        |
| virt.host\_type                                 |          | `virt.go`         |
| virt.is\_guest                                  | yes      | `virt.go`         |
| virt.uuid                                       | yes      | `virt.go`         |
