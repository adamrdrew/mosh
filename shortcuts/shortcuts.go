package shortcuts

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

//Get the shortcut for a given ID
func ReverseResolve(id string) string {
	shortcuts := loadShortcutFile()

	for key, value := range shortcuts.Map {
		if value == id {
			return key
		}
	}

	return id
}

//Gets an ID for a given shortcut
//If the shortcut isn't found just return the token
func Resolve(token string) string {
	retVal := token

	shortcuts := loadShortcutFile()

	value := shortcuts.Map[retVal]

	if value == "" {
		return retVal
	}

	return value
}

func Delete(key string) {
	shortcuts := loadShortcutFile()

	delete(shortcuts.Map, key)

	shortcuts.Save()
}

func Add(key string, value string) {
	shortcuts := loadShortcutFile()

	shortcuts.Map[key] = value

	shortcuts.Save()
}

func GetAll() map[string]string {
	shortcuts := loadShortcutFile()

	return shortcuts.Map
}

func loadShortcutFile() Shortcuts {
	sc := Shortcuts{}
	sc.Load()
	return sc
}

const UNINITIALIZED = "UNINITIALIZED"

const FILE = "config/shortcuts.yaml"

type Shortcuts struct {
	Map map[string]string
}

func (c *Shortcuts) Load() {
	c.createShortcutsFileIfNotThere()
	c.loadYAML()
}

func (c *Shortcuts) Save() {
	data, err := yaml.Marshal(&c)
	if err != nil {
		panic(err)
	}

	err2 := ioutil.WriteFile(FILE, data, 0)
	if err2 != nil {
		panic(err)
	}
}

func (c *Shortcuts) loadYAML() {
	yfile, err := ioutil.ReadFile(FILE)
	if err != nil {
		panic(err)
	}

	errorUnmarshal := yaml.Unmarshal(yfile, &c)
	if errorUnmarshal != nil {
		panic(errorUnmarshal)
	}
}

func (c *Shortcuts) createShortcutsFileIfNotThere() {
	_, statErr := os.Stat(FILE)
	if !os.IsNotExist(statErr) {
		return
	}

	defaultShortcuts := Shortcuts{
		Map: map[string]string{},
	}

	yamlData, err := yaml.Marshal(&defaultShortcuts)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(FILE, yamlData, 0755)
	if err != nil {
		panic(err)
	}

}

/*
//Gets a shortcut for a given token
//If the shortcut isn't found just return the token
func Resolve(token string) string {
	retVal := token

	shortcuts := loadShortcutFile()

	value := shortcuts[retVal]

	if value == "" {
		return retVal
	}

	return value
}

func Delete(key string) {
	shortcuts := loadShortcutFile()

	delete(shortcuts, key)

	shortcuts.Save()
}

func Add(key string, value string) {
	shortcuts := loadShortcutFile()

	shortcuts[key] = value

	shortcuts.Save()
}

//Loads the shortcuts file
func loadShortcutFile() map[string]string {
	decodeFile, err := os.Open(FILE)
	if err != nil {
		panic(err)
	}
	defer decodeFile.Close()

	decoder := json.NewDecoder(decodeFile)

	shortcuts := make(map[string]string)

	decoder.Decode(&shortcuts)

	return shortcuts
}

//Saves the shortcuts file
func saveShortcutFile(shortcuts map[string]string) {
	encodeFile, err := os.OpenFile(FILE, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(encodeFile)

	if err := encoder.Encode(shortcuts); err != nil {
		panic(err)
	}

	encodeFile.Close()
}
*/
