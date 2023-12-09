package providerconfig

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	adapterconfig "github.com/nebula-aac/kubernexus-mono/pkg/adapter-config"
	"github.com/spf13/viper"
)

const (
	FilePath = "filepath"
	FileType = "filetype"
	FileName = "filename"
)

// Viper implements the config interface Handler for a Viper configuration registry.
type Viper struct {
	instance *viper.Viper
}

func NewAdapterViper(opts ViperOptions) (adapterconfig.Handler, error) {
	v := viper.New()
	v.AddConfigPath(opts.FilePath)
	v.SetConfigType(opts.FileType)
	v.SetConfigName(opts.FileName)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err) // Debug print
		return nil, handleViperConfigError(err, v, opts)
	}

	viperInstance := &Viper{
		instance: v,
	}

	fmt.Println("Viper instance:", viperInstance) // Debug print

	return viperInstance, nil
}

func handleViperConfigError(err error, v *viper.Viper, opts ViperOptions) error {
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// Config file not found; ignore error
		// Hack until viper issue #433 is fixed
		err := v.WriteConfigAs(fmt.Sprintf("%s/%s.%s", opts.FilePath, opts.FileName, opts.FileType))
		if err != nil {
			return err
		}
		return v.WriteConfig()
	}
	// Config file was found but another error was produced
	return err
}

func (v *Viper) SetKey(key string, value string) {
	v.instance.Set(key, value)
	_ = v.instance.WriteConfig()
}

func (v *Viper) GetKey(key string) string {
	_ = v.instance.ReadInConfig()
	return v.instance.Get(key).(string)
}

func (v *Viper) GetObject(key string, result interface{}) error {
	// the actual implementation of the method
	// this is just a placeholder, replace with the actual implementation
	value := v.instance.Get(key) // assuming instance is a *viper.Viper and it has been properly initialized
	if value == nil {
		return fmt.Errorf("key %s not found", key)
	}
	err := mapstructure.Decode(value, result)
	if err != nil {
		return err
	}
	return nil
}

func (v *Viper) SetObject(key string, value interface{}) error {
	v.instance.Set(key, value)
	err := v.instance.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}
