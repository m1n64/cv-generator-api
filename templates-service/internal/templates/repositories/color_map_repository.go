package repositories

type colorMapRepository struct {
}

func NewColorMapRepository() ColorRepository {
	return &colorMapRepository{}
}

func (r *colorMapRepository) GetColors() (map[string]string, error) {
	return map[string]string{
		"blue":    "#1976d2",
		"cyan":    "#00bcd4",
		"green":   "#4caf50",
		"red":     "#f44336",
		"yellow":  "#b0a100",
		"black":   "#262626",
		"magenta": "#FF00FF",
		"orange":  "#FFA500",
		"purple":  "#a100a1",
		"navy":    "#000080",
		"teal":    "#008080",
		"maroon":  "#800000",
		"violet":  "#4B0082",
		"salmon":  "#FA8072",
		"coral":   "#FF7F50",
	}, nil
}
