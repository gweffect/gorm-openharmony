package utils

import (
	"strings"
)

// Regular 表示一个简单的字符串替换规则
type Regular struct {
	find    string
	replace string
}

// Irregular 表示一个不规则的单复数形式
type Irregular struct {
	singular string
	plural   string
}

// RegularSlice 是 Regular 的切片
type RegularSlice []Regular

// IrregularSlice 是 Irregular 的切片
type IrregularSlice []Irregular

var pluralInflections = RegularSlice{
	{"y", "ies"},
	{"s", "s"},
	{"x", "xes"},
	{"sh", "shes"},
	{"ch", "ches"},
	{"o", "oes"},
}

var singularInflections = RegularSlice{
	{"ies", "y"},
	{"s", ""},
	{"xes", "x"},
	{"shes", "sh"},
	{"ches", "ch"},
	{"oes", "o"},
}

var irregularInflections = IrregularSlice{
	{"person", "people"},
	{"man", "men"},
	{"child", "children"},
	{"sex", "sexes"},
	{"move", "moves"},
	{"mombie", "mombies"},
}

var uncountableInflections = []string{"equipment", "information", "rice", "money", "species", "series", "fish", "sheep", "jeans", "police"}

// AddPlural 添加一个复数规则
func AddPlural(find, replace string) {
	pluralInflections = append(pluralInflections, Regular{find, replace})
}

// AddSingular 添加一个单数规则
func AddSingular(find, replace string) {
	singularInflections = append(singularInflections, Regular{find, replace})
}

// AddIrregular 添加一个不规则形式
func AddIrregular(singular, plural string) {
	irregularInflections = append(irregularInflections, Irregular{singular, plural})
}

// AddUncountable 添加不可数名词
func AddUncountable(values ...string) {
	uncountableInflections = append(uncountableInflections, values...)
}

// GetPlural 获取复数规则
func GetPlural() RegularSlice {
	plurals := make(RegularSlice, len(pluralInflections))
	copy(plurals, pluralInflections)
	return plurals
}

// GetSingular 获取单数规则
func GetSingular() RegularSlice {
	singulars := make(RegularSlice, len(singularInflections))
	copy(singulars, singularInflections)
	return singulars
}

// GetIrregular 获取不规则形式
func GetIrregular() IrregularSlice {
	irregular := make(IrregularSlice, len(irregularInflections))
	copy(irregular, irregularInflections)
	return irregular
}

// GetUncountable 获取不可数名词
func GetUncountable() []string {
	uncountables := make([]string, len(uncountableInflections))
	copy(uncountables, uncountableInflections)
	return uncountables
}

// SetPlural 设置复数规则
func SetPlural(inflections RegularSlice) {
	pluralInflections = inflections
}

// SetSingular 设置单数规则
func SetSingular(inflections RegularSlice) {
	singularInflections = inflections
}

// SetIrregular 设置不规则形式
func SetIrregular(inflections IrregularSlice) {
	irregularInflections = inflections
}

// SetUncountable 设置不可数名词
func SetUncountable(inflections []string) {
	uncountableInflections = inflections
}

// Plural 将单词转换为复数形式
func Plural(str string) string {
	// 检查是否是不可数名词
	for _, uncountable := range uncountableInflections {
		if strings.EqualFold(str, uncountable) {
			return str
		}
	}

	// 检查是否是不规则形式
	for _, irregular := range irregularInflections {
		if strings.EqualFold(str, irregular.singular) {
			return irregular.plural
		}
	}

	// 应用常规规则
	for _, rule := range pluralInflections {
		if strings.HasSuffix(strings.ToLower(str), rule.find) {
			base := str[:len(str)-len(rule.find)]
			return base + rule.replace
		}
	}

	// 如果没有匹配的规则，直接加s
	return str + "s"
}

// Singular 将单词转换为单数形式
func Singular(str string) string {
	// 检查是否是不可数名词
	for _, uncountable := range uncountableInflections {
		if strings.EqualFold(str, uncountable) {
			return str
		}
	}

	// 检查是否是不规则形式
	for _, irregular := range irregularInflections {
		if strings.EqualFold(str, irregular.plural) {
			return irregular.singular
		}
	}

	// 应用常规规则
	for _, rule := range singularInflections {
		if strings.HasSuffix(strings.ToLower(str), rule.find) {
			base := str[:len(str)-len(rule.find)]
			return base + rule.replace
		}
	}

	// 如果没有匹配的规则，直接去掉s
	if strings.HasSuffix(str, "s") {
		return str[:len(str)-1]
	}
	return str
}
