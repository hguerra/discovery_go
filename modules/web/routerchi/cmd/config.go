/*
Copyright © 2023 Heitor Carneiro <heitorgcarneiro@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgPath           = "./configs"
	cfgType           = "yaml"
	cfgName           = "application"
	cfgEnvPrefix      = "APP"
	cfgProfileKey     = "profile"
	cfgDefaultProfile = "development"
)

func loadDefault() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(cfgPath)
		viper.SetConfigType(cfgType)
		viper.SetConfigName(cfgName)
	}

	viper.SetEnvPrefix(cfgEnvPrefix)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using base config file:", viper.ConfigFileUsed())
	}

	viper.SetDefault(cfgProfileKey, cfgDefaultProfile)
	err := viper.BindEnv(cfgProfileKey)
	cobra.CheckErr(err)
}

func loadProfile() {
	profile := viper.GetString(cfgProfileKey)
	viper.SetConfigName(fmt.Sprintf("%s-%s", cfgName, profile))
	err := viper.MergeInConfig()
	cobra.CheckErr(err)
	if err == nil {
		fmt.Fprintln(os.Stderr, "Using profile config file:", viper.ConfigFileUsed())
	}

	// .env should be the developers local env config.
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	err = viper.MergeInConfig()
	cobra.CheckErr(err)
	if err == nil {
		fmt.Fprintln(os.Stderr, "Using env file:", viper.ConfigFileUsed())
	}
}
