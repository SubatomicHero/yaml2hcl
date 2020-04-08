# yaml2hcl

Handy little tool that will convert YAML to HCL (or to string, if needed).

## Install

```bash
go get github.com/SubatomicHero/yaml2hcl
```

## Example

```golang
import (
    "github.com/spf13/viper"
    "github.com/SubatomicHero/yaml2hcl"
    "github.com/hashicorp/hcl/v2/hclwrite"
)

type Config struct {
    Variables map[interface{}]interface{} `yaml:"variables,omitempty"`
}

func main() {
    var yaml = []byte(`
        variables:
            some_key: some_value
            some_bool: true
            some_int: 44
            some_list:
                - a
                - b
                - c
            some_map:
                key: value
                keya: valuea  
    `)
    var conf Config
    viper.ReadConfig(bytes.NewBuffer(yaml))
    if err := viper.Unmarshal(&conf); err != nil {
        // handle error
    }
    hcl := yaml2hcl.Convert(conf.Variables) // converts to HCL (see HCL import)
    hclString := yaml2hcl.ConvertToString(conf.variables) // converts to HCL and then to string
}

```
