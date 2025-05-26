# Log Config

* This `log_config` provides utilities to log structured configuration data in Go.
It extracts public fields from complex configuration structures, converts them into a JSON format, and logs them using the `zap` logger.

## Key Features
* Logging of structured configurations in JSON format.
* Support for tags to manage field display in logs.
* Obfuscation of sensitive data.
* Support for nested structures for more organized output.

## Table of Contents
* [Usage](#usage)
  * [Logging Configurations](#logging-configurations)
  * [Tagging Fields](#tagging-fields)
  * [Handling Secret Fields](#handling-secret-fields)
  * [Handling Nested Structs](#handling-nested-structs)
  * [Configuration names](#Configuration-names)
* [Examples](#examples)

## Usage

### Logging Configurations

 * To log a configuration struct using the `LogConfig` function, define the struct with exported fields that represent the configuration data,
and pass the struct instance (or a pointer to it) to `LogConfig`.
 * non-exported fields will be ignored

**Example:**

```go
package main

import "path/to/your/log_config"

type AppConfig struct {
    Port    int    `logKey:"app_port"`
    Env     string `logKey:"environment"`
    Secret  string `logKey:"api_secret,secret"`
}

func main() {
    config := AppConfig{
        Port:   8080,
        Env:    "production",
        Secret: "super-secret-key",
    }

    log_config.LogConfig(config)
}
```
## Tagging Fields

Fields to be logged must include a `logKey` struct tag. The `logKey` defines:

* **Field Key**: A unique name to represent the field in the logs.
* **Optional `Secret` Attribute**: Use secret to obfuscate sensitive fields. Fields marked as secret will only indicate their presence 
(`"exists": true`) instead of logging their actual value.

### **Handling Secret Fields**

The `LogConfig` function obfuscates fields tagged as secrets. If a field is tagged with `,secret`, only a presence indicator will appear in the logs:
```go
type SecureConfig struct {
    Password string `logKey:"db_password,secret"`
}
```
In logs, a field tagged as `secret` will be shown as `"db_password_exists": "true"` instead of revealing the actual password.

### **Handling Nested Structs**
For complex configurations that include nested structs, the `nested` tag provides control over how these fields are logged. 
If a struct field has the `nested:"true"` tag, it will be logged in its entirety as a nested configuration block. 

**Example:**
```go
type ServerConfig struct {
    Port    int    `logKey:"server_port"`
    Address string `logKey:"server_address"`
}

type AppConfig struct {
    Server   ServerConfig `logKey:"server" nested:"true"` //the structure fields will be embedded in the main config
    ApiToken string       `logKey:"api_token,secret"`
}
```
**Expected log output with nested:"true":**
```go
{
  "config": {
    "server_port": "8080", 
    "server_address": "127.0.0.1",
    "api_token_exists": "true"
  }
}
```
### **Configuration-names**

Configuration names are taken from logKey tags, if specified. If the logKey tag is missing, 
the name will be automatically generated based on the camelCase structure name.

```go
type config struct {
	DBConfig DBConfig `logKey:"postgres_db"`
	RedisConfig RedisConfig //the name will be taken from the name of the structure field
}
```
**Expected:**

```JSON
{
  "postgres_db": {
    ...
  },
  "redis_config": {
    ...
  }
}
```
## Examples

```go
package main

import "path/to/your/log_config"

type (
	Config struct {
		Server  Server
		Storage storage
	}

	Server struct {
		Port    int    `logKey:"server_port"`
		Address string `logKey:"server_address"`
	}

	storage struct {
		DB    DB
		Cache Cache `nested:"true"`
	}

	DB struct {
		FirstInstance  Mongo `logKey:"first_instance"`
		SecondInstance Mongo `logKey:"second_instance"`
	}

	Mongo struct {
		Host     string `logKey:"host"`
		Port     int    `logKey:"port"`
		User     string `logKey:"user,secret"`
		Password string `logKey:"password,secret"`
		DBName   string `logKey:"db_name"`
	}

	Cache struct {
		RedisBase
		CacheKey string `logKey:"key"`
	}

	RedisBase struct {
		Hosts      string `logKey:"host"`
		MasterName string `logKey:"master_name"`
		Password   string `logKey:"password,secret"`
	}
)

func main() {
	serverConfig := Server{
		Port:    8080,
		Address: "0.0.0.0",
	}

	dbConfig := DB{
		FirstInstance: Mongo{
			Host:     "localhost",
			Port:     27015,
			User:     "root",
			Password: "root",
			DBName:   "first_instance",
		},
		SecondInstance: Mongo{
			Host:     "localhost",
			Port:     27015,
			User:     "root",
			Password: "root",
			DBName:   "second_instance",
		},
	}

	cacheConfig := Cache{
		RedisBase: RedisBase{
			Hosts:      "localhost",
			MasterName: "master",
			Password:   "root",
		},
		CacheKey: "cache_key",
	}

	cfg := Config{
		Server: serverConfig,
		Storage: storage{
			DB:    dbConfig,
			Cache: cacheConfig,
		},
	}

	log_config.LogConfig(cfg)
}
```
**Expected:**

```JSON
{
  "server": {
    "server_address": "0.0.0.0",
    "server_port": "8080"
  },
  "first_instance": {
    "db_name": "first_instance",
    "host": "localhost",
    "password_exists": "true",
    "port": "27015",
    "user_exists": "true"
  },
  "second_instance": {
    "db_name": "second_instance",
    "host": "localhost",
    "password_exists": "true",
    "port": "27015",
    "user_exists": "true"
  },
  "cache": {
    "host": "localhost",
    "key": "cache_key",
    "master_name": "master",
    "password_exists": "true"
  }
}
```