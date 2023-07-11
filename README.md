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
| cpu.cpu(s)                                      |          |
| cpu.cpu\_socket(s)                              | yes      |
| cpu.thread(s)\_per\_core                        |          |
| cpu.topology\_source                            |          |
| dev_sku                                         | yes      | 3rd party (?)     |
| distribution.id                                 |          | `distribution.go` |
| distribution.name                               | yes      | `distribution.go` |
| distribution.version                            | yes      | `distribution.go` |
| distribution.version.modifier                   |          |
| distributor_version                             | yes      | 3rd party (?)     |
| dmi.baseboard.chassis\_handle                   |          | `dmidecode.go`    |
| dmi.baseboard.contained\_object\_handles        |          | `dmidecode.go`    |
| dmi.baseboard.manufacturer                      | yes      | `dmidecode.go`    |
| dmi.baseboard.product\_name                     |          | `dmidecode.go`    |
| dmi.baseboard.serial\_number                    |          | `dmidecode.go`    |
| dmi.baseboard.type                              |          | `dmidecode.go`    |
| dmi.baseboard.version                           | yes      | `dmidecode.go`    |
| dmi.bios.address                                |          | `dmidecode.go`    |
| dmi.bios.bios\_revision                         |          | `dmidecode.go`    |
| dmi.bios.currently\_installed\_language         |          | `dmidecode.go`    |
| dmi.bios.firmware\_revision                     |          | `dmidecode.go`    |
| dmi.bios.installable\_languages                 |          |
| dmi.bios.language\_description\_format          |          | `dmidecode.go`    |
| dmi.bios.release\_date                          |          | `dmidecode.go`    |
| dmi.bios.rom\_size                              |          | `dmidecode.go`    |
| dmi.bios.runtime\_size                          |          | `dmidecode.go`    |
| dmi.bios.vendor                                 | yes      | `dmidecode.go`    |
| dmi.bios.version                                | yes      | `dmidecode.go`    |
| dmi.chassis.asset\_tag                          | yes      | `dmidecode.go`    |
| dmi.chassis.contained\_elements                 |          | `dmidecode.go`    |
| dmi.chassis.lock                                |          | `dmidecode.go`    |
| dmi.chassis.manufacturer                        | yes      | `dmidecode.go`    |
| dmi.chassis.oem\_information                    |          | `dmidecode.go`    |
| dmi.chassis.serial\_number                      | yes      | `dmidecode.go`    |
| dmi.chassis.type                                |          | `dmidecode.go`    |
| dmi.chassis.version                             | yes      | `dmidecode.go`    |
| dmi.connector.external\_connector\_type         |          | `dmidecode.go`    |
| dmi.connector.external\_reference\_designator   |          | `dmidecode.go`    |
| dmi.connector.internal\_connector\_type         |          | `dmidecode.go`    |
| dmi.connector.port\_type                        |          | `dmidecode.go`    |
| dmi.memory.array\_handle                        |          | `dmidecode.go`    |
| dmi.memory.asset\_tag                           |          | `dmidecode.go`    |
| dmi.memory.bank\_locator                        |          | `dmidecode.go`    |
| dmi.memory.cache\_size                          |          | `dmidecode.go`    |
| dmi.memory.configured\_memory\_speed            |          | `dmidecode.go`    |
| dmi.memory.configured\_voltage                  |          | `dmidecode.go`    |
| dmi.memory.data\_width                          |          | `dmidecode.go`    |
| dmi.memory.error\_correction\_type              |          | `dmidecode.go`    |
| dmi.memory.error\_information\_handle           |          | `dmidecode.go`    |
| dmi.memory.form\_factor                         |          | `dmidecode.go`    |
| dmi.memory.location                             |          | `dmidecode.go`    |
| dmi.memory.locator                              |          | `dmidecode.go`    |
| dmi.memory.logical\_size                        |          | `dmidecode.go`    |
| dmi.memory.manufacturer                         |          | `dmidecode.go`    |
| dmi.memory.maximum\_capacity                    |          | `dmidecode.go`    |
| dmi.memory.memory\_operating\_mode\_capability  |          | `dmidecode.go`    |
| dmi.memory.memory\_technology                   |          | `dmidecode.go`    |
| dmi.memory.module\_manufacturer\_id             |          | `dmidecode.go`    |
| dmi.memory.non-volatile\_size                   |          | `dmidecode.go`    |
| dmi.memory.number\_of\_devices                  |          | `dmidecode.go`    |
| dmi.memory.part\_number                         |          | `dmidecode.go`    |
| dmi.memory.rank                                 |          | `dmidecode.go`    |
| dmi.memory.serial\_number                       |          | `dmidecode.go`    |
| dmi.memory.set                                  |          | `dmidecode.go`    |
| dmi.memory.size                                 |          | `dmidecode.go`    |
| dmi.memory.speed                                |          | `dmidecode.go`    |
| dmi.memory.total\_width                         |          | `dmidecode.go`    |
| dmi.memory.type                                 |          | `dmidecode.go`    |
| dmi.memory.type\_detail                         |          | `dmidecode.go`    |
| dmi.memory.use                                  |          | `dmidecode.go`    |
| dmi.memory.volatile\_size                       |          | `dmidecode.go`    |
| dmi.meta.cpu\_socket\_count                     |          |
| dmi.processor.asset\_tag                        |          | `dmidecode.go`    |
| dmi.processor.core\_count                       |          | `dmidecode.go`    |
| dmi.processor.core\_enabled                     |          | `dmidecode.go`    |
| dmi.processor.current\_speed                    |          | `dmidecode.go`    |
| dmi.processor.external\_clock                   |          | `dmidecode.go`    |
| dmi.processor.family                            |          | `dmidecode.go`    |
| dmi.processor.id                                |          | `dmidecode.go`    |
| dmi.processor.l1\_cache\_handle                 |          | `dmidecode.go`    |
| dmi.processor.l2\_cache\_handle                 |          | `dmidecode.go`    |
| dmi.processor.l3\_cache\_handle                 |          | `dmidecode.go`    |
| dmi.processor.manufacturer                      |          | `dmidecode.go`    |
| dmi.processor.max\_speed                        |          | `dmidecode.go`    |
| dmi.processor.part\_number                      |          | `dmidecode.go`    |
| dmi.processor.serial\_number                    |          | `dmidecode.go`    |
| dmi.processor.signature                         |          | `dmidecode.go`    |
| dmi.processor.socket\_designation               |          | `dmidecode.go`    |
| dmi.processor.status                            |          | `dmidecode.go`    |
| dmi.processor.thread\_count                     |          | `dmidecode.go`    |
| dmi.processor.type                              |          | `dmidecode.go`    |
| dmi.processor.upgrade                           |          | `dmidecode.go`    |
| dmi.processor.version                           |          | `dmidecode.go`    |
| dmi.processor.voltage                           |          | `dmidecode.go`    |
| dmi.slot.bus\_address                           |          | `dmidecode.go`    |
| dmi.slot.characteristics                        |          | `dmidecode.go`    |
| dmi.slot.current\_usage                         |          | `dmidecode.go`    |
| dmi.slot.designation                            |          | `dmidecode.go`    |
| dmi.slot.length                                 |          | `dmidecode.go`    |
| dmi.slot.type                                   |          | `dmidecode.go`    |
| dmi.system.family                               |          | `dmidecode.go`    |
| dmi.system.manufacturer                         | yes      | `dmidecode.go`    |
| dmi.system.product\_name                        |          | `dmidecode.go`    |
| dmi.system.serial\_number                       | yes      | `dmidecode.go`    |
| dmi.system.sku\_number                          |          | `dmidecode.go`    |
| dmi.system.uuid                                 | yes      | `dmidecode.go`    |
| dmi.system.version                              |          | `dmidecode.go`    |
| dmi.system.wake-up\_type                        |          | `dmidecode.go`    |
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
| virt.host\_type                                 | yes      | `virt.go`         |
| virt.is\_guest                                  | yes      | `virt.go`         |
| virt.uuid                                       | yes      | `virt.go`         |
