package facts

import (
	"encoding/json"
	"fmt"
	"git.sr.ht/~spc/go-log"
	"strings"
)

const (
	biosInformation                    = 0
	systemInformation                  = 1
	baseboardInformation               = 2
	systemEnclosureOrChassis           = 3
	processorInformation               = 4
	memoryControllerInformation        = 5
	memoryModuleInformation            = 6
	cacheInformation                   = 7
	portConnectorInformation           = 8
	systemSlots                        = 9
	onBoardDevicesInformation          = 10
	oemStrings                         = 11
	systemConfigurationOptions         = 12
	biosLanguageInformation            = 13
	groupAssociations                  = 14
	systemEventLog                     = 15
	physicalMemoryArray                = 16
	memoryDevice                       = 17
	thirtyTwoBitMemoryErrorInformation = 18
	memoryArrayMappedDevices           = 19
	memoryDeviceMappedAddress          = 20
	builtInPointingDevice              = 21
	portableBattery                    = 22
	systemReset                        = 23
	hardwareSecurity                   = 24
	systemPowerControls                = 25
	voltageProbe                       = 26
	coolingDevice                      = 27
	temperatureProbe                   = 28
	electricalCurrentProbe             = 29
	outOfBandRemoteAccess              = 30
	bootIntegrityServicesEntryPoint    = 31
	systemBootInformation              = 32
	sixtyFourBitMemoryErrorInformation = 33
	managementDevice                   = 34
	managementDeviceComponent          = 35
	managementDeviceTresholdData       = 36
	memoryChannel                      = 37
	ipmiDeviceInformation              = 38
	systemPowerSupply                  = 39
	additionalInformation              = 40
	onboardDevicesExtendedInformation  = 41
	managementControllerHostInterface  = 42
	tpmDevice                          = 43
	processorAdditionalInformation     = 44
)

type DmidecodeSection struct {
	Handle      int            `json:"handle"`
	Type        int            `json:"type"`
	Length      int            `json:"length"`
	Description string         `json:"description"`
	Values      map[string]any `json:"values"`
}

type DmidecodeData struct {
	Data []DmidecodeSection `json:"data"`
}

func (d DmidecodeData) getSections(section int) ([]DmidecodeSection, error) {
	var sections []DmidecodeSection
	for _, dmiSection := range d.Data {
		if dmiSection.Type == section {
			sections = append(sections, dmiSection)
		}
	}
	if len(sections) == 0 {
		return sections, fmt.Errorf("no section of type %d found", section)
	}
	return sections, nil
}

type DmidecodeBiosFacts struct {
	Address                    string `json:"address"`
	Revision                   string `json:"revision"`
	CurrentlyInstalledLanguage string `json:"currently_installed_language"`
	FirmwareRevision           string `json:"firmware_revision"`
	LanguageDescriptionFormat  string `json:"language_description_format"`
	ReleaseDate                string `json:"release_date"`
	ROMSize                    string `json:"rom_size"`
	RuntimeSize                string `json:"runtime_size"`
	Vendor                     string `json:"vendor"`
	Version                    string `json:"version"`
}
type DmidecodeProcessorFacts struct{}
type DmidecodeBaseboardFacts struct{}
type DmidecodeChassisFacts struct{}
type DmidecodeSlotFacts struct{}
type DmidecodeSystemFacts struct{}
type DmidecodeMemoryFacts struct{}
type DmidecodeConnectorFacts struct{}

type DmidecodeFacts struct {
	BIOS      DmidecodeBiosFacts      `json:"bios"`
	Processor DmidecodeProcessorFacts `json:"processor"`
	Baseboard DmidecodeBaseboardFacts `json:"baseboard"`
	Chassis   DmidecodeChassisFacts   `json:"chassis"`
	Slot      DmidecodeSlotFacts      `json:"slot"`
	System    DmidecodeSystemFacts    `json:"system"`
	Memory    DmidecodeMemoryFacts    `json:"memory"`
	Connector DmidecodeConnectorFacts `json:"connector"`
}

type DmidecodeCollector struct {
	data      DmidecodeFacts
	collected bool
}

func (c *DmidecodeCollector) GetData(rescan bool) (DmidecodeFacts, error) {
	if rescan || !c.collected {
		c.data = DmidecodeFacts{}
	}

	err := c.collect()
	if err != nil {
		return DmidecodeFacts{}, err
	}
	return c.data, nil
}

func (c *DmidecodeCollector) collect() error {
	// NOTE this requires a patched dmidecode capable of JSON output:
	//  https://github.com/jirihnidek/dmidecode/tree/json-c
	stdout, stderr, err := getCommandOutput("/home/mhorky/.local/bin/dmidecode", "--json")

	if err != nil {
		log.Errorf("Could not collect DMI facts: %s (%s)", err, strings.Join(stderr, " \\n "))
		return err
	}

	var dmidecodeOutput DmidecodeData
	err = json.Unmarshal([]byte(strings.Join(stdout, "")), &dmidecodeOutput)
	if err != nil {
		log.Errorf("Could not decode output of dmidecode: %s", err)
		return err
	}

	// QUIRK: Specific collectors use the latest handle (likely the one with the highest value).
	//  This makes it compatible with subscription-manager and python-dmidecode library.
	fmt.Print("\u001B[2m")
	_ = c.collectBios(dmidecodeOutput)
	fmt.Print("\u001B[0m")

	c.collected = true
	return nil

}

func (c *DmidecodeCollector) collectBios(data DmidecodeData) error {
	biosSections, _ := data.getSections(biosInformation)
	if len(biosSections) > 0 {
		section := biosSections[len(biosSections)-1]
		for key, rawValue := range section.Values {
			value := fmt.Sprintf("%s", rawValue)

			switch key {
			case "address":
				// QUIRK: Addresses (0x) are reported as lowercase
				c.data.BIOS.Address = strings.ToLower(value)
			case "bios_revision":
				c.data.BIOS.Revision = value
			case "firmware_revision":
				c.data.BIOS.FirmwareRevision = value
			case "release_date":
				c.data.BIOS.ReleaseDate = value
			case "rom_size":
				c.data.BIOS.ROMSize = value
			case "runtime_size":
				c.data.BIOS.RuntimeSize = value
			case "vendor":
				c.data.BIOS.Vendor = value
			case "version":
				c.data.BIOS.Version = value
			case "characteristics":
				continue
			default:
				log.Debugf("Ignoring dmidecode BIOS key: %s = %s", key, value)
			}
		}
	}

	biosLanguageSections, _ := data.getSections(biosLanguageInformation)
	if len(biosLanguageSections) > 0 {
		section := biosLanguageSections[len(biosLanguageSections)-1]
		for key, rawValue := range section.Values {
			value := fmt.Sprintf("%s", rawValue)

			switch key {
			case "language_description_format":
				c.data.BIOS.LanguageDescriptionFormat = value
			case "currently_installed_language":
				c.data.BIOS.CurrentlyInstalledLanguage = value
			default:
				log.Debugf("Ignoring dmidecode BIOS key: %s = %s", key, value)
			}
		}
	}
	return nil
}
