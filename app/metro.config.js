const { withTamagui } = require("@tamagui/metro-plugin");
const { getDefaultConfig } = require("expo/metro-config");

module.exports = withTamagui(getDefaultConfig(__dirname));
