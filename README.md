# subscription-manager like fact collection

We don't know if there's more services that ingest our facts.
In the table below, there are just `candlepin` and `swatch`.

However, we don't know if they would still be needed for SCA mode rhc.next will operate in.

We plan on asking about:

- whether they will be required for SCA,
- in which types of connectivity they are required (CDN cert for downloading content, insights, ...).


## Running

```console
$ make run
```


## The table

| fact                                            | required by             | implemented             |
|-------------------------------------------------|-------------------------|-------------------------|
| aws\_account\_id                                | swatch                  |
| aws\_billing\_products                          | swatch                  |
| aws\_instance\_id                               | swatch                  |
| aws\_marketplace\_product\_codes                | swatch                  |
| azure\_instance\_id                             | swatch                  |
| azure\_offer                                    | swatch                  |
| azure\_sku                                      | swatch                  |
| gcp\_instance\_id                               | swatch                  |
| cpu.core(s)\_per\_socket                        | candlepin, swatch (HBI) |
| cpu.cpu(s)                                      |                         |
| cpu.cpu\_socket(s)                              | candlepin, swatch (HBI) |
| cpu.thread(s)\_per\_core                        |                         |
| cpu.topology\_source                            |                         |
| distribution.id                                 |                         | `distribution.go`       |
| distribution.name                               | swatch (HBI)            | `distribution.go`       |
| distribution.version                            | swatch (HBI)            | `distribution.go`       |
| distribution.version.modifier                   |                         |
| dmi.baseboard.chassis\_handle                   |                         |
| dmi.baseboard.contained\_object\_handles        |                         |
| dmi.baseboard.manufacturer                      |                         |
| dmi.baseboard.product\_name                     |                         |
| dmi.baseboard.serial\_number                    |                         |
| dmi.baseboard.type                              |                         |
| dmi.baseboard.version                           |                         |
| dmi.bios.address                                |                         |
| dmi.bios.bios\_revision                         |                         |
| dmi.bios.currently\_installed\_language         |                         |
| dmi.bios.firmware\_revision                     |                         |
| dmi.bios.installable\_languages                 |                         |
| dmi.bios.language\_description\_format          |                         |
| dmi.bios.release\_date                          |                         |
| dmi.bios.rom\_size                              |                         |
| dmi.bios.runtime\_size                          |                         |
| dmi.bios.vendor                                 | swatch                  |
| dmi.bios.version                                | swatch                  |
| dmi.chassis.asset\_tag                          | swatch                  |
| dmi.chassis.contained\_elements                 |                         |
| dmi.chassis.lock                                |                         |
| dmi.chassis.manufacturer                        | swatch                  |
| dmi.chassis.oem\_information                    |                         |
| dmi.chassis.serial\_number                      |                         |
| dmi.chassis.type                                |                         |
| dmi.chassis.version                             |                         |
| dmi.connector.external\_connector\_type         |                         |
| dmi.connector.external\_reference\_designator   |                         |
| dmi.connector.internal\_connector\_type         |                         |
| dmi.connector.port\_type                        |                         |
| dmi.memory.array\_handle                        |                         |
| dmi.memory.asset\_tag                           |                         |
| dmi.memory.bank\_locator                        |                         |
| dmi.memory.cache\_size                          |                         |
| dmi.memory.configured\_memory\_speed            |                         |
| dmi.memory.configured\_voltage                  |                         |
| dmi.memory.data\_width                          |                         |
| dmi.memory.error\_correction\_type              |                         |
| dmi.memory.error\_information\_handle           |                         |
| dmi.memory.form\_factor                         |                         |
| dmi.memory.location                             |                         |
| dmi.memory.locator                              |                         |
| dmi.memory.logical\_size                        |                         |
| dmi.memory.manufacturer                         |                         |
| dmi.memory.maximum\_capacity                    |                         |
| dmi.memory.memory\_operating\_mode\_capability  |                         |
| dmi.memory.memory\_technology                   |                         |
| dmi.memory.module\_manufacturer\_id             |                         |
| dmi.memory.non-volatile\_size                   |                         |
| dmi.memory.number\_of\_devices                  |                         |
| dmi.memory.part\_number                         |                         |
| dmi.memory.rank                                 |                         |
| dmi.memory.serial\_number                       |                         |
| dmi.memory.set                                  |                         |
| dmi.memory.size                                 |                         |
| dmi.memory.speed                                |                         |
| dmi.memory.total\_width                         |                         |
| dmi.memory.type                                 |                         |
| dmi.memory.type\_detail                         |                         |
| dmi.memory.use                                  |                         |
| dmi.memory.volatile\_size                       |                         |
| dmi.meta.cpu\_socket\_count                     |                         |
| dmi.processor.asset\_tag                        |                         |
| dmi.processor.core\_count                       |                         |
| dmi.processor.core\_enabled                     |                         |
| dmi.processor.current\_speed                    |                         |
| dmi.processor.external\_clock                   |                         |
| dmi.processor.family                            |                         |
| dmi.processor.id                                |                         |
| dmi.processor.l1\_cache\_handle                 |                         |
| dmi.processor.l2\_cache\_handle                 |                         |
| dmi.processor.l3\_cache\_handle                 |                         |
| dmi.processor.manufacturer                      |                         |
| dmi.processor.max\_speed                        |                         |
| dmi.processor.part\_number                      |                         |
| dmi.processor.serial\_number                    |                         |
| dmi.processor.signature                         |                         |
| dmi.processor.socket\_designation               |                         |
| dmi.processor.status                            |                         |
| dmi.processor.thread\_count                     |                         |
| dmi.processor.type                              |                         |
| dmi.processor.upgrade                           |                         |
| dmi.processor.version                           |                         |
| dmi.processor.voltage                           |                         |
| dmi.slot.bus\_address                           |                         |
| dmi.slot.characteristics                        |                         |
| dmi.slot.current\_usage                         |                         |
| dmi.slot.designation                            |                         |
| dmi.slot.length                                 |                         |
| dmi.slot.type                                   |                         |
| dmi.system.family                               |                         |
| dmi.system.manufacturer                         |                         |
| dmi.system.product\_name                        |                         |
| dmi.system.serial\_number                       |                         |
| dmi.system.sku\_number                          |                         |
| dmi.system.uuid                                 | swatch (HBI)            |
| dmi.system.version                              |                         |
| dmi.system.wake-up\_type                        |                         |
| last\_boot                                      |                         |
| lscpu.address\_sizes                            |                         |
| lscpu.architecture                              |                         |
| lscpu.bios\_cpu\_family                         |                         |
| lscpu.bios\_model\_name                         |                         |
| lscpu.bios\_vendor\_id                          |                         |
| lscpu.bogomips                                  |                         |
| lscpu.byte\_order                               |                         |
| lscpu.core(s)\_per\_socket                      |                         |
| lscpu.cpu(s)                                    |                         |
| lscpu.cpu(s)\_scaling\_mhz                      |                         |
| lscpu.cpu\_family                               |                         |
| lscpu.cpu\_max\_mhz                             |                         |
| lscpu.cpu\_min\_mhz                             |                         |
| lscpu.cpu\_op-mode(s)                           |                         |
| lscpu.flags                                     |                         |
| lscpu.l1d\_cache                                |                         |
| lscpu.l1i\_cache                                |                         |
| lscpu.l2\_cache                                 |                         |
| lscpu.l3\_cache                                 |                         |
| lscpu.model                                     |                         |
| lscpu.model\_name                               |                         |
| lscpu.numa\_node(s)                             |                         |
| lscpu.numa\_node0\_cpu(s)                       |                         |
| lscpu.on-line\_cpu(s)\_list                     |                         |
| lscpu.socket(s)                                 |                         |
| lscpu.stepping                                  |                         |
| lscpu.thread(s)\_per\_core                      |                         |
| lscpu.vendor\_id                                |                         |
| lscpu.virtualization                            |                         |
| lscpu.vulnerability\_itlb\_multihit             |                         |
| lscpu.vulnerability\_l1tf                       |                         |
| lscpu.vulnerability\_mds                        |                         |
| lscpu.vulnerability\_meltdown                   |                         |
| lscpu.vulnerability\_mmio\_stale\_data          |                         |
| lscpu.vulnerability\_retbleed                   |                         |
| lscpu.vulnerability\_spec\_store\_bypass        |                         |
| lscpu.vulnerability\_spectre\_v1                |                         |
| lscpu.vulnerability\_spectre\_v2                |                         |
| lscpu.vulnerability\_srbds                      |                         |
| lscpu.vulnerability\_tsx\_async\_abort          |                         |
| memory.memtotal                                 | candlepin, swatch (HBI) |
| memory.swaptotal                                |                         |
| net.interface.$IFACE.ipv4\_address              | swatch (HBI)            | `network.go`            |
| net.interface.$IFACE.ipv4\_address\_list        | swatch (HBI)            | `network.go`            |
| net.interface.$IFACE.ipv4\_broadcast            |                         |                         |
| net.interface.$IFACE.ipv4\_broadcast\_list      |                         |                         |
| net.interface.$IFACE.ipv4\_netmask              |                         | `network.go` (disabled) |
| net.interface.$IFACE.ipv4\_netmask\_list        |                         | `network.go` (disabled) |
| net.interface.$IFACE.ipv6\_address.global       | swatch (HBI)            | `network.go`            |
| net.interface.$IFACE.ipv6\_address.global\_list | swatch (HBI)            | `network.go`            |
| net.interface.$IFACE.ipv6\_address.host         |                         | `network.go` (disabled) |
| net.interface.$IFACE.ipv6\_address.host\_list   |                         | `network.go` (disabled) |
| net.interface.$IFACE.ipv6\_address.link         | swatch (HBI)            | `network.go`            |
| net.interface.$IFACE.ipv6\_address.link\_list   | swatch (HBI)            | `network.go`            |
| net.interface.$IFACE.ipv6\_netmask.global       |                         | `network.go` (disabled) |
| net.interface.$IFACE.ipv6\_netmask.global\_list |                         | `network.go` (disabled) |
| net.interface.$IFACE.ipv6\_netmask.host         |                         | `network.go` (disabled) |
| net.interface.$IFACE.ipv6\_netmask.host\_list   |                         | `network.go` (disabled) |
| net.interface.$IFACE.ipv6\_netmask.link         |                         | `network.go` (disabled) |
| net.interface.$IFACE.ipv6\_netmask.link\_list   |                         | `network.go` (disabled) |
| net.interface.$IFACE.mac\_address               |                         |                         |
| network.fqdn                                    | swatch (HBI)            | `network.go`            |
| network.hostname                                |                         | `network.go` (disabled) |
| network.ipv4\_address                           |                         | `network.go` (disabled) |
| network.ipv6\_address                           |                         | `network.go` (disabled) |
| proc\_cpuinfo.common.address\_sizes             |                         |
| proc\_cpuinfo.common.bogomips                   |                         |
| proc\_cpuinfo.common.bugs                       |                         |
| proc\_cpuinfo.common.cache\_alignment           |                         |
| proc\_cpuinfo.common.cache\_size                |                         |
| proc\_cpuinfo.common.clflush\_size              |                         |
| proc\_cpuinfo.common.cpu\_cores                 |                         |
| proc\_cpuinfo.common.cpu\_family                |                         |
| proc\_cpuinfo.common.cpuid\_level               |                         |
| proc\_cpuinfo.common.flags                      |                         |
| proc\_cpuinfo.common.fpu                        |                         |
| proc\_cpuinfo.common.fpu\_exception             |                         |
| proc\_cpuinfo.common.microcode                  |                         |
| proc\_cpuinfo.common.model                      |                         |
| proc\_cpuinfo.common.model\_name                |                         |
| proc\_cpuinfo.common.physical\_id               |                         |
| proc\_cpuinfo.common.power\_management          |                         |
| proc\_cpuinfo.common.siblings                   |                         |
| proc\_cpuinfo.common.stepping                   |                         |
| proc\_cpuinfo.common.vendor\_id                 |                         |
| proc\_cpuinfo.common.vmx\_flags                 |                         |
| proc\_cpuinfo.common.wp                         |                         |
| proc\_stat.btime                                |                         |
| system.certificate\_version                     |                         |
| system.default\_locale                          |                         |
| uname.machine                                   | swatch (HBI)            | `uname.go`              |
| uname.nodename                                  |                         | `uname.go`              |
| uname.release                                   |                         | `uname.go`              |
| uname.sysname                                   |                         | `uname.go`              |
| uname.version                                   |                         | `uname.go`              |
| virt.host\_type                                 |                         |
| virt.is\_guest                                  | swatch (HBI)            |
| virt.uuid                                       | candlepin, swatch       |
