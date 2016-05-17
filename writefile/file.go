package writefile

import (
	"crypto/sha256"
	"encoding/hex"
	// "github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"os"
	"path"
)

func fileResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"contents": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
		Create: fileCreateUpdate,
		Read:   fileRead,
		Delete: fileDelete,
		Exists: fileExists,
	}
}

func hash(s string) string {
	sha := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sha[:])
}

func fileCreateUpdate(d *schema.ResourceData, meta interface{}) error {

	filepath := d.Get("target").(string)
	contents := d.Get("contents").(string)

	if err := os.MkdirAll(path.Dir(filepath), 0755); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath, []byte(contents), 0666); err != nil {
		return err
	}

	// d.SetId(hashcode.String(contents))
	d.SetId(hash(contents))

	return nil
}

func fileRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func fileExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	filepath := d.Get("target").(string)

	var out []byte
	var err error
	if out, err = ioutil.ReadFile(filepath); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	if string(out) == d.Get("contents").(string) {
		d.SetId(hash(string(out)))
		return true, nil
	} else {
		return false, nil
	}
}

func fileDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
