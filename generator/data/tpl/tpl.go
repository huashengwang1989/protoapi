// Code generated by "esc -o generator/data/tpl/tpl.go -modtime 0 -pkg=tpl generator/template"; DO NOT EDIT.

package tpl

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/generator/template/echo_enum.gogo": {
		local:   "generator/template/echo_enum.gogo",
		size:    352,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/3yOsWrDQBBEa+1XDCaFBLHcO7iKE0hjB2LSGBcbaSNErNVxOhXm2H8PkmKBmnSPuZnb
t9nguS0Flah4DlLi6wbn29Cyq5+wP+JwPOFl/3bKiRwXP1wJYszfJzQjCjc3RgduxAy1BqKi1S4gJQCI
cQ3PWgny11quZQezv4f7JsaHO+6G+JOvvcytNURLM8qIvnstkBaD8bzO8BF8rVWaoRsBkRLlRjpsd2jY
nefqZSrE8eN/3RZ+W6xmXj0uxpNZYkSJl9B7xXj5PCheyOg3AAD///CkU9tgAQAA
`,
	},

	"/generator/template/echo_service.gogo": {
		local:   "generator/template/echo_service.gogo",
		size:    1635,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/6RU30/bSBB+9v4Vc9bpZKPgcHdvQbwAKaUVBJG8w8aeJFucXXd3TIus/d+rWTsmgYBQ
6ydr5pv55se3MxzCmSkQlqjRSsIC5k9QWUNGVuoYzidwPZnB+PxylglRyfxBLhGaJrtpf70XQq0rYwkS
EcUaabgiqmIhonipaFXPs9ysh6WcO5L5wxDzlYlFKsRwyFmu5Rq9B+WAVghKE9qFzBFyo0kq7UCWZXCx
wZqyROsEPVW4HdxHNSJqmkOwUi8RsiuklSkceC8AgANmikr0PjlomuxSVzXNnir0PoVgmdTUm5pGLUAj
ZGNrjWUbxLH3A2Bgb2Mc6sL7tGM4BNQFE3rRVhmgl1ygliU4snVOXGYYOQDX/v9/0H7335zRozg3Bcb3
IrpC53jWjqzSyx3IunUx6hxJqtI9D6HxG1TRuuL7vppbdJXRDvcUtHHtJHpmtJ2bKUNP7DvY7a6DIhtb
0v3LWNQ6h7t+f3efpS5KtImzj9A0f3fmFFgqWef8xDGNiCxSbTVwiiRvEWdGE/6kFBK0FgJ7ytBIaRid
gMYfyYt1C3YuGAonkGenSheJ0ulxsPx1AlqVIUHEXXOOl3MLznYOI/hnZwoNb3YE/x4dHQ2gW+GIE7eq
SVI/4GDfpg/N5NmX6eQ64WeTTUlS7U5lcYvfa3Q0AK4hFRwhosjUtFeZwNIMrXeC5Kqdfcy2RK80p+GV
7E/QD8X8+RA6VY7adB9pefJ1u9W2ztAJt/1uCRvjCExNg82kPkLEz3TrzfJJusWlcoR25zTVDgsgA3Ol
C7CmJj5CQcWv4AnCQVDlOF+ZAbSS7hXd9GfirRPVL2eK9lHl2K7nZjKdxRsQZhfjWRKHE0wr7+PBG48p
fXWV3iW5GG9zMOXvkby+hL8CAAD//1P1wApjBgAA
`,
	},

	"/generator/template/echo_struct.gogo": {
		local:   "generator/template/echo_struct.gogo",
		size:    235,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/1zNsa7CIBTG8Z2n+NL9wn7v2F4Tl9ahD1Asx6baAgIdCOHdDakx6vbPj/AdIVAbRZhI
k5OBFM4R1plgpJ3/0HRoux7/zbHnjFk53uRESImf9syZMSEK1Iv0vpVroRAtfRl8cNsYkBgApPQDJ/VE
4IeZFuWR8/OB93NYyoeS0ZYart7o3yolvm9VuBi3fsB9IxffZXjdIa3KemaPAAAA//9AUK296wAAAA==
`,
	},

	"/generator/template/php.gophp": {
		local:   "generator/template/php.gophp",
		size:    2425,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5xVTW/jNhC961cMBAOWAlvrbQ9dxI2LNEm3OSQput6TEwi0NLG5kEmFHCbrCPrvBfVl
SZa3RXkwTHjmzeObmedff0u3qfPh7MyBM1huuQaugcEzTxA2KFAxwhjWe0iVJMlSDt4rKs2lCIx5N0Ek
dx/qn3yL8blJYnQOWQbBku8Q8tz+eP0A9w9LuLm+XQYOnH1wHMF2qFMWIWRZcM92+MVe8nzuOEYj3HyP
MCUuxby4fjbv7wn+SZQ+XiUcBQHTYK/lbe5kmWJigxDcodZsgzrPnShhWkOWEacEoSiS507mAABk2RSq
jD84JrEuiAJAqvgrI4RRRcsyKhNQxHWMWSc8gmcjIksRXlnCY0boMaXYHkYKdSqFRr/MLD5/WNQe/gwe
1xrJa/JXbkPCffL9FlKJZnO4flh/w4gguGbElvvUSj6iLdfTRZMNFyDwrS/FqULzLANMNA7iDOcUKY0+
9uQFRI8ybZV8K6g0/fXccYMztjMoJAF+55pcf35A62h4qJRXvRluZa9NGilsanlZNqhenvdUqjt/GIh+
W4dU6k1P7gwy2nQY9XEVklHiCH7eefnpqSQZFvN4Cnf13+ayDDihVWsG4GJxRHW6OJCY9BDLCfsXgEn7
gcftt+ep1iN36uB68bsrb90KI2tQo23LOQa1C8NICk3KROSN1kzjV8XhAsYff/olmAWz4OP5p9mn2fjE
JBzgq8U7OJXXUWHVudkztsVCo/i4kKMqPTmOI75DaagI+3nWDXhqbn53/Cpva54ZsSS5TLk1gpcJjHZI
WxlPYGQUn8BIGkoNWQ2P3hkzYnDR4z/+pqUoeSt8Ge7902GnGyex+9IXbrpQ+GJQk9dnZUv7bRT2diUF
lWI3mNPFBul3Ge89v/hahWjPrzpe5EZNoqUexhjJGL0W5ARIGWyVa7l0lbxyUSmpBgx6yO2OslburvzH
ssY74HdDBetHDtS0ApiknrtWC70Weitwumj+vAbxu0m1I5W58/9h9ldMWIN/5iKGpv9rGe/PwYWg3c0j
NXpef1eMhT5hfm2fD25Faqg0rGI0f2y03aWwBmXX947iPHfL+9e/b4vvYXh/eXfz5a/Lq5swhADcx8cs
Cx4KzSt/bBaw8qbc+ScAAP//SLLRKXkJAAA=
`,
	},

	"/generator/template/spring_service.gojava": {
		local:   "generator/template/spring_service.gojava",
		size:    854,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/6yRv27CQAzG9zyFlSkMPR4gC6IgxABBVdTdSUw4Ab7rnSOEonv3KuFPGNpKRWSK7vNn
fz97PIZ3UxHUxORQqILiDNYZMWh1CrMM1lkO89kyV1FksdxjTdC2anP5DSGNIn20xgkYVytvneZ66/BI
J+P26kSFKjRXCpmNoGjDakGyQms11+l/rRvjn/Z+kLeGPU1NdX7C/NWQl4s3sk1x0CVg4cVhKVAe0Ptu
K2s8UghT9ARtBADQtm/gkGsCtSLZmcpDCHdFb4EJVGb7GZ94aAjixTyPb0WTB+Ak7tcuuxDi0UV9ROpf
rsHaVmWN2Ebys6UQhmRdu2TyANNJSx4qNY+uybvPkTSOB3uieZT26sBAXHVp/2YagIbjv4JnQZK8lOB+
0l9H/jAvvba6dQrRdwAAAP//oNlDdVYDAAA=
`,
	},

	"/generator/template/spring_struct.gojava": {
		local:   "generator/template/spring_struct.gojava",
		size:    585,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5SPwWoCMRCG73mKOdqD8QGWQkFbqBT1sC8w7o7baDYJk1mphLx7ibqtUmhpTjOB+f7v
n81g7luCjhwxCrWwPUFgLx6DqWCxhtW6hufFa62VCtgcsCNISW8uY86VUqYPngUa3+sdRiH+6K3eY3OI
3ml0zguK8U4vo3dzJhTP1b+ONuwDsZy+s/Z4RD2IsfrNRKmUCsPWmgYaizEWv3kZVthTzpAUAEBKU2B0
HYF+MWTbmPP5P7A5ohDsjENbTpd4xPoUzpcJdIFA6TlCyLWQszrvTzedLrSLx73BpKzeCQ+NbJCxz/nh
D6vy5N1EfaMAj78KlfWqldIdE6Zj1y+7m5IdScHWRmzhTka18phkYPczNl9jxuysPgMAAP//ioLHmkkC
AAA=
`,
	},

	"/generator/template/ts/data.gots": {
		local:   "generator/template/ts/data.gots",
		size:    584,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5yOQWsaQRSA7/MrHl6Epeq90IPFLRSkSJFeSinj7lMH1tntzGypDAOFGBIhBiGag4fk
FBAE9RRIzM9xd/0ZYdVFc8jFOc3wvvneV7IsYkG9zSQ0mYfAJLSQo6AKXWh0IR8IX/k0YPk3mONzRRmX
QD0PVBvBpYqCVCJ0VCgQGsh4C0KJLjC+BQ5WJUGi+MsclMSCwimHWLBZPMS3F+vVYzK6jy+H6+frrJRY
sJtEV+fRcFaufY36g2S22CzPktE07v+PXsbJaJpMevF4GQ/myeomvutF88n6qU+sEtG6AILyFkLR5mFH
gjEE/wW+UIA87IDWxW+0g8aAJgAAWmf4F4aeu+VhN8jAT+n9B/VCNOZD9gm5m6Jmu3D/OF5eoYrWuwEe
BzCuUDSpgydUfEwZJVPlQW6M1qwJ+AeKVdpAD3LV8me7+vu7XbPLdbuSM+bnL62Ru3vfO+GvAQAA//9s
n4weSAIAAA==
`,
	},

	"/generator/template/ts/helper.gots": {
		local:   "generator/template/ts/helper.gots",
		size:    2361,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xWbW/bNhD+7l9xE7ZKil0p2zqgUKamL+u2DkPTxcknwx8Y62SzpkmNopxobf77cKQk
y7a8JQhg6+655+G98OT47GwE9A+vC6bZBrZMAIPSaC6X1qzRVFqWwCSgXKgMs8YLZsUMLJiEOwRWFCgz
MApYwaHSYgRn8QgfCqUN5JVcGK5agmDLRNKQhO0X+DICAHBqDfD2+sM7tSmURGkoKIwsxuEKwRYYxN+9
OI+XfAL+a3/Q/eMb506G3T+8iJcT8L894X3ngicn3Oc2eDzs/emtC56dcP/i3HM/vBg9jkauE/C24iID
BrfXf8Jd3dSV6mPbU1KFzQoBZdZrWqUFJNZ+x0q0j7nS1lCi3vIF9sANkcNr/LvC0oC6+4wLEwFGy8jG
/o5CqOvGe68qkVGbKcRBQeVg6gL3gP2BSaCbHMgrIeBWi7bXz63Eb+9vKMU11vGWiQqhYFyXlgMf2KYQ
mNADpUSxKXgrY4okjoVaMLFSpUlenr889wjE9BJS+CLZBhPw7lEu75F7E5B8sXYGw5n0Hgnbnit9BYOE
l8SSdiTPiCPtKAYHe4kSNTP4V4W6vtXi55tXQaW7KZ80RU/g5nDgeQ7BN84bNqbePai0uLC2x5H9EGiI
ypSQwmx+MXJWKmZArjXWwCUc05FzyxquVvbK9bzQyihqZbRi5dW9/KRVgdrUwRrrsE9Cf7Qe0kZgtsZ6
vqNsjtiyW2SagqTef/1qh0Xl0Jr9SmaYc4mZf6jhch8ktklOiKTc+eOYBHkJTGtWHx4iMmpqCx6ETnjW
DPAbQs+P1NeQ2jqOwZ/N/d4pAEWJw+DDEtn2bJkYrg4Bolzp92yxCrY0hvukXT4ZM7jnsSmdTOhXLvA4
H6sIKVDch+lVG3pxBOpkHd2ev0me9Ns+Om2HPSn6x/TqY+Tmned1sD2Qfdx7soMdFVW5CpoXxTqkNqQ+
jLtXRxjuShq2d6MdjRI1Z4L/g9knt+TShvSz4jLwn9Gq7W7dIbifA23QcQp0hyMuM3y4ygP/0ncFf/49
XIJ/6UMCRAnjI93+uXY3ebflD1b3Exd3Y/nINthHAK2rHqxdSh2u21IdsFmvT96ufeUUPLvyp85m/XuS
KXhTVluM95R1G/fpojb0P9fs0IbtnXFn7B/sf35y2JaDF3uun12+Y/Aib7zHRO/rfwMAAP//4RQuczkJ
AAA=
`,
	},

	"/generator/template/ts/ts_service.gots": {
		local:   "generator/template/ts/ts_service.gots",
		size:    1554,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xUX28bxRd9309xZf0kJ5azm/wkquAoSBEClBeIROB9snttT1nPLDOzTaPVSA3BLZWc
JlKjCkigQVAU0SQGCUrsNvTL7K6dp34FNPsnTiqomKfZuefce86913ZqNasGq20qoUl9BCqhhQwFUejB
2gZUA8EVJwGtZjDMUS5nilAmoSk4U8g8WFpZBpd7CKpNFKxz8RmsU9UG1Ubw6ZogYgOq5Dblslo3jwKb
XGAdqKpKEPh5SAV6ObmAGSmUSUV8Hz2gLEsVCH4TXVVomSjNSlMJ64IqhczAVzcC/NgVNCjRGSYQ/Bb1
UAKBNSKpC6EkLYQmF7kF4vtAmAcdsgEM0QPi3Qyl6iBTQFyXC4+yFigOMkCXNqlbSipN5EjmgaQqJIpy
ZtVg5r8fqwbj/pP00b34+bPR3uP0q914+KAcglWDPJL0usnu0+T+9uhpf/zr1mjvaGllefTtl/HzH0eH
m69e9JLBs/j85WjvaHR8HJ/dTx8NkuHDrLGvXvQg2T9MT366OLgz/nkzfvnduL+ZhZKTr5ODo3j44OKH
wWj/ND47mRS8281z51mvNXfcf5IrLaA7v6Q7u/Ff+7m8pZXlXGHyeJgeHE8Unn5/8U033eomd/9Idvrj
rfNcT3o4SLdPk+6f8fnDXEd+N6HfvojPtpPd3vhOLz4/SO4N0/3f072BVXMs2gm4UJD5qEMES+ayIniH
SgRtFrVTbtZCCY4sAIAoEoS1EOz3Q+aagUmti8AM2MssCJWxCzNa14t3ZJ7WVpnVdjyiyJW0l4v5ifBh
gmqjH6CoLliWU1sjMosmO/3cQ82xXM6kgjKyCJW2UkHDcebe/r89d2Penpt7y74x25ifnZ+tLFhR9D/X
J1J+SDoIjUWw37380toy2v/BFt7OFDaLJ4gi2zC0ngqIIB3ZMC+XlrWeblxvZN4xHxWEwm+AVML8HBav
Gp4qDNQrVxVqXalD5bJcZXohy+Q4oFCq7C5QhYLlE7SNftoEhmB/FBip9qfEDxEqH7y3WgGtAy5VFKEv
UesWqmIkU6Hw6xBFb2RmRktuBPk36CLFdKbFHFu1kU0JlLD4TmG8PI4DbcI8HwGF4OJajDbBkGyzE3YW
nn6NXZ7CcNFbW6D5H3mde42p/0WFDF0XpQTDs95YQnL/Fpoak8R6esEyC5MvtfV3AAAA//+QlIQGEgYA
AA==
`,
	},

	"/generator/template/ts/ts_service.govue": {
		local:   "generator/template/ts/ts_service.govue",
		size:    2168,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xVXW/cRBR996+4WlXaD23svCBVG1oRSkGREIlomveJfXc9rXfGzEfSaLHUENJSKWki
NaqABBoERRHNB0hQkrShf8be3Tz1L6CxdzfeNNDOk2fm3DvnnHtn7FQqVgWmfSqhTgMEKqGBDAVR6MHs
AhRDwRUnIS2mMMxQLmeKUCahLjhTyDwYn5oAl3sIyicK5rm4DfNU+aB8hIDOCiIWoDij8XOUXAsXi1Wz
JbDOBVaBqqIEgV9oKtDLUgyBDS3KpCJBgB5QlqYNBb+FrurxOmOd0qAS5gVVCpmBTy+EeMMVNOyjU0wo
+Bz1UAKBWSKpC1qSBkKdi0wOCQIgzIMmWQCG6AHxbmmpmsgUENflwqOsAYqDDNGlder2KfWlZEjmgaRK
E0U5syow8u7DqkD34Gn78f34xfPOxpP2N+vx8cN+QawKZDvJynKy/ix5sNp5dtD9famzsTM+NdH5/uv4
xc+d7cXXL1eSo+fxyavOxk5ndzc+fNB+fJQcP8rZ+/rlCiSb2+29X0637nZ/XYxf/dA9WMwBkr1vk62d
+Pjh6U9Hnc39+HDv7PB7y9k52QlDRncPnmase9C139pr6/E/mxnV8amJjG3y5Li9tXvGdv/H0++W20vL
yb2/krWD7tJJxqq9fdRe3U+W/45PHmU8sm+z9cdX8eFqsr7SvbsSn2wl94/bm3+2N46simPRZsiFghZc
482QM2SqCjMaITK924TinMaRUPAQhVoY8dDlgiguimP9wJwNuQjR78wBrmUBALRagrAGgv2xZq4puIyi
3oY9wUKtjEFRBFWzMKlVbqUHQ+ZFkdVnZzseUSR3yqDPb4oAzlA+BiGKomXNaLS1xFKOdnnMsj4YqLfw
TprJwzrRgQI3IFIaNtfMx2ekacjgHXOpZWpUJsxxzKWXSmhXcZEu5ealcg9mhtQhilI5nUdWP/pDIhFu
iiCdh4LOEYXm2hkhNZBKmLt0BQq+UmHNcQLuksDnUtUuj14eLfTMueT2SULtCpxRhoHL/21/Jq0UEkGa
snauIOUaTAnepBI/pbfx/fPF+RKuC8HFVWhZA5lmBKhAD/HPlaekfCrtnsRqIc8+igpVKAxIFUyJ8nkd
B6YnP5qsgYcKRZMyBC3NG6eckMtzr495bFCqoQQClRbmmaTSvmQstVutEaB1YAj2ZGissWdIoBEKn1yf
LkAUmbytFgYSo6iBqteIJS2CqrHvLcGpqf3wVjaNejnKtvKRlYb4ZRwlXLmaa5z8MM3FA7QD3igJlOWx
C1GOA2gKAz5hXkBZ40IUrZsUtrlJdgovX3xmzjiG81nJz0deGBhZ/5NqkECgDDmTCESev/1vqouqbyxl
So1jRjc3v89sTYLUrm/Svjc6WgXGgaEyv+C3OptpGrPeyYnCDd5E5ZuumxecNQoXBA47kQNEw8+b9W8A
AAD//yT033t4CAAA
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/generator": {
		isDir: true,
		local: "generator",
	},

	"/generator/template": {
		isDir: true,
		local: "generator/template",
	},

	"/generator/template/ts": {
		isDir: true,
		local: "generator/template/ts",
	},
}
