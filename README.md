# NetUserEnum

The code gets a list of local users, using Windows API method [NetUserEnum](https://docs.microsoft.com/en-us/windows/win32/api/lmaccess/nf-lmaccess-netuserenum) from netapi32.dll

Structure USER_INFO_1 is defined according to [Microsoft docs](https://docs.microsoft.com/en-us/windows/win32/api/lmaccess/ns-lmaccess-user_info_1)


```golang

type USER_INFO_1 struct {
	Usri1_name         *uint16
	Usri1_password     *uint16
	Usri1_password_age uint32
	Usri1_priv         uint32
	Usri1_home_dir     *uint16
	Usri1_comment      *uint16
	Usri1_flags        uint32
	Usri1_script_path  *uint16
}
```

The const were also defined according to Microsoft docs.

## License

See the [LICENSE](LICENSE) file for license rights and limitations (MIT).
