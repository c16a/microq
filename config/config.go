package config

type Config struct {
	Storage *Storage `json:"storage"`
}

type Storage struct {
	RootDir string `json:"root_dir"`
}
