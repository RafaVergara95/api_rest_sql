package main

func createFrutas(Frutas Fruta) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO fruta (nombre, color) VALUES (?, ?)", Frutas.Nombre, Frutas.Color)
	return err
}

func deleteFrutas(id int64) error {

	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM fruta WHERE id = ?", id)
	return err
}

// It takes the ID to make the update
func updateFrutas(Frutas Fruta) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE fruta SET nombre = ?, color = ? WHERE id = ?", Frutas.Nombre, Frutas.Color, Frutas.Id)
	return err
}
func getFrutas() ([]Fruta, error) {
	//Declare an array because if there's error, we return it empty
	Frutas := []Fruta{}
	bd, err := getDB()
	if err != nil {
		return Frutas, err
	}
	// Get rows so we can iterate them
	rows, err := bd.Query("SELECT id, nombre, color FROM fruta")
	if err != nil {
		return Frutas, err
	}
	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var Fruta Fruta
		err = rows.Scan(&Fruta.Id, &Fruta.Nombre, &Fruta.Color)
		if err != nil {
			return Frutas, err
		}
		// and append it to the array
		Frutas = append(Frutas, Fruta)
	}
	return Frutas, nil
}

func getFrutasById(id int64) (Fruta, error) {
	var Frutas Fruta
	bd, err := getDB()
	if err != nil {
		return Frutas, err
	}
	row := bd.QueryRow("SELECT id, nombre, color FROM fruta WHERE id = ?", id)
	err = row.Scan(&Frutas.Id, &Frutas.Nombre, &Frutas.Color)
	if err != nil {
		return Frutas, err
	}
	// Success!
	return Frutas, nil
}
