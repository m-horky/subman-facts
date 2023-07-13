package facts

import (
	"encoding/json"
	"fmt"
	"git.sr.ht/~spc/go-log"
	"github.com/mitchellh/mapstructure"
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

type dmidecodeSection struct {
	Handle      int            `json:"handle"`
	Type        int            `json:"type"`
	Length      int            `json:"length"`
	Description string         `json:"description"`
	Values      map[string]any `json:"values"`
}

type dmidecodeData struct {
	Data []dmidecodeSection `json:"data"`
}

func (d dmidecodeData) getSections(section int) ([]dmidecodeSection, error) {
	var sections []dmidecodeSection
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

type DmidecodeBaseboardFacts struct {
	ChassisHandle          string `json:"chassis_handle" mapstructure:"chassis_handle"`
	ContainedObjectHandles int    `json:"contained_object_handles" mapstructure:"contained_object_handles"`
	Manufacturer           string `json:"manufacturer" mapstructure:"manufacturer"`
	ProductName            string `json:"product_name" mapstructure:"product_name"`
	SerialNumber           string `json:"serial_number" mapstructure:"serial_number"`
	Type                   string `json:"type" mapstructure:"type"`
	Version                string `json:"version" mapstructure:"version"`
}
type DmidecodeBiosFacts struct {
	Address                    string `json:"address" mapstructure:"address"`
	Revision                   string `json:"revision" mapstructure:"bios_revision"`
	CurrentlyInstalledLanguage string `json:"currently_installed_language" mapstructure:"currently_installed_language"`
	FirmwareRevision           string `json:"firmware_revision" mapstructure:"firmware_revision"`
	LanguageDescriptionFormat  string `json:"language_description_format" mapstructure:"language_description_format"`
	ReleaseDate                string `json:"release_date" mapstructure:"release_date"`
	ROMSize                    string `json:"rom_size" mapstructure:"rom_size"`
	RuntimeSize                string `json:"runtime_size" mapstructure:"runtime_size"`
	Vendor                     string `json:"vendor" mapstructure:"vendor"`
	Version                    string `json:"version" mapstructure:"version"`
}
type DmidecodeChassisFacts struct {
	AssetTag          string `json:"asset_tag" mapstructure:"asset_tag"`
	ContainedElements int    `json:"contained_elements" mapstructure:"contained_elements"`
	Lock              string `json:"lock" mapstructure:"lock"`
	Manufacturer      string `json:"manufacturer" mapstructure:"manufacturer"`
	OEMInformation    string `json:"oem_information" mapstructure:"oem_information"`
	SerialNumber      string `json:"serial_number" mapstructure:"serial_number"`
	Type              string `json:"type" mapstructure:"type"`
	Version           string `json:"version" mapstructure:"version"`
}
type DmidecodeConnectorFacts struct {
	ExternalConnectorType       string `json:"external_connector_type" mapstructure:"external_connector_type"`
	ExternalReferenceDesignator string `json:"external_reference_designator" mapstructure:"external_reference_designator"`
	InternalConnectorType       string `json:"internal_connector_type" mapstructure:"internal_connector_type"`
	PortType                    string `json:"port_type" mapstructure:"port_type"`
}
type DmidecodeMemoryFacts struct {
	ArrayHandle                   string `json:"array_handle" mapstructure:"array_handle"`
	AssetTag                      string `json:"asset_tag" mapstructure:"asset_tag"`
	BankLocator                   string `json:"bank_locator" mapstructure:"bank_locator"`
	CacheSize                     string `json:"cache_size" mapstructure:"cache_size"`
	ConfiguredMemorySpeed         string `json:"configured_memory_speed" mapstructure:"configured_memory_speed"`
	ConfiguredVoltage             string `json:"configured_voltage" mapstructure:"configured_voltage"`
	DataWidth                     string `json:"data_width" mapstructure:"data_width"`
	ErrorCorrectionType           string `json:"error_correction_type" mapstructure:"error_correction_type"`
	ErrorInformationHandle        string `json:"error_information_handle" mapstructure:"error_information_handle"`
	FormFactor                    string `json:"form_factor" mapstructure:"form_factor"`
	Location                      string `json:"location" mapstructure:"location"`
	Locator                       string `json:"locator" mapstructure:"locator"`
	LogicalSize                   string `json:"logical_size" mapstructure:"logical_size"`
	Manufacturer                  string `json:"manufacturer" mapstructure:"manufacturer"`
	MaximumCapacity               string `json:"maximum_capacity" mapstructure:"maximum_capacity"`
	MemoryOperatingModeCapability string `json:"memory_operating_mode_capability" mapstructure:"memory_operating_mode_capability"`
	MemoryTechnology              string `json:"memory_technology" mapstructure:"memory_technology"`
	ModuleManufacturerID          string `json:"module_manufacturer_id" mapstructure:"module_manufacturer_id"`
	NonVolatileSize               string `json:"non-volatile_size" mapstructure:"non-volatile_size"`
	NumberOfDevices               string `json:"number_of_devices" mapstructure:"number_of_devices"`
	PortNumber                    string `json:"part_number" mapstructure:"part_number"`
	Rank                          string `json:"rank" mapstructure:"rank"`
	SerialNumber                  string `json:"serial_number" mapstructure:"serial_number"`
	Set                           string `json:"set" mapstructure:"set"`
	Size                          string `json:"size" mapstructure:"size"`
	Speed                         string `json:"speed" mapstructure:"speed"`
	TotalWidth                    string `json:"total_width" mapstructure:"total_width"`
	Type                          string `json:"type" mapstructure:"type"`
	TypeDetail                    string `json:"type_detail" mapstructure:"type_detail"`
	Use                           string `json:"use" mapstructure:"use"`
	VolatileSize                  string `json:"volatile_size" mapstructure:"volatile_size"`
}
type DmidecodeProcessorFacts struct {
	AssetTag          string `json:"asset_tag" mapstructure:"asset_tag"`
	CoreCount         string `json:"core_count" mapstructure:"core_count"`
	CoreEnabled       string `json:"core_enabled" mapstructure:"core_enabled"`
	CurrentSpeed      string `json:"current_speed" mapstructure:"current_speed"`
	ExternalClock     string `json:"external_clock" mapstructure:"external_clock"`
	Family            string `json:"family" mapstructure:"family"`
	ID                string `json:"id" mapstructure:"id"`
	L1CacheHandle     string `json:"l1_cache_handle" mapstructure:"l1_cache_handle"`
	L2CacheHandle     string `json:"l2_cache_handle" mapstructure:"l2_cache_handle"`
	L3CacheHandle     string `json:"l3_cache_handle" mapstructure:"l3_cache_handle"`
	Manufacturer      string `json:"manufacturer" mapstructure:"manufacturer"`
	MaxSpeed          string `json:"max_speed" mapstructure:"max_speed"`
	PartNumber        string `json:"part_number" mapstructure:"part_number"`
	SerialNumber      string `json:"serial_number" mapstructure:"serial_number"`
	Signature         string `json:"signature" mapstructure:"signature"`
	SocketDesignation string `json:"socket_designation" mapstructure:"socket_designation"`
	Status            string `json:"status" mapstructure:"status"`
	ThreadCount       string `json:"thread_count" mapstructure:"thread_count"`
	Type              string `json:"type" mapstructure:"type"`
	Upgrade           string `json:"upgrade" mapstructure:"upgrade"`
	Version           string `json:"version" mapstructure:"version"`
	Voltage           string `json:"voltage" mapstructure:"voltage"`
}
type DmidecodeSlotFacts struct {
	BusAddress      string `json:"bus_address" mapstructure:"bus_address"`
	Characteristics string `json:"characteristics" mapstructure:"characteristics"`
	CurrentUsage    string `json:"current_usage" mapstructure:"current_usage"`
	Designation     string `json:"designation" mapstructure:"designation"`
	Length          string `json:"other" mapstructure:"length"`
	Type            string `json:"type" mapstructure:"type"`
}
type DmidecodeSystemFacts struct {
	Family       string `json:"family" mapstructure:"family"`
	Manufacturer string `json:"manufacturer" mapstructure:"manufacturer"`
	ProductName  string `json:"product_name" mapstructure:"product_name"`
	SerialNumber string `json:"serial_number" mapstructure:"serial_number"`
	SKUNumber    string `json:"sku_number" mapstructure:"sku_number"`
	UUID         string `json:"uuid" mapstructure:"uuid"`
	Version      string `json:"version" mapstructure:"version"`
	WakeUpType   string `json:"wake-up_type" mapstructure:"wake-up_type"`
}

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
	data             DmidecodeFacts
	collected        bool
	getCommandOutput func(command string, args ...string) ([]string, []string, error)
}

func NewDmidecodeCollector() DmidecodeCollector {
	return DmidecodeCollector{
		data:             DmidecodeFacts{},
		collected:        false,
		getCommandOutput: getCommandOutput,
	}
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

	var dmidecodeOutput dmidecodeData
	err = json.Unmarshal([]byte(strings.Join(stdout, "")), &dmidecodeOutput)
	if err != nil {
		log.Errorf("Could not decode output of dmidecode: %s", err)
		return err
	}

	// QUIRK: Specific collectors use the latest Handle (likely the one with the highest value).
	//  This makes it compatible with subscription-manager and python-dmidecode library.
	_ = c.collectBaseboard(dmidecodeOutput)
	_ = c.collectBios(dmidecodeOutput)
	_ = c.collectChassis(dmidecodeOutput)
	_ = c.collectConnector(dmidecodeOutput)
	_ = c.collectMemory(dmidecodeOutput)
	_ = c.collectProcessor(dmidecodeOutput)
	_ = c.collectSystem(dmidecodeOutput)
	_ = c.collectSlot(dmidecodeOutput)

	c.collected = true
	return nil
}

func (c *DmidecodeCollector) collectBaseboard(data dmidecodeData) error {
	sections, _ := data.getSections(baseboardInformation)
	if len(sections) > 0 {
		section := sections[len(sections)-1]

		facts := DmidecodeBaseboardFacts{}
		err := mapstructure.Decode(section.Values, &facts)
		if err != nil {
			log.Errorf("Failed to decode DMI slot facts: %s", err)
			return err
		}
		c.data.Baseboard = facts
	}
	return nil
}

func (c *DmidecodeCollector) collectBios(data dmidecodeData) error {
	sections, _ := data.getSections(biosInformation)
	if len(sections) > 0 {
		section := sections[len(sections)-1]

		facts := DmidecodeBiosFacts{}
		err := mapstructure.Decode(section.Values, &facts)
		if err != nil {
			log.Errorf("Failed to decode DMI bios facts: %s", err)
			return err
		}
		c.data.BIOS = facts
	}

	langSections, _ := data.getSections(biosLanguageInformation)
	if len(langSections) > 0 {
		section := langSections[len(langSections)-1]

		facts := DmidecodeBiosFacts{}
		err := mapstructure.Decode(section.Values, &facts)
		if err != nil {
			log.Errorf("Failed to decode DMI bios language facts: %s", err)
			return err
		}
		c.data.BIOS.LanguageDescriptionFormat = facts.LanguageDescriptionFormat
		c.data.BIOS.CurrentlyInstalledLanguage = facts.CurrentlyInstalledLanguage

	}

	// QUIRK: Address is lowercase
	c.data.BIOS.Address = strings.ToLower(c.data.BIOS.Address)
	return nil
}

func (c *DmidecodeCollector) collectChassis(data dmidecodeData) error {
	sections, _ := data.getSections(systemEnclosureOrChassis)
	if len(sections) > 0 {
		section := sections[len(sections)-1]

		facts := DmidecodeChassisFacts{}
		err := mapstructure.Decode(section.Values, &facts)
		if err != nil {
			log.Errorf("Failed to decode DMI system facts: %s", err)
			return err
		}
		c.data.Chassis = facts
	}
	return nil
}

func (c *DmidecodeCollector) collectConnector(data dmidecodeData) error {
	sections, _ := data.getSections(portConnectorInformation)
	if len(sections) > 0 {
		section := sections[len(sections)-1]

		facts := DmidecodeConnectorFacts{}
		err := mapstructure.Decode(section.Values, &facts)
		if err != nil {
			log.Errorf("Failed to decode DMI system facts: %s", err)
			return err
		}
		c.data.Connector = facts
	}
	return nil
}

func (c *DmidecodeCollector) collectMemory(data dmidecodeData) error {
	sections, _ := data.getSections(memoryDevice)
	if len(sections) > 0 {
		section := sections[len(sections)-1]

		facts := DmidecodeMemoryFacts{}
		err := mapstructure.Decode(section.Values, &facts)
		if err != nil {
			log.Errorf("Failed to decode DMI slot facts: %s", err)
			return err
		}
		c.data.Memory = facts
	}

	physSections, _ := data.getSections(physicalMemoryArray)
	if len(physSections) > 0 {
		section := physSections[len(physSections)-1]

		facts := DmidecodeMemoryFacts{}
		err := mapstructure.Decode(section.Values, &facts)
		if err != nil {
			log.Errorf("Failed to decode DMI slot facts: %s", err)
			return err
		}
		c.data.Memory.ErrorCorrectionType = facts.ErrorCorrectionType
		c.data.Memory.ErrorInformationHandle = facts.ErrorInformationHandle
		c.data.Memory.Location = facts.Location
		c.data.Memory.MaximumCapacity = facts.MaximumCapacity
		c.data.Memory.NumberOfDevices = facts.NumberOfDevices
		c.data.Memory.Use = facts.Use
	}
	return nil
}

func (c *DmidecodeCollector) collectProcessor(data dmidecodeData) error {
	sections, _ := data.getSections(processorInformation)
	if len(sections) > 0 {
		section := sections[len(sections)-1]

		facts := DmidecodeProcessorFacts{}
		err := mapstructure.Decode(section.Values, &facts)
		if err != nil {
			log.Errorf("Failed to decode DMI system facts: %s", err)
			return err
		}
		c.data.Processor = facts
	}
	return nil
}

func (c *DmidecodeCollector) collectSlot(data dmidecodeData) error {
	sections, _ := data.getSections(systemSlots)
	if len(sections) > 0 {
		section := sections[len(sections)-1]

		facts := DmidecodeSlotFacts{}
		err := mapstructure.Decode(section.Values, &facts)
		if err != nil {
			log.Errorf("Failed to decode DMI slot facts: %s", err)
			return err
		}
		c.data.Slot = facts
	}
	return nil
}

func (c *DmidecodeCollector) collectSystem(data dmidecodeData) error {
	// NOTE The original subscription-manager implementation also included
	//  systemConfigurationOptions section. However, it does not provide anything the
	//  systemInformation would not already provide.

	sections, _ := data.getSections(systemInformation)
	if len(sections) > 0 {
		section := sections[len(sections)-1]

		facts := DmidecodeSystemFacts{}
		err := mapstructure.Decode(section.Values, &facts)
		if err != nil {
			log.Errorf("Failed to decode DMI system facts: %s", err)
			return err
		}
		c.data.System = facts
	}

	// QUIRK: UUID is uppercase
	c.data.System.UUID = strings.ToUpper(c.data.System.UUID)
	return nil
}
